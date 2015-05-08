package main

import (
// "github.com/gopherjs/gopherjs/js"
// "github.com/gopherjs/webgl"
)

var (
	// The only instance of loadState we need
	mainLoadState = new(loadState)
)

// The state for loading the game's assets. Implements fsm.State and mainState.
type loadState struct {
}

func (s *loadState) Name() string {
	return "loadState"
}

func (s *loadState) OnEnter() {
	println("loadState.OnEnter")
	println("not implemented")
}

func (s *loadState) OnExit() {
	println("loadState.OnExit")
	println("not implemented")
}

func (s *loadState) Update() {
}

func (s *loadState) Draw() {
}
