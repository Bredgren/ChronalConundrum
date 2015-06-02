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
	buttons [2]*menuButton
}

func (s *menuState) Name() string {
	return "menuState"
}

func (s *menuState) OnEnter() {
	println("menuState.OnEnter")
	s.buttons[0] = newMenuButton(MENU_NEW_RECT, MENU_NEW_IMG)
	s.buttons[1] = newMenuButton(MENU_CONT_RECT, MENU_CONT_IMG)
}

func (s *menuState) OnExit() {
	println("menuState.OnExit")
	println("not implemented")
}

func (s *menuState) Update(timestamp float64) {
}

func (s *menuState) Draw(timestamp float64) {
	s.drawButtons()
	s.drawShip()
}

func (s *menuState) drawButtons() {
	gl.UseProgram(uiShader)

	for _, button := range s.buttons {
		button.draw()
	}
}

func (s *menuState) drawShip() {
	gl.UseProgram(shipShader)

	var vertAttr int = gl.GetAttribLocation(shipShader, "aVertexPosition")
	gl.EnableVertexAttribArray(vertAttr)
	var texAttr int = gl.GetAttribLocation(shipShader, "aUV")
	gl.EnableVertexAttribArray(texAttr)
	var normAttr int = gl.GetAttribLocation(shipShader, "aNormal")
	gl.EnableVertexAttribArray(normAttr)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))
	mvMatrix = mvMatrix.Mul4(mgl32.Scale3D(0.15, 0.15, 0.15))
	// mvMatrix = mvMatrix.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(20), mgl32.Vec3{1, 1, 0}))

	var pUniform *js.Object = gl.GetUniformLocation(shipShader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(shipShader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	// TODO cache attribs/uniforms, and precalculate stride/offset

	gl.BindBuffer(gl.ARRAY_BUFFER, ship1Verts)
	var position int = gl.GetAttribLocation(shipShader, "aVertexPosition")
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 4*(3+3+2), 0)
	var normal int = gl.GetAttribLocation(shipShader, "aNormal")
	gl.VertexAttribPointer(normal, 2, gl.FLOAT, false, 4*(3+3+2), 3*4)
	var uv int = gl.GetAttribLocation(shipShader, "aUV")
	gl.VertexAttribPointer(uv, 2, gl.FLOAT, false, 4*(3+3+2), (3+3)*4)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ship1Faces)
	gl.DrawElements(gl.TRIANGLES, ship1Len, gl.UNSIGNED_SHORT, 0)
}

type menuButton struct {
	vertBuffer     *js.Object
	texCoordBuffer *js.Object
	vertAttr       int
	texAttr        int
	img            mgl32.Vec4
}

func newMenuButton(rect, img mgl32.Vec4) *menuButton {
	vertices := []float32{
		rect.X() - rect[2], rect.Y() - rect[3],
		rect.X() + rect[2], rect.Y() - rect[3],
		rect.X() - rect[2], rect.Y() + rect[3],
		rect.X() + rect[2], rect.Y() + rect[3],
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

	var texAttr int = gl.GetAttribLocation(uiShader, "aUV")
	gl.EnableVertexAttribArray(texAttr)

	return &menuButton{buf, tex, vertAttr, texAttr, img}
}

func (b *menuButton) draw() {
	// TODO cache attribs/uniforms
	gl.BindBuffer(gl.ARRAY_BUFFER, b.vertBuffer)
	gl.VertexAttribPointer(b.vertAttr, 2, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, b.texCoordBuffer)
	gl.VertexAttribPointer(b.texAttr, 2, gl.FLOAT, false, 0, 0)

	mvMatrix = mgl32.Ident4()
	mvMatrix = mvMatrix.Mul4(mgl32.Translate3D(0.0, 0.0, -6.0))

	var pUniform *js.Object = gl.GetUniformLocation(uiShader, "uPMatrix")
	pm := [16]float32(perspectiveMatrix)
	gl.UniformMatrix4fv(pUniform, false, pm[:])

	var mvUniform *js.Object = gl.GetUniformLocation(uiShader, "uMVMatrix")
	mvm := [16]float32(mvMatrix)
	gl.UniformMatrix4fv(mvUniform, false, mvm[:])

	// gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, uiTexture)
	// TODO: does uiTexture.Int() actually work?
	gl.Uniform1i(gl.GetUniformLocation(uiShader, "uTexture"), uiTexture.Int())

	gl.Uniform4f(gl.GetUniformLocation(uiShader, "uRect"), b.img[0], b.img[1], b.img[2], b.img[3])

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
