package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	modelIndex = uint32(1)
	models     = make(map[uint32]framework.IModel)
)

func init() {
	// Sets zero model to cube.
	actionQueue = append(actionQueue, func() {
		models[0] = context.CreateModel(
			[]uint32{
				3, 1, 0, 2, 1, 3,
				4, 5, 7, 7, 5, 6,
				11, 9, 8, 10, 9, 11,
				12, 13, 15, 15, 13, 14,
				19, 17, 16, 18, 17, 19,
				20, 21, 23, 23, 21, 22,
			},
			[]float32{
				-1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0,
				-1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0,
				1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0,
				-1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0,
				-1.0, 1.0, 1.0, -1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0,
				-1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0,
			},
			[]float32{
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
				0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
			},
			[]float32{
				0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
				0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
				1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
				-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
				0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
				0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
			},
		)
	})
}

// Model is a gfx model.
type Model struct {
	id          uint32
	prefs       uniformLoader
	DepthTest   bool
	BackCulling bool
	WriteDepth  bool
}

// CreateModel queues a gfx action that creates a model using the input model data.
func CreateModel(indices []uint32, vertices, texCoords, normals []float32) *Model {
	model := &Model{
		id:          modelIndex,
		DepthTest:   true,
		BackCulling: true,
		WriteDepth:  true,
	}
	modelIndex++
	actionQueue = append(actionQueue, func() { models[model.id] = context.CreateModel(indices, vertices, texCoords, normals) })
	return model
}

// DeleteModel queues a gfx action that deletes the input model.
func DeleteModel(model *Model) {
	actionQueue = append(actionQueue, func() {
		iModel := models[model.id]
		iModel.Delete()
		delete(models, model.id)
	})
}

type Instance struct {
	uniformInts      map[string]int32
	uniformMatrix44s map[string]gmath.Matrix44
}

func NewInstance() *Instance {
	return &Instance{
		uniformInts:      make(map[string]int32),
		uniformMatrix44s: make(map[string]gmath.Matrix44),
	}
}

func (this *Instance) loadTo(iShader framework.IShader) {
	gfxMutex.RLock()
	for varName, value := range this.uniformInts {
		iShader.LoadUniform1I(varName, value)
	}
	for varName, value := range this.uniformMatrix44s {
		iShader.LoadUniformMatrix4fv(varName, value)
	}
	gfxMutex.RUnlock()
}

func (this *Instance) AddInt(varName string, val int32) {
	gfxMutex.Lock()
	this.uniformInts[varName] = val
	gfxMutex.Unlock()
}

func (this *Instance) AddMatrix44(varName string, val gmath.Matrix44) {
	gfxMutex.Lock()
	this.uniformMatrix44s[varName] = val
	gfxMutex.Unlock()
}
