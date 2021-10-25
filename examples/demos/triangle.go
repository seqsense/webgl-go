package main

import (
	"syscall/js"

	webgl "github.com/seqsense/webgl-go"
)

const vsSource = `
attribute vec3 position;
attribute vec3 color;
varying vec3 vColor;

void main(void) {
  gl_Position = vec4(position, 1.0);
  vColor = color;
}
`

const fsSource = `
precision mediump float;
varying vec3 vColor;
void main(void) {
  gl_FragColor = vec4(vColor, 1.);
}
`

var vertices = []float32{
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
	0, 0.5, 0,
}

var colors = []float32{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

func run() {
	canvas := js.Global().Get("document").Call("getElementById", "glcanvas")

	gl, err := webgl.New(canvas)
	if err != nil {
		panic(err)
	}

	width := gl.Canvas.ClientWidth()
	height := gl.Canvas.ClientHeight()

	vertexBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(vertices), gl.STATIC_DRAW)

	colorBuffer := gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, webgl.Float32ArrayBuffer(colors), gl.STATIC_DRAW)

	var vs, fs webgl.Shader
	if vs, err = initVertexShader(gl, vsSource); err != nil {
		panic(err)
	}

	if fs, err = initFragmentShader(gl, fsSource); err != nil {
		panic(err)
	}

	program, err := linkShaders(gl, nil, vs, fs)
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	position := gl.GetAttribLocation(program, "position")
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(position)

	gl.BindBuffer(gl.ARRAY_BUFFER, colorBuffer)
	color := gl.GetAttribLocation(program, "color")
	gl.VertexAttribPointer(color, 3, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(color)

	gl.ClearColor(0.5, 0.5, 0.5, 0.9)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Enable(gl.DEPTH_TEST)
	gl.Viewport(0, 0, width, height)
	gl.DrawArrays(gl.TRIANGLES, 0, len(vertices)/3)
}

func main() {
	go run()
	select {}
}
