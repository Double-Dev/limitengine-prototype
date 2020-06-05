package framework

type Context interface {
	Resize(width, height int)

	ClearScreen(r, g, b, a float32)
	CreateFramebuffer() IFramebuffer
	CreateRenderbuffer() IRenderbuffer
	CreateShader(vertSrc, fragSrc string) IShader

	CreateEmptyTexture() ITexture
	CreateTexture(image []byte, width, height int32) ITexture

	CreateMesh(indices []uint32, vertices, texCoords, normals []float32) IMesh
	CreateInstanceBuffer(instanceDataSize int) IInstanceBuffer

	GetMaxInstances() int
}
