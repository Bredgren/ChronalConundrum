package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	WINDOW_RATIO = 2.0 / 3.0
	VIEW_DEGREES = 45
)

var (
	VIEW_ANGLE = mgl32.DegToRad(VIEW_DEGREES)
)

// UI - vector values are x, y, w, h
var (
	MENU_NEW_RECT = mgl32.Vec4{-0.75, 1.0, 0.5, 0.125}
	MENU_NEW_IMG  = mgl32.Vec4{0.0, 0.75, 1.0, 0.25}

	MENU_CONT_RECT = mgl32.Vec4{0.75, 1.0, 0.5, 0.125}
	MENU_CONT_IMG  = mgl32.Vec4{0.0, 0.5, 1.0, 0.25}
)
