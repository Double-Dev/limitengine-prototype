package framework

type Context interface {
	Resize(width, height int)

	ClearScreen(r, g, b, a float32)
	CreateFrameBuffer()
	CreateRenderBuffer()
	CreateShader(vertSrc, fragSrc string) IShader
	CreateModel(indices []uint32, vertices, texCoords, normals []float32) IModel
	CreateTexture(image []byte, width, height int32) ITexture
}
