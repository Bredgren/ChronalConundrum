package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

var (
	input = new(inputType)
)

type inputType struct {
	mouse
}

type mouse struct {
	pos mgl32.Vec2
	leftDown bool
	rightDown bool
}

func onMouseMove(event *js.Object) {
	// TODO: figure out cross-browser support (at least for firefox)

	// client := mgl32.Vec2{
	// 	float32(event.Get("clientX").Int()),
	// 	float32(event.Get("clientY").Int()),
	// }
	layer := mgl32.Vec2{
		float32(event.Get("layerX").Int()),
		float32(event.Get("layerY").Int()),
	}
	// offset := mgl32.Vec2{
	// 	float32(event.Get("offsetX").Int()),
	// 	float32(event.Get("offsetY").Int()),
	// }
	// page := mgl32.Vec2{
	// 	float32(event.Get("pageX").Int()),
	// 	float32(event.Get("pageY").Int()),
	// }
	// screen := mgl32.Vec2{
	// 	float32(event.Get("screenX").Int()),
	// 	float32(event.Get("screenY").Int()),
	// }
	// xy := mgl32.Vec2{
	// 	float32(event.Get("x").Int()),
	// 	float32(event.Get("y").Int()),
	// }

	// println("mousemove", client, layer, offset, page, screen, xy)
	input.mouse.pos = layer
}

func onMouseDown(event *js.Object) {
	println(event)
	button := event.Get("which").Int()
	if button == 1 {
		input.mouse.leftDown = true
	} else if button == 3 {
		input.mouse.rightDown = true
	}
}

func onMouseUp(event *js.Object) {
	button := event.Get("which").Int()
	if button == 1 {
		input.mouse.leftDown = false
	} else if button == 2 {
		input.mouse.rightDown = false
	}
}
