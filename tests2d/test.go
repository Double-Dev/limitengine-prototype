package main

import (
	"image"
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/ui"

	"github.com/pkg/profile"
)

var (
	camera   *gfx.Camera
	shader   *gfx.Shader
	material *gfx.Material
	mesh     *gfx.Mesh
)

func main() {
	// Profile
	defer profile.Start().Stop()

	// Setup Window
	limitengine.AppView().SetTitle("2D Tests!")
	limitengine.AppView().SetPosition(100, 100)
	limitengine.AppView().SetAspectRatio(3, 2)
	limitengine.AppView().SetIcons([]image.Image{gio.LoadIcon("Test.png")})

	// Creating State
	state := limitengine.NewState()

	// Assets
	shader = gfx.CreateShader(gio.LoadAsString("testshader.lesl"))

	// Player
	playerTexture := gfx.CreateTexture(gio.LoadPNG("testsprite.png"))
	playerMaterial := gfx.CreateTextureMaterial(playerTexture)

	// Walls
	material = gfx.CreateMaterial(gmath.NewVector4(0.4, 0.4, 0.45, 1.0))

	mesh = gfx.SpriteMesh()

	// TODO: Fix issue where having a texture as a framebuffer depth attachment doesn't work.
	cam1Color := gfx.CreateRenderbuffer(true)
	cam1Depth := gfx.CreateRenderbuffer(true)
	camera = gfx.CreateCamera2D(cam1Color, cam1Depth)
	camera.SetClearColor(0.0, 0.25, 0.25, 1.0)

	cam2Color := gfx.CreateEmptyTexture()
	cam2Depth := gfx.CreateEmptyTexture()
	camera2 := gfx.CreateCamera(cam2Color, cam2Depth)

	camera.AddBlitCamera(camera2)

	fboShader := gfx.CreateShader(gio.LoadAsString("fboshader.lesl"))
	cam2Mat := gfx.CreateTextureMaterial(cam2Color)

	pos := gfx.NewInstance()
	pos.SetTransform(gmath.NewTransformMatrix(gmath.NewVector3(0.0, 0.0, 0.5), gmath.NewIdentityQuaternion(), gmath.NewVector3(1.0, 1.0, 1.0)))
	gfx.AddRenderable(gfx.DefaultCamera(), fboShader, cam2Mat, mesh, pos)

	// Controls
	xAxis := &ui.InputControl{}
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyA}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeft}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyD}, 1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyRight}, 1.0)
	yAxis := &ui.InputControl{}
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyUp}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyDown}, -1.0)

	// Entities
	state.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(-0.5, 0.0, -0.3),
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
			AABB: gmath.NewAABB(gmath.NewVector3(-0.075, -0.075, 0.0), gmath.NewVector3(0.075, 0.075, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: playerMaterial,
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
		&ControlComponent{
			XAxis: xAxis,
			YAxis: yAxis,
		},
	)

	// state.NewEntity(
	// 	&gmath.TransformComponent{
	// 		Position: gmath.NewVector3(-1.05, -0.25, -0.6),
	// 		Rotation: gmath.NewIdentityQuaternion(),
	// 		Scale:    gmath.NewVector3(0.35, 0.65, 1.0),
	// 	},
	// 	&gfx.RenderComponent{
	// 		Camera:   camera,
	// 		Shader:   shader,
	// 		Material: material,
	// 		Mesh:     mesh,
	// 		Instance: gfx.NewInstance(),
	// 	},
	// )

	state.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.5, 0.0, -0.5),
			Rotation: gmath.NewQuaternion(gmath.Pi/8.0, 0.0, 0.0, 1.0),
			Scale:    gmath.NewVector3(0.25, 0.05, 1.0),
		},
		&gmath.MotionComponent{
			Velocity:        gmath.NewVector3(-0.1, 0.0, 0.0),
			Acceleration:    gmath.NewZeroVector3(),
			AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.25, -0.05, 0.0), gmath.NewVector3(0.25, 0.05, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
	)

	// Left Wall
	CreateStaticEntity(state, gmath.NewVector3(-1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Right Wall
	CreateStaticEntity(state, gmath.NewVector3(1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Top Wall
	CreateStaticEntity(state, gmath.NewVector3(0.0, 1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))
	// Bottom Wall
	CreateStaticEntity(state, gmath.NewVector3(0.0, -1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))

	CreateStaticEntity(state, gmath.NewVector3(-0.6, -0.3, -0.4), gmath.NewVector3(0.1, 0.7, 1.0))

	// Systems
	interactionWorld := interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)

	myInteraction := TestInteraction{
		test: "Hello",
	}

	interactionWorld.AddInteraction(myInteraction)

	gfxListener := gfx.NewGFXListener()
	state.AddListener(gfxListener)

	state.AddSystem(gfx.NewRenderSystem())
	state.AddSystem(gmath.NewMotionSystem(1.0))
	state.AddSystem(limitengine.NewSystem(func(delta float32, entities [][]limitengine.Component) {
		for _, components := range entities {
			control := components[0].(*ControlComponent)
			motion := components[1].(*gmath.MotionComponent)

			speed := float32(3.0)
			maxSpeed := float32(1.0)
			if control.XAxis.Amount() > 0.01 {
				motion.Acceleration[0] = speed
			} else if control.XAxis.Amount() < -0.01 {
				motion.Acceleration[0] = -speed
			} else {
				motion.Acceleration[0] = 0.0
				if !control.gravityEnabled {
					motion.Velocity[0] *= 0.75
				}
			}

			if gmath.Abs(motion.Velocity[0]) > maxSpeed {
				motion.Velocity[0] = maxSpeed * gmath.Sign(motion.Velocity[0])
			}

			if control.YAxis.Amount() > 0.01 {
				if control.canJump {
					control.gravityEnabled = true
					control.canJump = false
					motion.Velocity[1] = 1.0
				} else if control.canWallJump {
					control.gravityEnabled = true
					control.canWallJump = false
					motion.Velocity[1] = 1.0
					if control.wallJumpLeft {
						motion.Velocity[0] = -2.0
					} else {
						motion.Velocity[0] = 2.0
					}
				}
			}

			if !control.canJump && !control.canWallJump {
				control.gravityEnabled = true
			}

			if control.gravityEnabled {
				motion.Acceleration[1] = -2.0
				if motion.Velocity[1] < 0.0 {
					motion.Acceleration[1] = -3.0
				}
			} else {
				motion.Acceleration[1] = -0.25
			}
		}
	}, (*ControlComponent)(nil), (*gmath.MotionComponent)(nil)))
	state.AddListener(interactionWorld)

	// Launch!
	limitengine.Launch(state)
}

// STATIC ENTITY FUNC
func CreateStaticEntity(state *limitengine.State, position, scale gmath.Vector3) {
	state.NewEntity(
		&gmath.TransformComponent{
			Position: position,
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    scale,
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-1.0*scale[0], -1.0*scale[1], 0.0), gmath.NewVector3(scale[0], scale[1], 0.0)),
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

// INPUT TESTS
type ControlComponent struct {
	XAxis, YAxis              *ui.InputControl
	canJump                   bool
	canWallJump, wallJumpLeft bool
	gravityEnabled            bool
}

func (controlComponent *ControlComponent) Delete() {}

// INTERACTION TESTS
type TestInteraction struct {
	test string
}

func (test TestInteraction) StartInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3, penetration float32) {
	// fmt.Println("BEGIN INTERACT")
	control := interactor.Entity.GetComponent((*ControlComponent)(nil)).(*ControlComponent)
	if !interactee.Collider.IsTrigger {
		if normal[1] < -0.5 {
			interactor.Motion.Velocity[1] = 0.0
			control.canJump = true
			control.gravityEnabled = false
		} else if gmath.Abs(normal[0]) > 0.9 {
			interactor.Motion.Velocity[0] = 0.0
			interactor.Motion.Velocity[1] = -0.1
			control.canWallJump = true
			control.gravityEnabled = false
			if normal[0] < 0.0 {
				control.wallJumpLeft = false
			} else {
				control.wallJumpLeft = true
			}
		}
	}
}

func (test TestInteraction) EndInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3) {
	// fmt.Println("end interact")
	control := interactor.Entity.GetComponent((*ControlComponent)(nil)).(*ControlComponent)
	if !interactee.Collider.IsTrigger {
		if gmath.Abs(normal[0]) > 0.9 {
			control.canWallJump = false
		}
	}
}

func (test TestInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*ControlComponent)(nil)),
		reflect.TypeOf((*gmath.MotionComponent)(nil)),
	}
}

func (test TestInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{}
}
