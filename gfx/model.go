package gfx

import "double-dev/limitengine/gfx/framework"

var (
	modelIndex = uint32(0)
	models     = make(map[uint32]framework.IModel)
)

// Model is a gfx model.
type Model struct {
	id          uint32
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
