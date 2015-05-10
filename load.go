package main

import (
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

	s.loadChannel = make(chan string)

	s.totalAssets = len(shaderAssets)
	// s.totalAssets += len(textureAssets)
	// s.totalAssets += len(soundAssets)
	// s.totalAssets += len(modelAssets)

	for _, asset := range shaderAssets {
		go loadShaderAsset(s.loadChannel, &asset)
	}
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	close(s.loadChannel)
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
	println("loading... ", percent, "%")
}

func loadShaderAsset(done chan<- string, asset *shaderAsset) {
	// TODO: cache files if more than one program needs the same shader
	vert := make(chan string)
	frag := make(chan string)

	defer close(vert)
	defer close(frag)

	retrieveFile(asset.vertFile, vert)
	retrieveFile(asset.fragFile, frag)

	vertSource := <-vert
	if vertSource == "" {
		mainFailedState.reason = "Failed to load asset " + asset.vertFile
		mainSm.GotoState(mainFailedState)
		return
	}
	fragSource := <-frag
	if fragSource == "" {
		mainFailedState.reason = "Failed to load asset " + asset.fragFile
		mainSm.GotoState(mainFailedState)
		return
	}

	var vertShader *js.Object = gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertShader, vertSource)
	gl.CompileShader(vertShader)
	if !gl.GetShaderParameterb(vertShader, gl.COMPILE_STATUS) {
		js.Global.Call("alert", "Error compiling vertex shaders: "+gl.GetShaderInfoLog(vertShader))
		vertShader = nil
	}

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
	// TODO: exception is thrown about sending to closed channel because this triggers
	//       the loadState to exit
	done <- asset.vertFile + " " + asset.fragFile
}
