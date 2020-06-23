package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
)

type FontMaterial struct {
	texture *Texture
	prefs   UniformLoader
}

func NewFontMaterial(font *gio.Font) *FontMaterial {
	fontMaterial := &FontMaterial{
		// texture: New,
		prefs: NewUniformLoader(),
	}
	fontMaterial.prefs.AddVector3("fragtintColor", gmath.NewZeroVector3())
	fontMaterial.prefs.AddFloat("fragtintAmount", 0.0)
	return fontMaterial
}

func (fontMaterial *FontMaterial) SetTint(color gmath.Vector3, amount float32) {
	fontMaterial.prefs.AddVector3("fragtintColor", color)
	fontMaterial.prefs.AddFloat("fragtintAmount", amount)
}

func (fontMaterial *FontMaterial) Texture() *Texture    { return fontMaterial.texture }
func (fontMaterial *FontMaterial) Prefs() UniformLoader { return fontMaterial.prefs }
func (fontMaterial *FontMaterial) Transparency() bool   { return true }

type TextComponent struct {
	Text        string
	camera      *Camera
	shader      *Shader
	material    Material
	renderables []*Renderable
}

func NewTextComponent(camera *Camera, shader *Shader, material Material) *TextComponent {
	return &TextComponent{
		"",
		camera,
		shader,
		material,
		[]*Renderable{},
	}
}

func (textComponent *TextComponent) Renderables() []*Renderable { return textComponent.renderables }

func NewTextSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		// for _, components := range entities {
		// transform := components[1].(*gmath.TransformComponent)

		// transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

		// text := components[0].(*TextComponent)
		// text.Instance().SetTransform(transformMat)
		// }
	}, (*TextComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewTextListener() *GFXListener { return NewGFXListener((*TextComponent)(nil)) }
