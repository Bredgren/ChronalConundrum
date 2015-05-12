package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

const (
	LEFT_BUTTON = 0
	RIGHT_BUTTON = 2
)

var (
	input = new(inputType)
)

type inputType struct {
	mouse
}

type mouse struct {
	pos       mgl32.Vec2
	leftDown  bool
	rightDown bool
}

func onMouseMove(event *js.Object) {
	// Find mouse position relative to the canvas
	objLeft := 0
	objTop := 0
	for obj := canvas; obj.Get("offsetParent") != nil; obj = obj.Get("offsetParent") {
		objLeft += obj.Get("offsetLeft").Int()
		objTop += obj.Get("offsetTop").Int()
	}
	input.mouse.pos = mgl32.Vec2{
		float32(event.Get("pageX").Int() - objLeft),
		float32(event.Get("pageY").Int() - objTop),
	}
}

func onMouseDown(event *js.Object) {
	button := event.Get("button").Int()
	if button == LEFT_BUTTON {
		input.mouse.leftDown = true
	} else if button == RIGHT_BUTTON {
		input.mouse.rightDown = true
	}
}

func onMouseUp(event *js.Object) {
	button := event.Get("button").Int()
	if button == LEFT_BUTTON {
		input.mouse.leftDown = false
	} else if button == RIGHT_BUTTON {
		input.mouse.rightDown = false
	}
}
