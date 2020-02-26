package gl

import "github.com/go-gl/gl/v3.2-core/gl"

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

func (vbo *vbo) bind()   { gl.BindBuffer(vbo.vboType, vbo.id) }
func (vbo *vbo) unbind() { gl.BindBuffer(vbo.vboType, 0) }

func (vbo *vbo) storeUIntData(data []uint32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo *vbo) storeIntData(data []int32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo *vbo) storeFloatData(data []float32) {
	gl.BufferData(vbo.vboType, 4*len(data), gl.Ptr(data), gl.STATIC_DRAW)
}

func (vbo *vbo) delete() {
	gl.DeleteBuffers(1, &vbo.id)
}
