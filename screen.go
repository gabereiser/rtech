package rtech

import mgl "github.com/go-gl/mathgl/mgl64"

type Screen interface {
	Project(x, y float64) mgl.Vec3
	Unproject(p mgl.Vec3) (x, y float64)

	Begin2D()
	End2D()

	RenderFullscreenQuad(t *RTexture2D, s *RShader)
}
