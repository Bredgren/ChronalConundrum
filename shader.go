package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// createShader creates a new gl shader and compiles it the returns the shader object
func createShader(source string, kind int) *js.Object {
	var shader *js.Object = gl.CreateShader(kind)
	gl.ShaderSource(shader, source)
	gl.CompileShader(shader)
	if !gl.GetShaderParameterb(shader, gl.COMPILE_STATUS) {
		fail("Error compiling shader: " + gl.GetShaderInfoLog(shader))
		return nil
	}
	return shader
}

// loadShaderAsset takes the given shaderAsset, loads the source files it specifies
// and returns the files names separated by a space via the done channel.
func loadShaderAsset(asset *shaderAsset, done chan<- string) {
	// TODO: cache files if more than one program needs the same shader
	vert := make(chan string)
	frag := make(chan string)

	defer close(vert)
	defer close(frag)

	retrieveFile(asset.vertFile, vert)
	retrieveFile(asset.fragFile, frag)

	vertSource := <-vert
	if vertSource == "" {
		fail("Failed to load asset " + asset.vertFile)
		return
	}
	fragSource := <-frag
	if fragSource == "" {
		fail("Failed to load asset " + asset.fragFile)
		return
	}

	var vertShader *js.Object = createShader(vertSource, gl.VERTEX_SHADER)
	var fragShader *js.Object = createShader(fragSource, gl.FRAGMENT_SHADER)

	shader := gl.CreateProgram()
	gl.AttachShader(shader, vertShader)
	gl.AttachShader(shader, fragShader)
	gl.LinkProgram(shader)

	if !gl.GetProgramParameterb(shader, gl.LINK_STATUS) {
		fail("Unable to initiaize shader program")
		return
	}

	*asset.shader = shader
	// TODO: exception is thrown about sending to closed channel because this triggers
	//       the loadState to exit and the js code created here access the channel
	//       afterward. (may not be a serious problem though)
	done <- asset.vertFile + " " + asset.fragFile
}
