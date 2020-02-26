package gl

import (
	"github.com/go-gl/gl/v3.2-core/gl"
)

type vao struct {
	id         uint32
	vertexNum  int32
	attributes []uint32
	vbos       []vbo
}

func newVAO() *vao {
	var id uint32
	gl.GenVertexArrays(1, &id)
	return &vao{
		id:         id,
		attributes: []uint32{},
		vbos:       []vbo{},
	}
}

func (vao *vao) bind() { gl.BindVertexArray(vao.id) }
func (*vao) unbind()   { gl.BindVertexArray(0) }

func (vao *vao) Enable() {
	vao.bind()
	for _, index := range vao.attributes {
		gl.EnableVertexAttribArray(index)
	}
}

func (vao *vao) Render() {
	gl.DrawElements(gl.TRIANGLES, vao.vertexNum, gl.UNSIGNED_INT, gl.PtrOffset(0))
}

func (vao *vao) Disable() {
	for _, index := range vao.attributes {
		gl.DisableVertexAttribArray(index)
	}
	vao.unbind()
}

func (vao *vao) addIndicesArr(indices []uint32) {
	vbo := newVBO(vboElementArrayBufferType)
	vbo.bind()
	vbo.storeUIntData(indices)
	vao.vertexNum = int32(len(indices))
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addIntAttribArr(data []int32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := newVBO(vboArrayBufferType)
	vbo.bind()
	vbo.storeIntData(data)
	gl.VertexAttribPointer(index, varsPerVertex, gl.INT, normalized, 0, nil)
	vbo.unbind()
	vao.attributes = append(vao.attributes, index)
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addFloatAttribArr(data []float32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := newVBO(vboArrayBufferType)
	vbo.bind()
	vbo.storeFloatData(data)
	gl.VertexAttribPointer(index, varsPerVertex, gl.FLOAT, normalized, 0, nil)
	vbo.unbind()
	vao.attributes = append(vao.attributes, index)
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) Delete() {
	for _, vbo := range vao.vbos {
		vbo.delete()
	}
	gl.DeleteVertexArrays(1, &vao.id)
}
