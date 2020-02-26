package main

import (
	"double-dev/limitengine/core"
	"double-dev/limitengine/ecs"
	"double-dev/limitengine/gfx"
	"double-dev/limitengine/gio"
	"double-dev/limitengine/gmath"
	"fmt"
	"time"
)

type TestComponent struct {
	test1 string
	test2 int
}

func main() {
	// indices := []uint32{
	// 	0, 1, 2, 2, 1, 3,
	// }

	// vertices := []float32{
	// 	-0.75, 0.75, -1.5,
	// 	-0.75, -0.75, -1.5,
	// 	0.75, 0.75, -1.5,
	// 	0.75, -0.75, -1.5,
	// }

	// texCoords := []float32{
	// 	0.0, 0.0,
	// 	0.0, 1.0,
	// 	1.0, 0.0,
	// 	1.0, 1.0,
	// }

	// normals := []float32{
	// 	0.0, 0.0, 0.0,
	// 	0.0, 0.0, 0.0,
	// 	0.0, 0.0, 0.0,
	// 	0.0, 0.0, 0.0,
	// }

	// model := gfx.CreateModel(indices, vertices, texCoords, normals)
	model := gfx.CreateModel(gio.LoadOBJ("lamp.obj"))
	shader := gfx.CreateShader(
		`#version 400 core
layout (location = 0) in vec3 coord;
layout (location = 1) in vec2 texCoord;
layout (location = 2) in vec3 norm;
out vec2 textureCoords;
uniform mat4 projMat;
uniform mat4 viewMat;
void main()
{
	vec4 pos = projMat * vec4(coord, 1.0);
	textureCoords = texCoord;
	gl_Position = pos;
}`,
		`#version 400 core
in vec3 pos;
in vec2 textureCoords;
out vec4 fragColor;
uniform sampler2D tex;
void main()
{
	vec4 texColor = texture(tex, textureCoords);
	texColor.a = 1.0;
	if (texColor.a < 0.5) {
		discard;
	}
	fragColor = texColor;
}`,
	)

	texture := gfx.CreateTexture(gio.LoadPNG("testIcon.png"))

	component := &TestComponent{
		test1: "hello",
		test2: 0,
	}
	entity := ecs.NewEntity()
	entity.AddComponent(component)

	system := ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		fmt.Println(entity.GetComponent((*TestComponent)(nil)).(*TestComponent))
	}, (*TestComponent)(nil))

	fmt.Println(system.GetEntities())

	currentTime := time.Now().UnixNano()
	t := float32(0.0)
	for core.Running {
		lastTime := currentTime
		currentTime = time.Now().UnixNano()
		delta := float32(currentTime-lastTime) / 1000000000.0
		// fmt.Println("Update FPS:", 1.0/delta)

		camera := gmath.NewMatrix(4, 4)
		camera.Translate(gmath.NewVector(0.0, 0.0, -10.0))

		t += 1.5 * delta
		gfx.ClearScreen(gmath.Sin(t)/2.0+0.5, gmath.Cos(t)/2.0+0.5, 0.5, 1.0)
		gfx.Render(camera, shader, model, texture)
		gfx.RenderSweep()
		time.Sleep(time.Millisecond * 10)
	}
}
