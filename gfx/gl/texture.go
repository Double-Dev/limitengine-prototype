package gl

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	textureType2D = gl.TEXTURE_2D
)

type texture struct {
	id, texType uint32
}

func newTexture(image []uint8, width, height int32) *texture {
	var id uint32
	gl.GenTextures(1, &id)
	texture := &texture{
		id:      id,
		texType: textureType2D,
	}
	texture.Bind()

	gl.TexParameterf(texture.texType, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameterf(texture.texType, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.TexImage2D(texture.texType, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image))

	// gl.GenerateMipmap(texture.texType)
	// gl.TexParameteri(texture.texType, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	// gl.TexParameterf(texture.texType, gl.MAX_TEXTURE_LOD_BIAS, 0.0)

	// var data float32
	// gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
	// amount := gmath.Min(4.0, data)
	// gl.TexParameterf(texture.texType, gl.TEXTURE_MAX_ANISOTROPY, amount)

	texture.Unbind()
	return texture
}

func newTextureEmpty(texType uint32) *texture {
	var id uint32
	gl.GenTextures(1, &id)
	return &texture{
		id:      id,
		texType: textureType2D,
	}
}

func (texture *texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.id)
}

func (texture *texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (texture *texture) Delete() {

}
