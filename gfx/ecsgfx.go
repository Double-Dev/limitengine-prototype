package gfx

import (
	"fmt"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type RenderComponent struct {
	Camera   *Camera
	Shader   *Shader
	Material *Material
	Mesh     *Mesh
	Instance *Instance
}

func NewRenderSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
			render.Instance.SetTransform(transformMat)

			Render(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
		}
		fmt.Println("Sweeping...")
		Sweep()
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}

type CameraComponent struct {
	Camera      *Camera
	PosOffset   gmath.Vector3
	RotOffset   gmath.Quaternion
	ScaleOffset gmath.Vector3
}

func NewCameraMotionSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			camera := entity.GetComponent((*CameraComponent)(nil)).(*CameraComponent)
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)

			transformViewMat := gmath.NewViewMatrix(
				transform.Position,
				transform.Rotation,
				transform.Scale,
			)
			offsetViewMat := gmath.NewViewMatrix(
				camera.PosOffset,
				camera.RotOffset,
				camera.ScaleOffset,
			)

			viewMat := offsetViewMat.MulM(transformViewMat)

			camera.Camera.SetViewMat(viewMat)
		}
	}, (*CameraComponent)(nil), (*gmath.TransformComponent)(nil))
}
