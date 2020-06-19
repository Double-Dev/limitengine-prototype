package tests3d

import (
	"math/rand"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
	"github.com/pkg/profile"
)

func main() {
	// Profile
	defer profile.Start().Stop()

	// Setup Window
	limitengine.AppView().SetTitle("3D Tests!")
	limitengine.AppView().SetPosition(100, 100)
	limitengine.AppView().SetAspectRatio(3, 2)

	// Creating State
	state := limitengine.NewState(nil)

	// Assets
	shader := gfx.NewShader(gio.LoadAsString("testShader.lesl"))
	texture := gfx.NewTexture(gio.LoadPNG("lamp.png"))
	material := gfx.NewTextureMaterial(texture)
	mesh := gfx.NewMesh(gio.LoadOBJ("monkey.obj"))

	camColor := gfx.NewRenderbuffer(true)
	camDepth := gfx.NewRenderbuffer(true)
	camera := gfx.NewCamera3D(camColor, camDepth, 0.001, 1000.0, 60.0)
	camera.SetClearColor(0.0, 0.1, 0.25, 1.0)
	camera.AddBlitCamera(gfx.DefaultCamera())

	// Controls
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

	state.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.0, -10.0),
			Rotation: gmath.NewIdentityQuaternion(),
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
		&CameraComponent{
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
		state.NewEntity(
			&gmath.TransformComponent{
				Position: gmath.NewVector3(rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-750.0),
				Rotation: gmath.NewQuaternion(rand.Float32()*gmath.Pi, rand.Float32(), rand.Float32(), rand.Float32()),
				Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
			},
			&gfx.RenderComponent{
				Camera:   camera,
				Shader:   shader,
				Material: material,
				Mesh:     mesh,
				Instance: gfx.NewInstance(),
			},
		)
	}

	// System & Listeners
	gfxListener := gfx.NewGFXListener()
	state.AddListener(gfxListener)

	state.AddSystem(limitengine.NewSystem(func(delta float32, entities [][]limitengine.Component) {
		for _, components := range entities {
			control := components[0].(*MotionControlComponent)
			motion := components[1].(*gmath.MotionComponent)
			transform := components[2].(*gmath.TransformComponent)

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
	state.AddSystem(gmath.NewMotionSystem(0.95))
	state.AddSystem(gfx.NewRenderSystem())
	state.AddSystem(NewCameraMotionSystem())

	limitengine.Launch(state)
}

type CameraComponent struct {
	Camera      *gfx.Camera
	PosOffset   gmath.Vector3
	RotOffset   gmath.Quaternion
	ScaleOffset gmath.Vector3
}

func (cameraComponent *CameraComponent) Delete() {}

func NewCameraMotionSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.Component) {
		for _, components := range entities {
			camera := components[0].(*CameraComponent)
			transform := components[1].(*gmath.TransformComponent)

			transformViewMat := gmath.NewViewMatrix(
				transform.Position,
				transform.Rotation,
				transform.Scale,
			)
			offsetViewMat := gmath.NewViewMatrix(
				camera.PosOffset,
				camera.RotOffset,
				camera.ScaleOffset,
			)

			viewMat := offsetViewMat.MulM(transformViewMat)

			camera.Camera.SetViewMat(viewMat)
		}
	}, (*CameraComponent)(nil), (*gmath.TransformComponent)(nil))
}

type MotionControlComponent struct {
	Axis  []*ui.InputControl
	Speed float32
}

func (motionControlComponent *MotionControlComponent) Delete() {}
