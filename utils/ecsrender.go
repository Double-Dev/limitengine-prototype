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
	Mesh     *gfx.Mesh
	Instance *gfx.Instance
}

func NewRenderSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)

			transformMat := gmath.NewTransformMatrix3D(transform.Position, transform.Rotation, transform.Scale)

			render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
			render.Instance.SetTransform(transformMat)

			gfx.Render(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
		}
		gfx.Sweep()
	}, (*RenderComponent)(nil), (*TransformComponent)(nil))
}

type CameraComponent struct {
	Camera      *gfx.Camera
	PosOffset   gmath.Vector3
	RotOffset   gmath.Quaternion
	ScaleOffset gmath.Vector3
}

func NewCameraMotionSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		for _, entity := range entities {
			camera := entity.GetComponent((*CameraComponent)(nil)).(*CameraComponent)
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)

			transformViewMat := gmath.NewViewMatrix3D(
				transform.Position,
				transform.Rotation,
				transform.Scale,
			)
			offsetViewMat := gmath.NewViewMatrix3D(
				camera.PosOffset,
				camera.RotOffset,
				camera.ScaleOffset,
			)

			viewMat := offsetViewMat.MulM(transformViewMat)

			camera.Camera.SetViewMat(viewMat)
		}
	}, (*CameraComponent)(nil), (*TransformComponent)(nil))
}
