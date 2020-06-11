package gfx

import (
	"github.com/double-dev/limitengine/gmath"
)

// TODO: Think about turning Material into interface to allow easy creation of variations.
type Material interface {
	Texture() *Texture
	Prefs() UniformLoader
	Transparency() bool
}

type ColorMaterial struct {
	prefs        UniformLoader
	transparency bool
}

func CreateColorMaterial(color gmath.Vector4) *ColorMaterial {
	colorMaterial := &ColorMaterial{
		prefs:        NewUniformLoader(),
		transparency: color[3] < 1.0,
	}
	colorMaterial.prefs.AddVector4("fragtexture0Bounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	colorMaterial.prefs.AddVector4("fragtintColor", color)
	colorMaterial.prefs.AddFloat("fragtintAmount", 1.0)
	return colorMaterial
}

func (colorMaterial *ColorMaterial) Texture() *Texture    { return nilTexture }
func (colorMaterial *ColorMaterial) Prefs() UniformLoader { return colorMaterial.prefs }
func (colorMaterial *ColorMaterial) Transparency() bool   { return colorMaterial.transparency }

type TextureMaterial struct {
	texture *Texture
	prefs   UniformLoader
}

func CreateTextureMaterial(texture *Texture) *TextureMaterial {
	textureMaterial := &TextureMaterial{
		texture: texture,
		prefs:   NewUniformLoader(),
	}
	textureMaterial.prefs.AddVector4("fragtexture0Bounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	textureMaterial.prefs.AddVector4("fragtintColor", gmath.NewZeroVector4())
	textureMaterial.prefs.AddFloat("fragtintAmount", 0.0)
	// textureMaterial.prefs.AddVector4("fragtintColor", gmath.NewVector4(0.0, 0.25, 0.75, 1.0))
	// textureMaterial.prefs.AddFloat("fragtintAmount", 0.5)
	return textureMaterial
}

func (textureMaterial *TextureMaterial) Texture() *Texture    { return textureMaterial.texture }
func (textureMaterial *TextureMaterial) Prefs() UniformLoader { return textureMaterial.prefs }
func (textureMaterial *TextureMaterial) Transparency() bool   { return true }
