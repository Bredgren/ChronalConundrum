package main

import (
	"encoding/json"
	"github.com/gopherjs/gopherjs/js"
)

type jsonModel struct {
	Name     string
	Vertices []float64
	Indices  []float64
}

// loadModelAsset takes the given modelAsset, loads the json file it specifies,
// turns it into a usable model object and returns the file name via the done channel.
func loadModelAsset(asset *modelAsset, done chan<- string) {
	file := make(chan string)
	defer close(file)

	retrieveFile(asset.jsonFile, file)
	fileContents := <-file
	fileBytes := []byte(fileContents)

	var contents jsonModel
	err := json.Unmarshal(fileBytes, &contents)
	if err != nil {
		fail("Failed to decode model's json data: " + err.Error())
	}

	vertices := make([]float32, len(contents.Vertices))
	for i, vert := range contents.Vertices {
		vertices[i] = float32(vert)
	}

	indices := make([]uint16, len(contents.Indices))
	for i, index := range contents.Indices {
		indices[i] = uint16(index)
	}

	var vBuf *js.Object = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vBuf)
	gl.BufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW)

	var fBuf *js.Object = gl.CreateBuffer()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, fBuf)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW)

	*asset.verts = vBuf
	*asset.faces = fBuf
	*asset.len = len(contents.Indices)

	// *asset.texture = texture
	done <- asset.jsonFile
}
