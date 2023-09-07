package gl

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type GLBindable interface {
	Bind()
	Unbind()
}

type GLBuffer interface {
	GetGLBuffer() uint32
}

type GLFramebuffer interface {
	GLBindable
	GLBuffer

	Begin()
	End()
}

type GLTexture interface {
	GLBindable
	GLBuffer
}

type GLTexture2D interface {
	GLTexture
}

type GLTexture3D interface {
	GLTexture
}

type GLTextureCube interface {
	GLTexture
}

type GLVertexBuffer interface {
	GLBindable
	GLBuffer
}

type GLIndexBuffer interface {
	GLBindable
	GLBuffer
}

type GLContext interface {
	Clear(bit int)
}

func Error() error {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		e := gl.GoStr(gl.GetString(err))
		fmt.Printf("gl error: %s\n", e)
		return fmt.Errorf("gl error: %s", e)
	}
	return nil
}

const (
	DEPTH_BUFFER_BIT   = gl.DEPTH_BUFFER_BIT
	COLOR_BUFFER_BIT   = gl.COLOR_BUFFER_BIT
	STENCIL_BUFFER_BIT = gl.STENCIL_BUFFER_BIT
)

func Clear(r, g, b, a float32, bit int) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(uint32(bit))
}
func ClearAll(r, g, b, a float32) {
	Clear(r, g, b, a, DEPTH_BUFFER_BIT|COLOR_BUFFER_BIT|STENCIL_BUFFER_BIT)
}

type DRAW_MODE = uint32
type BLEND_MODE = uint32
type BLEND_FUNC = uint32

const (
	DRAW_MODE_POINTS         DRAW_MODE = gl.POINTS
	DRAW_MODE_LINES          DRAW_MODE = gl.LINES
	DRAW_MODE_LINE_STRIP     DRAW_MODE = gl.LINE_STRIP
	DRAW_MODE_LINE_LOOP      DRAW_MODE = gl.LINE_LOOP
	DRAW_MODE_TRIANGLES      DRAW_MODE = gl.TRIANGLES
	DRAW_MODE_TRIANGLE_STRIP DRAW_MODE = gl.TRIANGLE_STRIP
	DRAW_MODE_TRIANGLE_FAN   DRAW_MODE = gl.TRIANGLE_FAN
)
const (
	BLEND_ZERO              BLEND_MODE = gl.ZERO
	BLEND_ONE               BLEND_MODE = gl.ONE
	BLEND_SRC_COLOR         BLEND_MODE = gl.SRC_COLOR
	BLEND_INVERSE_SRC_COLOR BLEND_MODE = gl.ONE_MINUS_SRC_COLOR
	BLEND_SRC_ALPHA         BLEND_MODE = gl.SRC_ALPHA
	BLEND_INVERSE_SRC_ALPHA BLEND_MODE = gl.ONE_MINUS_SRC_ALPHA
	BLEND_DST_COLOR         BLEND_MODE = gl.DST_COLOR
	BLEND_INVERSE_DST_COLOR BLEND_MODE = gl.ONE_MINUS_DST_COLOR
	BLEND_DST_ALPHA         BLEND_MODE = gl.DST_ALPHA
	BLEND_INVERSE_DST_ALPHA BLEND_MODE = gl.ONE_MINUS_DST_ALPHA
)
const (
	BLEND_FUNC_ADD     BLEND_FUNC = gl.FUNC_ADD
	BLEND_FUNC_SUB     BLEND_FUNC = gl.FUNC_SUBTRACT
	BLEND_FUNC_REV_SUB BLEND_FUNC = gl.FUNC_REVERSE_SUBTRACT
	BLEND_FUNC_MIN     BLEND_FUNC = gl.MIN
	BLEND_FUNC_MAX     BLEND_FUNC = gl.MAX
)

func Blend(src BLEND_MODE, dst BLEND_MODE) {
	gl.BlendFunc(src, dst)
}
func BlendColor(color [4]float32) {
	gl.BlendColor(color[0], color[1], color[2], color[3])
}
func BlendFunc(blend BLEND_FUNC) {
	gl.BlendEquation(blend)
}
func BlendAdd() {
	gl.BlendFunc(BLEND_SRC_ALPHA, BLEND_ONE)
}
func BlendAlpha() {
	gl.BlendFunc(BLEND_ONE, BLEND_INVERSE_SRC_ALPHA)
}
func BlendNonPreMultiplied() {
	gl.BlendFunc(BLEND_SRC_ALPHA, BLEND_INVERSE_SRC_ALPHA)
}
func BlendNone() {
	gl.BlendFunc(BLEND_ONE, BLEND_ZERO)
}

func BindVertexData() {

}
func DrawElements(mode DRAW_MODE, count int32, indices []uint32) {
	ptr := gl.Ptr(indices)
	gl.DrawElements(uint32(mode), count, gl.UNSIGNED_INT, ptr)
}
func DrawArrays(mode DRAW_MODE, first, count int) {
	gl.DrawArrays(uint32(mode), int32(first), int32(count))
}
