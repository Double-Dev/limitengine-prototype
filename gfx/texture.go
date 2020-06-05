package gfx

import "github.com/double-dev/limitengine/gfx/framework"

var (
	textureIndex = uint32(1)
	textures     = make(map[uint32]framework.ITexture)
)

func init() { textures[0] = nil }

type Texture struct {
	id uint32
}

func CreateEmptyTexture() *Texture {
	texture := &Texture{
		id: textureIndex,
	}
	textureIndex++
	actionQueue = append(actionQueue, func() {
		textures[texture.id] = context.CreateEmptyTexture()
	})
	return texture
}

func CreateTexture(image []uint8, width, height int32) *Texture {
	texture := &Texture{
		id: textureIndex,
	}
	textureIndex++
	actionQueue = append(actionQueue, func() {
		textures[texture.id] = context.CreateTexture(image, width, height)
	})
	return texture
}

// func (texture *Texture) SetPointFilter() framework.IAttachment {
// 	return textures[texture.id]
// }

// Attachment function:
func (texture *Texture) getFrameworkAttachment() framework.IAttachment {
	return textures[texture.id]
}
