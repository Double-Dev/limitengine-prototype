package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
)

type SpriteSheet struct {
	texture                   *gfx.Texture
	spriteWidth, spriteHeight float32
	columns, rows             uint32
	index                     uint32
}

func CreateSpriteSheet(texture *gfx.Texture, spriteWidth, spriteHeight float32) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		texture:      texture,
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		columns:      uint32(1.0 / spriteWidth),
		rows:         uint32(1.0 / spriteHeight),
		index:        0,
	}
	return spriteSheet
}

func (spriteSheet *SpriteSheet) Texture() *gfx.Texture {
	return spriteSheet.texture
}

func (spriteSheet *SpriteSheet) SpriteX() float32 {
	return spriteSheet.spriteWidth * float32(spriteSheet.index%spriteSheet.rows)
}

func (spriteSheet *SpriteSheet) SpriteY() float32 {
	return spriteSheet.spriteHeight * float32(spriteSheet.index/spriteSheet.rows)
}

func (spriteSheet *SpriteSheet) SpriteWidth() float32 {
	return spriteSheet.spriteWidth
}

func (spriteSheet *SpriteSheet) SpriteHeight() float32 {
	return spriteSheet.spriteHeight
}
