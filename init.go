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

	canvas.Call("addEventListener", "mousemove", onMouseMove)
	canvas.Call("addEventListener", "mousedown", onMouseDown)
	canvas.Call("addEventListener", "mouseup", onMouseUp)
	// TODO: disable context menu
	// canvas.Call("addEventListener", "oncontextmenu", func(e *js.Object) {
	// 	e.Call("preventDefault")
	// })
}

func (s *initState) OnExit() {
	println("initState.OnExit")
	println("not implemented")
}

func (s *initState) Update() {
	// Must do this here instead of the end of OnEnter because mainSm isn't initalized
	mainSm.GotoState(mainLoadState)
}

func (s *initState) Draw() {
	// Don't need to draw anything
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
