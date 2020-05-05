package interaction

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type ColliderComponent struct {
	AABB gmath.AABB
}

type PhysicsComponent struct {
	awake bool

	Velocity        gmath.Vector3
	Acceleration    gmath.Vector3
	AngVelocity     gmath.Quaternion
	AngAcceleration gmath.Quaternion
}

func NewPhysicsSystem(damping float32) *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
			motion := entity.GetComponent((*PhysicsComponent)(nil)).(*PhysicsComponent)
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

			motion.AngVelocity.MulQ(gmath.NewIdentityQuaternion().Slerp(motion.AngAcceleration, delta))

			motion.AngVelocity = gmath.NewIdentityQuaternion().Slerp(motion.AngVelocity, 0.95)

			transform.Rotation.MulQ(gmath.NewIdentityQuaternion().Slerp(motion.AngVelocity, delta))
		}
	}, (*PhysicsComponent)(nil), (*gmath.TransformComponent)(nil))
}
