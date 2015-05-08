package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

var (
	// The only instance of initState we need
	mainInitState = new(initState)
)

// The state for initializing the program. Implements fsm.State and mainState.
type initState struct {
}

func (s *initState) Name() string {
	return "initState"
}

func (s *initState) OnEnter() {
	println("initState.OnEnter")
	document = js.Global.Get("document")
	initCanvas()
	initWebGl()
}

func (s *initState) OnExit() {
	println("initState.OnExit")
	println("not implemented")
}

func (s *initState) Update() {
}

func (s *initState) Draw() {
	drawTest()
}

func initCanvas() {
	println("initCanvas")
	canvas = document.Call("createElement", "canvas")
	canvas.Get("style").Set("margin", "auto")
	canvas.Get("style").Set("display", "inherit")
	document.Get("body").Call("appendChild", canvas)
}

func initWebGl() {
	println("initWebGl")
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
