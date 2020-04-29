package tests3d

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
)

type MotionControlComponent struct {
	Axis  []*ui.InputControl
	Speed float32
}

func NewMotionControlSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			control := entity.GetComponent((*MotionControlComponent)(nil)).(*MotionControlComponent)
			motion := entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)

			motion.Acceleration.Set(0.0, 0.0, 0.0)
			for i, axis := range control.Axis {
				direction := gmath.NewZeroVector3()
				direction[i] = axis.Amount()
				direction = transform.Rotation.RotateV(direction)
				motion.Acceleration.AddV(direction.MulSc(control.Speed))
			}
		}
	}, (*MotionControlComponent)(nil), (*gmath.MotionComponent)(nil), (*gmath.TransformComponent)(nil))
}
