package gl

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	RTEXTURE_FILTER_LINEAR  = int32(gl.LINEAR)
	RTEXTURE_FILTER_NEAREST = int32(gl.NEAREST)

	RTEXTURE_WRAP_CLAMP_TO_EDGE        = int32(gl.CLAMP_TO_EDGE)
	RTEXTURE_WRAP_CLAMP_TO_BORDER      = int32(gl.CLAMP_TO_BORDER)
	RTEXTURE_WRAP_CLAMP_READ_COLOR     = int32(gl.CLAMP_READ_COLOR)
	RTEXTURE_WRAP_MIRRORED_REPEAT      = int32(gl.MIRRORED_REPEAT)
	RTEXTURE_WRAP_MIRROR_CLAMP_TO_EDGE = int32(gl.MIRROR_CLAMP_TO_EDGE)
	RTEXTURE_WRAP_REPEAT               = int32(gl.REPEAT)
)

type Texture2D struct {
	image     *image.RGBA
	textureID uint32
}

func NewTextureFromPath(file string) (*Texture2D, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	defer imgFile.Close()
	return NewTextureFromFile(imgFile)
}
func NewTextureFromFile(file *os.File) (*Texture2D, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	rt := &Texture2D{
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

func (tex *Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tex.textureID)
}
func (tex *Texture2D) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (tex *Texture2D) SetTextureMinFilter(filter int32) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, filter)
	tex.Unbind()
}
func (tex *Texture2D) SetTextureMagFilter(filter int32) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, filter)
	tex.Unbind()
}
func (tex *Texture2D) SetTextureWrapS(wrap int32) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, wrap)
	tex.Unbind()
}
func (tex *Texture2D) SetTextureWrapT(wrap int32) {
	gl.ActiveTexture(gl.TEXTURE0)
	tex.Bind()
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, wrap)
	tex.Unbind()
}
