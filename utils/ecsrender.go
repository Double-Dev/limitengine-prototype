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
	Instance *gfx.Instance
}

func NewRenderSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)

			transformMat := gmath.NewIdentityMatrix(4, 4)
			transformMat.Translate(transform.Position)

			render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
			render.Instance.AddMatrix44("transformMat", transformMat.ToMatrix44())

			gfx.Render(render.Camera, render.Shader, render.Material, render.Model, render.Instance)
		}
		gfx.Sweep()
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
