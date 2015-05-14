package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

const (
	WINDOW_RATIO = 2.0 / 3.0
	VIEW_DEGREES = 45
	// UI
	MENU_NEW_POS_X float32 = 0.0
	MENU_NEW_POS_Y float32 = 0.0
)

var (
	VIEW_ANGLE = mgl32.DegToRad(VIEW_DEGREES)
)
