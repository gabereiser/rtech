package rtech

import "github.com/gabereiser/rtech/gl"

type RID = uint32

type RBlendFunc struct {
	value gl.BLEND_MODE
}

func (RBlendFunc) Zero() RBlendFunc {
	return RBlendFunc{gl.BLEND_ZERO}
}
func (RBlendFunc) Add() RBlendFunc {
	return RBlendFunc{gl.BLEND_SRC_ALPHA}
}

type RBlendState struct {
	colorSourceBlend RBlendFunc
	alphaSourceBlend RBlendFunc
	colorDestBlend   RBlendFunc
	alphaDestBlend   RBlendFunc
}

func (this *RBlendState) Bind() {
	gl.Blend(this.colorSourceBlend.value, this.colorDestBlend.value)
}
