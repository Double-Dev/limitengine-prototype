package interaction

import (
	"fmt"
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type ColliderComponent struct {
	AABB gmath.AABB
}

type interactEntity struct {
	entity      limitengine.ECSEntity
	interactors []Interaction
	interactees []Interaction
}

var (
	targets = []reflect.Type{
		reflect.TypeOf((*ColliderComponent)(nil)),
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
	}
)

type World struct {
	entities         []limitengine.ECSEntity
	entitiesToRemove []limitengine.ECSEntity
	interactions     []Interaction
}

func NewWorld() *World {
	world := World{}
	return &world
}

func (world *World) OnAddEntity(entity limitengine.ECSEntity) {
	world.entities = append(world.entities, entity)
	fmt.Println(world.entities)
}

func (world *World) OnAddComponent(entity limitengine.ECSEntity, component interface{}) {
	if (reflect.TypeOf(component) == targets[0] && entity.HasComponent(targets[1])) ||
		(reflect.TypeOf(component) == targets[1] && entity.HasComponent(targets[0])) {
		world.entities = append(world.entities, entity)
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {

	}
}

func (world *World) OnRemoveComponent(entity limitengine.ECSEntity, component interface{}) {
	if reflect.TypeOf(component) == targets[0] || reflect.TypeOf(component) == targets[1] {
		world.entitiesToRemove = append(world.entitiesToRemove, entity)
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {
		// Update in entities
	}
}

func (world *World) OnRemoveEntity(entity limitengine.ECSEntity) {
	world.entitiesToRemove = append(world.entitiesToRemove, entity)
}

func (world *World) ProcessInteractions(delta float32) {
	for _, removeEntity := range world.entitiesToRemove {
		for i, entity := range world.entities {
			if entity == removeEntity {
				world.entities[i] = world.entities[len(world.entities)-1]
				world.entities = world.entities[:len(world.entities)-1]
				break
			}
		}
	}

	// Update entities

	for i := 0; i < len(world.entities); i++ {
		aabbA := world.entities[i].GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
		transformA := world.entities[i].GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
		colliderA := gmath.NewAABB(aabbA.Min.Clone().AddV(transformA.Position), aabbA.Max.Clone().AddV(transformA.Position))
		for j := i + 1; j < len(world.entities); j++ {
			aabbB := world.entities[j].GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
			transformB := world.entities[j].GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
			colliderB := gmath.NewAABB(aabbB.Min.Clone().AddV(transformB.Position), aabbB.Max.Clone().AddV(transformB.Position))
			if colliderA.IntersectsAABB2D(colliderB) ||
				colliderA.ContainsAABB2D(colliderB) ||
				colliderB.ContainsAABB2D(colliderA) {

				// world.entities[i].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)
				// world.entities[j].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)

				world.entities[i].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.MulSc(-1.0)
				world.entities[j].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.MulSc(-1.0)
			}
		}
	}
}

func (world *World) AddInteraction(interaction Interaction) {
	world.interactions = append(world.interactions, interaction)
}

func (world *World) RemoveInteraction(interaction Interaction) {
	for i, loopInteraction := range world.interactions {
		if interaction == loopInteraction {
			world.interactions[i] = world.interactions[len(world.interactions)-1]
			world.interactions = world.interactions[:len(world.interactions)-1]
			break
		}
	}
}

func (world *World) GetTargetComponents() []reflect.Type { return targets }
func (world *World) GetEntities() []limitengine.ECSEntity {
	return world.entities
}
func (world *World) ShouldListenForAllComponents() bool { return true }
