package framework

import (
	"github.com/double-dev/limitengine/gmath"
)

type IShader interface {
	Start()
	Stop()

	LoadUniform1I(uniformName string, i int32)
	LoadUniform2I(uniformName string, i, j int32)
	LoadUniform3I(uniformName string, i, j, k int32)
	LoadUniform4I(uniformName string, i, j, k, l int32)
	LoadUniform1F(uniformName string, x float32)
	LoadUniform2F(uniformName string, v gmath.Vector2)
	LoadUniform3F(uniformName string, v gmath.Vector3)
	LoadUniform4F(uniformName string, v gmath.Vector4)

	LoadUniformMatrix4fv(uniformName string, matrix gmath.Matrix44)
	// TODO: Add other uniform methods.

	Delete()
}
