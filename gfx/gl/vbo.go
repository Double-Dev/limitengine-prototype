package gl

import "github.com/double-dev/limitengine/dependencies/gl/v3.3-core/gl"

const (
	vboArrayBufferType               = gl.ARRAY_BUFFER
	vboElementArrayBufferType        = gl.ELEMENT_ARRAY_BUFFER
	vboArrayBufferBindingType        = gl.ARRAY_BUFFER_BINDING
	vboElementArrayBufferBindingType = gl.ELEMENT_ARRAY_BUFFER_BINDING
)

type vbo struct {
	id      uint32
	vboType uint32
}

func newVBO(vboType uint32) vbo {
	var id uint32
	gl.GenBuffers(1, &id)
	return vbo{
		id:      id,
		vboType: vboType,
	}
}

func (vbo vbo) Bind()   { gl.BindBuffer(vbo.vboType, vbo.id) }
func (vbo vbo) Unbind() { gl.BindBuffer(vbo.vboType, 0) }

func (vbo *vbo) setEmpty(capacity int) {
	gl.BufferData(vbo.vboType, 4*capacity, nil, gl.STREAM_DRAW)
}

func (vbo *vbo) storeUIntData(data []uint32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo *vbo) storeIntData(data []int32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo *vbo) storeFloatData(data []float32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo vbo) StoreInstancedData(data []float32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STREAM_DRAW)
}

func (vbo *vbo) delete() {
	gl.DeleteBuffers(1, &vbo.id)
}
