package limitengine

import (
	"reflect"
	"sync"
)

var (
	// TargetUpdatesPerSecond determine the amount of updates ECSSystems will attempt to perform per second.
	TargetUpdatesPerSecond = float32(100.0)
	entityIndex            = uint32(0)
)

type ECS struct {
	ecs       map[ECSEntity]map[reflect.Type]Component
	mutex     sync.RWMutex
	listeners []ECSListener
}

func NewECS() *ECS {
	return &ECS{
		ecs:   make(map[ECSEntity]map[reflect.Type]Component),
		mutex: sync.RWMutex{},
	}

	// go func() {
	// 	currentTime := time.Now().UnixNano()
	// 	for Running() {
	// 		if time.Now().UnixNano()-currentTime > int64((1.0/TargetUpdatesPerSecond)*1000000000.0) {
	// 			lastTime := currentTime
	// 			currentTime = time.Now().UnixNano()
	// 			delta := float32(currentTime-lastTime) / 1000000000.0
	// 			UpdateSystems(delta)
	// 		} else {
	// 			time.Sleep(time.Millisecond * 10)
	// 		}

	// 	}
	// }()
	// TODO: Sort out ECS threading.
}

type ECSEntity struct {
	id  uint32
	ecs *ECS
}

type Component interface {
	Delete()
}

type ECSListener interface {
	OnAddEntity(entity ECSEntity)
	OnAddComponent(entity ECSEntity, component Component)
	OnRemoveComponent(entity ECSEntity, component Component)
	OnRemoveEntity(entity ECSEntity)
	GetTargetComponents() []reflect.Type
	GetEntities() []ECSEntity
	ShouldListenForAllComponents() bool
}

func (ecs *ECS) NewEntity(components ...Component) ECSEntity {
	entity := ECSEntity{
		id:  entityIndex,
		ecs: ecs,
	}
	ecs.mutex.Lock()
	ecs.ecs[entity] = make(map[reflect.Type]Component)
	entityIndex++
	for _, component := range components {
		ecs.ecs[entity][reflect.TypeOf(component)] = component
	}
	ecs.mutex.Unlock()
	for _, listener := range ecs.listeners {
		listener.GetTargetComponents()
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(entity)
		}
	}
	return entity
}

func (entity ECSEntity) AddComponent(component Component) {
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

func (entity ECSEntity) RemoveComponent(nilComponent interface{}) bool {
	componentType := reflect.TypeOf(nilComponent)
	if entity.HasComponent(componentType) {
		entity.ecs.mutex.RLock()
		component := entity.ecs.ecs[entity][componentType]
		entity.ecs.mutex.RUnlock()
		for _, listener := range entity.ecs.listeners {
			if listener.ShouldListenForAllComponents() {
				listener.OnAddComponent(entity, component)
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
	if &entity != nil && ecs.ecs[entity] != nil {
		for _, listener := range ecs.listeners {
			for _, listenEntity := range listener.GetEntities() {
				if entity == listenEntity {
					listener.OnRemoveEntity(entity)
				}
				break
			}
		}
		ecs.mutex.Lock()
		delete(ecs.ecs, entity)
		ecs.mutex.Unlock()
		ecs.mutex.RUnlock()
		return true
	}
	ecs.mutex.RUnlock()
	return false
}

func (entity ECSEntity) GetComponent(nilComponent interface{}) interface{} {
	entity.ecs.mutex.RLock()
	if entity.ecs.ecs[entity] == nil {
		entity.ecs.mutex.RUnlock()
		return nil
	}
	component := entity.ecs.ecs[entity][reflect.TypeOf(nilComponent)]
	entity.ecs.mutex.RUnlock()
	return component
}

func (entity ECSEntity) HasComponent(targets ...reflect.Type) bool {
	entity.ecs.mutex.RLock()
	if entity.ecs.ecs[entity] == nil {
		return false
	}
	for _, target := range targets {
		if entity.ecs.ecs[entity][target] == nil {
			return false
		}
	}
	entity.ecs.mutex.RUnlock()
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
	for i := 0; i < len(ecs.listeners); i++ {
		if ecs.listeners[i] == listener {
			ecs.listeners[i] = ecs.listeners[len(ecs.listeners)-1]
			ecs.listeners = ecs.listeners[:len(ecs.listeners)-1]
		}
	}
}
