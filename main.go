package main

import (
	"github.com/Bredgren/fsm"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

// Global state
var (
	document *js.Object
	canvas   *js.Object
	gl       *webgl.Context
	// mainSm is the overarching state machine for the game. Possible states are
	// initState, loadState, menuState, playState, failedState
	mainSm *fsm.Fsm
)

func clearWindow() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func onWindowResize() {
	height := js.Global.Get("innerHeight").Int()
	width := int(float64(height) * WINDOW_RATIO)
	canvas.Set("width", width)
	canvas.Set("height", height)

	gl.Viewport(0, 0, width, height)
	clearWindow()
}

func mainLoop() {
	currentState := mainSm.CurrentState.(mainState)
	currentState.Update()
	if mainSm.CurrentState.(mainState) != currentState {
		// We switched states in the update
		currentState = mainSm.CurrentState.(mainState)
		currentState.Update()
	}
	clearWindow()
	currentState.Draw()

	if currentState != mainFailedState {
		js.Global.Call("requestAnimationFrame", mainLoop)
	}
}

func onBodyLoad() {
	mainSm = fsm.NewFsm(mainInitState)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)

	onWindowResize()

	js.Global.Call("requestAnimationFrame", mainLoop)
}

func main() {
	js.Global.Set("onBodyLoad", onBodyLoad)
	js.Global.Call("addEventListener", "resize", onWindowResize)
}

// For getting started

var (
	squareVerticesBuffer *js.Object
	vPositionAttr        int
	perspectiveMatrix    mgl32.Mat4
	mvMatrix             mgl32.Mat4
	shader *js.Object
)

func initTest() {
	println("initTest")
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

	shader = gl.CreateProgram()
	gl.AttachShader(shader, vertShader)
	gl.AttachShader(shader, fragShader)
	gl.LinkProgram(shader)

	if !gl.GetProgramParameterb(shader, gl.LINK_STATUS) {
		fail("Unable to initiaize shader program")
		return
	}

	gl.UseProgram(shader)

	vPositionAttr = gl.GetAttribLocation(shader, "aVertexPosition")
	gl.EnableVertexAttribArray(vPositionAttr)
}

func drawTest() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareVerticesBuffer)
	gl.VertexAttribPointer(vPositionAttr, 3, gl.FLOAT, false, 0, 0)

	perspectiveMatrix = mgl32.Perspective(VIEW_ANGLE, WINDOW_RATIO, 0.1, 100.0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(shader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(shader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 3)
}
