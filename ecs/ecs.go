package ecs

import (
	"reflect"
)

var (
	entityIndex = uint32(0)
	ecs         = make(map[ECSEntity]map[reflect.Type]interface{})
	listeners   = []ECSListener{}
)

type ECSEntity uint32

type ECSListener interface {
	OnAddEntity(entity *ECSEntity)
	OnAddComponent(entity *ECSEntity, component interface{})
	OnRemoveComponent(entity *ECSEntity, component interface{})
	OnRemoveEntity(entity *ECSEntity)
	GetTargetComponents() []reflect.Type
	GetEntities() []ECSEntity
}

func NewEntity(components ...interface{}) *ECSEntity {
	entity := ECSEntity(entityIndex)
	ecs[entity] = make(map[reflect.Type]interface{})
	entityIndex++
	for _, component := range components {
		ecs[entity][reflect.TypeOf(component)] = component
	}
	for _, listener := range listeners {
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(&entity)
		}
	}
	return &entity
}

func (entity *ECSEntity) AddComponent(component interface{}) {
	componentType := reflect.TypeOf(component)
	ecs[*entity][componentType] = component
	for _, listener := range listeners {
		for _, target := range listener.GetTargetComponents() {
			if target == componentType {
				listener.OnAddComponent(entity, component)
				break
			}
		}
	}
}

func (entity *ECSEntity) RemoveComponent(nilComponent interface{}) bool {
	componentType := reflect.TypeOf(nilComponent)
	if entity.HasComponent(componentType) {
		component := ecs[*entity][componentType]
		for _, listener := range listeners {
			for _, target := range listener.GetTargetComponents() {
				if target == componentType {
					listener.OnRemoveComponent(entity, component)
					break
				}
			}
		}
		delete(ecs[*entity], componentType)
		return true
	}
	return false
}

func RemoveEntity(entity *ECSEntity) bool {
	if entity != nil && ecs[*entity] != nil {
		for _, listener := range listeners {
			for _, listenEntity := range listener.GetEntities() {
				if entity == &listenEntity {
					listener.OnRemoveEntity(entity)
				}
				break
			}
		}
		delete(ecs, *entity)
		return true
	}
	return false
}

func (entity *ECSEntity) GetComponent(nilComponent interface{}) interface{} {
	if entity.HasComponent(reflect.TypeOf(nilComponent)) {
		return ecs[*entity][reflect.TypeOf(nilComponent)]
	}
	return nil
}

func (entity *ECSEntity) HasComponent(targets ...reflect.Type) bool {
	if ecs[*entity] == nil {
		return false
	}
	for _, target := range targets {
		if ecs[*entity][target] == nil {
			return false
		}
	}
	return true
}

func AddECSListener(listener ECSListener) {
	for entity := range ecs {
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(&entity)
		}
	}
	listeners = append(listeners, listener)
}

func (entity *ECSEntity) RemoveECSListener(listener ECSListener) {
	for i := 0; i < len(listeners); i++ {
		if listeners[i] == listener {
			listeners[i] = listeners[len(listeners)-1]
			listeners = listeners[:len(listeners)-1]
		}
	}
}
