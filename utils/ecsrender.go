package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type RenderComponent struct {
	Camera   *gfx.Camera
	Shader   *gfx.Shader
	Material *gfx.Material
	Model    *gfx.Model
}

func NewRenderSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		// for _, entity := range entities {
		// 	transform := entity.GetComponent((*TranformationComponent)(nil)).(*TranformationComponent)
		// 	render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
		// }
	}, (*RenderComponent)(nil), (*TransformComponent)(nil))
}

type CameraComponent struct {
	Camera    *gfx.Camera
	PosOffset gmath.Vector
}

func NewCameraMotionSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		for _, entity := range entities {
			camera := entity.GetComponent((*CameraComponent)(nil)).(*CameraComponent)
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)

			viewMat := gmath.NewIdentityMatrix(4, 4)
			viewMat.Translate(transform.Position.Clone().Add(camera.PosOffset...).MulSc(-1.0))

			camera.Camera.SetViewMat(viewMat.ToMatrix44())
		}
	}, (*CameraComponent)(nil), (*TransformComponent)(nil))
}
