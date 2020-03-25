package gl

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
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

func (vao *vao) Render(instanceBuffer framework.IInstanceBuffer, instanceData []float32, numInstances int32) {
	instanceBuffer.Bind()
	instanceBuffer.StoreInstancedData(instanceData)

	gl.VertexAttribPointer(3, 4, gl.FLOAT, false, 4*16, gl.PtrOffset(0))
	gl.VertexAttribDivisor(3, 1)

	gl.VertexAttribPointer(4, 4, gl.FLOAT, false, 4*16, gl.PtrOffset(4*4))
	gl.VertexAttribDivisor(4, 1)

	gl.VertexAttribPointer(5, 4, gl.FLOAT, false, 4*16, gl.PtrOffset(8*4))
	gl.VertexAttribDivisor(5, 1)

	gl.VertexAttribPointer(6, 4, gl.FLOAT, false, 4*16, gl.PtrOffset(12*4))
	gl.VertexAttribDivisor(6, 1)

	instanceBuffer.Unbind()

	gl.EnableVertexAttribArray(3)
	gl.EnableVertexAttribArray(4)
	gl.EnableVertexAttribArray(5)
	gl.EnableVertexAttribArray(6)

	gl.DrawElementsInstanced(gl.TRIANGLES, vao.vertexNum, gl.UNSIGNED_INT, gl.PtrOffset(0), numInstances)

	gl.DisableVertexAttribArray(3)
	gl.DisableVertexAttribArray(4)
	gl.DisableVertexAttribArray(5)
	gl.DisableVertexAttribArray(6)

	// gl.DrawElements(gl.TRIANGLES, vao.vertexNum, gl.UNSIGNED_INT, gl.PtrOffset(0))
}

func (vao *vao) Disable() {
	for _, index := range vao.attributes {
		gl.DisableVertexAttribArray(index)
	}
	vao.unbind()
}

func (vao *vao) addIndices(indices []uint32) {
	vbo := newVBO(vboElementArrayBufferType)
	vbo.Bind()
	vbo.storeUIntData(indices)
	vao.vertexNum = int32(len(indices))
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addIntAttrib(data []int32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := newVBO(vboArrayBufferType)
	vbo.Bind()
	vbo.storeIntData(data)
	gl.VertexAttribPointer(index, varsPerVertex, gl.INT, normalized, 0, nil)
	vbo.Unbind()
	vao.attributes = append(vao.attributes, index)
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addFloatAttrib(data []float32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := newVBO(vboArrayBufferType)
	vbo.Bind()
	vbo.storeFloatData(data)
	gl.VertexAttribPointer(index, varsPerVertex, gl.FLOAT, normalized, 0, nil)
	vbo.Unbind()
	vao.attributes = append(vao.attributes, index)
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) Delete() {
	for _, vbo := range vao.vbos {
		vbo.delete()
	}
	gl.DeleteVertexArrays(1, &vao.id)
}
