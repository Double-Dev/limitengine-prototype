package gl

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// TODO: Clean up texture file.
const (
	textureType2D = gl.TEXTURE_2D
)

type texture struct {
	id, textureType uint32
}

func newTexture(textureType uint32) *texture {
	var id uint32
	gl.GenTextures(1, &id)
	texture := &texture{
		id:          id,
		textureType: textureType,
	}
	return texture
}

func (texture *texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.id)
}

func (texture *texture) NearestFilter(mipmap, antisotrophic bool) {
	if mipmap {
		gl.GenerateMipmap(texture.textureType)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)

		if antisotrophic {
			gl.TexParameterf(texture.textureType, gl.MAX_TEXTURE_LOD_BIAS, 0.0)
			var data float32
			gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
			amount := gmath.Min(4.0, data)
			gl.TexParameterf(texture.textureType, gl.TEXTURE_MAX_ANISOTROPY, amount)
		}
	} else {
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	}
}

func (texture *texture) LinearFilter(mipmap, antisotrophic bool) {
	if mipmap {
		gl.GenerateMipmap(texture.textureType)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)

		if antisotrophic {
			gl.TexParameterf(texture.textureType, gl.TEXTURE_LOD_BIAS, 0.0)
			var data float32
			gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
			amount := gmath.Min(4.0, data)
			gl.TexParameterf(texture.textureType, gl.TEXTURE_MAX_ANISOTROPY, amount)
		}
	} else {
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	}
}

func (texture *texture) TextureData(image []uint8, width, height int32) {
	if texture.textureType == textureType2D {
		gl.TexImage2D(texture.textureType, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image))
	}
}

func (texture *texture) Delete() {
	gl.DeleteTextures(1, &texture.id)
}

// Attachment functions:
func (texture *texture) AttachToFramebufferColor(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.RGBA, framebuffer.Width(), framebuffer.Height(), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, uint32(gl.COLOR_ATTACHMENT0), texture.textureType, texture.id, 0)
	}
}

func (texture *texture) ResizeFramebufferColor(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.RGBA, framebuffer.Width(), framebuffer.Height(), 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	}
}

func (texture *texture) AttachToFramebufferDepth(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.DEPTH_COMPONENT, framebuffer.Width(), framebuffer.Height(), 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, texture.textureType, texture.id, 0)
	}
}

func (texture *texture) ResizeFramebufferDepth(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.DEPTH_COMPONENT, framebuffer.Width(), framebuffer.Height(), 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
	}
}

func (texture *texture) ID() *uint32 {
	return &texture.id
}
