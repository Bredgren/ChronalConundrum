package main

import (
	"github.com/gopherjs/gopherjs/js"
)

// loadTextureAsset takes the given textureAsset, loads the image it specifies and
// returns the file name via the done channel.
func loadTextureAsset(asset *textureAsset, done chan<- string) {
	img := make(chan *js.Object)
	defer close(img)

	retrieveImage(asset.textureFile, img)
	image := <-img

	var texture *js.Object = gl.CreateTexture()
	// TODO: what if I want some textures configured differently?
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, gl.RGBA, gl.UNSIGNED_BYTE, image)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_NEAREST)
	gl.GenerateMipmap(gl.TEXTURE_2D)
	gl.BindTexture(gl.TEXTURE_2D, nil)

	*asset.texture = texture
	done <- asset.textureFile
}
