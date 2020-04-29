package interaction

import (
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

func init() {

}

type ColliderComponent struct {
	AABB gmath.AABB
}

var (
	interactTargets = []reflect.Type{
		reflect.TypeOf((*ColliderComponent)(nil)),
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
	}
)

type InteractionWorld struct {
	entities []limitengine.ECSEntity
}

func (interactionWorld InteractionWorld) OnAddEntity(entity limitengine.ECSEntity) {
	interactionWorld.entities = append(interactionWorld.entities, entity)
}

func (interactionWorld InteractionWorld) OnAddComponent(entity limitengine.ECSEntity, component interface{}) {
	if (reflect.TypeOf(component) == interactTargets[0] && entity.HasComponent(interactTargets[1])) ||
		(reflect.TypeOf(component) == interactTargets[1] && entity.HasComponent(interactTargets[0])) {
		interactionWorld.entities = append(interactionWorld.entities, entity)
	} else if entity.HasComponent(interactTargets[0]) && entity.HasComponent(interactTargets[1]) {

	}
}

func (interactionWorld InteractionWorld) OnRemoveComponent(entity limitengine.ECSEntity, component interface{}) {
	if reflect.TypeOf(component) == interactTargets[0] || reflect.TypeOf(component) == interactTargets[1] {
		// Remove from entities
	} else if entity.HasComponent(interactTargets[0]) && entity.HasComponent(interactTargets[1]) {
		// Update in entities
	}
}

func (interactionWorld InteractionWorld) OnRemoveEntity(entity limitengine.ECSEntity) {
	// Remove from entities
}

func (interactionWorld InteractionWorld) processInteractions(delta float32) {
	// Remove entities
	// Update entities

}

func (interactionWorld InteractionWorld) GetTargetComponents() []reflect.Type { return interactTargets }
func (interactionWorld InteractionWorld) GetEntities() []limitengine.ECSEntity {
	return interactionWorld.entities
}
func (interactionWorld InteractionWorld) ShouldListenForAllComponents() bool { return true }
