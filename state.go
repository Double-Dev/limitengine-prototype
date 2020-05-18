package limitengine

type State struct {
	ecs     *ECS
	systems []*ECSSystem
}

func NewState() *State {
	return &State{
		ecs: NewECS(),
	}
}

func (state *State) NewEntity(components ...Component) ECSEntity {
	return state.ecs.NewEntity(components...)
}

func (state *State) RemoveEntity(entity ECSEntity) bool {
	return state.ecs.RemoveEntity(entity)
}

func (state *State) Update(delta float32) {
	for _, system := range state.systems {
		system.Update(delta)
	}
}

func (state *State) AddListener(listener ECSListener) {
	state.ecs.AddECSListener(listener)
}

func (state *State) RemoveListener(listener ECSListener) {
	state.ecs.RemoveECSListener(listener)
}

func (state *State) AddSystem(system *ECSSystem) {
	state.ecs.AddECSListener(system)
	state.systems = append(state.systems, system)
}

func (state *State) RemoveSystem(system *ECSSystem) {
	state.ecs.RemoveECSListener(system)
	// TODO: Replace with quicker search algorithm?
	for i, stateSystem := range state.systems {
		if stateSystem == system {
			state.systems[i] = state.systems[len(state.systems)-1]
			state.systems[len(state.systems)-1] = nil
			state.systems = state.systems[:len(state.systems)-1]
			break
		}
	}
}
