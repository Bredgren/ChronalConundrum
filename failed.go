package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	mainFailedState = new(failedState)
)

type failedState struct {
	reason string
}

func (s *failedState) Name() string {
	return "failedState"
}

func (s *failedState) OnEnter() {
	println("failedState.OnEnter")
	js.Global.Call("alert", "Oh dear, the game seems to have failed.\n\nReason:\n"+s.reason)
}

func (s *failedState) OnExit() {
	println("failedState.OnExit")
}

func (s *failedState) Update(timestamp float64) {
}

func (s *failedState) Draw(timestamp float64) {
}

func fail(message string) {
	mainFailedState.reason = message
	mainSm.GotoState(mainFailedState)
}
