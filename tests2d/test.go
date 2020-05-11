package main

import (
	"image"
	"math/rand"
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

	// Setup Window
	limitengine.AppView().SetTitle("2D Tests!")
	limitengine.AppView().SetPosition(0, 50)
	limitengine.AppView().SetAspectRatio(3, 2)
	limitengine.AppView().SetIcons([]image.Image{gio.LoadIcon("Test.png")})

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
		&ControlComponent{
			randomTestVar: "kasdfkjsdahflksahndnhslk",
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

	for i := 0; i < 2; i++ {
		limitengine.NewEntity(
			&gmath.TransformComponent{
				Position: gmath.NewVector3(rand.Float32()*2.0-1.0, rand.Float32()*2.0-1.0, -0.1),
				Rotation: gmath.NewIdentityQuaternion(),
				Scale:    gmath.NewVector3(0.2, 0.2, 1.0),
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

	// Left Wall
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(-1.5, 0.0, -0.45),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.1, 1.0, 1.0),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -1.0, 0.0), gmath.NewVector3(0.1, 1.0, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)
	// Right Wall
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(1.5, 0.0, -0.45),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.1, 1.0, 1.0),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -1.0, 0.0), gmath.NewVector3(0.1, 1.0, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)
	// Top Wall
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 1.0, -0.4),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(1.5, 0.1, 1.0),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-1.5, -0.1, 0.0), gmath.NewVector3(1.5, 0.1, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)
	// Bottom Wall
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, -1.0, -0.4),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(1.5, 0.1, 1.0),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-1.5, -0.1, 0.0), gmath.NewVector3(1.5, 0.1, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)

	// Systems
	interactionWorld := interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)

	myInteraction := TestInteraction{
		test: "Hello",
	}

	interactionWorld.AddInteraction(myInteraction)

	limitengine.AddSystem(gfx.NewRenderSystem())
	limitengine.AddSystem(gmath.NewMotionSystem(0.99))
	limitengine.AddSystem(NewPlatformControlSystem())
	limitengine.AddSystem(limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			motion := entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)

			speed := float32(0.5)
			if xAxis.Amount() > 0.01 {
				motion.Acceleration[0] = speed
			} else if xAxis.Amount() < -0.01 {
				motion.Acceleration[0] = -speed
			} else {
				motion.Acceleration[0] = 0.0
			}
			if yAxis.Amount() > 0.01 {
				motion.Acceleration[1] = speed * 30.0
			} else {
				motion.Acceleration[1] = -2.45
			}
		}
	}, (*ControlComponent)(nil), (*gmath.MotionComponent)(nil), (*gmath.TransformComponent)(nil)))
	limitengine.AddECSListener(interactionWorld)

	// Launch!
	limitengine.Launch()
}

type ControlComponent struct {
	randomTestVar string
}

type TestInteraction struct {
	test string
}

func (test TestInteraction) StartInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3, penetration float32) {
}

func (test TestInteraction) EndInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3) {
}

func (test TestInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*gmath.MotionComponent)(nil)),
	}
}

func (test TestInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{}
}
