package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gio"
)

var (
	textureIndex = uint32(1)
	textures     = make(map[uint32]framework.ITexture)

	nilTexture = &Texture{}
)

func init() { textures[0] = nil }

type Texture struct {
	id uint32
}

func NewEmptyTexture() *Texture {
	texture := &Texture{
		id: textureIndex,
	}
	textureIndex++
	actionQueue = append(actionQueue, func() {
		textures[texture.id] = context.NewEmptyTexture()
	})
	return texture
}

func NewTexture(image *gio.Image) *Texture {
	texture := &Texture{
		id: textureIndex,
	}
	textureIndex++
	actionQueue = append(actionQueue, func() {
		textures[texture.id] = context.NewTexture(image.Data(), image.Width(), image.Height())
	})
	return texture
}

func (texture *Texture) SetPointFilter(mipmap, antisotrophic bool) {
	actionQueue = append(actionQueue, func() {
		textures[texture.id].Bind()
		textures[texture.id].NearestFilter(mipmap, antisotrophic)
	})
}

func (texture *Texture) SetLinearFilter(mipmap, antisotrophic bool) {
	actionQueue = append(actionQueue, func() {
		textures[texture.id].Bind()
		textures[texture.id].LinearFilter(mipmap, antisotrophic)
	})
}

// Attachment function:
func (texture *Texture) getFrameworkAttachment() framework.IAttachment {
	return textures[texture.id]
}
