package limitengine

import (
	"reflect"
)

// TODO: Store component sets rather than entities to avoid map queries.
type ECSSystem struct {
	targetComponents []reflect.Type
	entities         []ECSEntity
	update           func(delta float32, entities []ECSEntity)
}

func NewSystem(update func(delta float32, entities []ECSEntity), nilTargetComponents ...interface{}) *ECSSystem {
	system := ECSSystem{
		targetComponents: []reflect.Type{},
		entities:         []ECSEntity{},
		update:           update,
	}
	for _, nilTargetComponent := range nilTargetComponents {
		system.targetComponents = append(system.targetComponents, reflect.TypeOf(nilTargetComponent))
	}
	return &system
}

func (system *ECSSystem) Update(delta float32) {
	system.update(delta, system.GetEntities())
}

func (system *ECSSystem) OnAddEntity(entity ECSEntity) {
	system.entities = append(system.entities, entity)
}

func (system *ECSSystem) OnAddComponent(entity ECSEntity, component Component) {
	if entity.HasComponent(system.GetTargetComponents()...) {
		system.entities = append(system.entities, entity)
	}
}

func (system *ECSSystem) OnRemoveComponent(entity ECSEntity, component Component) {
	for i, target := range system.entities {
		if target == entity {
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
		}
	}
}

func (system *ECSSystem) OnRemoveEntity(entity ECSEntity) {
	for i, target := range system.entities {
		if target == entity {
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
		}
	}
}

func (system *ECSSystem) GetTargetComponents() []reflect.Type {
	return system.targetComponents
}

func (system *ECSSystem) GetEntities() []ECSEntity {
	return system.entities
}

func (system *ECSSystem) ShouldListenForAllComponents() bool { return false }
