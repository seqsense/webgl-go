package main

// Taken with minor modifications from https://github.com/seqsense/pcdeditor/blob/master/shader_js.go

import (
	"errors"
	"fmt"
	"syscall/js"

	webgl "github.com/seqsense/webgl-go"
)

func initShader(gl *webgl.WebGL, src string, sType webgl.ShaderType) (webgl.Shader, error) {
	s := gl.CreateShader(sType)
	gl.ShaderSource(s, src)
	gl.CompileShader(s)
	if !gl.GetShaderParameter(s, gl.COMPILE_STATUS).(bool) {
		compilationLog := gl.GetShaderInfoLog(s)
		var sTypeStr string
		if sType == gl.VERTEX_SHADER {
			sTypeStr = "VERTEX_SHADER"
		} else {
			sTypeStr = "FRAGMENT_SHADER"
		}
		return webgl.Shader(js.Null()), fmt.Errorf("compile failed (%s) %v", sTypeStr, compilationLog)
	}
	return s, nil
}

func linkShaders(gl *webgl.WebGL, fbVarings []string, shaders ...webgl.Shader) (webgl.Program, error) {
	program := gl.CreateProgram()
	for _, s := range shaders {
		gl.AttachShader(program, s)
	}
	if len(fbVarings) > 0 {
		gl.TransformFeedbackVaryings(program, fbVarings, gl.SEPARATE_ATTRIBS)
	}
	gl.LinkProgram(program)
	if !gl.GetProgramParameter(program, gl.LINK_STATUS).(bool) {
		return webgl.Program(js.Null()), errors.New("link failed: " + gl.GetProgramInfoLog(program))
	}
	return program, nil
}
