package gfx

import (
	"github.com/double-dev/limitengine/gmath"
)

// TODO: Think about turning Material into interface to allow easy creation of variations.
type Material struct {
	texture      *Texture
	prefs        uniformLoader
	Transparency bool
}

func CreateMaterial(color gmath.Vector4) *Material {
	material := &Material{
		texture:      &Texture{},
		prefs:        newUniformLoader(),
		Transparency: color[3] < 1.0,
	}
	material.prefs.AddVector4("spriteBounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	material.prefs.AddVector4("tintColor", color)
	material.prefs.AddFloat("tintAmount", 1.0)
	return material
}

func CreateTextureMaterial(texture *Texture) *Material {
	material := &Material{
		texture:      texture,
		prefs:        newUniformLoader(),
		Transparency: true,
	}
	material.prefs.AddVector4("spriteBounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	material.prefs.AddVector4("tintColor", gmath.NewVector4(0.0, 0.25, 0.75, 1.0))
	material.prefs.AddFloat("tintAmount", 0.5)
	return material
}
