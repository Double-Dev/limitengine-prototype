package interaction

import "reflect"

type Interaction interface {
	Interact(delta float32, interactorComponents, interacteeComponents []reflect.Type)
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
