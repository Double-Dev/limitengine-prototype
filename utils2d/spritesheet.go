package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type SpriteSheet struct {
	texture                   *gfx.Texture
	prefs                     gfx.UniformLoader
	spriteWidth, spriteHeight float32
	padding                   float32
	columns, rows             uint32
}

func CreateSpriteSheet(texture *gfx.Texture, spriteWidth, spriteHeight, padding float32) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		texture:      texture,
		prefs:        gfx.NewUniformLoader(),
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		padding:      padding,
		columns:      uint32(1.0 / spriteWidth),
		rows:         uint32(1.0 / spriteHeight),
	}
	spriteSheet.prefs.AddVector3("fragtintColor", gmath.NewVector3(0.0, 0.0, 1.0))
	spriteSheet.prefs.AddFloat("fragtintAmount", 0.5)
	return spriteSheet
}

func (spriteSheet *SpriteSheet) ApplyToInstance(instance *gfx.Instance, index uint32) {
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

func (spriteSheet *SpriteSheet) Texture() *gfx.Texture    { return spriteSheet.texture }
func (spriteSheet *SpriteSheet) Prefs() gfx.UniformLoader { return spriteSheet.prefs }
func (spriteSheet *SpriteSheet) Transparency() bool       { return true }
