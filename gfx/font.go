package gfx

import (
	"fmt"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
)

var (
	textPlugin = CreateLESLPlugin(`
frag{
	vars{
		uniform vec3 color;
		uniform float width;
		uniform float edge;
		uniform vec3 borderColor;
		uniform float borderWidth;
		uniform float borderEdge;
	},
	main{
		float distance = 1.0 - lesl.outColor.a;
		float outlineAlpha = 1.0 - smoothstep(borderWidth, borderWidth+borderEdge, distance);
		if (outlineAlpha == 0.0) {
			discard;
		}
		float alpha = 1.0 - smoothstep(width, width+edge, distance);
		vec3 color = mix(borderColor, color, alpha);
		lesl.outColor = vec4(color, outlineAlpha);
	},
},`)
)

type TextShader struct {
	renderProgram *RenderProgram
	uniformLoader UniformLoader
}

func NewTextShader(leslPlugins ...*LESLPlugin) *TextShader {
	return &TextShader{
		NewRenderProgram(append([]*LESLPlugin{textureBoundsPlugin, textPlugin}, leslPlugins...)...),
		NewUniformLoader(),
	}
}

func (shader *TextShader) RenderProgram() *RenderProgram { return shader.renderProgram }
func (shader *TextShader) UniformLoader() UniformLoader  { return shader.uniformLoader }

type TextMaterial struct {
	texture *Texture
	prefs   UniformLoader
}

func NewTextMaterial(texture *Texture, color gmath.Vector3, width, edge float32, borderColor gmath.Vector3, borderWidth, borderEdge float32) *TextMaterial {
	textMaterial := &TextMaterial{
		texture: texture,
		prefs:   NewUniformLoader(),
	}
	textMaterial.prefs.AddVector3("color", color)
	textMaterial.prefs.AddFloat("width", width)
	textMaterial.prefs.AddFloat("edge", edge)
	textMaterial.prefs.AddVector3("borderColor", borderColor)
	textMaterial.prefs.AddFloat("borderWidth", borderWidth)
	textMaterial.prefs.AddFloat("borderEdge", borderEdge)
	return textMaterial
}

func (textMaterial *TextMaterial) Texture() *Texture    { return textMaterial.texture }
func (textMaterial *TextMaterial) Prefs() UniformLoader { return textMaterial.prefs }
func (textMaterial *TextMaterial) Transparency() bool   { return true }

type Font struct {
	textMaterials []*TextMaterial
	font          *gio.Font
}

func NewFont(font *gio.Font, color gmath.Vector3, width, edge float32, borderColor gmath.Vector3, borderWidth, borderEdge float32) *Font {
	var textures []*TextMaterial
	for _, image := range font.Pages() {
		material := NewTextMaterial(NewTexture(image), color, width, edge, borderColor, borderWidth, borderEdge)
		textures = append(textures, material)
	}
	return &Font{textures, font}
}

type TextComponent struct {
	camera             *Camera
	shader             *TextShader
	font               *Font
	text               string
	lineWidth          float32
	relativeTransforms []gmath.Matrix4
	renderables        []*Renderable
}

func NewTextComponent(layer int32, camera *Camera, shader *TextShader, font *Font, text string, lineWidth float32) *TextComponent {
	fontSize := float32(1.0)
	var renderables []*Renderable
	var relativeTransforms []gmath.Matrix4

	var xTotal, yTotal float32
	for _, character := range text {
		char := font.font.GetChar(character)
		if char == nil {
			fmt.Println(string(character)+":", character)
			continue
		}
		charAdvance := char.Advance()
		if character != 32 {
			charBounds := char.Bounds()
			charSize := gmath.NewVector2(charBounds[2], charBounds[3]).MulSc(limitengine.AspectRatio())
			charOffset := char.Offset().Mul(limitengine.AspectRatio(), 1.0)
			xTotal += (charAdvance) * fontSize
			yTotal = gmath.Max(yTotal, (charOffset[1] + charSize[1]))
		} else {
			xTotal += (charAdvance * 2.0) * fontSize
		}
	}
	var xOffset float32
	for _, character := range text {
		char := font.font.GetChar(character)
		if char == nil {
			continue
		}
		charAdvance := char.Advance()
		// fmt.Println(string(character)+":", character)

		if character != 32 {
			charBounds := char.Bounds()
			charSize := gmath.NewVector2(charBounds[2], charBounds[3]).MulSc(limitengine.AspectRatio())
			charOffset := char.Offset().Mul(limitengine.AspectRatio(), 1.0)
			instance := NewInstance()
			instance.SetData("textureBounds", charBounds)

			transform := gmath.NewTransformMatrix(
				gmath.NewVector3(xOffset+((charOffset[0]+charSize[0]*0.5)*fontSize)-xTotal/2.0, (-charOffset[1]-charSize[1]*0.5)+yTotal/2.0, 0.0),
				gmath.NewIdentityQuaternion(),
				gmath.NewVector3(charSize[0]*fontSize, charSize[1]*fontSize, 1.0),
			)

			instance.SetTransform(transform)
			xOffset += (charAdvance) * fontSize

			relativeTransforms = append(relativeTransforms, transform)
			renderables = append(renderables, &Renderable{
				Layer:    layer,
				Camera:   camera,
				Shader:   shader,
				Material: font.textMaterials[char.Page()],
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
		lineWidth,
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
