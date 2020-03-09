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

type TestComponent struct {
	t1 string
	t2 int
}

func main() {
	model := gfx.CreateModel(gio.LoadOBJ("monkey.obj"))
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
	
	interpNormal = normalize(normal);
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
	fragColor = vec4(diffuse, 1.0) * vec4(interpNormal, 1.0);
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
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeftShift}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyRightShift}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeySpace}, 1.0)
	zAxis := ui.InputControl{}
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, -1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyUp}, -1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, 1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyDown}, 1.0)

	camera := gfx.CreateCamera()
	ecs.NewEntity(
		&utils.TransformComponent{
			Position: gmath.Vector{0.0, 1.0, -10.0, 1.0},
			Rotation: gmath.Vector{0.0, 0.0, 0.0},
			Scale:    gmath.Vector{1.0, 1.0, 1.0},
		},
		&utils.MotionComponent{
			Velocity:     gmath.Vector{0.0, 0.0, 0.0},
			Acceleration: gmath.Vector{0.0, 0.0, 0.0},
		},
		&utils.MotionControlComponent{
			Axis:  []*ui.InputControl{&xAxis, &yAxis, &zAxis},
			Speed: 100.0,
		},
		&utils.CameraComponent{
			Camera:    camera,
			PosOffset: gmath.Vector{0.0, 2.0, 15.0},
		},
		&TestComponent{},
	)

	for i := 0; i < 1000; i++ {
		randVelocity := gmath.Vector{rand.Float32()*500.0 - 250.0, rand.Float32()*500.0 - 250.0, rand.Float32()*500.0 - 250.0}
		ecs.NewEntity(
			&utils.TransformComponent{
				Position: gmath.Vector{rand.Float32()*10.0 - 5.0, rand.Float32()*10.0 - 5.0, rand.Float32()*10.0 - 25.0, 1.0},
				Rotation: gmath.Vector{0.0, 0.0, 0.0},
				Scale:    gmath.Vector{1.0, 1.0, 1.0},
			},
			&utils.MotionComponent{
				Velocity:     randVelocity,
				Acceleration: randVelocity.Clone().MulSc(-0.5),
			},
			&TestComponent{},
		)
	}

	testSystem := ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)
			tMat := gmath.NewIdentityMatrix(4, 4)
			tMat.Translate(transform.Position)

			instance := gfx.NewUniformLoader()
			instance.AddMatrix44("transformMat", tMat.ToMatrix44())

			gfx.Render(camera, shader, material, model, instance)
		}
		gfx.Sweep()
	}, (*utils.TransformComponent)(nil), (*TestComponent)(nil))
	ecs.AddSystem(utils.NewMotionControlSystem())
	ecs.AddSystem(utils.NewMotionSystem(0.95))
	ecs.AddSystem(utils.NewCameraMotionSystem())
	ecs.AddSystem(testSystem)

	limitengine.Launch()
}
