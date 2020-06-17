package gfx

import (
	"github.com/double-dev/limitengine/gmath"
)

type SpriteSheet struct {
	spriteWidth, spriteHeight float32
	padding                   float32
	columns, rows             uint32
}

func CreateSpriteSheet(spriteWidth, spriteHeight, padding float32) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		padding:      padding,
		columns:      uint32(1.0 / spriteWidth),
		rows:         uint32(1.0 / spriteHeight),
	}
	return spriteSheet
}

func (spriteSheet *SpriteSheet) Apply(instance *Instance, index uint32) {
	instance.SetTextureBounds(
		spriteSheet.spriteWidth*float32(index%spriteSheet.rows)+spriteSheet.padding,
		spriteSheet.spriteHeight*float32(index/spriteSheet.rows)+spriteSheet.padding,
		spriteSheet.spriteWidth-(2.0*spriteSheet.padding),
		spriteSheet.spriteHeight-(2.0*spriteSheet.padding),
	)
}

func (spriteSheet *SpriteSheet) GetBounds(index uint32) gmath.Vector4 {
	return gmath.NewVector4(
		spriteSheet.spriteWidth*float32(index%spriteSheet.rows)+spriteSheet.padding,
		spriteSheet.spriteHeight*float32(index/spriteSheet.rows)+spriteSheet.padding,
		spriteSheet.spriteWidth-(2.0*spriteSheet.padding),
		spriteSheet.spriteHeight-(2.0*spriteSheet.padding),
	)
}
