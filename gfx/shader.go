package gfx

import (
	"double-dev/limitengine/gfx/framework"
)

var (
	shaderIndex = uint32(0)
	shaders     = make(map[uint32]framework.IShader)
)

type Shader struct {
	id            uint32
	uniformLoader func(framework.IShader, framework.IModel)
}

func CreateShader(vertSrc, fragSrc string) *Shader {
	shader := &Shader{
		id: shaderIndex,
	}
	shaderIndex++
	actionQueue = append(actionQueue, func() { shaders[shader.id] = context.CreateShader(vertSrc, fragSrc) })
	return shader
}
