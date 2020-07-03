package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
)

var (
	meshIndex  = uint32(2)
	meshes     = make(map[uint32]framework.IMesh)
	spriteMesh = &Mesh{
		id:          0,
		DepthTest:   true, // Enabling depth test with write depth false causes a depth buffer issue with the POST shader in the tests.
		BackCulling: false,
		WriteDepth:  true,
	}
	cubeMesh = &Mesh{
		id:          1,
		DepthTest:   true,
		BackCulling: true,
		WriteDepth:  true,
	}
)

func init() {
	// Sets [0] mesh to plane.
	actionQueue = append(actionQueue, func() {
		meshes[0] = context.NewMesh(
			[]uint32{0, 1, 3, 3, 1, 2},
			[]float32{-1.0, 1.0, 0.0, -1.0, -1.0, 0.0, 1.0, -1.0, 0.0, 1.0, 1.0, 0.0},
			[]float32{0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0},
			[]float32{0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0},
		)
	})
	// Sets [1] mesh to cube.
	actionQueue = append(actionQueue, func() {
		meshes[1] = context.NewMesh(
			[]uint32{3, 1, 0, 2, 1, 3, 4, 5, 7, 7, 5, 6, 11, 9, 8, 10, 9, 11, 12, 13, 15, 15, 13, 14, 19, 17, 16, 18, 17, 19, 20, 21, 23, 23, 21, 22},
			[]float32{-1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, -1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, -1.0, -1.0, -1.0, 1.0, -1.0, -1.0, 1.0, -1.0, 1.0},
			[]float32{0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 1.0, 1.0, 1.0, 0.0},
			[]float32{0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0},
		)
	})
}

func deleteMeshes() {
	meshIndex = 0
	for _, iMesh := range meshes {
		iMesh.Delete()
	}
	meshes = nil
	spriteMesh = nil
	cubeMesh = nil
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
	mesh := &Mesh{id: meshIndex, DepthTest: true, BackCulling: true, WriteDepth: true}
	meshIndex++
	actionQueue = append(actionQueue, func() { meshes[mesh.id] = context.NewMesh(indices, vertices, texCoords, normals) })
	return mesh
}

// DeleteMesh queues a gfx action that deletes the input mesh.
func DeleteMesh(mesh *Mesh) {
	actionQueue = append(actionQueue, func() {
		iMesh := meshes[mesh.id]
		iMesh.Delete()
		delete(meshes, mesh.id)
	})
}

func SpriteMesh() *Mesh { return spriteMesh }
func CubeMesh() *Mesh   { return cubeMesh }
