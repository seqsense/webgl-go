package main

import (
	"math"
	"syscall/js"

	webgl "github.com/seqsense/webgl-go"
)

const vsSource = `
attribute vec3 position;
attribute vec2 coord;

varying vec2 texCoord;

void main(void) {
  gl_Position = vec4(position, 1.0);
  texCoord = coord;
}
`
const fsSource = `
precision mediump float;
uniform sampler2D texture0;
varying vec2 texCoord;

void main(void) {
  gl_FragColor = texture2D(texture0, texCoord);
}
`

var vertices = []float32{
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0, 0.5, 0,
}

var texCoords = []float32{
	0.0, 0.0,
	1.0, 0.0,
	0.0, 1.0,
}

func run() {
	canvas := js.Global().Get("document").Call("getElementById", "glcanvas")

	gl, err := webgl.New(canvas)
	if err != nil {
		panic(err)
	}

	width := gl.Canvas.ClientWidth()
	height := gl.Canvas.ClientHeight()

	var vs, fs webgl.Shader
	if vs, err = initShader(gl, vsSource, gl.VERTEX_SHADER); err != nil {
		panic(err)
	}

	if fs, err = initShader(gl, fsSource, gl.FRAGMENT_SHADER); err != nil {
		panic(err)
	}

	program, err := linkShaders(gl, nil, vs, fs)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(vertices), gl.STATIC_DRAW)

	positionLoc := gl.GetAttribLocation(program, "position")
	gl.VertexAttribPointer(positionLoc, 3, gl.FLOAT, false, 0, 0)

	texCoordsBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, texCoordsBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(texCoords), gl.STATIC_DRAW)
	texCoordsLoc := gl.GetAttribLocation(program, "coord")
	gl.VertexAttribPointer(texCoordsLoc, 2, gl.FLOAT, false, 0, 0)

	texture := gl.CreateTexture()
	gl.BindTexture(gl.TEXTURE_2D, texture)
	textureULoc := gl.GetUniformLocation(program, "texture0")

	texWidth := 256
	texHeight := 256
	texData := make([]uint8, 4*texWidth*texHeight)
	// fill the texture with some colors
	firstColor := [3]float64{255, 0, 0}
	secondColor := [3]float64{0, 0, 255}
	for i := 0; i < texWidth; i += 1 {
		for j := 0; j < texHeight; j += 1 {
			index := (j*texWidth + i) * 4
			d := float64(i) / math.Max(float64(texWidth), float64(texHeight))
			for k := 0; k < 3; k++ {
				texData[index+k] = uint8(firstColor[k] + d*(secondColor[k]-firstColor[k]))
			}
			texData[index+3] = 255
		}
	}
	gl.TexImage2D2(gl.TEXTURE_2D, 0, gl.RGBA, texWidth, texHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, texData)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	gl.ClearColor(0.5, 0.5, 0.5, 0.9)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Viewport(0, 0, width, height)

	gl.EnableVertexAttribArray(positionLoc)
	gl.EnableVertexAttribArray(texCoordsLoc)
	gl.Uniform1i(textureULoc, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, len(vertices)/3)
}

func main() {
	go run()
	select {}
}
