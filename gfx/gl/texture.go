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

func createTexture(textureType uint32) *texture {
	var id uint32
	gl.GenTextures(1, &id)
	texture := &texture{
		id:          id,
		textureType: textureType,
	}
	return texture
}

func (texture *texture) NearestFilter(mipmap, antisotrophic bool) {
	if mipmap {
		gl.GenerateMipmap(texture.textureType)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.NEAREST_MIPMAP_NEAREST)
		gl.TexParameterf(texture.textureType, gl.MAX_TEXTURE_LOD_BIAS, 0.0)

		if antisotrophic {
			var data float32
			gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
			amount := gmath.Min(4.0, data)
			gl.TexParameterf(texture.textureType, gl.TEXTURE_MAX_ANISOTROPY, amount)
		}
	} else {
		gl.TexParameterf(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameterf(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	}
}

func (texture *texture) LinearFilter(mipmap, antisotrophic bool) {
	if mipmap {
		gl.GenerateMipmap(texture.textureType)
		gl.TexParameteri(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
		gl.TexParameterf(texture.textureType, gl.MAX_TEXTURE_LOD_BIAS, 0.0)

		if antisotrophic {
			var data float32
			gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
			amount := gmath.Min(4.0, data)
			gl.TexParameterf(texture.textureType, gl.TEXTURE_MAX_ANISOTROPY, amount)
		}
	} else {
		gl.TexParameterf(texture.textureType, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
		gl.TexParameterf(texture.textureType, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	}
}

func (texture *texture) TextureData(image []uint8, width, height int32) {
	if texture.textureType == textureType2D {
		gl.TexImage2D(texture.textureType, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image))
	}
}

// func createTexture(image []uint8, width, height int32) *texture {
// 	var id uint32
// 	gl.GenTextures(1, &id)
// 	texture := &texture{
// 		id:      id,
// 		texType: textureType2D,
// 	}
// 	texture.Bind()

// 	gl.TexParameterf(texture.texType, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
// 	gl.TexParameterf(texture.texType, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

// 	gl.TexImage2D(texture.texType, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(image))

// 	// gl.GenerateMipmap(texture.texType)
// 	// gl.TexParameteri(texture.texType, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
// 	// gl.TexParameterf(texture.texType, gl.MAX_TEXTURE_LOD_BIAS, 0.0)

// 	// var data float32
// 	// gl.GetFloatv(gl.MAX_TEXTURE_MAX_ANISOTROPY, &data)
// 	// amount := gmath.Min(4.0, data)
// 	// gl.TexParameterf(texture.texType, gl.TEXTURE_MAX_ANISOTROPY, amount)

// 	texture.Unbind()
// 	return texture
// }

// func newTextureEmpty(texType uint32) *texture {
// 	var id uint32
// 	gl.GenTextures(1, &id)
// 	return &texture{
// 		id:      id,
// 		texType: textureType2D,
// 	}
// }

func (texture *texture) Bind() {
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture.id)
}

func (texture *texture) Unbind() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (texture *texture) Delete() {

}

// Attachment functions:
func (texture *texture) AttachToFramebufferColor(framebuffer framework.IFramebuffer, attachment int) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.RGB, framebuffer.Width(), framebuffer.Height(), 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, uint32(gl.COLOR_ATTACHMENT0+attachment), texture.textureType, texture.id, 0)
	}
}

func (texture *texture) AttachToFramebufferDepth(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.DEPTH_COMPONENT, framebuffer.Width(), framebuffer.Height(), 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, texture.textureType, texture.id, 0)
	}
}

func (texture *texture) AttachToFramebufferStencil(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.STENCIL_INDEX, framebuffer.Width(), framebuffer.Height(), 0, gl.STENCIL_INDEX, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.STENCIL_ATTACHMENT, texture.textureType, texture.id, 0)
	}
}

func (texture *texture) AttachToFramebufferDepthStencil(framebuffer framework.IFramebuffer) {
	if texture.textureType == textureType2D {
		texture.Bind()
		gl.TexImage2D(texture.textureType, 0, gl.DEPTH_STENCIL, framebuffer.Width(), framebuffer.Height(), 0, gl.DEPTH_STENCIL, gl.UNSIGNED_BYTE, nil)
		gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, texture.textureType, texture.id, 0)
	}
}
