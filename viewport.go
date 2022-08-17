package rtech

import "github.com/go-gl/gl/v4.1-core/gl"

type RViewport struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (v *RViewport) AspectRatio() float32 {
	return float32(v.Width) / float32(v.Height)
}

func (v *RViewport) Bind() {
	gl.Viewport(int32(v.X), int32(v.Y), int32(v.Width), int32(v.Height))
	gl.Scissor(int32(v.X), int32(v.Y), int32(v.Width), int32(v.Height))
}
func (v *RViewport) Unbind() {
	//no-op as the next viewport bind will set the viewport for the draw.
}
