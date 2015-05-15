package main

import (
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
}

func (s *loadState) Name() string {
	return "loadState"
}

func (s *loadState) OnEnter() {
	println("loadState.OnEnter")
	initLoadBar()

	if s.loadChannel == nil {
		s.loadChannel = make(chan string)
	}

	s.totalAssets = len(shaderAssets)
	s.totalAssets += len(textureAssets)
	// s.totalAssets += len(soundAssets)
	// s.totalAssets += len(modelAssets)

	for _, asset := range shaderAssets {
		go loadShaderAsset(&asset, s.loadChannel)
	}

	for _, asset := range textureAssets {
		go loadTextureAsset(&asset, s.loadChannel)
	}
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	// TODO: closing this causes an exception to be thrown about sending to closed
	//       channel by who ever is the last to load.
	// close(s.loadChannel)
}

func (s *loadState) Update() {
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

func (s *loadState) Draw() {
	percent := float64(s.assetsLoaded) / float64(s.totalAssets) * 100.0
	// println("loading... ", percent, "%")
	drawLoadBar(percent)
}

var (
	squareVerticesBuffer *js.Object
	vPositionAttr        int
	mvMatrix             mgl32.Mat4
	loadShader               *js.Object
)

func initLoadBar() {
	squareVerticesBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, squareVerticesBuffer)

	vertices := []float32{
		1.0, 1.0, 0.0,
		-1.0, 1.0, 0.0,
		-1.0, -1.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	var vertShader *js.Object = createShader(`attribute vec3 aVertexPosition;

uniform mat4 uMVMatrix;
uniform mat4 uPMatrix;

void main(void) {
  gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
}`, gl.VERTEX_SHADER)
	var fragShader *js.Object = createShader(`void main(void) {
  gl_FragColor = vec4(0.5, 1.0, 1.0, 1.0);
}`, gl.FRAGMENT_SHADER)

	loadShader = gl.CreateProgram()
	gl.AttachShader(loadShader, vertShader)
	gl.AttachShader(loadShader, fragShader)
	gl.LinkProgram(loadShader)

	if !gl.GetProgramParameterb(loadShader, gl.LINK_STATUS) {
		fail("Unable to initiaize shader program")
		return
	}

	gl.UseProgram(loadShader)

	vPositionAttr = gl.GetAttribLocation(loadShader, "aVertexPosition")
	gl.EnableVertexAttribArray(vPositionAttr)
}

func drawLoadBar(percent float64) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareVerticesBuffer)
	gl.VertexAttribPointer(vPositionAttr, 3, gl.FLOAT, false, 0, 0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(loadShader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(loadShader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 3)
}
