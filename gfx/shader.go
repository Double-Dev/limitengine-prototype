package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	shaderIndex = uint32(1)
	shaders     = make(map[uint32]framework.IShader)
)

func init() { shaders[0] = nil }

type Shader struct {
	id             uint32
	uniformLoader  uniformLoader
	instanceBuffer framework.IInstanceBuffer
}

func CreateShader(vertSrc, fragSrc string) *Shader {
	shader := &Shader{
		id: shaderIndex,
	}
	shaderIndex++
	actionQueue = append(actionQueue, func() {
		shaders[shader.id] = context.CreateShader(vertSrc, fragSrc)
		totalInstanceSize := 0
		_, varMap := shader.GetInstanceVarsSize()
		for _, size := range varMap {
			totalInstanceSize += size
		}
		shader.instanceBuffer = context.CreateInstanceBuffer(totalInstanceSize)
	})
	return shader
}

func (shader *Shader) GetInstanceVarsSize() ([]string, map[string]int) {
	return []string{
			"transformMat0",
			"transformMat1",
			"transformMat2",
			"transformMat3",
		}, map[string]int{
			"transformMat0": 4,
			"transformMat1": 4,
			"transformMat2": 4,
			"transformMat3": 4,
		}
}

// TODO: Add support for more variables + array uniforms.
type uniformLoader struct {
	uniformInts      map[string]int32
	uniformMatrix44s map[string]gmath.Matrix
}

func newUniformLoader() uniformLoader {
	return uniformLoader{
		uniformInts:      make(map[string]int32),
		uniformMatrix44s: make(map[string]gmath.Matrix),
	}
}

func (this uniformLoader) loadTo(iShader framework.IShader) {
	for varName, value := range this.uniformInts {
		iShader.LoadUniform1I(varName, value)
	}
	for varName, value := range this.uniformMatrix44s {
		iShader.LoadUniformMatrix4fv(varName, value.ToArray())
	}
}

func (this uniformLoader) AddInt(varName string, val int32) {
	this.uniformInts[varName] = val
}

func (this uniformLoader) AddMatrix44(varName string, val gmath.Matrix) {
	if val.IsSize(4, 4) {
		this.uniformMatrix44s[varName] = val
	}
}
