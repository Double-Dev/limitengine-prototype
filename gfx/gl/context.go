package gl

import (
	"fmt"

	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type context struct{}

func NewContext() (context, error) {
	var err error
	if err = gl.Init(); err != nil {
		err = fmt.Errorf("Error starting OpenGL instance: %s", err)
	}
	// TODO: Add options for opengl features.
	gl.Enable(gl.MULTISAMPLE)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthMask(true)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return context{}, err
}

func (context context) Resize(width, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
}

func (context context) ClearScreen(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (context context) DepthTest(depthTest bool) {
	if depthTest {
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(gl.BACK)
	} else {
		gl.Disable(gl.CULL_FACE)
	}
}

func (context context) BackCulling(backCulling bool) {
	if backCulling {
		gl.Enable(gl.DEPTH_TEST)
	} else {
		gl.Disable(gl.DEPTH_TEST)
	}
}

func (context context) WriteDepth(writeDepth bool) {
	gl.DepthMask(writeDepth)
}

func (context context) NewFramebuffer(colorAttachment, depthAttachment framework.IAttachment, width, height float32, samples int32) framework.IFramebuffer {
	framebuffer := newFBO(width, height, samples)
	framebuffer.SetColorAttachment(colorAttachment)
	framebuffer.SetDepthAttachment(depthAttachment)
	return framebuffer
}

func (context context) UnbindFramebuffers() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
}

func (context context) NewRenderbuffer(multisample bool) framework.IRenderbuffer {
	renderbuffer := newRBO(multisample)
	return renderbuffer
}

func (context context) NewShader(vertSrc, fragSrc string) framework.IShader {
	shader := newShader(vertSrc, fragSrc)
	return shader
}

func (context context) NewEmptyTexture() framework.ITexture {
	texture := newTexture(textureType2D)
	texture.Bind()
	texture.LinearFilter(false, false)
	context.UnbindTextures()
	return texture
}

func (context context) NewTexture(image []uint8, width, height int32) framework.ITexture {
	texture := newTexture(textureType2D)
	texture.Bind()
	texture.TextureData(image, width, height)
	texture.LinearFilter(true, false)
	context.UnbindTextures()
	return texture
}

func (context context) UnbindTextures() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
}

func (context context) NewMesh(indices []uint32, vertices, texCoords, normals []float32) framework.IMesh {
	mesh := newVAO()
	mesh.bind()
	mesh.addIndices(indices)
	mesh.addFloatAttrib(vertices, 0, 3, false)
	if len(texCoords) > 0 {
		mesh.addFloatAttrib(texCoords, 1, 2, false)
	}
	if len(normals) > 0 {
		mesh.addFloatAttrib(normals, 2, 3, false)
	}
	// mesh.addInstancedFloatAttrib(
	// 	context.GetMaxInstances()*context.GetMaxInstanceData(),
	// 	3,
	// 	4,
	// 	int32(context.GetMaxInstanceData()),
	// 	false,
	// )
	mesh.unbind()
	return mesh
}

func (context context) NewInstanceBuffer(instanceDataSize int) framework.IInstanceBuffer {
	instanceBuffer := newVBO(vboArrayBufferType)
	instanceBuffer.Bind()
	instanceBuffer.setEmpty(context.GetMaxInstances() * instanceDataSize)
	instanceBuffer.Unbind()
	return instanceBuffer
}

func (context context) GetMaxInstances() int {
	return 10000
}
