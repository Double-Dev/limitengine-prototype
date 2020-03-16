package gfx

import (
	"strings"

	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	shaderIndex = uint32(1)
	shaders     = make(map[uint32]framework.IShader)
)

func init() { shaders[0] = nil }

type Shader struct {
	id            uint32
	uniformLoader uniformLoader
}

func CreateShader(vertSrc, fragSrc string) *Shader {
	shader := &Shader{
		id: shaderIndex,
	}
	shaderIndex++
	actionQueue = append(actionQueue, func() { shaders[shader.id] = context.CreateShader(vertSrc, fragSrc) })
	return shader
}

func processVertShader(vertSrc string) string { // TODO: Parse custom shader
	header := `#version 330 core
layout(location = 0) in vec3 vertPosition;
layout(location = 1) in vec2 vertTextureCoord;
layout(location = 2) in vec3 vertNormal;
uniform mat4 projMat;
uniform mat4 viewMat;
uniform mat4 transformMat;
out vec3 fragPosition;
out vec2 fragTextureCoord;
out vec3 fragNormal;
vec4 worldPos;
`
	footer := `void main()
{
	worldPos = transformMat * vec4(coord, 1.0);
	gl_Position = projMat * viewMat * worldPos;
	vertTextureCoord = vec2(texCoord.x, 1.0 - texCoord.y);
	fragTextureCoord = vertTextureCoord; 
	fragNormal = vertNormal;
	vert();
}`
	vertSrc = strings.ReplaceAll(vertSrc, "this.position", "vertPosition")
	vertSrc = strings.ReplaceAll(vertSrc, "this.textureCoord", "vertTextureCoord")
	vertSrc = strings.ReplaceAll(vertSrc, "this.normal", "vertNormal")
	return header + vertSrc + footer
}

// TODO: Add support for more variables + array uniforms.
type uniformLoader struct {
	uniformInts      map[string]int32
	uniformMatrix44s map[string]gmath.Matrix
}

func newUniformLoader() uniformLoader {
	return uniformLoader{
		uniformInts:      make(map[string]int32),
		uniformMatrix44s: make(map[string]gmath.Matrix),
	}
}

func (this uniformLoader) loadTo(iShader framework.IShader) {
	for varName, value := range this.uniformInts {
		iShader.LoadUniform1I(varName, value)
	}
	for varName, value := range this.uniformMatrix44s {
		iShader.LoadUniformMatrix4fv(varName, value.ToArray())
	}
}

func (this uniformLoader) AddInt(varName string, val int32) {
	this.uniformInts[varName] = val
}

func (this uniformLoader) AddMatrix44(varName string, val gmath.Matrix) {
	if val.IsSize(4, 4) {
		this.uniformMatrix44s[varName] = val
	}
}
