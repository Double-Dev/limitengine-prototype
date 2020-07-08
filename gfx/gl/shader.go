package gl

import (
	"fmt"
	"strings"

	"github.com/double-dev/limitengine/dependencies/gl/v3.3-core/gl"
)

const (
	vertexShaderType   = gl.VERTEX_SHADER
	fragmentShaderType = gl.FRAGMENT_SHADER
)

type shader struct {
	id       uint32
	vertID   uint32
	fragID   uint32
	uniforms map[uint32]map[string]int32
}

func newShader(vertSrc, fragSrc string) *shader {
	id := gl.CreateProgram()

	vertID, vertErr := compileShader(vertSrc, vertexShaderType)
	if vertErr != nil {
		panic(vertErr)
	}

	fragID, fragErr := compileShader(fragSrc, fragmentShaderType)
	if fragErr != nil {
		panic(fragErr)
	}

	gl.AttachShader(id, vertID)
	gl.AttachShader(id, fragID)
	gl.LinkProgram(id)

	shader := &shader{
		id:       id,
		vertID:   vertID,
		fragID:   fragID,
		uniforms: make(map[uint32]map[string]int32),
	}

	shader.Start()
	var numUniforms int32
	gl.GetProgramiv(shader.id, gl.ACTIVE_UNIFORMS, &numUniforms)
	for i := uint32(0); i < uint32(numUniforms); i++ {
		var name [256]byte
		var xtype uint32
		var length, size int32
		gl.GetActiveUniform(shader.id, i, 256, &length, &size, &xtype, &name[0])
		if shader.uniforms[xtype] == nil {
			shader.uniforms[xtype] = make(map[string]int32)
		}
		shader.uniforms[xtype][gl.GoStr(&name[0])] = gl.GetUniformLocation(shader.id, &name[0])
	}
	shader.Stop()

	return shader
}

func compileShader(src string, shaderType uint32) (uint32, error) {
	id := gl.CreateShader(shaderType)
	srcPtr, free := gl.Strs(src + "\x00")
	gl.ShaderSource(id, 1, srcPtr, nil)
	free()
	gl.CompileShader(id)

	var status int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(id, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("Shader failed to compile %v: %v", src, log)
	}

	return id, nil
}

func (shader *shader) Start() {
	gl.UseProgram(shader.id)
}

func (*shader) Stop() {
	gl.UseProgram(0)
}

func (shader *shader) LoadUniformTextureSampler2D(uniformName string, i int32) {
	gl.Uniform1i(shader.uniforms[gl.IMAGE_2D][uniformName], i)
}
func (shader *shader) LoadUniform1I(uniformName string, i int32) {
	gl.Uniform1i(shader.uniforms[gl.INT][uniformName], i)
}
func (shader *shader) LoadUniform2I(uniformName string, i, j int32) {
	gl.Uniform2i(shader.uniforms[gl.INT_VEC2][uniformName], i, j)
}
func (shader *shader) LoadUniform3I(uniformName string, i, j, k int32) {
	gl.Uniform3i(shader.uniforms[gl.INT_VEC3][uniformName], i, j, k)
}
func (shader *shader) LoadUniform4I(uniformName string, i, j, k, l int32) {
	gl.Uniform4i(shader.uniforms[gl.INT_VEC4][uniformName], i, j, k, l)
}
func (shader *shader) LoadUniform1F(uniformName string, x float32) {
	gl.Uniform1f(shader.uniforms[gl.FLOAT][uniformName], x)
}
func (shader *shader) LoadUniform2F(uniformName string, x, y float32) {
	gl.Uniform2f(shader.uniforms[gl.FLOAT_VEC2][uniformName], x, y)
}
func (shader *shader) LoadUniform3F(uniformName string, x, y, z float32) {
	gl.Uniform3f(shader.uniforms[gl.FLOAT_VEC3][uniformName], x, y, z)
}
func (shader *shader) LoadUniform4F(uniformName string, x, y, z, w float32) {
	gl.Uniform4f(shader.uniforms[gl.FLOAT_VEC4][uniformName], x, y, z, w)
}

func (shader *shader) LoadUniformMatrix4fv(uniformName string, matrix []float32) {
	gl.UniformMatrix4fv(shader.uniforms[gl.FLOAT_MAT4][uniformName], 1, false, &matrix[0])
}

// TODO: Add other uniform load functions.

func (shader *shader) Delete() {
	gl.DetachShader(shader.id, shader.vertID)
	gl.DeleteShader(shader.vertID)
	gl.DetachShader(shader.id, shader.fragID)
	gl.DeleteShader(shader.fragID)
	gl.DeleteProgram(shader.id)
}
