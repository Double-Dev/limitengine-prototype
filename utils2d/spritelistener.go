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

func NewSpriteComponent(layer int32, camera *gfx.Camera, shader *SpriteShader, material gfx.Material, instance *gfx.Instance) *SpriteComponent {
	if instance.GetData("textureBounds") == nil {
		instance.SetData("textureBounds", gmath.NewVector4(0.0, 0.0, 1.0, 1.0))
	}
	renderable := &gfx.Renderable{
		Layer:    layer,
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
			if transform.IsAwake() {
				transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

				sprite := components[0].(*SpriteComponent)
				sprite.Renderable.Instance.SetTransform(transformMat)
				transform.SetAwake(false)
			}
		}
	}, (*SpriteComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewSpriteListener() *gfx.GFXListener { return gfx.NewGFXListener((*SpriteComponent)(nil)) }
