package utils

import (
	"github.com/double-dev/limitengine/ecs"
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
			for i, axis := range control.Axis {
				motion.Acceleration[i] = axis.Amount() * control.Speed
			}
		}
	}, (*MotionControlComponent)(nil), (*MotionComponent)(nil))
}
