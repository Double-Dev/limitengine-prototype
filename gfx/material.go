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

func NewColorMaterial(color gmath.Vector3) *ColorMaterial {
	colorMaterial := &ColorMaterial{
		prefs: NewUniformLoader(),
	}
	colorMaterial.prefs.AddVector3("fragtintColor", color)
	colorMaterial.prefs.AddFloat("fragtintAmount", 1.0)
	return colorMaterial
}

func (colorMaterial *ColorMaterial) SetColor(color gmath.Vector3) {
	colorMaterial.prefs.AddVector3("fragtintColor", color)
}

func (colorMaterial *ColorMaterial) Texture() *Texture    { return nilTexture }
func (colorMaterial *ColorMaterial) Prefs() UniformLoader { return colorMaterial.prefs }
func (colorMaterial *ColorMaterial) Transparency() bool   { return false }

type TextureMaterial struct {
	texture *Texture
	prefs   UniformLoader
}

func NewTextureMaterial(texture *Texture) *TextureMaterial {
	textureMaterial := &TextureMaterial{
		texture: texture,
		prefs:   NewUniformLoader(),
	}
	textureMaterial.prefs.AddVector3("fragtintColor", gmath.NewZeroVector3())
	textureMaterial.prefs.AddFloat("fragtintAmount", 0.0)
	// textureMaterial.prefs.AddVector4("fragtintColor", gmath.NewVector4(0.0, 0.25, 0.75, 1.0))
	// textureMaterial.prefs.AddFloat("fragtintAmount", 0.5)
	return textureMaterial
}

func (textureMaterial *TextureMaterial) SetTint(color gmath.Vector3, amount float32) {
	textureMaterial.prefs.AddVector3("fragtintColor", color)
	textureMaterial.prefs.AddFloat("fragtintAmount", amount)
}

func (textureMaterial *TextureMaterial) Texture() *Texture    { return textureMaterial.texture }
func (textureMaterial *TextureMaterial) Prefs() UniformLoader { return textureMaterial.prefs }
func (textureMaterial *TextureMaterial) Transparency() bool   { return true }
