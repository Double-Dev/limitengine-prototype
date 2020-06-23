package gfx

import (
	"fmt"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
)

type Font struct {
	textureMaterials []*TextureMaterial
	font             *gio.Font
}

func NewFont(font *gio.Font, color gmath.Vector3) *Font {
	var textures []*TextureMaterial
	for _, image := range font.Pages() {
		material := NewTextureMaterial(NewTexture(image))
		material.SetTint(color, 1.0)
		textures = append(textures, material)
	}
	return &Font{textures, font}
}

type TextComponent struct {
	camera             *Camera
	shader             *Shader
	font               *Font
	text               []string
	relativeTransforms []gmath.Matrix4
	renderables        []*Renderable
}

func NewTextComponent(camera *Camera, shader *Shader, font *Font, text []string) *TextComponent {
	fontSize := float32(1.0)
	var renderables []*Renderable
	var relativeTransforms []gmath.Matrix4
	xOffset := float32(0.0)
	for _, character := range text {
		char := font.font.GetChar(character)
		charAdvance := char.Advance()
		fmt.Println(character+":", character[0])

		if character != " " {
			charBounds := char.Bounds()
			charSize := gmath.NewVector2(charBounds[2], charBounds[3]).MulSc(limitengine.AspectRatio())
			charOffset := char.Offset().Mul(limitengine.AspectRatio(), 1.0)
			instance := NewInstance()
			instance.SetTextureBoundsV(charBounds)

			transform := gmath.NewTransformMatrix(
				gmath.NewVector3(xOffset+((charOffset[0]+charSize[0]*0.5)*fontSize), (-charOffset[1]-charSize[1]*0.5), 0.0),
				gmath.NewIdentityQuaternion(),
				gmath.NewVector3(charSize[0]*fontSize, charSize[1]*fontSize, 1.0),
			)

			instance.SetTransform(transform)
			xOffset += (charAdvance) * fontSize

			relativeTransforms = append(relativeTransforms, transform)
			renderables = append(renderables, &Renderable{
				Camera:   camera,
				Shader:   shader,
				Material: font.textureMaterials[char.Page()],
				Mesh:     SpriteMesh(),
				Instance: instance,
			})
		} else {
			xOffset += (charAdvance * 2.0) * fontSize
		}
	}

	return &TextComponent{
		camera,
		shader,
		font,
		text,
		relativeTransforms,
		renderables,
	}
}

func (textComponent *TextComponent) Renderables() []*Renderable { return textComponent.renderables }

func NewTextSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			text := components[0].(*TextComponent)
			for i, renderable := range text.Renderables() {
				renderable.Instance.SetTransform(
					transformMat.MulM(text.relativeTransforms[i]),
				)
			}
		}
	}, (*TextComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewTextListener() *GFXListener { return NewGFXListener((*TextComponent)(nil)) }
