package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shaders
var (
	uiShader     *js.Object
	shipShader   *js.Object
	shaderAssets = []shaderAsset{
		{"ui.vert", "ui.frag", &uiShader},
		{"ship.vert", "ship.frag", &shipShader},
	}
)

type shaderAsset struct {
	vertFile string
	fragFile string
	shader   **js.Object
}

// Textures
var (
	uiTexture     *js.Object
	ship1Texture  *js.Object
	textureAssets = []textureAsset{
		{"ui.png", &uiTexture},
		{"model/ship1.png", &ship1Texture},
	}
)

type textureAsset struct {
	textureFile string
	texture     **js.Object
}

// Sounds

// Models
var (
	ship1Verts  *js.Object
	ship1Faces  *js.Object
	ship1Len    int
	modelAssets = []modelAsset{
		{"model/ship.json", &ship1Verts, &ship1Faces, &ship1Len},
	}
)

type modelAsset struct {
	jsonFile string
	verts    **js.Object
	faces    **js.Object
	len      *int
}

// retrieveFile asyncronously gets the contents of the given file and returns it
// through the given channel.
func retrieveFile(fileName string, contents chan<- string) {
	println("retrieving file", fileName)
	var xmlHttp *js.Object = js.Global.Get("XMLHttpRequest").New()
	xmlHttp.Call("open", "GET", fileName, true)
	xmlHttp.Set("onload", func() {
		go func() {
			if xmlHttp.Get("readyState").Int() == 4 && xmlHttp.Get("status").Int() == 200 {
				contents <- xmlHttp.Get("responseText").String()
			} else {
				fail("Failed to retrieve file " + fileName)
			}
		}()
	})
	xmlHttp.Call("send")
}

// retrieveImage asyncronously gets the image specified and returns it through the
// image channel.
func retrieveImage(fileName string, image chan<- *js.Object) {
	println("retrieving image", fileName)
	var img *js.Object = js.Global.Get("Image").New()
	img.Set("onload", func(*js.Object) {
		go func() { image <- img }()
	})
	img.Set("onerror", func(o *js.Object) {
		go func() {
			fail("Failed to load texture " + fileName + ": " + o.String())
		}()
	})
	img.Set("src", fileName)
}
