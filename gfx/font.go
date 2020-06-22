package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
)

type TextComponent struct {
	Font      gio.Font
	Test      string
	LineWidth float32
}

func NewTextSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := components[0].(*RenderComponent)
			render.Instance.SetTransform(transformMat)
		}

		Sweep()
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}
