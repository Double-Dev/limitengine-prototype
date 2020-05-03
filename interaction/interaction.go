package interaction

import (
	"reflect"

	"github.com/double-dev/limitengine"
)

type Interaction interface {
	Interact(delta float32, interactor, interactee limitengine.ECSEntity)
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
