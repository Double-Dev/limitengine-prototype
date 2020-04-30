package interaction

import (
	"fmt"
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

func NewInteractionWorld() *InteractionWorld {
	interactionWorld := InteractionWorld{
		entities: []limitengine.ECSEntity{},
	}
	return &interactionWorld
}

func (interactionWorld *InteractionWorld) OnAddEntity(entity limitengine.ECSEntity) {
	interactionWorld.entities = append(interactionWorld.entities, entity)
	fmt.Println(interactionWorld.entities)
}

func (interactionWorld *InteractionWorld) OnAddComponent(entity limitengine.ECSEntity, component interface{}) {
	if (reflect.TypeOf(component) == interactTargets[0] && entity.HasComponent(interactTargets[1])) ||
		(reflect.TypeOf(component) == interactTargets[1] && entity.HasComponent(interactTargets[0])) {
		interactionWorld.entities = append(interactionWorld.entities, entity)
	} else if entity.HasComponent(interactTargets[0]) && entity.HasComponent(interactTargets[1]) {

	}
}

func (interactionWorld *InteractionWorld) OnRemoveComponent(entity limitengine.ECSEntity, component interface{}) {
	if reflect.TypeOf(component) == interactTargets[0] || reflect.TypeOf(component) == interactTargets[1] {
		// Remove from entities
	} else if entity.HasComponent(interactTargets[0]) && entity.HasComponent(interactTargets[1]) {
		// Update in entities
	}
}

func (interactionWorld *InteractionWorld) OnRemoveEntity(entity limitengine.ECSEntity) {
	// Remove from entities
}

func (interactionWorld *InteractionWorld) ProcessInteractions(delta float32) {
	// Remove entities
	// Update entities

	for i := 0; i < len(interactionWorld.entities); i++ {
		aabbA := interactionWorld.entities[i].GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
		transformA := interactionWorld.entities[i].GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
		colliderA := gmath.NewAABB(aabbA.Min.Clone().AddV(transformA.Position), aabbA.Max.Clone().AddV(transformA.Position))
		for j := i + 1; j < len(interactionWorld.entities); j++ {
			aabbB := interactionWorld.entities[j].GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
			transformB := interactionWorld.entities[j].GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
			colliderB := gmath.NewAABB(aabbB.Min.Clone().AddV(transformB.Position), aabbB.Max.Clone().AddV(transformB.Position))
			if colliderA.IntersectsAABB2D(colliderB) ||
				colliderA.ContainsAABB2D(colliderB) ||
				colliderB.ContainsAABB2D(colliderA) {

				interactionWorld.entities[i].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)
				interactionWorld.entities[j].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)
			}
		}
	}
}

func (interactionWorld *InteractionWorld) GetTargetComponents() []reflect.Type { return interactTargets }
func (interactionWorld *InteractionWorld) GetEntities() []limitengine.ECSEntity {
	return interactionWorld.entities
}
func (interactionWorld *InteractionWorld) ShouldListenForAllComponents() bool { return true }
