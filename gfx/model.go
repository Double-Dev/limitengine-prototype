package gfx

import "github.com/double-dev/limitengine/gfx/framework"

var (
	modelIndex = uint32(1)
	models     = make(map[uint32]framework.IModel)
)

func init() {
	// Sets zero model to cube.
	actionQueue = append(actionQueue, func() {
		models[0] = context.CreateModel(
			[]uint32{
				0, 1, 3, 3, 1, 2,
				4, 5, 7, 7, 5, 6,
				8, 9, 11, 11, 9, 10,
				12, 13, 15, 15, 13, 14,
				16, 17, 19, 19, 17, 18,
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
