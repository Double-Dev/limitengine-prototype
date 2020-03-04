package gl

import (
	"fmt"
	"strings"

	"github.com/double-dev/limitengine/gmath"
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
		uniforms: make(map[string]int32),
	}

	shader.Start()
	var numUniforms int32
	gl.GetProgramiv(shader.id, gl.ACTIVE_UNIFORMS, &numUniforms)
	for i := uint32(0); i < uint32(numUniforms); i++ {
		var name [256]byte
		gl.GetActiveUniformName(shader.id, i, 256, nil, &name[0])
		shader.uniforms[gl.GoStr(&name[0])] = gl.GetUniformLocation(shader.id, &name[0])
	}
	// var numAttribs int32
	// gl.GetProgramiv(shader.id, gl.ACTIVE_ATTRIBUTES, &numAttribs)
	// for i := uint32(0); i < uint32(numAttribs); i++ {
	// 	var name [256]byte
	// 	gl.GetActiveAttrib(shader.id, i, 256, nil, nil, nil, &name[0])
	// 	gl.BindAttribLocation(shader.id, i, &name[0])
	// }
	shader.Stop()
	for key, num := range shader.uniforms {
		fmt.Println(key, num)
	}

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
