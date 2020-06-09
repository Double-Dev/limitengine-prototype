package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

func NewRenderSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.Component) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := components[0].(*RenderComponent)
			render.Instance.SetTransform(transformMat)
		}

		Sweep()
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}
