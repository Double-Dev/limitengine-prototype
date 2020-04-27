package gfx

type VertexLESL string
type FragmentLESL string

const (
	vertHeader = `
#version 330 core
layout(location = 0) in vec3 vertPosition;
layout(location = 1) in vec2 vertTextureCoord;
layout(location = 2) in vec3 vertNormal;
uniform mat4 projMat;
uniform mat4 viewMat;

layout(location = 3) in vec4 transformMat;

out vec3 fragPosition;
out vec2 fragTextureCoord;
out vec3 fragNormal;
vec4 worldPos;
`
	vertFooter = `
void main()
{
	worldPos = transformMat * vec4(coord, 1.0);
	gl_Position = projMat * viewMat * worldPos;
	vertTextureCoord = vec2(texCoord.x, 1.0 - texCoord.y);
	fragTextureCoord = vertTextureCoord; 
	fragNormal = vertNormal;
	vert();
}`

	fragHeader = `
#version 330 core
in vec3 fragPosition;
in vec2 fragTextureCoord;
in vec3 fragNormal;
out outColor;
`
	fragFooter = `
void main()
{
	frag();
}`
)

func processLESL(src string) Shader { // TODO: Parse custom shader

	return *(*Shader)(nil)
}
