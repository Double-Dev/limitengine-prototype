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
	model := gfx.CreateModel(gio.LoadOBJ("lamp.obj"))
	shader := gfx.CreateShader(`#version 150 core
in vec3 coord;
in vec2 texCoord;
in vec3 norm;
out vec2 textureCoord;
// out vec3 normal;
// out vec3 toLight;
uniform mat4 projMat;
uniform mat4 mvMat;
// const vec3 lightPos = vec3(-5.0, -10.0, 0.0);
void main()
{
	// toLight = normalize(normalize(lightPos) - normalize(coord));
	textureCoord = vec2(texCoord.x, 1.0 - texCoord.y);
	// normal = normalize(norm);
	gl_Position = projMat * mvMat * vec4(coord, 1.0);
}`, `#version 150 core
in vec2 textureCoord;
// in vec3 normal;
// in vec3 toLight;
out vec4 fragColor;
uniform sampler2D tex;
// const vec3 lightColor = vec3(1.0, 1.0, 1.0);
// const float ambient = 0.1;
void main()
{
	// float lightDot = dot(normal, toLight);
	// float brightness = max(lightDot, ambient);
	// vec3 diffuse = brightness * lightColor;
	fragColor = texture(tex, textureCoord);
	// fragColor = vec4(diffuse, 1.0) * vec4(0.6, 0.6, 0.1, 1.0);
}`,
	)
	texture := gfx.CreateTexture(gio.LoadPNG("lamp.png"))

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
			Position: gmath.NewVector(0.0, 1.0, -10.0, 1.0),
			Rotation: gmath.NewVector(0.0, 0.0, 0.0),
			Scale:    gmath.NewVector(1.0, 1.0, 1.0),
		},
		&utils.MotionComponent{
			Velocity:     gmath.NewVector(0.0, 0.0, 0.0),
			Acceleration: gmath.NewVector(0.0, 0.0, 0.0),
		},
		&utils.MotionControlComponent{
			Axis:  []*ui.InputControl{&xAxis, &yAxis, &zAxis},
			Speed: 100.0,
		},
		&utils.CameraComponent{
			Camera:    camera,
			PosOffset: gmath.NewVector(0.0, 5.0, 10.0),
		},
		&TestComponent{},
	)

	for i := 0; i < 50; i++ {
		ecs.NewEntity(
			&utils.TransformComponent{
				Position: gmath.NewVector(rand.Float32()*10.0-15.0, rand.Float32()*10.0-15.0, rand.Float32()*10.0-15.0, 1.0),
				Rotation: gmath.NewVector(0.0, 0.0, 0.0),
				Scale:    gmath.NewVector(1.0, 1.0, 1.0),
			},
			&TestComponent{},
		)
	}

	testSystem := ecs.NewSystem(func(delta float32, entities []ecs.ECSEntity) {
		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		for _, entity := range entities {
			transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)
			tMat := gmath.NewMatrix(4, 4)
			tMat.Translate(transform.Position)

			gfx.Render(camera, shader, model, texture, tMat)
		}
		gfx.Sweep()
	}, (*utils.TransformComponent)(nil), (*TestComponent)(nil))
	ecs.AddSystem(testSystem)
	ecs.AddSystem(utils.NewMotionControlSystem())
	ecs.AddSystem(utils.NewMotionSystem(0.95))
	ecs.AddSystem(utils.NewCameraMotionSystem())

	limitengine.Launch()
}
