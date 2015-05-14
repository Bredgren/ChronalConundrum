package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gopherjs/gopherjs/js"
)

var (
	// The only instance of menuState we need
	mainMenuState = new(menuState)
)

// The state for handling the game's main menu. Implements fsm.State and mainState.
type menuState struct {
	newButton *menuNewButton
}

func (s *menuState) Name() string {
	return "menuState"
}

func (s *menuState) OnEnter() {
	println("menuState.OnEnter")
	s.newButton = newMenuNewButton()
}

func (s *menuState) OnExit() {
	println("menuState.OnExit")
	println("not implemented")
}

func (s *menuState) Update() {
}

func (s *menuState) Draw() {
	s.newButton.draw()
}

type menuNewButton struct {
	vertBuffer     *js.Object
	texCoordBuffer *js.Object
	vertAttr       int
	texAttr        int
}

func newMenuNewButton() *menuNewButton {
	var w float32 = 1.0 / 2.0
	var h float32 = w / 4.0
	vertices := []float32{
		MENU_NEW_POS_X - w, MENU_NEW_POS_Y - h,
		MENU_NEW_POS_X + w, MENU_NEW_POS_Y - h,
		MENU_NEW_POS_X - w, MENU_NEW_POS_Y + h,
		MENU_NEW_POS_X + w, MENU_NEW_POS_Y + h,
	}

	texCoords := []float32{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 1.0,
		1.0, 1.0,
	}

	var buf *js.Object = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, buf)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	var tex *js.Object = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, tex)
	gl.BufferData(gl.ARRAY_BUFFER, texCoords, gl.STATIC_DRAW)

	var vertAttr int = gl.GetAttribLocation(uiShader, "aVertexPosition")
	gl.EnableVertexAttribArray(vertAttr)

	var texAttr int = gl.GetAttribLocation(uiShader, "aTextureCoord")
	gl.EnableVertexAttribArray(texAttr)

	return &menuNewButton{buf, tex, vertAttr, texAttr}
}

func (b *menuNewButton) draw() {
	gl.UseProgram(uiShader)

	gl.BindBuffer(gl.ARRAY_BUFFER, b.vertBuffer)
	gl.VertexAttribPointer(b.vertAttr, 2, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, b.texCoordBuffer)
	gl.VertexAttribPointer(b.texAttr, 2, gl.FLOAT, false, 0, 0)

	perspectiveMatrix = mgl32.Perspective(VIEW_ANGLE, WINDOW_RATIO, 0.1, 100.0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(uiShader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(uiShader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, uiTexture)
	gl.Uniform1i(gl.GetUniformLocation(uiShader, "uSampler"), 0)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
