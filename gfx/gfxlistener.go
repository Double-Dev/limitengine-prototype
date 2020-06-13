package gfx

import (
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
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

func (renderComponent *RenderComponent) Delete() {}

func NewRenderSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.Component) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := components[0].(*RenderComponent)
			render.Instance.SetTransform(transformMat)
		}

		Sweep()
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
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
