package main

import (
	"math/rand"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/ui"
	"github.com/double-dev/limitengine/utils"
)

func main() {
	model := gfx.CreateModel(gio.LoadOBJ("monkey.obj"))
	// model := &gfx.Model{}
	shader := gfx.CreateShader(`#version 330 core
layout(location = 0) in vec3 coord;
layout(location = 1) in vec2 texCoord;
layout(location = 2) in vec3 normal;
uniform mat4 projMat;
uniform mat4 viewMat;
uniform mat4 transformMat;

out vec2 textureCoord;
out vec3 interpNormal;
out vec3 toLight;
const vec3 lightPos = vec3(-5.0, 10.0, 5.0);
void main()
{
	vec4 worldPos = transformMat * vec4(coord, 1.0);
	gl_Position = projMat * viewMat * worldPos;

	textureCoord = vec2(texCoord.x, 1.0 - texCoord.y);

	interpNormal = normalize((transformMat * vec4(normal, 0.0)).xyz);
	toLight = normalize(lightPos - worldPos.xyz);
}`, `#version 330 core
in vec2 textureCoord;
in vec3 interpNormal;
in vec3 toLight;
uniform sampler2D tex;

out vec4 fragColor;
const vec3 lightColor = vec3(1.0, 1.0, 1.0);
const float ambient = 0.1;
void main()
{
	float lightDot = dot(interpNormal, toLight);
	float brightness = max(lightDot, ambient);
	vec3 diffuse = brightness * lightColor;
	// fragColor = vec4(diffuse, 1.0) * texture(tex, textureCoord);
	fragColor = vec4(diffuse, 1.0) * vec4(abs(interpNormal), 1.0);
	// fragColor = vec4(diffuse, 1.0) * vec4(1.0, 1.0, 1.0, 1.0);
}`,
	)
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
	boost.AddTrigger(ui.InputEvent{Key: ui.KeyRightShift}, -0.15)
	boost.AddTrigger(ui.InputEvent{Key: ui.KeyLeftShift}, -0.15)

	camera := gfx.CreateCamera()
	ecs.NewEntity(
		&utils.TransformComponent{
			Position: gmath.NewVector3(0.0, 1.0, -10.0),
			Rotation: gmath.NewQuaternion(rand.Float32()*gmath.Pi, 0.0, 0.0, 1.0),
			Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
		},
		&utils.MotionComponent{
			Velocity:        gmath.NewVector3(0.0, 0.0, 0.0),
			Acceleration:    gmath.NewVector3(0.0, 0.0, 0.0),
			AngVelocity:     gmath.NewIdentityQuaternion(),
			AngAcceleration: gmath.NewIdentityQuaternion(),
		},
		&utils.MotionControlComponent{
			Axis:  []*ui.InputControl{&xAxis, &yAxis, &zAxis, &boost},
			Speed: 600.0,
		},
		&utils.CameraComponent{
			Camera:      camera,
			PosOffset:   gmath.NewVector3(0.0, 0.0, 15.0),
			RotOffset:   gmath.NewIdentityQuaternion(),
			ScaleOffset: gmath.NewVector3(1.0, 1.0, 1.0),
		},
		&utils.RenderComponent{
			Camera:   camera,
			Shader:   shader,
			Material: material,
			Model:    model,
			Instance: gfx.NewInstance(),
		},
	)

	for i := 0; i < 2000; i++ {
		// randVelocity := gmath.NewVector(rand.Float32()*500.0-250.0, rand.Float32()*500.0-250.0, rand.Float32()*500.0-250.0)
		ecs.NewEntity(
			&utils.TransformComponent{
				Position: gmath.NewVector3(rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-500.0, rand.Float32()*1000.0-750.0),
				Rotation: gmath.NewQuaternion(rand.Float32()*gmath.Pi, rand.Float32(), rand.Float32(), rand.Float32()),
				Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
			},
			// &utils.MotionComponent{
			// 	Velocity:     randVelocity,
			// 	Acceleration: randVelocity.Clone().MulSc(-0.5),
			// },
			&utils.RenderComponent{
				Camera:   camera,
				Shader:   shader,
				Material: material,
				Model:    model,
				Instance: gfx.NewInstance(),
			},
		)
	}

	// ecs.AddSystem(utils.NewMotionControlSystem())
	ecs.AddSystem(ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		for _, entity := range entities {
			control := entity.GetComponent((*utils.MotionControlComponent)(nil)).(*utils.MotionControlComponent)
			motion := entity.GetComponent((*utils.MotionComponent)(nil)).(*utils.MotionComponent)
			transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)

			direction := gmath.NewVector3(control.Axis[2].Amount(), control.Axis[0].Amount(), control.Axis[1].Amount())

			if direction.LenSq() > 0.5 {
				motion.AngAcceleration = gmath.NewQuaternion(gmath.Pi, direction[0], direction[1], direction[2])
			} else {
				motion.AngAcceleration = gmath.NewIdentityQuaternion()
			}

			var speed float32
			if control.Axis[3].Amount() == 0.0 {
				speed = -control.Speed * 0.25
			} else {
				speed = -control.Speed * control.Axis[3].Amount()
			}

			motion.Acceleration = transform.Rotation.RotateV(gmath.NewVector3(0.0, 0.0, speed))
		}
	}, (*utils.MotionControlComponent)(nil), (*utils.MotionComponent)(nil), (*utils.TransformComponent)(nil)))
	ecs.AddSystem(utils.NewMotionSystem(0.95))
	ecs.AddSystem(utils.NewCameraMotionSystem())
	ecs.AddSystem(utils.NewRenderSystem())

	limitengine.Launch()
}
