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

func createVAO() *vao {
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

func (vao *vao) Render(instanceBuffer framework.IInstanceBuffer, instanceDefs []struct {
	Name  string
	Size  int
	Index int
}, instanceData []float32, numInstances int32) {
	instanceBuffer.Bind()
	instanceBuffer.StoreInstancedData(instanceData)

	stride := 0
	for _, instanceDef := range instanceDefs {
		stride += instanceDef.Size
	}

	for i, instanceDef := range instanceDefs {
		gl.VertexAttribPointer(uint32(i+3), int32(instanceDef.Size), gl.FLOAT, false, int32(stride*4), gl.PtrOffset(instanceDef.Index*4))
		gl.VertexAttribDivisor(uint32(i+3), 1)
	}

	instanceBuffer.Unbind()

	for i := range instanceDefs {
		gl.EnableVertexAttribArray(uint32(i + 3))
	}

	if numInstances > 1 {
		gl.DrawElementsInstanced(gl.TRIANGLES, vao.vertexNum, gl.UNSIGNED_INT, gl.PtrOffset(0), numInstances)
	} else {
		gl.DrawElements(gl.TRIANGLES, vao.vertexNum, gl.UNSIGNED_INT, gl.PtrOffset(0))
	}

	for i := range instanceDefs {
		gl.DisableVertexAttribArray(uint32(i + 3))
	}
}

func (vao *vao) Disable() {
	for _, index := range vao.attributes {
		gl.DisableVertexAttribArray(index)
	}
	vao.unbind()
}

func (vao *vao) addIndices(indices []uint32) {
	vbo := createVBO(vboElementArrayBufferType)
	vbo.Bind()
	vbo.storeUIntData(indices)
	vao.vertexNum = int32(len(indices))
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addIntAttrib(data []int32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := createVBO(vboArrayBufferType)
	vbo.Bind()
	vbo.storeIntData(data)
	gl.VertexAttribPointer(index, varsPerVertex, gl.INT, normalized, 0, nil)
	vbo.Unbind()
	vao.attributes = append(vao.attributes, index)
	vao.vbos = append(vao.vbos, vbo)
}

func (vao *vao) addFloatAttrib(data []float32, index uint32, varsPerVertex int32, normalized bool) {
	vbo := createVBO(vboArrayBufferType)
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
