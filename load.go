package main

import (
	"github.com/gopherjs/gopherjs/js"
// "github.com/gopherjs/webgl"
)

var (
	// The only instance of loadState we need
	mainLoadState = new(loadState)
)

// The state for loading the game's assets. Implements fsm.State and mainState.
type loadState struct {
	totalAssets int
	assetsLoaded int
}

func (s *loadState) Name() string {
	return "loadState"
}

func (s *loadState) OnEnter() {
	println("loadState.OnEnter")

	shaderAssets := []shaderAsset{
		shaderAsset{TEST_SHADER_FILE, &testShader},
	}
	s.totalAssets = len(shaderAssets)

	// textureAssets := []textureAsset{...}
	// s.totalAssets += len(textureAssets)

	// soundAssets := []soundAsset{...}
	// s.totalAssets += len(soundAssets)

	for _, asset := range shaderAssets {
		// TODO asyncronous
		loadShaderAsset(s, &asset)
	}
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	println("not implemented")
}

func (s *loadState) Update() {
	percent := float64(s.assetsLoaded) / float64(s.totalAssets) * 100.0
	println("loading... ", percent, "%")

	if s.assetsLoaded == s.totalAssets {
		mainSm.GotoState(mainMenuState)
		return
	}
}

func (s *loadState) Draw() {
}


type shaderAsset struct {
	spec shaderSpec
	shader **js.Object
}

func loadShaderAsset(s *loadState, asset *shaderAsset) {
	// TODO load from files
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
		js.Global.Call("alert", "Error compiling vertex shaders: "+gl.GetShaderInfoLog(vertShader))
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
		js.Global.Call("alert", "Error compiling fragment shaders: "+gl.GetShaderInfoLog(fragShader))
		fragShader = nil
	}

	shader := gl.CreateProgram()
	gl.AttachShader(shader, vertShader)
	gl.AttachShader(shader, fragShader)
	gl.LinkProgram(shader)

	if !gl.GetProgramParameterb(shader, gl.LINK_STATUS) {
		js.Global.Call("alert", "Unable to initialize the shader program.")
	}

	*asset.shader = shader
	s.assetsLoaded += 1
}
