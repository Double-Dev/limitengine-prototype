package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/ui"
)

type MotionControlComponent struct {
	Axis []ui.InputControl
}

func NewMotionControlSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		control := entity.GetComponent((*MotionControlComponent)(nil)).(*MotionControlComponent)
		motion := entity.GetComponent((*MotionComponent)(nil)).(*MotionComponent)
		for i, axis := range control.Axis {
			motion.Acceleration[i] = axis.Amount()
		}
	}, (*MotionControlComponent)(nil), (*MotionComponent)(nil))
}
