package interaction

import "reflect"

type Interaction interface {
	Interact(delta float32, interactorComponents, interacteeComponents []reflect.Type)
	GetInteractorComponents() []reflect.Type
	GetInteracteeComponents() []reflect.Type
}
