package math

import mgl "github.com/go-gl/mathgl/mgl64"

type BoundingBox struct {
	Origin mgl.Vec3
	Min    mgl.Vec3
	Max    mgl.Vec3
}

type BoundingBoxAxisAligned struct {
	Origin      mgl.Vec3
	Min         mgl.Vec3
	Max         mgl.Vec3
	Orientation mgl.Quat
}
