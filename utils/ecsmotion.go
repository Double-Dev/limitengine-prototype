package utils

import (
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gmath"
)

type MotionComponent struct {
	Velocity        gmath.Vector3
	Acceleration    gmath.Vector3
	AngVelocity     gmath.Quaternion
	AngAcceleration gmath.Quaternion
}

func NewMotionSystem(damping float32) *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		for _, entity := range entities {
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)
			motion := entity.GetComponent((*MotionComponent)(nil)).(*MotionComponent)
			motion.Velocity.AddV(motion.Acceleration.Clone().MulSc(delta))

			// TODO: Implement proper damping.
			motion.Velocity.MulSc(damping)

			transform.Position.AddV(motion.Velocity.Clone().MulSc(delta))

			motion.AngVelocity.MulQ(gmath.NewIdentityQuaternion().Slerp(motion.AngAcceleration, delta))

			motion.AngVelocity = gmath.NewIdentityQuaternion().Slerp(motion.AngVelocity, 0.95)

			transform.Rotation.MulQ(gmath.NewIdentityQuaternion().Slerp(motion.AngVelocity, delta))
		}
	}, (*MotionComponent)(nil), (*TransformComponent)(nil))
}
