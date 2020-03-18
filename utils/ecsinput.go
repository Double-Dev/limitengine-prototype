package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
)

type MotionControlComponent struct {
	Axis  []*ui.InputControl
	Speed float32
}

func NewMotionControlSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		for _, entity := range entities {
			control := entity.GetComponent((*MotionControlComponent)(nil)).(*MotionControlComponent)
			motion := entity.GetComponent((*MotionComponent)(nil)).(*MotionComponent)
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)

			motion.Acceleration.Set(0.0, 0.0, 0.0)
			for i, axis := range control.Axis {
				direction := gmath.NewZeroVector3()
				direction[i] = axis.Amount()
				direction = transform.Rotation.RotateV(direction)
				motion.Acceleration.AddV(direction.MulSc(control.Speed))
			}
		}
	}, (*MotionControlComponent)(nil), (*MotionComponent)(nil), (*TransformComponent)(nil))
}
