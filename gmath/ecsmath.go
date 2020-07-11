package gmath

import (
	"github.com/double-dev/limitengine"
)

type TransformComponent struct {
	notAwake bool

	Position Vector3
	Rotation Quaternion
	Scale    Vector3
}

func (transform *TransformComponent) SetAwake(awake bool) {
	transform.notAwake = !awake
}

func (transform *TransformComponent) IsAwake() bool {
	return !transform.notAwake
}

type MotionComponent struct {
	awake bool

	Velocity        Vector3
	Acceleration    Vector3
	AngVelocity     Quaternion
	AngAcceleration Quaternion
}

func (motion *MotionComponent) IsAwake() bool {
	return motion.awake
}

func NewMotionSystem(damping float32) *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			transform := components[1].(*TransformComponent)
			motion := components[0].(*MotionComponent)
			if motion.Velocity.LenSq() > 0.0 || motion.Velocity.LenSq() > 0.0 {
				transform.notAwake = false
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
			} else {
				transform.notAwake = true
			}
		}
	}, (*MotionComponent)(nil), (*TransformComponent)(nil))
}
