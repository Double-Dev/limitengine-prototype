package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gmath"
)

type MotionComponent struct {
	Velocity        gmath.Vector
	Acceleration    gmath.Vector
	AngVelocity     gmath.Vector
	AngAcceleration gmath.Vector
}

func NewMotionSystem(damping float32) *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		transform := entity.GetComponent((*TranformComponent)(nil)).(*TranformComponent)
		motion := entity.GetComponent((*MotionComponent)(nil)).(*MotionComponent)
		motion.Velocity.AddV(motion.Acceleration.Clone().MulSc(delta))

		// TODO: Implement proper damping.
		motion.Velocity.MulSc(damping)

		transform.Position.AddV(motion.Velocity.Clone().MulSc(delta))
	}, (*MotionComponent)(nil), (*TranformComponent)(nil))
}
