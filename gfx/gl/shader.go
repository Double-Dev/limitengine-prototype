package gl

import (
	"doubledev/limitengine/gmath"

	"github.com/go-gl/gl/v3.2-core/gl"
)

const (
	vertexShaderType   = gl.VERTEX_SHADER
	fragmentShaderType = gl.FRAGMENT_SHADER
)

type shader struct {
	id       uint32
	vertID   uint32
	fragID   uint32
	uniforms map[string]int32
}

func newShader(vertSrc, fragSrc string) *shader {
	id := gl.CreateProgram()

	vertID := gl.CreateShader(vertexShaderType)
	vertPtr, free := gl.Strs(vertSrc + "\x00")
	gl.ShaderSource(vertID, 1, vertPtr, nil)
	free()
	gl.CompileShader(vertID)

	fragID := gl.CreateShader(fragmentShaderType)
	fragPtr, free := gl.Strs(fragSrc + "\x00")
	gl.ShaderSource(fragID, 1, fragPtr, nil)
	free()
	gl.CompileShader(fragID)

	gl.AttachShader(id, vertID)
	gl.AttachShader(id, fragID)
	gl.LinkProgram(id)

	shader := &shader{
		id:       id,
		vertID:   vertID,
		fragID:   fragID,
		uniforms: make(map[string]int32),
	}

	shader.Start()
	var count int32
	gl.GetProgramiv(shader.id, gl.ACTIVE_UNIFORMS, &count)
	for i := uint32(0); i < uint32(count); i++ {
		var name [256]byte
		gl.GetActiveUniformName(shader.id, i, 256, nil, &name[0])
		shader.uniforms[gl.GoStr(&name[0])] = gl.GetUniformLocation(shader.id, &name[0])
	}
	shader.Stop()

	return shader
}

func (shader *shader) Start() {
	gl.UseProgram(shader.id)
}

func (*shader) Stop() {
	gl.UseProgram(0)
}

func (shader *shader) LoadUniform1I(uniformName string, i int32) {
	gl.Uniform1i(shader.uniforms[uniformName], i)
}
func (shader *shader) LoadUniform2I(uniformName string, i, j int32) {
	gl.Uniform2i(shader.uniforms[uniformName], i, j)
}
func (shader *shader) LoadUniform3I(uniformName string, i, j, k int32) {
	gl.Uniform3i(shader.uniforms[uniformName], i, j, k)
}
func (shader *shader) LoadUniform4I(uniformName string, i, j, k, l int32) {
	gl.Uniform4i(shader.uniforms[uniformName], i, j, k, l)
}
func (shader *shader) LoadUniform1F(uniformName string, x int32) {
	gl.Uniform1i(shader.uniforms[uniformName], x)
}
func (shader *shader) LoadUniform2F(uniformName string, x, y int32) {
	gl.Uniform2i(shader.uniforms[uniformName], x, y)
}
func (shader *shader) LoadUniform3F(uniformName string, x, y, z int32) {
	gl.Uniform3i(shader.uniforms[uniformName], x, y, z)
}
func (shader *shader) LoadUniform4F(uniformName string, x, y, z, w int32) {
	gl.Uniform4i(shader.uniforms[uniformName], x, y, z, w)
}

func (shader *shader) LoadUniformMatrix4fv(uniformName string, matrix gmath.Matrix) {
	gl.UniformMatrix4fv(shader.uniforms[uniformName], 1, false, &matrix.ToArray()[0])
}

// TODO: Add other uniform load functions.

func (shader *shader) Delete() {
	gl.DetachShader(shader.id, shader.vertID)
	gl.DeleteShader(shader.vertID)
	gl.DetachShader(shader.id, shader.fragID)
	gl.DeleteShader(shader.fragID)
	gl.DeleteProgram(shader.id)
}
