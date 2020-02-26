package ecs

import "reflect"

var (
	systems = []ecsSystem{}
)

type ecsSystem struct {
	targetComponents []reflect.Type
	entities         []ECSEntity
	update           func(delta float32, entity ECSEntity)
}

func NewSystem(update func(delta float32, entity ECSEntity), nilTargetComponents ...interface{}) *ecsSystem {
	system := ecsSystem{
		targetComponents: []reflect.Type{},
		entities:         []ECSEntity{},
		update:           update,
	}
	for _, nilTargetComponent := range nilTargetComponents {
		system.targetComponents = append(system.targetComponents, reflect.TypeOf(nilTargetComponent))
	}
	AddECSListener(&system)
	return &system
}

func UpdateSystems(delta float32) {
	for _, system := range systems {
		for _, entity := range system.entities {
			system.update(delta, entity)
		}
	}
}

func (system *ecsSystem) OnAddEntity(entity *ECSEntity) {
	system.entities = append(system.entities, *entity)
}

func (system *ecsSystem) OnAddComponent(entity *ECSEntity, component interface{}) {
	if entity.HasComponent(system.GetTargetComponents()...) {
		system.entities = append(system.entities, *entity)
	}
}

func (system *ecsSystem) OnRemoveComponent(entity *ECSEntity, component interface{}) {
	for i, target := range system.entities {
		if target == *entity {
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
		}
	}
}

func (system *ecsSystem) OnRemoveEntity(entity *ECSEntity) {
	for i, target := range system.entities {
		if target == *entity {
			system.entities[i] = system.entities[len(system.entities)-1]
			system.entities = system.entities[:len(system.entities)-1]
		}
	}
}

func (system *ecsSystem) GetTargetComponents() []reflect.Type {
	return system.targetComponents
}

func (system *ecsSystem) GetEntities() []ECSEntity {
	return system.entities
}
