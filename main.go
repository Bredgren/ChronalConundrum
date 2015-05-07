package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	document *js.Object
	canvas   *js.Object
	gl       *webgl.Context
)

func initCanvas() {
	canvas = document.Call("createElement", "canvas")
	canvas.Get("style").Set("margin", "auto")
	canvas.Get("style").Set("display", "inherit")
	document.Get("body").Call("appendChild", canvas)
}

func initWebGl() {
	attrs := webgl.DefaultAttributes()
	attrs.Alpha = false

	glcontext, err := webgl.NewContext(canvas, attrs)
	if err != nil {
		js.Global.Call("alert", "Error: "+err.Error())
	}
	gl = glcontext

	println(gl.GetParameter(gl.VERSION))
	println(gl.GetParameter(gl.SHADING_LANGUAGE_VERSION))
}

func onWindowResize() {
	height := js.Global.Get("innerHeight").Int()
	width := int(float64(height) * VIEW_RATIO)
	canvas.Set("width", width)
	canvas.Set("height", height)

	gl.Viewport(0, 0, width, height)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func onBodyLoad() {
	document = js.Global.Get("document")
	initCanvas()
	initWebGl()
	onWindowResize()

	initTest()
	js.Global.Call("requestAnimationFrame", drawTest)
}

func main() {
	js.Global.Set("onBodyLoad", onBodyLoad)
	js.Global.Call("addEventListener", "resize", onWindowResize)
}

// For getting started

var (
	squareVerticesBuffer *js.Object
	vPositionAttr int
	perspectiveMatrix mgl32.Mat4
	shaderProgram *js.Object
	mvMatrix mgl32.Mat4
)

func initTest() {
	squareVerticesBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, squareVerticesBuffer)

	vertices := []float32{
		1.0, 1.0, 0.0,
		-1.0, 1.0, 0.0,
		-1.0, -1.0, 0.0,
	}

	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	vertSource := `
		attribute vec3 aVertexPosition;

		uniform mat4 uMVMatrix;
		uniform mat4 uPMatrix;

		void main(void) {
		  gl_Position = uPMatrix * uMVMatrix * vec4(aVertexPosition, 1.0);
		}`
	var vertShader *js.Object = gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertShader, vertSource)
	gl.CompileShader(vertShader)
	if !gl.GetShaderParameterb(vertShader, gl.COMPILE_STATUS) {
		js.Global.Call("alert", "Error compiling vertex shaders: " + gl.GetShaderInfoLog(vertShader))
		vertShader = nil
	}

	fragSource := `
		void main(void) {
  		gl_FragColor = vec4(0.5, 1.0, 1.0, 1.0);
		}`
	var fragShader *js.Object = gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragShader, fragSource)
	gl.CompileShader(fragShader)
	if !gl.GetShaderParameterb(fragShader, gl.COMPILE_STATUS) {
		js.Global.Call("alert", "Error compiling fragment shaders: " + gl.GetShaderInfoLog(fragShader))
		fragShader = nil
	}

	shaderProgram = gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertShader)
	gl.AttachShader(shaderProgram, fragShader)
	gl.LinkProgram(shaderProgram)

	if !gl.GetProgramParameterb(shaderProgram, gl.LINK_STATUS) {
		js.Global.Call("alert", "Unable to initialize the shader program.")
	}

	gl.UseProgram(shaderProgram)

	vPositionAttr = gl.GetAttribLocation(shaderProgram, "aVertexPosition")
	gl.EnableVertexAttribArray(vPositionAttr)
}

func drawTest() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.BindBuffer(gl.ARRAY_BUFFER, squareVerticesBuffer)
	gl.VertexAttribPointer(vPositionAttr, 3, gl.FLOAT, false, 0, 0)

	perspectiveMatrix = mgl32.Perspective(mgl32.DegToRad(45), 640.0 / 480.0, 0.1, 100.0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(shaderProgram, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(shaderProgram, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 3)

	js.Global.Call("requestAnimationFrame", drawTest)
}
