package gl

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
