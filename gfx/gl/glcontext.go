package gl

import (
	"fmt"

	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type glContext struct{}

func NewGLContext() (glContext, error) {
	var err error
	if err = gl.Init(); err != nil {
		err = fmt.Errorf("Error starting OpenGL instance: %s", err)
	}
	// TODO: Add options for opengl features.
	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return glContext{}, err
}

func (glContext glContext) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (glContext glContext) ClearScreen(r, g, b, a float32) {
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (glContext glContext) CreateFrameBuffer() {

}

func (glContext glContext) CreateRenderBuffer() {

}

func (glContext glContext) CreateShader(vertSrc, fragSrc string) framework.IShader {
	shader := newShader(vertSrc, fragSrc)
	return shader
}

func (glContext glContext) CreateTexture(image []uint8, width, height int32) framework.ITexture {
	texture := newTexture(image, width, height)
	return texture
}

func (glContext glContext) CreateModel(indices []uint32, vertices, texCoords, normals []float32) framework.IModel {
	model := newVAO()
	model.bind()
	model.addIndices(indices)
	model.addFloatAttrib(vertices, 0, 3, false)
	if len(texCoords) > 0 {
		model.addFloatAttrib(texCoords, 1, 2, false)
	}
	if len(normals) > 0 {
		model.addFloatAttrib(normals, 2, 3, false)
	}
	// model.addInstancedFloatAttrib(
	// 	glContext.GetMaxInstances()*glContext.GetMaxInstanceData(),
	// 	3,
	// 	4,
	// 	int32(glContext.GetMaxInstanceData()),
	// 	false,
	// )
	model.unbind()
	return model
}

func (glContext glContext) CreateInstanceBuffer(instanceDataSize int) framework.IInstanceBuffer {
	instanceBuffer := newVBO(vboArrayBufferType)
	instanceBuffer.Bind()
	instanceBuffer.setEmpty(glContext.GetMaxInstances() * instanceDataSize)
	instanceBuffer.Unbind()
	return instanceBuffer
}

func (glContext glContext) GetMaxInstances() int {
	return 10000
}
