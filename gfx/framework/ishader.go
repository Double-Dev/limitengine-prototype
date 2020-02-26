package framework

import (
	"doubledev/limitengine/gmath"
)

type IShader interface {
	Start()
	Stop()

	LoadUniform1I(uniformName string, i int32)
	LoadUniform2I(uniformName string, i, j int32)
	LoadUniform3I(uniformName string, i, j, k int32)
	LoadUniform4I(uniformName string, i, j, k, l int32)
	LoadUniform1F(uniformName string, x int32)
	LoadUniform2F(uniformName string, x, y int32)
	LoadUniform3F(uniformName string, x, y, z int32)
	LoadUniform4F(uniformName string, x, y, z, w int32)

	LoadUniformMatrix4fv(uniformName string, matrix gmath.Matrix)
	// TODO: Add other uniform methods.

	Delete()
}
