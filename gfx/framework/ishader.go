package framework

type IShader interface {
	Start()
	Stop()

	LoadUniform1I(uniformName string, i int32)
	LoadUniform2I(uniformName string, i, j int32)
	LoadUniform3I(uniformName string, i, j, k int32)
	LoadUniform4I(uniformName string, i, j, k, l int32)
	LoadUniform1F(uniformName string, x float32)
	LoadUniform2F(uniformName string, x, y float32)
	LoadUniform3F(uniformName string, x, y, z float32)
	LoadUniform4F(uniformName string, x, y, z, w float32)

	LoadUniformMatrix4fv(uniformName string, matrix []float32)
	// TODO: Add other uniform methods.

	Delete()
}
