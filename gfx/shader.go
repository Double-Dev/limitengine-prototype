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

func CreateShader(leslSrc string) *Shader {
	shader := &Shader{
		id: shaderIndex,
	}
	shaderIndex++
	actionQueue = append(actionQueue, func() {
		vertSrc, fragSrc, textureVars := processLESL(leslSrc)
		shaders[shader.id] = context.CreateShader(vertSrc, fragSrc)
		totalInstanceSize := 0
		instanceDefs := shader.GetInstanceDefs()
		for _, instanceDef := range instanceDefs {
			totalInstanceSize += instanceDef.Size
		}
		shader.uniformLoader = newUniformLoader()
		for key, value := range textureVars {
			shader.uniformLoader.AddInt(key, value)
		}
		shader.instanceBuffer = context.CreateInstanceBuffer(totalInstanceSize)
	})
	return shader
}

func (shader *Shader) GetInstanceDefs() []struct {
	Name  string
	Size  int
	Index int
} {
	return []struct {
		Name  string
		Size  int
		Index int
	}{
		{"verttransformMat0", 4, 0},
		{"verttransformMat1", 4, 4},
		{"verttransformMat2", 4, 8},
		{"verttransformMat3", 4, 12},
	}
}

func (shader *Shader) GetUniformLoader() uniformLoader {
	return shader.uniformLoader
}

// TODO: Add support for more variables + array uniforms.
type uniformLoader struct {
	uniformInts     map[string]int32
	uniformFloats   map[string]float32
	uniformVector4s map[string]gmath.Vector4
	uniformMatrix4s map[string]gmath.Matrix4
}

func newUniformLoader() uniformLoader {
	return uniformLoader{
		uniformInts:     make(map[string]int32),
		uniformFloats:   make(map[string]float32),
		uniformVector4s: make(map[string]gmath.Vector4),
		uniformMatrix4s: make(map[string]gmath.Matrix4),
	}
}

func (uniformLoader uniformLoader) loadTo(iShader framework.IShader) {
	for varName, value := range uniformLoader.uniformInts {
		iShader.LoadUniform1I(varName, value)
	}
	for varName, value := range uniformLoader.uniformFloats {
		iShader.LoadUniform1F(varName, value)
	}
	for varName, value := range uniformLoader.uniformVector4s {
		iShader.LoadUniform4F(varName, value[0], value[1], value[2], value[3])
	}
	for varName, value := range uniformLoader.uniformMatrix4s {
		iShader.LoadUniformMatrix4fv(varName, value.ToArray())
	}
}

func (uniformLoader uniformLoader) AddInt(varName string, val int32) {
	uniformLoader.uniformInts[varName] = val
}

func (uniformLoader uniformLoader) AddFloat(varName string, val float32) {
	uniformLoader.uniformFloats[varName] = val
}

func (uniformLoader uniformLoader) AddVector4(varName string, val gmath.Vector4) {
	uniformLoader.uniformVector4s[varName] = val
}

func (uniformLoader uniformLoader) AddMatrix4(varName string, val gmath.Matrix4) {
	uniformLoader.uniformMatrix4s[varName] = val
}
