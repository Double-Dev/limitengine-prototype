package limitengine

import (
	"reflect"
	"sync"
)

var (
	entityIndex = uint32(0)
)

type ECS struct {
	ecs       map[ECSEntity]map[reflect.Type]ECSComponent
	mutex     sync.RWMutex
	listeners []ECSListener
	systems   []*ECSSystem
}

func NewECS() *ECS {
	return &ECS{
		ecs: make(map[ECSEntity]map[reflect.Type]ECSComponent), // TODO: Figure out an optimal 'max' size/number to entities needed in the program.
	}
}

type ECSEntity struct {
	id  uint32
	ecs *ECS
}

type ECSComponent interface{}

type ECSListener interface {
	OnAddEntity(entity ECSEntity)
	OnAddComponent(entity ECSEntity, component ECSComponent)
	OnRemoveComponent(entity ECSEntity, component ECSComponent)
	OnRemoveEntity(entity ECSEntity)
	GetTargetComponents() []reflect.Type
	GetEntities() []ECSEntity
	ShouldListenForAllComponents() bool
}

func (ecs *ECS) NewEntity(components ...ECSComponent) ECSEntity {
	entity := ECSEntity{
		id:  entityIndex,
		ecs: ecs,
	}
	entityIndex++
	ecsEntity := make(map[reflect.Type]ECSComponent, len(components))
	for _, component := range components {
		ecsEntity[reflect.TypeOf(component)] = component
	}
	ecs.ecs[entity] = ecsEntity
	for _, listener := range ecs.listeners {
		listener.GetTargetComponents()
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(entity)
		}
	}
	return entity
}

func (entity ECSEntity) AddComponent(component ECSComponent) {
	componentType := reflect.TypeOf(component)
	entity.ecs.mutex.Lock()
	entity.ecs.ecs[entity][componentType] = component
	entity.ecs.mutex.Unlock()
	for _, listener := range entity.ecs.listeners {
		if listener.ShouldListenForAllComponents() {
			listener.OnAddComponent(entity, component)
		} else {
			for _, target := range listener.GetTargetComponents() {
				if target == componentType {
					listener.OnAddComponent(entity, component)
					break
				}
			}
		}
	}
}

func (entity ECSEntity) RemoveComponent(nilComponent interface{}) bool { // Consider taking in reflect.Type instead of nil components (would potentially be faster).
	componentType := reflect.TypeOf(nilComponent)
	entity.ecs.mutex.RLock()
	component := entity.ecs.ecs[entity][componentType]
	entity.ecs.mutex.RUnlock()
	if component != nil {
		for _, listener := range entity.ecs.listeners {
			if listener.ShouldListenForAllComponents() {
				listener.OnRemoveComponent(entity, component)
			} else {
				for _, target := range listener.GetTargetComponents() {
					if target == componentType {
						listener.OnRemoveComponent(entity, component)
						break
					}
				}
			}
		}
		entity.ecs.mutex.Lock()
		delete(entity.ecs.ecs[entity], componentType)
		entity.ecs.mutex.Unlock()
		return true
	}
	return false
}

func (ecs *ECS) RemoveEntity(entity ECSEntity) bool {
	ecs.mutex.RLock()
	components := ecs.ecs[entity]
	ecs.mutex.RUnlock()
	if components != nil {
		for _, listener := range ecs.listeners {
			for _, listenEntity := range listener.GetEntities() {
				if entity == listenEntity {
					listener.OnRemoveEntity(entity)
					break
				}
			}
		}
		components = nil
		// ecs.mutex.Lock()
		delete(ecs.ecs, entity)
		// ecs.mutex.Unlock()
		return true
	}
	return false
}

func (entity ECSEntity) GetComponent(nilComponent interface{}) ECSComponent {
	return entity.GetComponentOfType(reflect.TypeOf(nilComponent))
}

func (entity ECSEntity) GetComponentOfType(componentType reflect.Type) ECSComponent {
	entity.ecs.mutex.RLock()
	defer entity.ecs.mutex.RUnlock()
	if entity.ecs.ecs[entity] == nil {
		return nil
	}
	component := entity.ecs.ecs[entity][componentType]
	return component
}

func (entity ECSEntity) HasComponent(targets ...reflect.Type) bool {
	entity.ecs.mutex.RLock()
	defer entity.ecs.mutex.RUnlock()
	if entity.ecs.ecs[entity] == nil {
		return false
	}
	for _, target := range targets {
		if entity.ecs.ecs[entity][target] == nil {
			return false
		}
	}
	return true
}

func (ecs *ECS) AddECSListener(listener ECSListener) {
	for entity := range ecs.ecs {
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(entity)
		}
	}
	ecs.listeners = append(ecs.listeners, listener)
}

func (ecs *ECS) RemoveECSListener(listener ECSListener) {
	for _, entity := range listener.GetEntities() {
		listener.OnRemoveEntity(entity)
	}
	for i := 0; i < len(ecs.listeners); i++ {
		if ecs.listeners[i] == listener {
			ecs.listeners[i] = ecs.listeners[len(ecs.listeners)-1]
			ecs.listeners[len(ecs.listeners)-1] = nil
			ecs.listeners = ecs.listeners[:len(ecs.listeners)-1]
		}
	}
}

func (ecs *ECS) AddECSSystem(system *ECSSystem) {
	ecs.AddECSListener(system)
	ecs.systems = append(ecs.systems, system)
}

func (ecs *ECS) UpdateSystems(delta float32) {
	for _, system := range ecs.systems {
		system.Update(delta)
	}
}

func (ecs *ECS) RemoveECSSystem(system *ECSSystem) {
	ecs.RemoveECSListener(system)
	for i := 0; i < len(ecs.systems); i++ {
		if ecs.systems[i] == system {
			ecs.systems[i] = ecs.systems[len(ecs.systems)-1]
			ecs.systems[len(ecs.systems)-1] = nil
			ecs.systems = ecs.systems[:len(ecs.systems)-1]
		}
	}
}
