package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shaders
var (
	uiShader   *js.Object
	shaderAssets = []shaderAsset{
		{"ui.vert", "ui.frag", &uiShader},
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
	textureAssets = []textureAsset{
		{"ui.png", &uiTexture},
	}
)

type textureAsset struct {
	textureFile string
	texture     **js.Object
}

// Sounds

// Models

// retrieveFile asyncronously gets the contents of the given file and returns it
// through the given channel.
func retrieveFile(fileName string, contents chan<- string) {
	println("retrieving shader", fileName)
	var xmlHttp *js.Object = js.Global.Get("XMLHttpRequest").New()
	xmlHttp.Call("open", "GET", fileName, true)
	xmlHttp.Set("onload", func() {
		go func() {
			if xmlHttp.Get("readyState").Int() == 4 && xmlHttp.Get("status").Int() == 200 {
				contents <- xmlHttp.Get("responseText").String()
			} else {
				fail("Failed to load shader " + fileName)
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
