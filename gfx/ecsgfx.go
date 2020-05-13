package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

func NewRenderSystem() *limitengine.ECSSystem {
	// TODO: Create render listener that only adds/removes entities from the batch when altered.
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
			render.Instance.SetTransform(transformMat)
		}
		Sweep()
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}
