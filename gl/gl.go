package gl

import (
	"github.com/gabereiser/rtech/types"
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

/******** CLEAR ********/
const (
	DEPTH_BUFFER_BIT   = gl.DEPTH_BUFFER_BIT
	COLOR_BUFFER_BIT   = gl.COLOR_BUFFER_BIT
	STENCIL_BUFFER_BIT = gl.STENCIL_BUFFER_BIT
)

func Clear(color types.RColor, bit int) {
	gl.ClearColor(color.RedF(), color.GreenF(), color.BlueF(), color.AlphaF())
	gl.Clear(uint32(bit))
}
func ClearAll(clearColor types.RColor) {
	Clear(clearColor, DEPTH_BUFFER_BIT|COLOR_BUFFER_BIT|STENCIL_BUFFER_BIT)
}

/********* END CLEAR *********/
