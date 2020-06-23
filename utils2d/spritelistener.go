package utils2d

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type SpriteComponent struct {
	Renderable  *gfx.Renderable
	renderables []*gfx.Renderable
}

func NewSpriteComponent(camera *gfx.Camera, shader *gfx.Shader, material gfx.Material, instance *gfx.Instance) *SpriteComponent {
	renderable := &gfx.Renderable{
		Camera:   camera,
		Shader:   shader,
		Material: material,
		Mesh:     gfx.SpriteMesh(),
		Instance: instance,
	}
	return &SpriteComponent{
		renderable,
		[]*gfx.Renderable{
			renderable,
		},
	}
}

func (spriteComponent *SpriteComponent) Renderables() []*gfx.Renderable {
	return spriteComponent.renderables
}

func NewSpriteSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			sprite := components[0].(*SpriteComponent)
			sprite.Renderable.Instance.SetTransform(transformMat)
		}
	}, (*SpriteComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewSpriteListener() *gfx.GFXListener { return gfx.NewGFXListener((*SpriteComponent)(nil)) }
