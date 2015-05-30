package main

import (
	// "math/rand"
	// "strconv"
	// "time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

var (
	// The only instance of loadState we need
	mainLoadState = new(loadState)
)

// The state for loading the game's assets. Implements fsm.State and mainState.
type loadState struct {
	totalAssets  int
	assetsLoaded int
	loadChannel  chan string
	bar          *loadBar
}

func (s *loadState) Name() string {
	return "loadState"
}

func (s *loadState) OnEnter() {
	println("loadState.OnEnter")
	s.bar = newLoadBar()

	if s.loadChannel == nil {
		s.loadChannel = make(chan string)
	}

	s.totalAssets = len(shaderAssets)
	s.totalAssets += len(textureAssets)
	// s.totalAssets += len(soundAssets)
	s.totalAssets += len(modelAssets)
	// fakeCount := 10
	// s.totalAssets += fakeCount

	for _, asset := range shaderAssets {
		go loadShaderAsset(&asset, s.loadChannel)
	}

	for _, asset := range textureAssets {
		go loadTextureAsset(&asset, s.loadChannel)
	}

	for _, asset := range modelAssets {
		go loadModelAsset(&asset, s.loadChannel)
	}

	// for i := 0; i < fakeCount; i++ {
	// 	go loadFake(i, s.loadChannel)
	// }
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	// TODO: closing this causes an exception to be thrown about sending to closed
	//       channel by who ever is the last to load.
	// close(s.loadChannel)
}

func (s *loadState) Update(timestamp float64) {
	select {
	case loaded := <-s.loadChannel:
		println("loaded", loaded)
		s.assetsLoaded += 1
	default:
	}

	if s.assetsLoaded == s.totalAssets {
		mainSm.GotoState(mainMenuState)
		return
	}
}

func (s *loadState) Draw(timestamp float64) {
	percent := float64(s.assetsLoaded) / float64(s.totalAssets)
	s.bar.draw(percent, timestamp)
}

// func loadFake(i int, done chan<- string) {
// 	time.Sleep(time.Duration(rand.Float64() * 3.0 + 0.5) * time.Second)
// 	done <- strconv.Itoa(i)
// }

type loadBar struct {
	vertBuffer *js.Object
	shader     *js.Object
	posAttr    int
}

func newLoadBar() *loadBar {
	w := float32(canvas.Get("width").Float())
	h := float32(canvas.Get("height").Float())
	rect := mgl32.Vec4{-1.0, -1.0, w, h}
	vertices := []float32{
		rect.X() - rect[2]/2.0, rect.Y() - rect[3]/2.0,
		rect.X() + rect[2]/2.0, rect.Y() - rect[3]/2.0,
		rect.X() - rect[2]/2.0, rect.Y() + rect[3]/2.0,
		rect.X() + rect[2]/2.0, rect.Y() + rect[3]/2.0,
	}

	var buf *js.Object = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, buf)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	var vertShader *js.Object = createShader(`
attribute vec3 aVertexPosition;

uniform mat4 uMVMatrix;
uniform mat4 uPMatrix;

void main(void) {
  gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
}`, gl.VERTEX_SHADER)
	var fragShader *js.Object = createShader(`
precision mediump float;

uniform float uPercent;
uniform vec2 uResolution;
uniform float uTime;

vec3 COLOR = vec3(1.0, 1.0, 1.0);

float TAU = 6.283185307179586476;

float distanceIntensity(float distance, float angle) {
  distance = distance + 5.0 * sin(uTime * 0.11 + angle * 10.0);
  distance = distance + (2.0 * uPercent + 1.0) * sin(uTime * 0.01);
  float radius = 100.0;
  float offset = abs(radius - distance);
  float thickness = 10.0;
  return offset / thickness;
}

float angleIntensity(float angle) {
  float edge = uPercent * TAU;
  return step(edge, angle);
}

void main(void) {
  vec2 pos = gl_FragCoord.xy;
  vec2 center = uResolution / 2.0;
  float dist = abs(length(pos - center));
  vec2 dir = pos - center;

  float speed = 0.01;
  float angle = mod(TAU - atan(dir.y, dir.x) - uTime * speed, TAU);

  float distInt = distanceIntensity(dist, angle);
  float angleInt = angleIntensity(angle);
  vec3 color = vec3(1.0) - (vec3(1.0) * distInt) - (vec3(1.0) * angleInt);

  gl_FragColor = vec4(color, 1.0);
}`, gl.FRAGMENT_SHADER)

	var shader *js.Object = gl.CreateProgram()
	gl.AttachShader(shader, vertShader)
	gl.AttachShader(shader, fragShader)
	gl.LinkProgram(shader)

	if !gl.GetProgramParameterb(shader, gl.LINK_STATUS) {
		fail("Unable to initiaize shader program")
		return &loadBar{}
	}

	var posAttr int = gl.GetAttribLocation(shader, "aVertexPosition")
	gl.EnableVertexAttribArray(posAttr)

	return &loadBar{buf, shader, posAttr}
}

func (b *loadBar) draw(percent, timestamp float64) {
	gl.UseProgram(b.shader)

	gl.BindBuffer(gl.ARRAY_BUFFER, b.vertBuffer)
	gl.VertexAttribPointer(b.posAttr, 2, gl.FLOAT, false, 0, 0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(b.shader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(b.shader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	gl.Uniform1f(gl.GetUniformLocation(b.shader, "uPercent"), float32(percent))
	w := float32(canvas.Get("width").Float())
	h := float32(canvas.Get("height").Float())
	gl.Uniform2f(gl.GetUniformLocation(b.shader, "uResolution"), w, h)
	gl.Uniform1f(gl.GetUniformLocation(b.shader, "uTime"), float32(timestamp))

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
