package main

import (
	"fmt"

	"github.com/double-dev/limitengine/ui"

	"github.com/double-dev/limitengine/gmath"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/ecs"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
)

type TestComponent struct {
	t1 string
	t2 int
}

func main() {
	indices := []uint32{
		0, 1, 2, 2, 1, 3,
	}

	vertices := []float32{
		-0.75, 0.75, -2.0,
		-0.75, -0.75, -2.0,
		0.75, 0.75, -2.0,
		0.75, -0.75, -2.0,
	}

	texCoords := []float32{
		0.0, 0.0,
		0.0, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	}

	normals := []float32{
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
		0.0, 0.0, 0.0,
	}

	model := gfx.CreateModel(indices, vertices, texCoords, normals)
	fmt.Println(model)
	// model := gfx.CreateModel(gio.LoadOBJ("lamp.obj"))
	shader := gfx.CreateShader(`#version 150 core
in vec3 coord;
in vec2 texCoord;
in vec3 norm;
uniform mat4 projMat;
uniform mat4 viewMat;
void main()
{
	gl_Position = projMat * viewMat * vec4(coord, 1.0);
}`, `#version 150 core
out vec4 fragColor;
void main()
{
	fragColor = vec4(0.0, 1.0, 0.0, 1.0);
}`,
	)
	fmt.Println(shader)
	texture := gfx.CreateTexture(gio.LoadPNG("../DefaultIcon.png"))
	fmt.Println(texture)

	camPos := gmath.NewVector(0.0, 0.0, -2.0, 1.0)
	limitengine.AddKeyCallback(
		func(key limitengine.Key, scancode int, action limitengine.Action, mods limitengine.ModKey) {
			if key == ui.KeyW {
				camPos[1] += 0.05
			} else if key == ui.KeyS {
				camPos[1] -= 0.05
			}
			fmt.Println(camPos)
		},
	)

	testSystem := ecs.NewSystem(func(delta float32, entity ecs.ECSEntity) {
		camera := gmath.NewMatrix(4, 4)
		camera.Translate(camPos)
		fmt.Println(camera)
		// fmt.Println(camera.MulV(gmath.NewVector(0.0, 1.0, 0.0, 1.0)))

		gfx.ClearScreen(0.0, 0.1, 0.25, 1.0)
		gfx.Render(camera, shader, model, texture)
		gfx.RenderSweep()
	})

	ecs.AddSystem(testSystem)
	ecs.NewEntity()

	limitengine.Launch()
}
