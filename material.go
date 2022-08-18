package rtech

import "github.com/gabereiser/rtech/gl"

type Material interface {
	gl.GLBindable
}

type RMaterial struct {
	shader RShader
}

type RShader struct {
	shader gl.Shader
}
