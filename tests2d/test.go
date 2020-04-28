package main

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
	"github.com/double-dev/limitengine/utils"
)

type PlatformControlComponent struct {
	WalkControl                 *ui.InputControl
	WalkSpeed, WalkAcceleration float32
	JumpControl                 *ui.InputControl
	JumpSpeed, JumpAcceleration float32
	DashControl                 *ui.InputControl
	DashSpeed, DashAcceleration float32
}

func NewPlatformControlSystem() *ecs.ECSSystem {
	return ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		// for _, entity := range entities {
		// control := entity.GetComponent((*PlatformControlComponent)(nil)).(*PlatformControlComponent)
		// motion := entity.GetComponent((*utils.MotionComponent)(nil)).(*utils.MotionComponent)
		// transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)
		// }
	}, (*PlatformControlComponent)(nil), (*utils.MotionComponent)(nil), (*utils.TransformComponent)(nil))
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
	ecs.NewEntity(
		&utils.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.0, -0.5),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.1, 0.1, 1.0),
		},
		&utils.MotionComponent{
			Velocity:        gmath.NewZeroVector3(),
			Acceleration:    gmath.NewZeroVector3(),
			AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&utils.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     &gfx.Mesh{},
			Instance: gfx.NewInstance(),
		},
	)

	// Systems
	ecs.AddSystem(utils.NewRenderSystem())
	ecs.AddSystem(NewPlatformControlSystem())

	// Launch!
	limitengine.Launch()
}
