package gfx

import (
	"sync"

	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	shaderIndex = uint32(1)
	shaders     = make(map[uint32]framework.IShader)
)

func init() { shaders[0] = nil }

func deleteShaders() {
	for _, iShader := range shaders {
		if iShader != nil {
			iShader.Delete()
		}
	}
	shaders = nil
}

type Shader struct {
	id             uint32
	uniformLoader  UniformLoader
	instanceBuffer framework.IInstanceBuffer
}

func NewShader(leslSrc string) *Shader {
	shader := &Shader{
		id: shaderIndex,
	}
	shaderIndex++
	actionQueue = append(actionQueue, func() {
		vertSrc, fragSrc, textureVars := processLESL(leslSrc)
		shaders[shader.id] = context.NewShader(vertSrc, fragSrc)
		totalInstanceSize := 0
		instanceDefs := shader.InstanceDefs()
		for _, instanceDef := range instanceDefs {
			totalInstanceSize += instanceDef.Size
		}
		shader.uniformLoader = NewUniformLoader()
		for key, value := range textureVars {
			shader.uniformLoader.AddInt(key, value)
		}
		shader.instanceBuffer = context.NewInstanceBuffer(totalInstanceSize)
	})
	return shader
}

func (shader *Shader) InstanceDefs() []struct {
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
		{"verttextureBounds", 4, 16},
	}
}

func (shader *Shader) UniformLoader() UniformLoader {
	return shader.uniformLoader
}

// TODO: Add support for more variables + array uniforms.
type UniformLoader struct {
	uniformInts     map[string]int32
	uniformFloats   map[string]float32
	uniformVector3s map[string]gmath.Vector3
	uniformVector4s map[string]gmath.Vector4
	uniformMatrix4s map[string]gmath.Matrix4
	mutex           sync.RWMutex
}

func NewUniformLoader() UniformLoader {
	return UniformLoader{
		uniformInts:     make(map[string]int32),
		uniformFloats:   make(map[string]float32),
		uniformVector3s: make(map[string]gmath.Vector3),
		uniformVector4s: make(map[string]gmath.Vector4),
		uniformMatrix4s: make(map[string]gmath.Matrix4),
	}
}

func (uniformLoader UniformLoader) loadTo(iShader framework.IShader) {
	uniformLoader.mutex.RLock()
	for varName, value := range uniformLoader.uniformInts {
		iShader.LoadUniform1I(varName, value)
	}
	for varName, value := range uniformLoader.uniformFloats {
		iShader.LoadUniform1F(varName, value)
	}
	for varName, value := range uniformLoader.uniformVector3s {
		iShader.LoadUniform3F(varName, value[0], value[1], value[2])
	}
	for varName, value := range uniformLoader.uniformVector4s {
		iShader.LoadUniform4F(varName, value[0], value[1], value[2], value[3])
	}
	for varName, value := range uniformLoader.uniformMatrix4s {
		iShader.LoadUniformMatrix4fv(varName, value.ToArray())
	}
	uniformLoader.mutex.RUnlock()
}

func (uniformLoader UniformLoader) AddInt(varName string, val int32) {
	uniformLoader.mutex.Lock()
	uniformLoader.uniformInts[varName] = val
	uniformLoader.mutex.Unlock()
}

func (uniformLoader UniformLoader) AddFloat(varName string, val float32) {
	uniformLoader.mutex.Lock()
	uniformLoader.uniformFloats[varName] = val
	uniformLoader.mutex.Unlock()
}

func (uniformLoader UniformLoader) AddVector3(varName string, val gmath.Vector3) {
	uniformLoader.mutex.Lock()
	uniformLoader.uniformVector3s[varName] = val
	uniformLoader.mutex.Unlock()
}

func (uniformLoader UniformLoader) AddVector4(varName string, val gmath.Vector4) {
	uniformLoader.mutex.Lock()
	uniformLoader.uniformVector4s[varName] = val
	uniformLoader.mutex.Unlock()
}

func (uniformLoader UniformLoader) AddMatrix4(varName string, val gmath.Matrix4) {
	uniformLoader.mutex.Lock()
	uniformLoader.uniformMatrix4s[varName] = val
	uniformLoader.mutex.Unlock()
}

// DeleteShader queues a gfx action that deletes the input shader.
func DeleteShader(shader *Shader) {
	actionQueue = append(actionQueue, func() {
		iShader := shaders[shader.id]
		iShader.Delete()
		delete(shaders, shader.id)
	})
}
