package interaction

import (
	"reflect"
	"time"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

var (
	targets = []reflect.Type{
		reflect.TypeOf((*ColliderComponent)(nil)),
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
	}
)

type World struct {
	spacialStructure SpacialStructure
	entities         map[limitengine.ECSEntity]*InteractEntity
	entitiesToRemove []limitengine.ECSEntity
	interactions     []Interaction
}

func NewWorld(spacialStructure SpacialStructure, targetUpdatesPerSecond float32) *World {
	world := World{
		spacialStructure: spacialStructure,
		entities:         make(map[limitengine.ECSEntity]*InteractEntity),
	}
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
	interactEntity := world.createInteractEntity(entity)
	world.entities[entity] = interactEntity
	world.spacialStructure.Add(interactEntity)
}

func (world *World) OnAddComponent(entity limitengine.ECSEntity, component interface{}) {
	if (reflect.TypeOf(component) == targets[0] && entity.HasComponent(targets[1])) ||
		(reflect.TypeOf(component) == targets[1] && entity.HasComponent(targets[0])) {
		interactEntity := world.createInteractEntity(entity)
		world.entities[entity] = interactEntity
		world.spacialStructure.Add(interactEntity)
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

func (world *World) createInteractEntity(entity limitengine.ECSEntity) *InteractEntity {
	interactEntity := &InteractEntity{
		entity:    entity,
		transform: entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent),
		collider:  entity.GetComponent((*ColliderComponent)(nil)).(*ColliderComponent),
		physics:   entity.GetComponent((*PhysicsComponent)(nil)).(*PhysicsComponent),
	}
	for _, interaction := range world.interactions {
		world.updateInteraction(interactEntity, interaction)
	}
	return interactEntity
}

func (world *World) updateInteraction(entity *InteractEntity, interaction Interaction) {
	isInteractor := true
	for _, target := range interaction.GetInteractorComponents() {
		if !entity.entity.HasComponent(target) {
			isInteractor = false
			break
		}
	}
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
		world.spacialStructure.Remove(world.entities[removeEntity])
		delete(world.entities, removeEntity)
	}
	world.entitiesToRemove = nil

	for _, interactEntity := range world.entities {
		if interactEntity.physics.awake {
			world.spacialStructure.Update(interactEntity)
		}
	}

	for _, interactEntityA := range world.entities {
		if interactEntityA.physics != nil && interactEntityA.physics.awake {
			potentialCollisions := world.spacialStructure.Query(gmath.NewAABB(
				interactEntityA.collider.AABB.Min.Clone().AddV(interactEntityA.transform.Position),
				interactEntityA.collider.AABB.Max.Clone().AddV(interactEntityA.transform.Position),
			))
			for _, interactEntityB := range potentialCollisions {
				if interactEntityA == interactEntityB {
					continue
				}
				// TODO: For other shapes, additional collision checks would be made here.
				// Additionally, the collision normal and penetration should be calculated
				// here for every collision and passed to the interactions.

				for _, interactor := range interactEntityA.interactors {
					for _, interactee := range interactEntityB.interactees {
						if interactor == interactee {
							interactor.Interact(delta, interactEntityA.entity, interactEntityB.entity)
							break
						}
					}
				}
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
