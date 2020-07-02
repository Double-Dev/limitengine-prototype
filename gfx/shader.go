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

type Shader interface {
	RenderProgram() *RenderProgram
	UniformLoader() UniformLoader
}

type GenericShader struct {
	renderProgram *RenderProgram
	uniformLoader UniformLoader
}

func NewGenericShader(renderProgram *RenderProgram) *GenericShader {
	return &GenericShader{
		renderProgram,
		NewUniformLoader(),
	}
}

func (shader *GenericShader) RenderProgram() *RenderProgram { return shader.renderProgram }
func (shader *GenericShader) UniformLoader() UniformLoader  { return shader.uniformLoader }

type RenderProgram struct {
	id             uint32
	instanceDefs   []framework.InstanceDef
	instanceBuffer framework.IInstanceBuffer
}

func NewRenderProgram(leslPlugins ...*LESLPlugin) *RenderProgram {
	program := &RenderProgram{id: shaderIndex}
	shaderIndex++
	vertSrc, fragSrc, instanceDefs, textureVars := processLESL(leslPlugins)
	program.instanceDefs = append([]framework.InstanceDef{
		framework.InstanceDef{Name: "verttransformMat0", Size: 4, Index: 0},
		framework.InstanceDef{Name: "verttransformMat1", Size: 4, Index: 4},
		framework.InstanceDef{Name: "verttransformMat2", Size: 4, Index: 8},
		framework.InstanceDef{Name: "verttransformMat3", Size: 4, Index: 12},
	}, instanceDefs...)
	totalInstanceSize := 0
	for _, instanceDef := range program.InstanceDefs() {
		totalInstanceSize += instanceDef.Size
	}
	actionQueue = append(actionQueue, func() {
		iShader := context.NewShader(vertSrc, fragSrc)
		iShader.Start()
		for key, value := range textureVars {
			iShader.LoadUniform1I(key, value)
		}
		iShader.Stop()
		shaders[program.id] = iShader
		program.instanceBuffer = context.NewInstanceBuffer(totalInstanceSize)
	})
	return program
}

func (program *RenderProgram) InstanceDefs() []framework.InstanceDef { return program.instanceDefs }

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
	for key, value := range uniformLoader.uniformInts {
		iShader.LoadUniform1I(key, value)
	}
	for key, value := range uniformLoader.uniformFloats {
		iShader.LoadUniform1F(key, value)
	}
	for key, value := range uniformLoader.uniformVector3s {
		iShader.LoadUniform3F(key, value[0], value[1], value[2])
	}
	for key, value := range uniformLoader.uniformVector4s {
		iShader.LoadUniform4F(key, value[0], value[1], value[2], value[3])
	}
	for key, value := range uniformLoader.uniformMatrix4s {
		iShader.LoadUniformMatrix4fv(key, value.ToArray())
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

// DeleteRenderProgram queues a gfx action that deletes the input shader.
func DeleteRenderProgram(program *RenderProgram) {
	actionQueue = append(actionQueue, func() {
		iShader := shaders[program.id]
		iShader.Delete()
		delete(shaders, program.id)
	})
}
