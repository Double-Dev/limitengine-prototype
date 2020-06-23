package gfx

import (
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type GFXComponent interface {
	Renderables() []*Renderable
}

type GFXListener struct {
	entities map[limitengine.ECSEntity]GFXComponent
	target   []reflect.Type
}

func NewGFXListener(nilGFXComponent GFXComponent) *GFXListener {
	return &GFXListener{
		entities: make(map[limitengine.ECSEntity]GFXComponent),
		target:   []reflect.Type{reflect.TypeOf(nilGFXComponent)},
	}
}

func (gfxListener *GFXListener) OnAddEntity(entity limitengine.ECSEntity) {
	gfx := entity.GetComponentOfType(gfxListener.target[0]).(GFXComponent)
	gfxListener.entities[entity] = gfx
	for _, renderable := range gfx.Renderables() {
		AddRenderable(renderable)
	}
}

func (gfxListener *GFXListener) OnAddComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	gfx := component.(GFXComponent)
	gfxListener.entities[entity] = gfx
	for _, renderable := range gfx.Renderables() {
		AddRenderable(renderable)
	}
}

func (gfxListener *GFXListener) OnRemoveComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	gfx := gfxListener.entities[entity]
	for _, renderable := range gfx.Renderables() {
		RemoveRenderable(renderable)
	}
	delete(gfxListener.entities, entity)
}

func (gfxListener *GFXListener) OnRemoveEntity(entity limitengine.ECSEntity) {
	gfx := gfxListener.entities[entity]
	for _, renderable := range gfx.Renderables() {
		RemoveRenderable(renderable)
	}
	delete(gfxListener.entities, entity)
}

func (gfxListener *GFXListener) GetTargetComponents() []reflect.Type  { return gfxListener.target }
func (gfxListener *GFXListener) GetEntities() []limitengine.ECSEntity { return nil }
func (gfxListener *GFXListener) ShouldListenForAllComponents() bool   { return false }

// Render component, system, and listener for generic renders:

type RenderComponent struct {
	Renderable  *Renderable
	renderables []*Renderable
}

func NewRenderComponent(camera *Camera, shader *Shader, material Material, mesh *Mesh, instance *Instance) *RenderComponent {
	renderable := &Renderable{
		Camera:   camera,
		Shader:   shader,
		Material: material,
		Mesh:     mesh,
		Instance: instance,
	}
	return &RenderComponent{
		renderable,
		[]*Renderable{renderable},
	}
}

func (renderComponent *RenderComponent) Renderables() []*Renderable {
	return renderComponent.renderables
}

func NewRenderSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			transform := components[1].(*gmath.TransformComponent)

			transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

			render := components[0].(*RenderComponent)
			render.Renderable.Instance.SetTransform(transformMat)
		}
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewRenderListener() *GFXListener { return NewGFXListener((*RenderComponent)(nil)) }
