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

type ColliderComponent struct {
	IsTrigger bool

	AABB    gmath.AABB
	InvMass float32
}

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
	interactEntity := world.newInteractEntity(entity)
	world.entities[entity] = interactEntity
	world.spacialStructure.Add(interactEntity)
}

func (world *World) OnAddComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	if (reflect.TypeOf(component) == targets[0] && entity.HasComponent(targets[1])) ||
		(reflect.TypeOf(component) == targets[1] && entity.HasComponent(targets[0])) {
		interactEntity := world.newInteractEntity(entity)
		world.entities[entity] = interactEntity
		world.spacialStructure.Add(interactEntity)
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {
		for _, interactEntity := range world.entities {
			if interactEntity.Entity == entity {
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

func (world *World) newInteractEntity(entity limitengine.ECSEntity) *InteractEntity {
	var motion *gmath.MotionComponent
	if entity.GetComponent((*gmath.MotionComponent)(nil)) != nil {
		motion = entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)
	}
	interactEntity := &InteractEntity{
		Entity:                    entity,
		Transform:                 entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent),
		Collider:                  entity.GetComponent((*ColliderComponent)(nil)).(*ColliderComponent),
		Motion:                    motion,
		previousCollidingEntities: make(map[*InteractEntity]gmath.Vector3),
		collidingEntities:         make(map[*InteractEntity]gmath.Vector3),
	}
	for _, interaction := range world.interactions {
		world.updateInteraction(interactEntity, interaction)
	}
	return interactEntity
}

func (world *World) updateInteraction(entity *InteractEntity, interaction Interaction) {
	isInteractor := true
	for _, target := range interaction.GetInteractorComponents() {
		if !entity.Entity.HasComponent(target) {
			isInteractor = false
			break
		}
	}
	if isInteractor {
		entity.interactors = append(entity.interactors, interaction)
	}
	isInteractee := true
	for _, target := range interaction.GetInteracteeComponents() {
		if !entity.Entity.HasComponent(target) {
			isInteractee = false
			break
		}
	}
	if isInteractee {
		entity.interactees = append(entity.interactees, interaction)
	}
}

func (world *World) OnRemoveComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	if reflect.TypeOf(component) == targets[0] || reflect.TypeOf(component) == targets[1] {
		world.entitiesToRemove = append(world.entitiesToRemove, entity)
	} else if entity.HasComponent(targets[0]) && entity.HasComponent(targets[1]) {
		for _, interactEntity := range world.entities {
			if interactEntity.Entity == entity {
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

func (world *World) OnActive()   {}
func (world *World) OnInactive() {}

func (world *World) ProcessInteractions(delta float32) {
	for _, removeEntity := range world.entitiesToRemove {
		world.spacialStructure.Remove(world.entities[removeEntity])
		delete(world.entities, removeEntity)
	}
	world.entitiesToRemove = nil

	for _, interactEntity := range world.entities {
		if interactEntity.Motion != nil && interactEntity.Motion.IsAwake() {
			world.spacialStructure.Update(interactEntity)
		}
	}

	for _, interactEntityA := range world.entities {
		if interactEntityA.Motion != nil && interactEntityA.Motion.IsAwake() {
			potentialCollisions := world.spacialStructure.Query(gmath.NewAABB(
				interactEntityA.Collider.AABB.Min.Clone().AddV(interactEntityA.Transform.Position),
				interactEntityA.Collider.AABB.Max.Clone().AddV(interactEntityA.Transform.Position),
			))
			for _, interactEntityB := range potentialCollisions {
				if interactEntityA == interactEntityB {
					continue
				}

				// TODO: For other shapes, additional collision checks would be made here.
				// Additionally, the collision normal and penetration should be calculated
				// here for every collision and passed to the interactions.
				// TEMPORARY CODE TO ALLOW FOR BETTER COLLISIONS IN THE DEMOS
				normal := gmath.NewZeroVector3()
				penetration := float32(0.0)
				AABB1 := gmath.NewAABB(
					interactEntityA.Collider.AABB.Min.Clone().AddV(interactEntityA.Transform.Position),
					interactEntityA.Collider.AABB.Max.Clone().AddV(interactEntityA.Transform.Position),
				)
				AABB2 := gmath.NewAABB(
					interactEntityB.Collider.AABB.Min.Clone().AddV(interactEntityB.Transform.Position),
					interactEntityB.Collider.AABB.Max.Clone().AddV(interactEntityB.Transform.Position),
				)
				xDiffMin := gmath.Min(gmath.Abs(AABB1.Min[0]-AABB2.Max[0]), gmath.Abs(AABB1.Max[0]-AABB2.Min[0]))
				yDiffMin := gmath.Min(gmath.Abs(AABB1.Min[1]-AABB2.Max[1]), gmath.Abs(AABB1.Max[1]-AABB2.Min[1]))
				if xDiffMin < yDiffMin {
					penetration = xDiffMin
					if gmath.Abs(AABB1.Min[0]-AABB2.Max[0]) < gmath.Abs(AABB1.Max[0]-AABB2.Min[0]) {
						normal.Set(-1.0, 0.0, 0.0)
					} else {
						normal.Set(1.0, 0.0, 0.0)
					}
				} else {
					penetration = yDiffMin
					if gmath.Abs(AABB1.Min[1]-AABB2.Max[1]) < gmath.Abs(AABB1.Max[1]-AABB2.Min[1]) {
						normal.Set(0.0, -1.0, 0.0)
					} else {
						normal.Set(0.0, 1.0, 0.0)
					}
				}
				// END TEMPORARY CODE

				if !interactEntityA.Collider.IsTrigger && !interactEntityB.Collider.IsTrigger {
					// TODO: Fix bouncing bug.
					var otherVel gmath.Vector3
					if interactEntityB.Motion != nil {
						otherVel = interactEntityB.Motion.Velocity.Clone()
					} else {
						otherVel = gmath.NewZeroVector3()
					}
					rv := otherVel.SubV(interactEntityA.Motion.Velocity)
					normVelocity := rv.Dot(normal)
					if normVelocity > 0 {
						continue
					}
					e := float32(1.0) // Restitution
					j := -(1.0 + e) * normVelocity
					j /= interactEntityA.Collider.InvMass + interactEntityB.Collider.InvMass
					impulse := normal.Clone().MulSc(j)
					interactEntityA.Motion.Velocity.SubV(impulse.MulSc(interactEntityA.Collider.InvMass))
					correction := normal.Clone().MulSc(penetration / (interactEntityA.Collider.InvMass + interactEntityB.Collider.InvMass) * 0.8)
					interactEntityA.Transform.Position.SubV(correction.MulSc(interactEntityA.Collider.InvMass))
				}

				interactEntityA.collidingEntities[interactEntityB] = normal
				if _, ok := interactEntityA.previousCollidingEntities[interactEntityB]; !ok {
					for _, interactor := range interactEntityA.interactors {
						for _, interactee := range interactEntityB.interactees {
							if interactor == interactee {
								interactor.StartInteract(delta, *interactEntityA, *interactEntityB, normal, penetration)
								break
							}
						}
					}
				}
				interactEntityB.collidingEntities[interactEntityA] = normal.Clone().MulSc(-1.0)
				if _, ok := interactEntityB.previousCollidingEntities[interactEntityA]; !ok {
					for _, interactor := range interactEntityB.interactors {
						for _, interactee := range interactEntityA.interactees {
							if interactor == interactee {
								interactor.StartInteract(delta, *interactEntityB, *interactEntityA, normal.Clone().MulSc(-1.0), penetration)
								break
							}
						}
					}
				}
			}
			for previousEntity, normal := range interactEntityA.previousCollidingEntities {
				if _, ok := interactEntityA.collidingEntities[previousEntity]; !ok {
					for _, interactor := range interactEntityA.interactors {
						for _, interactee := range previousEntity.interactees {
							if interactor == interactee {
								interactor.EndInteract(delta, *interactEntityA, *previousEntity, normal.Clone())
								break
							}
						}
					}
				}
			}
			interactEntityA.previousCollidingEntities = interactEntityA.collidingEntities
			interactEntityA.collidingEntities = make(map[*InteractEntity]gmath.Vector3)
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
		entities = append(entities, interactEntity.Entity)
	}
	return entities
}
func (world *World) ShouldListenForAllComponents() bool { return true }
