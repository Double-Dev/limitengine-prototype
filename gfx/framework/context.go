package framework

type Context interface {
	Resize(width, height int)

	ClearScreen(r, g, b, a float32)

	NewFramebuffer(colorAttachment, depthAttachment IAttachment, width, height float32, samples int32) IFramebuffer
	UnbindFramebuffers()

	NewRenderbuffer(multisample bool) IRenderbuffer
	NewShader(vertSrc, fragSrc string) IShader

	NewEmptyTexture() ITexture
	NewTexture(image []byte, width, height int32) ITexture
	UnbindTextures()

	NewMesh(indices []uint32, vertices, texCoords, normals []float32) IMesh
	NewInstanceBuffer(instanceDataSize int) IInstanceBuffer

	GetMaxInstances() int
}
