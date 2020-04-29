package limitengine

import (
	"reflect"
	"sync"
	"time"
)

var (
	entityIndex = uint32(0)
	ecs         = make(map[ECSEntity]map[reflect.Type]interface{})
	ecsMutex    = sync.RWMutex{}
	listeners   = []ECSListener{}
)

func initECS() {
	go func() {
		currentTime := time.Now().UnixNano()
		for Running() {
			lastTime := currentTime
			currentTime = time.Now().UnixNano()
			delta := float32(currentTime-lastTime) / 1000000000.0
			UpdateSystems(delta)
		}
	}()
	// TODO: Sort out ECS threading.
}

type ECSEntity uint32

type ECSListener interface {
	OnAddEntity(entity ECSEntity)
	OnAddComponent(entity ECSEntity, component interface{})
	OnRemoveComponent(entity ECSEntity, component interface{})
	OnRemoveEntity(entity ECSEntity)
	GetTargetComponents() []reflect.Type
	GetEntities() []ECSEntity
}

func NewEntity(components ...interface{}) ECSEntity {
	entity := ECSEntity(entityIndex)
	ecsMutex.Lock()
	ecs[entity] = make(map[reflect.Type]interface{})
	entityIndex++
	for _, component := range components {
		ecs[entity][reflect.TypeOf(component)] = component
	}
	ecsMutex.Unlock()
	for _, listener := range listeners {
		listener.GetTargetComponents()
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(entity)
		}
	}
	return entity
}

func (entity ECSEntity) AddComponent(component interface{}) {
	componentType := reflect.TypeOf(component)
	ecsMutex.Lock()
	ecs[entity][componentType] = component
	ecsMutex.Unlock()
	for _, listener := range listeners {
		for _, target := range listener.GetTargetComponents() {
			if target == componentType {
				listener.OnAddComponent(entity, component)
				break
			}
		}
	}
}

func (entity ECSEntity) RemoveComponent(nilComponent interface{}) bool {
	componentType := reflect.TypeOf(nilComponent)
	if entity.HasComponent(componentType) {
		ecsMutex.RLock()
		component := ecs[entity][componentType]
		ecsMutex.RUnlock()
		for _, listener := range listeners {
			for _, target := range listener.GetTargetComponents() {
				if target == componentType {
					listener.OnRemoveComponent(entity, component)
					break
				}
			}
		}
		ecsMutex.Lock()
		delete(ecs[entity], componentType)
		ecsMutex.Unlock()
		return true
	}
	return false
}

func RemoveEntity(entity ECSEntity) bool {
	ecsMutex.RLock()
	if &entity != nil && ecs[entity] != nil {
		for _, listener := range listeners {
			for _, listenEntity := range listener.GetEntities() {
				if entity == listenEntity {
					listener.OnRemoveEntity(entity)
				}
				break
			}
		}
		ecsMutex.Lock()
		delete(ecs, entity)
		ecsMutex.Unlock()
		ecsMutex.RUnlock()
		return true
	}
	ecsMutex.RUnlock()
	return false
}

func (entity ECSEntity) GetComponent(nilComponent interface{}) interface{} {
	ecsMutex.RLock()
	if ecs[entity] == nil {
		ecsMutex.RUnlock()
		return nil
	}
	component := ecs[entity][reflect.TypeOf(nilComponent)]
	ecsMutex.RUnlock()
	return component
}

func (entity ECSEntity) HasComponent(targets ...reflect.Type) bool {
	ecsMutex.RLock()
	if ecs[entity] == nil {
		return false
	}
	for _, target := range targets {
		if ecs[entity][target] == nil {
			return false
		}
	}
	ecsMutex.RUnlock()
	return true
}

func AddECSListener(listener ECSListener) {
	for entity := range ecs {
		if entity.HasComponent(listener.GetTargetComponents()...) {
			listener.OnAddEntity(entity)
		}
	}
	listeners = append(listeners, listener)
}

func RemoveECSListener(listener ECSListener) {
	for i := 0; i < len(listeners); i++ {
		if listeners[i] == listener {
			listeners[i] = listeners[len(listeners)-1]
			listeners = listeners[:len(listeners)-1]
		}
	}
}
