package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
)

type AnimationMaterial struct {
	material *gfx.Material
}

func CreateAnimationMaterial(spriteSheet *gfx.Texture) *AnimationMaterial {
	return nil
}

func (animationMaterial *AnimationMaterial) Material() *gfx.Material {
	return animationMaterial.material
}
