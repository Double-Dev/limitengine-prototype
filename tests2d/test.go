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
	mesh := gfx.GetSpriteMesh()
	camera := gfx.CreateCamera2D()

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
	limitengine.NewEntity(
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
			AABB: gmath.NewAABB(gmath.NewVector3(-0.1, -0.1, 0.0), gmath.NewVector3(0.1, 0.1, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
		&ControlComponent{
			XAxis: xAxis,
			YAxis: yAxis,
		},
	)

	limitengine.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.5, -0.25, -0.4),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.5, 0.5, 1.0),
		},
		&gfx.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Mesh:     mesh,
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
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
	)

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
			Mesh:     mesh,
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
			Mesh:     mesh,
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
			Mesh:     mesh,
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
			Mesh:     mesh,
			Instance: gfx.NewInstance(),
		},
	)

	// Systems
	interactionWorld := interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)

	myInteraction := TestInteraction{
		test: "Hello",
	}

	interactionWorld.AddInteraction(myInteraction)

	gfxListener := gfx.NewGFXListener()
	limitengine.AddECSListener(gfxListener)

	limitengine.AddSystem(gfx.NewRenderSystem())
	limitengine.AddSystem(gmath.NewMotionSystem(1.0))
	limitengine.AddSystem(limitengine.NewSystem(func(delta float32, entities []limitengine.ECSEntity) {
		for _, entity := range entities {
			control := entity.GetComponent((*ControlComponent)(nil)).(*ControlComponent)
			motion := entity.GetComponent((*gmath.MotionComponent)(nil)).(*gmath.MotionComponent)

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
					motion.Velocity[1] += 2.0
				} else if control.canWallJump {
					control.gravityEnabled = true
					control.canWallJump = false
					motion.Velocity[1] += 2.0
					if control.wallJumpLeft {
						motion.Velocity[0] = -4.0
					} else {
						motion.Velocity[0] = 4.0
					}
				}
			}

			if !control.canJump && !control.canWallJump {
				control.gravityEnabled = true
			}

			if control.gravityEnabled {
				motion.Acceleration[1] = -3.45
				if motion.Velocity[1] < 0.0 {
					motion.Acceleration[1] = -4.45
				}
			} else {
				motion.Acceleration[1] = -0.5
			}
		}
	}, (*ControlComponent)(nil), (*gmath.MotionComponent)(nil), (*gmath.TransformComponent)(nil)))
	limitengine.AddECSListener(interactionWorld)

	// Launch!
	limitengine.Launch()
}

type ControlComponent struct {
	XAxis, YAxis              *ui.InputControl
	canJump                   bool
	canWallJump, wallJumpLeft bool
	gravityEnabled            bool
}

type TestInteraction struct {
	test string
}

func (test TestInteraction) StartInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3, penetration float32) {
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
