package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type SpriteAnimationMaterial struct {
	texture *gfx.Texture
	prefs   gfx.UniformLoader
}

func CreateSpriteAnimationMaterial(texture *gfx.Texture) *SpriteAnimationMaterial {
	material := &SpriteAnimationMaterial{
		texture: texture,
		prefs:   gfx.NewUniformLoader(),
	}
	material.prefs.AddVector4("fragtexture0Bounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	material.prefs.AddVector4("fragtintColor", gmath.NewVector4(0.0, 0.0, 0.0, 1.0))
	material.prefs.AddFloat("fragtintAmount", 0.0)
	// material.prefs.AddVector4("fragtintColor", gmath.NewVector4(0.0, 0.25, 0.75, 1.0))
	// material.prefs.AddFloat("fragtintAmount", 0.5)
	return material
}

func (spriteAnimationMaterial *SpriteAnimationMaterial) Texture() *gfx.Texture {
	return spriteAnimationMaterial.texture
}
func (spriteAnimationMaterial *SpriteAnimationMaterial) Prefs() gfx.UniformLoader {
	return spriteAnimationMaterial.prefs
}
func (spriteAnimationMaterial *SpriteAnimationMaterial) Transparency() bool { return true }
