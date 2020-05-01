package main

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/ui"
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

	// Systems
	interactionWorld := interaction.NewWorld()

	limitengine.AddSystem(gfx.NewRenderSystem())
	limitengine.AddSystem(gmath.NewMotionSystem(1.0))
	limitengine.AddSystem(NewPlatformControlSystem())
	limitengine.AddECSListener(interactionWorld)

	go func() {
		for limitengine.Running() {
			interactionWorld.ProcessInteractions(1.0)
		}
	}()

	// Launch!
	limitengine.Launch()
}
