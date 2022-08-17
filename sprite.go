package rtech

import (
	"github.com/gabereiser/rtech/gl"
	mgl "github.com/go-gl/mathgl/mgl64"
)

type RSprite struct {
	position mgl.Vec2
	texture  *gl.Texture2D
}
