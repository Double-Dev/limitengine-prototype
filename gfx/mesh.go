package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
)

var (
	meshIndex  = uint32(1)
	meshes     = make(map[uint32]framework.IMesh)
	spriteMesh = &Mesh{
		id:          0,
		DepthTest:   true,
		BackCulling: false,
		WriteDepth:  true,
	}
)

func init() {
	// Sets zero mesh to plane.
	actionQueue = append(actionQueue, func() {
		meshes[0] = context.NewMesh(
			[]uint32{
				0, 1, 3, 3, 1, 2,
			},
			[]float32{
				-1.0, 1.0, 0.0,
				-1.0, -1.0, 0.0,
				1.0, -1.0, 0.0,
				1.0, 1.0, 0.0,
			},
			[]float32{
				0.0, 0.0,
				0.0, 1.0,
				1.0, 1.0,
				1.0, 0.0,
			},
			[]float32{
				0.0, 0.0, -1.0,
				0.0, 0.0, -1.0,
				0.0, 0.0, -1.0,
				0.0, 0.0, -1.0,
			},
		)
	})

	// Sets zero mesh to cube.
	// actionQueue = append(actionQueue, func() {
	// 	meshes[0] = context.NewMesh(
	// 		[]uint32{
	// 			3, 1, 0, 2, 1, 3,
	// 			4, 5, 7, 7, 5, 6,
	// 			11, 9, 8, 10, 9, 11,
	// 			12, 13, 15, 15, 13, 14,
	// 			19, 17, 16, 18, 17, 19,
	// 			20, 21, 23, 23, 21, 22,
	// 		},
	// 		[]float32{
	// 			-1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0,
	// 			-1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0,
	// 			1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0,
	// 			-1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0,
	// 			-1.0, 1.0, 1.0, -1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0,
	// 			-1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0,
	// 		},
	// 		[]float32{
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 			0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0,
	// 		},
	// 		[]float32{
	// 			0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
	// 			0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	// 			1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	// 			-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	// 			0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	// 			0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
	// 		},
	// 	)
	// })
}

// TODO: Think about turning Mesh into interface to allow easy creation of variations.
// Mesh is a gfx mesh.
type Mesh struct {
	id          uint32
	prefs       UniformLoader
	DepthTest   bool
	BackCulling bool
	WriteDepth  bool
}

// NewMesh queues a gfx action that news a mesh using the input mesh data.
func NewMesh(indices []uint32, vertices, texCoords, normals []float32) *Mesh {
	mesh := &Mesh{
		id:          meshIndex,
		DepthTest:   true,
		BackCulling: true,
		WriteDepth:  true,
	}
	meshIndex++
	actionQueue = append(actionQueue, func() { meshes[mesh.id] = context.NewMesh(indices, vertices, texCoords, normals) })
	return mesh
}

func SpriteMesh() *Mesh {
	return spriteMesh
}

// DeleteMesh queues a gfx action that deletes the input mesh.
func DeleteMesh(mesh *Mesh) {
	actionQueue = append(actionQueue, func() {
		iMesh := meshes[mesh.id]
		iMesh.Delete()
		delete(meshes, mesh.id)
	})
}
