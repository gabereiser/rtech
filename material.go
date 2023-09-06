package rtech

import "github.com/gabereiser/rtech/gl"

type Material interface {
	gl.GLBindable
}
type RTEXTURE_SLOT int

const (
	RTEXTURE_SLOT_ALBEIDO   RTEXTURE_SLOT = 0
	RTEXTURE_SLOT_NORMAL    RTEXTURE_SLOT = 1
	RTEXTURE_SLOT_HEIGHT    RTEXTURE_SLOT = 2
	RTEXTURE_SLOT_ROUGHNESS RTEXTURE_SLOT = 3
	RTEXTURE_SLOT_OCCLUSION RTEXTURE_SLOT = 4
)

type RMaterial struct {
	Color      RColor
	Metallic   float64
	Roughtness float64
	Refraction float64
	Textures   map[RTEXTURE_SLOT]*RTexture2D
	Shader     RShader
}

type RShader struct {
	shader gl.Shader
}
