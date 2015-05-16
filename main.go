package main

import (
	"time"

	"github.com/Bredgren/fsm"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
)

// Global state
var (
	document *js.Object
	canvas   *js.Object
	gl       *webgl.Context
	prevTime time.Time
	// mainSm is the overarching state machine for the game. Possible states are
	// initState, loadState, menuState, playState, failedState
	mainSm *fsm.Fsm
)

func clearWindow() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func onWindowResize() {
	height := js.Global.Get("innerHeight").Int()
	width := int(float64(height) * WINDOW_RATIO)
	canvas.Set("width", width)
	canvas.Set("height", height)

	gl.Viewport(0, 0, width, height)
	clearWindow()
}

func mainLoop(timestamp float64) {
	currentState := mainSm.CurrentState.(mainState)
	currentState.Update(timestamp)
	if mainSm.CurrentState.(mainState) != currentState {
		// We switched states in the update
		currentState = mainSm.CurrentState.(mainState)
		currentState.Update(timestamp)
	}
	clearWindow()
	currentState.Draw(timestamp)
	gl.Flush()

	if currentState != mainFailedState {
		js.Global.Call("requestAnimationFrame", mainLoop)
	}
}

func onBodyLoad() {
	mainSm = fsm.NewFsm(mainInitState)

	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)

	onWindowResize()

	js.Global.Call("requestAnimationFrame", mainLoop)
}

func main() {
	js.Global.Set("onBodyLoad", onBodyLoad)
	js.Global.Call("addEventListener", "resize", onWindowResize)
}
