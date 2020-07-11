package gfx

import (
	"reflect"
	"sync"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type GFXComponent interface {
	Renderables() []*Renderable
}

type GFXListener struct {
	target     []reflect.Type
	entities   []limitengine.ECSEntity
	components map[limitengine.ECSEntity]GFXComponent
	mutex      sync.RWMutex
}

func NewGFXListener(nilGFXComponent GFXComponent) *GFXListener {
	return &GFXListener{
		target:     []reflect.Type{reflect.TypeOf(nilGFXComponent)},
		components: make(map[limitengine.ECSEntity]GFXComponent),
	}
}

func (gfxListener *GFXListener) OnAddEntity(entity limitengine.ECSEntity) {
	gfxListener.entities = append(gfxListener.entities, entity)
	gfx := entity.GetComponentOfType(gfxListener.target[0]).(GFXComponent)
	gfxListener.mutex.Lock()
	gfxListener.components[entity] = gfx
	gfxListener.mutex.Unlock()
	for _, renderable := range gfx.Renderables() {
		AddRenderable(renderable)
	}
}

func (gfxListener *GFXListener) OnAddComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	gfxListener.entities = append(gfxListener.entities, entity)
	gfx := component.(GFXComponent)
	gfxListener.mutex.Lock()
	gfxListener.components[entity] = gfx
	gfxListener.mutex.Unlock()
	for _, renderable := range gfx.Renderables() {
		AddRenderable(renderable)
	}
}

func (gfxListener *GFXListener) OnRemoveComponent(entity limitengine.ECSEntity, component limitengine.ECSComponent) {
	gfxListener.mutex.RLock()
	gfx := gfxListener.components[entity]
	gfxListener.mutex.RUnlock()
	for _, renderable := range gfx.Renderables() {
		RemoveRenderable(renderable)
	}
	delete(gfxListener.components, entity)
	for i, potentialEntity := range gfxListener.entities {
		if potentialEntity == entity {
			gfxListener.entities[i] = gfxListener.entities[len(gfxListener.entities)-1]
			gfxListener.entities = gfxListener.entities[:len(gfxListener.entities)-1]
			break
		}
	}
}

func (gfxListener *GFXListener) OnRemoveEntity(entity limitengine.ECSEntity) {
	gfxListener.mutex.RLock()
	gfx := gfxListener.components[entity]
	gfxListener.mutex.RUnlock()
	for _, renderable := range gfx.Renderables() {
		RemoveRenderable(renderable)
	}
	delete(gfxListener.components, entity)
	for i, potentialEntity := range gfxListener.entities {
		if potentialEntity == entity {
			gfxListener.entities[i] = gfxListener.entities[len(gfxListener.entities)-1]
			gfxListener.entities = gfxListener.entities[:len(gfxListener.entities)-1]
			break
		}
	}
}

func (gfxListener *GFXListener) GetTargetComponents() []reflect.Type  { return gfxListener.target }
func (gfxListener *GFXListener) GetEntities() []limitengine.ECSEntity { return gfxListener.entities }
func (gfxListener *GFXListener) ShouldListenForAllComponents() bool   { return false }

// Render component, system, and listener for generic renders:

type RenderComponent struct {
	Renderable  *Renderable
	renderables []*Renderable
}

func NewRenderComponent(layer int32, camera *Camera, shader Shader, material Material, mesh *Mesh, instance *Instance) *RenderComponent {
	renderable := &Renderable{
		Layer:    layer,
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
			if transform.IsAwake() {
				transformMat := gmath.NewTransformMatrix(transform.Position, transform.Rotation, transform.Scale)

				render := components[0].(*RenderComponent)
				render.Renderable.Instance.SetTransform(transformMat)
			}
		}
	}, (*RenderComponent)(nil), (*gmath.TransformComponent)(nil))
}

func NewRenderListener() *GFXListener { return NewGFXListener((*RenderComponent)(nil)) }
