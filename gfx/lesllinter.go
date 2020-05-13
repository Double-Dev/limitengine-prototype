package gfx

import (
	"fmt"
	"strings"
)

const (
	vertHeader = `#version 330 core
layout(location = 0) in vec3 vertposition;
layout(location = 1) in vec2 verttextureCoord;
layout(location = 2) in vec3 vertnormal;
uniform mat4 vertprojMat;
uniform mat4 vertviewMat;

layout(location = 3) in vec4 verttransformMat0;
layout(location = 4) in vec4 verttransformMat1;
layout(location = 5) in vec4 verttransformMat2;
layout(location = 6) in vec4 verttransformMat3;
mat4 verttransformMat;

out vec3 fragposition;
out vec2 fragtextureCoord;
out vec3 fragnormal;
vec4 vertworldPos;
`
	vertFooter = `
void main()
{
	verttransformMat = mat4(verttransformMat0, verttransformMat1, verttransformMat2, verttransformMat3);
	vertworldPos = verttransformMat * vec4(vertposition, 1.0);
	gl_Position = vertprojMat * vertviewMat * vertworldPos;
	fragposition = vertposition;
	fragtextureCoord = verttextureCoord;
	fragnormal = vertnormal;
	vert();
}`

	fragHeader = `#version 330 core
in vec3 fragposition;
in vec2 fragtextureCoord;
in vec3 fragnormal;
out vec4 fragoutColor;
`
	fragFooter = `
void main()
{
	frag();
	if (fragoutColor.a < 0.01) {
		discard;
	}
}`
)

var (
	vertReservedVars = []string{
		"vertposition", "verttextureCoord", "vertnormal",
		"vertprojMat", "vertviewMat", "verttransformMat", "fragPosition",
		"fragTextureCoord", "fragNormal", "vertworldPos",
	}
	fragReservedVars = []string{
		"fragposition", "fragtextureCoord", "fragnormal", "fragoutColor",
	}
)

func processLESL(src string) (string, string, map[string]int32) { // TODO: Parse custom shader
	textureVars := make(map[string]int32)
	for i := int32(0); i < 10; i++ {
		if strings.Contains(src, fmt.Sprintf("texture%d", i)) {
			varNameStart := strings.Index(src, fmt.Sprintf("texture%d", i)) + 9
			varNameEnd := strings.Index(src[varNameStart:], ";")
			textureVars[src[varNameStart:varNameStart+varNameEnd]] = i
			src = strings.Replace(src, fmt.Sprintf("texture%d", i), "uniform sampler2D", 1)
			if strings.Contains(src, fmt.Sprintf("texture%d", i)) {
				log.ForceErr(fmt.Sprintf("LESL: Multiple uses of type 'texture%d' not allowed.", i))
			}
		}
	}
	fmt.Println(textureVars)

	vertStart := strings.Index(src, "#vert")
	vertEnd := len(src[:vertStart]) + strings.Index(src[vertStart:], "#end")

	vertCode := src[vertStart+5 : vertEnd]
	if strings.Contains(vertCode, "#") {
		log.ForceErr("LESL: Invalid block declaration in vertex code.")
	}
	for _, varName := range vertReservedVars {
		vertCodeCopy := vertCode
		for strings.Contains(vertCodeCopy, varName) {
			if vertCode[strings.Index(vertCodeCopy, varName)-5:strings.Index(vertCodeCopy, varName)] != "lesl." {
				log.ForceErr("LESL: Variable name '" + varName + "' not allowed in vertex code.")
			}
			vertCodeCopy = vertCodeCopy[strings.Index(vertCodeCopy, varName)+len(varName):]
		}
	}
	vertCode = vertHeader + strings.ReplaceAll(vertCode, "lesl.", "vert") + vertFooter
	src = src[:vertStart] + src[vertEnd+4:]

	fragStart := strings.Index(src, "#frag")
	fragEnd := len(src[:fragStart]) + strings.Index(src[fragStart:], "#end")

	fragCode := src[fragStart+5 : fragEnd]
	if strings.Contains(fragCode, "#") {
		log.ForceErr("LESL: Invalid block declaration in fragment code.")
	}
	for _, varName := range fragReservedVars {
		fragCodeCopy := fragCode
		for strings.Contains(fragCodeCopy, varName) {
			if fragCode[strings.Index(fragCodeCopy, varName)-5:strings.Index(fragCodeCopy, varName)] != "lesl." {
				log.ForceErr("LESL: Variable name '" + varName + "' not allowed in fragment code.")
			}
			fragCodeCopy = fragCodeCopy[strings.Index(fragCodeCopy, varName)+len(varName):]
		}
	}
	fragCode = fragHeader + strings.ReplaceAll(fragCode, "lesl.", "frag") + fragFooter
	src = src[:fragStart] + src[fragEnd+4:]

	// fmt.Println("VERTEX CODE:\n" + vertCode)
	// fmt.Println("FRAGMENT CODE:\n" + fragCode)

	return vertCode, fragCode, textureVars
}
