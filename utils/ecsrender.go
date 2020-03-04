package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gfx"
)

type RenderComponent struct {
	shader  gfx.Shader
	model   gfx.Model
	texture gfx.Texture
}

func NewRenderSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		// transform := entity.GetComponent((*TranformationComponent)(nil)).(*TranformationComponent)
		// render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
	}, (*RenderComponent)(nil), (*TranformComponent)(nil))
}
