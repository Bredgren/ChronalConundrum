package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
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
}

func main() {
	js.Global.Set("onBodyLoad", onBodyLoad)
	js.Global.Call("addEventListener", "resize", onWindowResize)
}
