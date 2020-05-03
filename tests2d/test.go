package main

import (
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/ui"

	"github.com/pkg/profile"
)

type PlatformControlComponent struct {
	WalkControl                 *ui.InputControl
	WalkSpeed, WalkAcceleration float32
	JumpControl                 *ui.InputControl
	JumpSpeed, JumpAcceleration float32
	DashControl                 *ui.InputControl
	DashSpeed, DashAcceleration float32
}

func NewPlatformControlSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		// for _, entity := range entities {
		// control := entity.GetComponent((*PlatformControlComponent)(nil)).(*PlatformControlComponent)
		// motion := entity.GetComponent((*utils.MotionComponent)(nil)).(*utils.MotionComponent)
		// transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)
		// }
	}, (*PlatformControlComponent)(nil), (*gmath.MotionComponent)(nil), (*gmath.TransformComponent)(nil))
}

func main() {
	// Profile
	defer profile.Start().Stop()

	// Assets
	shader := gfx.CreateShader(gio.LoadAsString("testshader.lesl"))
	texture := gfx.CreateTexture(gio.LoadPNG("testsprite.png"))
	material := gfx.CreateTextureMaterial(texture)
	camera := gfx.CreateCamera2D()

	// Controls
	xAxis := ui.InputControl{}
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyA}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeft}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyD}, 1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyRight}, 1.0)
	yAxis := ui.InputControl{}
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyUp}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyDown}, -1.0)
	dash := ui.InputControl{}
	dash.AddTrigger(ui.InputEvent{Key: ui.KeySpace}, 1.0)
	dash.AddTrigger(ui.InputEvent{Key: ui.KeyRightShift}, 1.0)
	dash.AddTrigger(ui.InputEvent{Key: ui.KeyLeftShift}, 1.0)

	// Entities
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(-0.5, 0.0, -0.5),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.1, 0.1, 1.0),
		},
		&gmath.MotionComponent{
			Velocity:        gmath.NewVector3(0.1, 0.0, 0.0),
			Acceleration:    gmath.NewZeroVector3(),
			AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -0.1, 0.0), gmath.NewVector3(0.1, 0.1, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)

	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.5, 0.0, -0.5),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.1, 0.1, 1.0),
		},
		&gmath.MotionComponent{
			Velocity:        gmath.NewVector3(-0.1, 0.0, 0.0),
			Acceleration:    gmath.NewZeroVector3(),
			AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -0.1, 0.0), gmath.NewVector3(0.1, 0.1, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)

	// Stress test...
	for i := 0; i < 2000; i++ {
		limitengine.NewEntity(
			&gmath.TransformComponent{
				Position: gmath.NewVector3(-0.5, -0.5, -0.5),
				Rotation: gmath.NewIdentityQuaternion(),
				Scale:    gmath.NewVector3(0.1, 0.1, 1.0),
			},
			&gmath.MotionComponent{
				Velocity:        gmath.NewVector3(0.1, 0.1, 0.0),
				Acceleration:    gmath.NewZeroVector3(),
				AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
				AngAcceleration: gmath.NewIdentityQuaternion(),
			},
			&interaction.ColliderComponent{
				AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -0.1, 0.0), gmath.NewVector3(0.1, 0.1, 0.0)),
			},
			&gfx.RenderComponent{
				Camera:   camera,
				Shader:   shader,
				Material: material,
				Mesh:     &gfx.Mesh{},
				Instance: gfx.NewInstance(),
			},
		)
	}

	// Systems
	interactionWorld := interaction.NewWorld(20.0)

	interaction := TestInteraction{
		test: "Hello",
	}

	interactionWorld.AddInteraction(interaction)

	limitengine.AddSystem(gfx.NewRenderSystem())
	limitengine.AddSystem(gmath.NewMotionSystem(1.0))
	limitengine.AddSystem(NewPlatformControlSystem())
	limitengine.AddECSListener(interactionWorld)

	// Launch!
	limitengine.Launch()
}

type TestInteraction struct {
	test string
}

func (test TestInteraction) Interact(delta float32, interactor, interactee limitengine.ECSEntity) {
	interactor.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent).Velocity.MulSc(-1.0)
}

func (test TestInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
		reflect.TypeOf((*interaction.ColliderComponent)(nil)),
	}
}

func (test TestInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*gmath.TransformComponent)(nil)),
		reflect.TypeOf((*interaction.ColliderComponent)(nil)),
	}
}
