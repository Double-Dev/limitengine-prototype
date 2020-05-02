package main

import (
	"fmt"
	"reflect"

	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
)

type TestInteraction struct {
	test string
}

func (test TestInteraction) Interact(delta float32, interactorComponents, interacteeComponents []reflect.Type) {
	fmt.Println("This should print twice!")
}

func (test TestInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
		reflect.TypeOf((*interaction.ColliderComponent)(nil)),
	}
}

func (test TestInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
		reflect.TypeOf((*interaction.ColliderComponent)(nil)),
	}
}
