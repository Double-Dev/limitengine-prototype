package interaction

import (
	"reflect"

	"github.com/double-dev/limitengine/gmath"
)

type Interaction interface {
	Interact(delta float32, interactor, interactee InteractEntity, normal gmath.Vector3, penetration float32)
	GetInteractorComponents() []reflect.Type
	GetInteracteeComponents() []reflect.Type
}

func interactionHasComponent(interaction Interaction, componentType reflect.Type) bool {
	for _, interactionTarget := range interaction.GetInteractorComponents() {
		if interactionTarget == componentType {
			return true
		}
	}
	for _, interactionTarget := range interaction.GetInteracteeComponents() {
		if interactionTarget == componentType {
			return true
		}
	}
	return false
}
