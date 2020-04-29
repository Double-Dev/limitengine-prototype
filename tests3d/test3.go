package tests3d

import (
	"math/rand"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
)

func main() {
	mesh := gfx.CreateMesh(gio.LoadOBJ("monkey.obj"))
	// mesh := &gfx.Mesh{}
	shader := gfx.CreateShader(gio.LoadAsString("testShader.lesl"))
	texture := gfx.CreateTexture(gio.LoadPNG("lamp.png"))
	material := gfx.CreateTextureMaterial(texture)

	xAxis := ui.InputControl{}
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyA}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeft}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyD}, 1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyRight}, 1.0)
	yAxis := ui.InputControl{}
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyQ}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyE}, 1.0)
	zAxis := ui.InputControl{}
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, -1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyUp}, -1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, 1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyDown}, 1.0)

	boost := ui.InputControl{}
	boost.AddTrigger(ui.InputEvent{Key: ui.KeySpace}, 1.0)
	boost.AddTrigger(ui.InputEvent{Key: ui.KeyRightShift}, 0.05)
	boost.AddTrigger(ui.InputEvent{Key: ui.KeyLeftShift}, 0.05)

	camera := gfx.CreateCamera3D(0.001, 1000.0, 60.0)
	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 1.0, -10.0),
			Rotation: gmath.NewQuaternion(0.0, 0.0, 0.0, 1.0),
			Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
		},
		&gmath.MotionComponent{
			Velocity:        gmath.NewZeroVector3(),
			Acceleration:    gmath.NewZeroVector3(),
			AngVelocity:     gmath.NewIdentityQuaternion(),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&MotionControlComponent{
			Axis:  []*ui.InputControl{&xAxis, &yAxis, &zAxis, &boost},
			Speed: 600.0,
		},
		&gfx.CameraComponent{
			Camera:      camera,
			PosOffset:   gmath.NewVector3(0.0, 0.0, 15.0),
			RotOffset:   gmath.NewIdentityQuaternion(),
			ScaleOffset: gmath.NewVector3(1.0, 1.0, 1.0),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
	)

	for i := 0; i < 2000; i++ {
		// randAxis := gmath.NewVector3(rand.Float32()-0.5, rand.Float32()-0.5, rand.Float32()-0.5).Normalize()
		limitengine.NewEntity(
			&gmath.TransformComponent{
				Position: gmath.NewVector3(rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-750.0),
				Rotation: gmath.NewQuaternion(rand.Float32()*gmath.Pi, rand.Float32(), rand.Float32(), rand.Float32()),
				Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
			},
			// &gmath.MotionComponent{
			// 	Velocity:        gmath.NewZeroVector3(),
			// 	Acceleration:    gmath.NewZeroVector3(),
			// 	AngVelocity:     gmath.NewIdentityQuaternion(),
			// 	AngAcceleration: gmath.NewQuaternionV(rand.Float32(), randAxis),
			// },
			&gfx.RenderComponent{
				Camera:   camera,
				Shader:   shader,
				Material: material,
				Mesh:     mesh,
				Instance: gfx.NewInstance(),
			},
		)
	}

	// limitengine.AddSystem(utils.NewMotionControlSystem())
	limitengine.AddSystem(limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			control := entity.GetComponent((*MotionControlComponent)(nil)).(*MotionControlComponent)
			motion := entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)
			transform := entity.GetComponent((*gmath.TransformComponent)(nil)).(*gmath.TransformComponent)

			direction := gmath.NewVector3(control.Axis[2].Amount(), control.Axis[0].Amount(), control.Axis[1].Amount())

			if direction.LenSq() > 0.5 {
				motion.AngAcceleration = gmath.NewQuaternion(gmath.Pi, direction[0], direction[1], direction[2])
			} else {
				motion.AngAcceleration = gmath.NewIdentityQuaternion()
			}

			var speed float32
			if control.Axis[3].Amount() == 0.0 {
				speed = -control.Speed * 0.2
			} else {
				speed = -control.Speed * control.Axis[3].Amount()
			}

			motion.Acceleration = transform.Rotation.RotateV(gmath.NewVector3(0.0, 0.0, speed))
		}
	}, (*MotionControlComponent)(nil), (*gmath.MotionComponent)(nil), (*gmath.TransformComponent)(nil)))
	limitengine.AddSystem(gmath.NewMotionSystem(0.95))
	limitengine.AddSystem(gfx.NewCameraMotionSystem())
	limitengine.AddSystem(gfx.NewRenderSystem())

	limitengine.Launch()
}
