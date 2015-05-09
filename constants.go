package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

const (
	WINDOW_RATIO = 2.0 / 3.0
	VIEW_DEGREES = 45
)

var (
	VIEW_ANGLE = mgl32.DegToRad(VIEW_DEGREES)
)


// Assets

type shaderSpec struct {
	vertFile string
	fragFile string
}

var (
	TEST_SHADER_FILE = shaderSpec{"test.vert", "test.frag"}
	testShader *js.Object
)
