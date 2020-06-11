package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type SpriteSheet struct {
	texture                   *gfx.Texture
	prefs                     gfx.UniformLoader
	spriteWidth, spriteHeight float32
	columns, rows             uint32
	index                     uint32
}

func CreateSpriteSheet(texture *gfx.Texture, spriteWidth, spriteHeight float32) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		texture:      texture,
		prefs:        gfx.NewUniformLoader(),
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		columns:      uint32(1.0 / spriteWidth),
		rows:         uint32(1.0 / spriteHeight),
		index:        0,
	}
	spriteSheet.prefs.AddVector4("fragtexture0Bounds", gmath.NewVector4(spriteSheet.SpriteX(), spriteSheet.SpriteY(), spriteSheet.SpriteWidth(), spriteSheet.SpriteHeight()))
	spriteSheet.prefs.AddVector4("fragtintColor", gmath.NewZeroVector4())
	spriteSheet.prefs.AddFloat("fragtintAmount", 0.0)
	return spriteSheet
}

func (spriteSheet *SpriteSheet) SetIndex(index uint32) {
	spriteSheet.index = index
	spriteSheet.prefs.AddVector4("fragtexture0Bounds", gmath.NewVector4(spriteSheet.SpriteX(), spriteSheet.SpriteY(), spriteSheet.SpriteWidth(), spriteSheet.SpriteHeight()))
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

func (spriteSheet *SpriteSheet) Texture() *gfx.Texture    { return spriteSheet.texture }
func (spriteSheet *SpriteSheet) Prefs() gfx.UniformLoader { return spriteSheet.prefs }
func (spriteSheet *SpriteSheet) Transparency() bool       { return true }
