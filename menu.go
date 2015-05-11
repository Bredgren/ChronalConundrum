package main

import (
// "github.com/gopherjs/gopherjs/js"
// "github.com/gopherjs/webgl"
)

var (
	// The only instance of menuState we need
	mainMenuState = new(menuState)
)

// The state for handling the game's main menu. Implements fsm.State and mainState.
type menuState struct {
}

func (s *menuState) Name() string {
	return "menuState"
}

func (s *menuState) OnEnter() {
	println("menuState.OnEnter")
	initTest()
}

func (s *menuState) OnExit() {
	println("menuState.OnExit")
	println("not implemented")
}

func (s *menuState) Update() {
	println(input.mouse.leftDown, input.mouse.rightDown)
}

func (s *menuState) Draw() {
	drawTest()
}
