package main

import (
	"fmt"

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
	model := gfx.CreateModel(gio.LoadOBJ("dragon.obj"))
	fmt.Println(model)
	shader := gfx.CreateShader(`#version 150 core
in vec3 coord;
in vec2 texCoord;
in vec3 norm;
out vec2 textureCoord;
out vec3 normal;
// out vec3 toLightVec;
uniform mat4 projMat;
uniform mat4 viewMat;
// const vec3 lightPos = vec3(0.0, 0.0, -10.0);
void main()
{
	// vec3 toLight = lightPos - coord;
	textureCoord = vec2(texCoord.x, 1.0 - texCoord.y);
	normal = norm;
	gl_Position = projMat * viewMat * vec4(coord, 1.0);
}`, `#version 150 core
in vec2 textureCoord;
in vec3 normal;
// in vec3 toLightVec;
out vec4 fragColor;
uniform sampler2D tex;
const vec3 lightColor = vec3(1.0, 1.0, 1.0);
const float ambient = 0.1;
void main()
{
	float lightDot = dot(normalize(normal), vec3(0.0, 1.0, 0.0));
	float brightness = max(lightDot, ambient);
	vec3 diffuse = brightness * lightColor;
	// fragColor = texture(tex, textureCoord);
	fragColor = vec4(diffuse, 1.0) * vec4(0.6, 0.6, 0.1, 1.0);
}`,
	)
	fmt.Println(shader)
	texture := gfx.CreateTexture(gio.LoadPNG("lamp.png"))
	fmt.Println(texture)

	xAxis := ui.InputControl{}
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyA}, 1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyD}, -1.0)
	yAxis := ui.InputControl{}
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeftShift}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeySpace}, -1.0)
	zAxis := ui.InputControl{}
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, 1.0)
	zAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, -1.0)

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
	)

	testSystem := ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		transform := entity.GetComponent((*utils.TransformComponent)(nil)).(*utils.TransformComponent)
		camera := gmath.NewMatrix(4, 4)
		camera.Translate(transform.Position)

		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		gfx.Render(camera, shader, model, texture)
		gfx.RenderSweep()
	}, (*utils.TransformComponent)(nil))
	ecs.AddSystem(testSystem)
	ecs.AddSystem(utils.NewMotionControlSystem())
	ecs.AddSystem(utils.NewMotionSystem(0.95))

	limitengine.Launch()
}
