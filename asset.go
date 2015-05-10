package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shaders
var (
	testShader   *js.Object
	tmpShader    *js.Object
	shaderAssets = []shaderAsset{
		{"test.vert", "test.frag", &testShader},
		// {"tmp.vert", "tmp.frag", &tmpShader},
	}
)

type shaderAsset struct {
	vertFile string
	fragFile string
	shader   **js.Object
}

// Textures

// Sounds

// Models
