package logic

import (
	"reflect"

	"github.com/double-dev/limitengine"

	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
)

type ParticleTrailComponent struct {
	Particles []limitengine.ECSEntity
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
