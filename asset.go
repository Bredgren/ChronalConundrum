package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// Shaders
var (
	testShader   *js.Object
	tmpShader    *js.Object
	shaderAssets = []shaderAsset{
		{"test.vert", "test.frag", &testShader},
		// {"tmp.vert", "tmp.frag", &tmpShader},
	}
)

type shaderAsset struct {
	vertFile string
	fragFile string
	shader   **js.Object
}

// Textures

// Sounds

// Models

// retrieveFile asyncronously gets the contents of the given file and returns it
// through the given channel. If an empty string was recieved then there was an error.
func retrieveFile(fileName string, contents chan<- string) {
	println("retrieving", fileName)
	var xmlHttp *js.Object = js.Global.Get("XMLHttpRequest").New()
	xmlHttp.Call("open", "GET", fileName, true)
	xmlHttp.Set("onload", func() {
		go func() {
			if xmlHttp.Get("readyState").Int() == 4 && xmlHttp.Get("status").Int() == 200 {
				contents <- xmlHttp.Get("responseText").String()
			} else {
				contents <- ""
			}
		}()
	})
	xmlHttp.Call("send")
}
