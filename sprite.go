package rtech

import (
	mgl "github.com/go-gl/mathgl/mgl64"
)

type RSprite struct {
	position mgl.Vec2
	texture  *RTexture2D
}
