package limitengine

import (
	"reflect"
)

type ECSSystem struct {
	targetComponents []reflect.Type
	entities         []ECSEntity
	components       [][]ECSComponent
	update           func(delta float32, entities [][]ECSComponent)
}

func NewSystem(update func(delta float32, entities [][]ECSComponent), nilTargetComponents ...interface{}) *ECSSystem {
	system := ECSSystem{
		targetComponents: []reflect.Type{},
		entities:         []ECSEntity{},
		components:       [][]ECSComponent{},
		update:           update,
	}
	for _, nilTargetComponent := range nilTargetComponents {
		system.targetComponents = append(system.targetComponents, reflect.TypeOf(nilTargetComponent))
	}
	return &system
}

func (system *ECSSystem) Update(delta float32) {
	system.update(delta, system.components)
}

func (system *ECSSystem) OnAddEntity(entity ECSEntity) {
	var components []ECSComponent
	for _, target := range system.targetComponents {
		components = append(components, entity.getComponentOfType(target))
	}
	system.components = append(system.components, components)
	system.entities = append(system.entities, entity)
}

func (system *ECSSystem) OnAddComponent(entity ECSEntity, component ECSComponent) {
	if entity.HasComponent(system.GetTargetComponents()...) {
		var components []ECSComponent
		for _, target := range system.targetComponents {
			components = append(components, entity.getComponentOfType(target))
		}
		system.components = append(system.components, components)
		system.entities = append(system.entities, entity)
	}
}

func (system *ECSSystem) OnRemoveComponent(entity ECSEntity, component ECSComponent) {
	for i, sysEntity := range system.entities {
		if sysEntity == entity {
			for j := 0; j < len(system.components); j++ {
				for _, sysComponent := range system.components[j] {
					if sysComponent == component {
						copy(system.components[j:], system.components[j+1:])
						system.components = system.components[:len(system.components)-1]
						j = len(system.components)
						break
					}
				}
			}
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
			break
		}
	}
}

func (system *ECSSystem) OnRemoveEntity(entity ECSEntity) {
	for i, sysEntity := range system.entities {
		if sysEntity == entity {
			component := entity.getComponentOfType(system.targetComponents[0])
			for j := 0; j < len(system.components); j++ {
				for _, sysComponent := range system.components[j] {
					if sysComponent == component {
						copy(system.components[j:], system.components[j+1:])
						system.components = system.components[:len(system.components)-1]
						j = len(system.components)
						break
					}
				}
			}
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
			break
		}
	}
}

func (system *ECSSystem) OnActive()   {}
func (system *ECSSystem) OnInactive() {}

func (system *ECSSystem) GetTargetComponents() []reflect.Type {
	return system.targetComponents
}

func (system *ECSSystem) GetEntities() []ECSEntity {
	return system.entities
}

func (system *ECSSystem) ShouldListenForAllComponents() bool { return false }
