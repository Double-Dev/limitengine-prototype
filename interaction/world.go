package interaction

import (
	"fmt"
	"reflect"
	"time"

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
	entities         []*interactEntity
	entitiesToRemove []limitengine.ECSEntity
	interactions     []Interaction
}

func NewWorld(targetUpdatesPerSecond float32) *World {
	world := World{}
	go func() {
		currentTime := time.Now().UnixNano()
		for limitengine.Running() {
			if time.Now().UnixNano()-currentTime > int64((1.0/targetUpdatesPerSecond)*1000000000.0) {
				lastTime := currentTime
				currentTime = time.Now().UnixNano()
				delta := float32(currentTime-lastTime) / 1000000000.0

				world.ProcessInteractions(delta)
			} else {
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()
	return &world
}

func (world *World) OnAddEntity(entity limitengine.ECSEntity) {
	world.entities = append(world.entities, world.createInteractEntity(entity))
}

func (world *World) OnAddComponent(entity limitengine.ECSEntity, component interface{}) {
	if (reflect.TypeOf(component) == targets[0] && entity.HasComponent(targets[1])) ||
		(reflect.TypeOf(component) == targets[1] && entity.HasComponent(targets[0])) {
		world.entities = append(world.entities, world.createInteractEntity(entity))
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {
		for _, interactEntity := range world.entities {
			if interactEntity.entity == entity {
				for _, interaction := range world.interactions {
					if interactionHasComponent(interaction, reflect.TypeOf(component)) {
						world.updateInteraction(interactEntity, interaction)
					}
				}
				break
			}
		}
	}
}

func (world *World) createInteractEntity(entity limitengine.ECSEntity) *interactEntity {
	interactEntity := &interactEntity{
		entity: entity,
	}
	for _, interaction := range world.interactions {
		world.updateInteraction(interactEntity, interaction)
	}
	return interactEntity
}

func (world *World) updateInteraction(entity *interactEntity, interaction Interaction) {
	isInteractor := true
	for _, target := range interaction.GetInteractorComponents() {
		if !entity.entity.HasComponent(target) {
			isInteractor = false
			break
		}
	}
	fmt.Println(isInteractor)
	if isInteractor {
		entity.interactors = append(entity.interactors, interaction)
	}
	isInteractee := true
	for _, target := range interaction.GetInteracteeComponents() {
		if !entity.entity.HasComponent(target) {
			isInteractee = false
			break
		}
	}
	if isInteractee {
		fmt.Println(interaction)
		entity.interactees = append(entity.interactees, interaction)
	}
}

func (world *World) OnRemoveComponent(entity limitengine.ECSEntity, component interface{}) {
	if reflect.TypeOf(component) == targets[0] || reflect.TypeOf(component) == targets[1] {
		world.entitiesToRemove = append(world.entitiesToRemove, entity)
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {
		for _, interactEntity := range world.entities {
			if interactEntity.entity == entity {
				for i, interaction := range interactEntity.interactors {
					if interactionHasComponent(interaction, reflect.TypeOf(component)) {
						interactEntity.interactors[i] = interactEntity.interactors[len(interactEntity.interactors)-1]
						interactEntity.interactors = interactEntity.interactors[:len(interactEntity.interactors)-1]
					}
				}
				for i, interaction := range interactEntity.interactees {
					if interactionHasComponent(interaction, reflect.TypeOf(component)) {
						interactEntity.interactees[i] = interactEntity.interactees[len(interactEntity.interactees)-1]
						interactEntity.interactees = interactEntity.interactees[:len(interactEntity.interactees)-1]
					}
				}
				break
			}
		}
	}
}

func (world *World) OnRemoveEntity(entity limitengine.ECSEntity) {
	world.entitiesToRemove = append(world.entitiesToRemove, entity)
}

func (world *World) ProcessInteractions(delta float32) {
	for _, removeEntity := range world.entitiesToRemove {
		for i, interactEntity := range world.entities {
			if interactEntity.entity == removeEntity {
				world.entities[i] = world.entities[len(world.entities)-1]
				world.entities = world.entities[:len(world.entities)-1]
				break
			}
		}
	}
	world.entitiesToRemove = nil

	// TODO: Optimize collision loops.
	for i := 0; i < len(world.entities); i++ {
		aabbA := world.entities[i].entity.GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
		transformA := world.entities[i].entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
		colliderA := gmath.NewAABB(aabbA.Min.Clone().AddV(transformA.Position), aabbA.Max.Clone().AddV(transformA.Position))
		for j := i + 1; j < len(world.entities); j++ {
			aabbB := world.entities[j].entity.GetComponent((*ColliderComponent)(nil)).(*ColliderComponent).AABB
			transformB := world.entities[j].entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)
			colliderB := gmath.NewAABB(aabbB.Min.Clone().AddV(transformB.Position), aabbB.Max.Clone().AddV(transformB.Position))
			if colliderA.IntersectsAABB2D(colliderB) ||
				colliderA.ContainsAABB2D(colliderB) ||
				colliderB.ContainsAABB2D(colliderA) {

				for _, interactEntity := range world.entities {
					fmt.Println(len(interactEntity.interactees))
				}

				// TODO: Perform interactions
				// world.entities[i]
				// world.entities[j]

				// world.entities[i].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)
				// world.entities[j].GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.Set(0.0, 0.0, 0.0)

				world.entities[i].entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.MulSc(-1.0)
				world.entities[j].entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.MulSc(-1.0)
			}
		}
	}
}

func (world *World) AddInteraction(interaction Interaction) {
	world.interactions = append(world.interactions, interaction)
	for _, interactEntity := range world.entities {
		world.updateInteraction(interactEntity, interaction)
	}
}

func (world *World) RemoveInteraction(interaction Interaction) {
	for _, interactEntity := range world.entities {
		for i, entityInteraction := range interactEntity.interactors {
			if entityInteraction == interaction {
				interactEntity.interactors[i] = interactEntity.interactors[len(interactEntity.interactors)-1]
				interactEntity.interactors = interactEntity.interactors[:len(interactEntity.interactors)-1]
				break
			}
		}
		for i, entityInteraction := range interactEntity.interactees {
			if entityInteraction == interaction {
				interactEntity.interactees[i] = interactEntity.interactees[len(interactEntity.interactees)-1]
				interactEntity.interactees = interactEntity.interactees[:len(interactEntity.interactees)-1]
				break
			}
		}
	}
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
	entities := []limitengine.ECSEntity{}
	for _, interactEntity := range world.entities {
		entities = append(entities, interactEntity.entity)
	}
	return entities
}
func (world *World) ShouldListenForAllComponents() bool { return true }
