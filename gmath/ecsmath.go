package gmath

import (
	"github.com/double-dev/limitengine"
)

type TransformComponent struct {
	Position Vector3
	Rotation Quaternion
	Scale    Vector3
}

func (transformComponent *TransformComponent) Delete() {}

type MotionComponent struct {
	awake bool

	Velocity        Vector3
	Acceleration    Vector3
	AngVelocity     Quaternion
	AngAcceleration Quaternion
}

func (motionComponent *MotionComponent) Delete() {}

func (motion *MotionComponent) IsAwake() bool {
	return motion.awake
}

func NewMotionSystem(damping float32) *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			transform := entity.GetComponent((*TransformComponent)(nil)).(*TransformComponent)
			motion := entity.GetComponent((*MotionComponent)(nil)).(*MotionComponent)
			motion.Velocity.AddV(motion.Acceleration.Clone().MulSc(delta))

			// TODO: Optimize/fix awake system.
			if motion.Velocity.LenSq() <= 0.0001 {
				motion.awake = false
			} else {
				motion.awake = true
			}

			// TODO: Implement proper damping.
			motion.Velocity.MulSc(damping)

			transform.Position.AddV(motion.Velocity.Clone().MulSc(delta))

			motion.AngVelocity.MulQ(NewIdentityQuaternion().Slerp(motion.AngAcceleration, delta))

			motion.AngVelocity = NewIdentityQuaternion().Slerp(motion.AngVelocity, 0.95)

			transform.Rotation.MulQ(NewIdentityQuaternion().Slerp(motion.AngVelocity, delta))
		}
	}, (*MotionComponent)(nil), (*TransformComponent)(nil))
}
