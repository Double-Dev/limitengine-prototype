package gfx

import (
	"reflect"

	"github.com/double-dev/limitengine"
)

var (
	targets = []reflect.Type{
		reflect.TypeOf((*RenderComponent)(nil)),
	}
)

type RenderComponent struct {
	Camera   *Camera
	Shader   *Shader
	Material Material
	Mesh     *Mesh
	Instance *Instance
}

func (renderComponent *RenderComponent) Delete() {
	// TODO: Finish cleaning up render component.
	DeleteMesh(renderComponent.Mesh)
}

type GFXListener struct {
	entities map[limitengine.ECSEntity]RenderComponent
}

func NewGFXListener() GFXListener {
	return GFXListener{
		entities: make(map[limitengine.ECSEntity]RenderComponent),
	}
}

func (gfxListener GFXListener) OnAddEntity(entity limitengine.ECSEntity) {
	render := entity.GetComponent((*RenderComponent)(nil)).(*RenderComponent)
	gfxListener.entities[entity] = *render
	AddRenderable(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
}

func (gfxListener GFXListener) OnAddComponent(entity limitengine.ECSEntity, component limitengine.Component) {
	render := component.(*RenderComponent)
	gfxListener.entities[entity] = *render
	AddRenderable(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
}

func (gfxListener GFXListener) OnRemoveComponent(entity limitengine.ECSEntity, component limitengine.Component) {
	render := gfxListener.entities[entity]
	RemoveRenderable(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
	delete(gfxListener.entities, entity)
}

func (gfxListener GFXListener) OnRemoveEntity(entity limitengine.ECSEntity) {
	render := gfxListener.entities[entity]
	RemoveRenderable(render.Camera, render.Shader, render.Material, render.Mesh, render.Instance)
	delete(gfxListener.entities, entity)
}

func (gfxListener GFXListener) GetTargetComponents() []reflect.Type  { return targets }
func (gfxListener GFXListener) GetEntities() []limitengine.ECSEntity { return nil }
func (gfxListener GFXListener) ShouldListenForAllComponents() bool   { return false }
