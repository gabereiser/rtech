package rtech

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type RTEXTURE_FILTER int
type RTEXTURE_WRAP int

const (
	RTEXTURE_FILTER_BILINEAR  RTEXTURE_FILTER = gl.LINEAR
	RTEXTURE_FILTER_TRILINEAR RTEXTURE_FILTER = gl.LINEAR_MIPMAP_LINEAR
	RTEXTURE_FILTER_NONE      RTEXTURE_FILTER = gl.NEAREST

	RTEXTURE_WRAP_CLAMP_TO_EDGE        RTEXTURE_WRAP = gl.CLAMP_TO_EDGE
	RTEXTURE_WRAP_CLAMP_TO_BORDER      RTEXTURE_WRAP = gl.CLAMP_TO_BORDER
	RTEXTURE_WRAP_CLAMP_READ_COLOR     RTEXTURE_WRAP = gl.CLAMP_READ_COLOR
	RTEXTURE_WRAP_MIRRORED_REPEAT      RTEXTURE_WRAP = gl.MIRRORED_REPEAT
	RTEXTURE_WRAP_MIRROR_CLAMP_TO_EDGE RTEXTURE_WRAP = gl.MIRROR_CLAMP_TO_EDGE
	RTEXTURE_WRAP_REPEAT               RTEXTURE_WRAP = gl.REPEAT
)

type RTexture2D struct {
	image     *image.RGBA
	textureID RID
}

func NewTextureFromPath(file string) (*RTexture2D, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	defer imgFile.Close()
	return NewTextureFromFile(imgFile)
}
func NewTextureFromFile(file *os.File) (*RTexture2D, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	rt := &RTexture2D{
		image:     rgba,
		textureID: 0,
	}
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	rt.textureID = texture
	rt.Bind()
	borderColor := [4]float32{0.0, 0.0, 0.0, 1.0}
	gl.TexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, &borderColor[0])
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rt.image.Rect.Size().X),
		int32(rt.image.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rt.image.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	rt.Unbind()

	return rt, nil
}

func (tex *RTexture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tex.textureID)
}
func (tex *RTexture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (tex *RTexture2D) SetTextureMinFilter(filter RTEXTURE_FILTER) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, int32(filter))
	tex.Unbind()
}
func (tex *RTexture2D) SetTextureMagFilter(filter RTEXTURE_FILTER) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, int32(filter))
	tex.Unbind()
}
func (tex *RTexture2D) SetTextureWrapS(wrap RTEXTURE_WRAP) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, int32(wrap))
	tex.Unbind()
}
func (tex *RTexture2D) SetTextureWrapT(wrap RTEXTURE_WRAP) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, int32(wrap))
	tex.Unbind()
}
