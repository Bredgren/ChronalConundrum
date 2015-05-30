package main

import (
// "github.com/gopherjs/gopherjs/js"
)

// loadModelAsset takes the given modelAsset, loads the json file it specifies,
// turns it into a usable model object and returns the file name via the done channel.
func loadModelAsset(asset *modelAsset, done chan<- string) {
	file := make(chan string)
	defer close(file)

	retrieveFile(asset.jsonFile, file)
	jsonFile := <-file
	_ = jsonFile

	// var texture *js.Object = gl.CreateTexture()
	// // TODO: what if I want some textures configured differently?
	// gl.BindTexture(gl.TEXTURE_2D, texture)
	// gl.PixelStorei(gl.UNPACK_FLIP_Y_WEBGL, 1)
	// gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, image)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	// gl.GenerateMipmap(gl.TEXTURE_2D)
	// gl.BindTexture(gl.TEXTURE_2D, nil)

	// *asset.texture = texture
	done <- asset.jsonFile
}
