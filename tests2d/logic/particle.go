package logic

import (
	"math/rand"
	"reflect"

	"github.com/double-dev/limitengine"

	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
)

type ParticleTrailComponent struct {
	Particles []limitengine.ECSEntity
	index     int
	timer     float32
}

func NewParticleTrailSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			trail := components[0].(*ParticleTrailComponent)
			transform := components[1].(*gmath.TransformComponent)
			motion := components[2].(*gmath.MotionComponent)
			if gmath.Abs(motion.Velocity[0]) > 0.25 && gmath.Abs(motion.Velocity[1]) < 0.01 {
				trail.timer += delta
				if trail.timer > 0.05 {
					sign := float32(1.0)
					if motion.Velocity[0] < 0.0 {
						sign = -1.0
					}

					particle := trail.Particles[trail.index]
					particleTransform := particle.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
					particleMotion := particle.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)
					particleTransform.Position.SetV(transform.Position)
					particleTransform.Position[0] += rand.Float32() * 0.05
					particleTransform.Position[1] += rand.Float32() * 0.025
					particleMotion.Velocity.Set(rand.Float32()*sign, rand.Float32()+0.5, 0.0)
					particleMotion.Acceleration.Set(0.0, -6.0, 0.0)
					trail.index++
					if trail.index >= len(trail.Particles) {
						trail.index = 0.0
					}
					trail.timer = 0.0
				}
			}
		}
	}, (*ParticleTrailComponent)(nil), (*gmath.TransformComponent)(nil), (*gmath.MotionComponent)(nil))
}

type ParticleComponent struct {
	resetPos gmath.Vector3
}

type ParticleInteraction struct{}

func (interation *ParticleInteraction) StartInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3, penetration float32) {
	interactor.Motion.Acceleration.Set(0.0, 0.0, 0.0)
	interactor.Motion.Velocity.Set(0.0, 0.0, 0.0)
	particle := interactor.Entity.GetComponent((*ParticleComponent)(nil)).(*ParticleComponent)
	interactor.Transform.Position.SetV(particle.resetPos)
}

func (interation *ParticleInteraction) EndInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3) {
}

func (interation *ParticleInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*ParticleComponent)(nil)),
	}
}

func (interation *ParticleInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{}
}
