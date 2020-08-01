package gfx

import (
	"fmt"
	"strings"

	"github.com/double-dev/limitengine/gfx/framework"
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
	vertMain = `
void main()
{
	verttransformMat = mat4(verttransformMat0, verttransformMat1, verttransformMat2, verttransformMat3);
	vertworldPos = verttransformMat * vec4(vertposition, 1.0);
	gl_Position = vertprojMat * vertviewMat * vertworldPos;
	fragposition = vertposition;
	fragtextureCoord = verttextureCoord;
	fragnormal = vertnormal;
`

	fragHeader = `#version 330 core
in vec3 fragposition;
in vec2 fragtextureCoord;
in vec3 fragnormal;
out vec4 fragoutColor;

uniform sampler2D fragtexture;
`
	fragMain = `
void main()
{
`
	footer = `}`

	vertex   = 0
	fragment = 1

	vars  = 0
	funcs = 1
	main  = 2
)

var (
	vertReservedVars = []string{
		"vertposition", "verttextureCoord", "vertnormal",
		"verttransformMat0", "verttransformMat1", "verttransformMat2", "verttransformMat3",
		"vertprojMat", "vertviewMat", "verttransformMat",
		"fragPosition", "fragTextureCoord", "fragNormal",
		"vertworldPos",
	}
	fragReservedVars = []string{
		"fragposition", "fragtextureCoord", "fragnormal", "fragoutColor",
		"fragtexture",
	}
)

func CreateLESLPlugin(src string, dependencies ...*LESLPlugin) *LESLPlugin {
	leslPlugin := &LESLPlugin{textures: make(map[string]int32)}

	for _, dependency := range dependencies {
		if !leslPlugin.hasDependency(dependency) {
			leslPlugin.dependencies = append(leslPlugin.dependencies, dependency)
		}
	}

	layer := 0
	var shaderType int
	var shaderSection int
	for _, line := range strings.Split(src, "\n") {
		for i := 0; i < layer; i++ {
			line = strings.TrimPrefix(line, "    ")
		}
		if strings.Contains(line, "},") {
			layer--
		}
		if layer == 0 {
			if strings.Contains(line, "vert{") {
				layer++
				shaderType = vertex
			} else if strings.Contains(line, "frag{") {
				layer++
				shaderType = fragment
			}
		} else if layer == 1 {
			if strings.Contains(line, "vars{") {
				layer++
				shaderSection = vars
			} else if strings.Contains(line, "funcs{") {
				layer++
				shaderSection = funcs
			} else if strings.Contains(line, "main{") {
				layer++
				shaderSection = main
			}
		} else if layer == 2 {
			if shaderSection == vars {
				if shaderType == vertex {
					leslPlugin.vertVars += line + "\n"
				} else if shaderType == fragment {
					leslPlugin.fragVars += line + "\n"
				}
			} else if shaderSection == funcs {
				if shaderType == vertex {
					leslPlugin.vertFuncs += line + "\n"
				} else if shaderType == fragment {
					leslPlugin.fragFuncs += line + "\n"
				}
			} else if shaderSection == main {
				if shaderType == vertex {
					leslPlugin.vertMain += "    " + line + "\n"
				} else if shaderType == fragment {
					leslPlugin.fragMain += "    " + line + "\n"
				}
			}
		}
	}

	leslPlugin.vertVars = processTextures(leslPlugin.vertVars, leslPlugin.textures)
	leslPlugin.vertFuncs = strings.ReplaceAll(leslPlugin.vertFuncs, "lesl.", "vert")
	leslPlugin.vertMain = strings.ReplaceAll(leslPlugin.vertMain, "lesl.", "vert")

	leslPlugin.fragVars = processTextures(leslPlugin.fragVars, leslPlugin.textures)
	leslPlugin.fragFuncs = strings.ReplaceAll(leslPlugin.fragFuncs, "lesl.", "frag")
	leslPlugin.fragMain = strings.ReplaceAll(leslPlugin.fragMain, "lesl.", "frag")

	return leslPlugin
}

func processTextures(src string, textureVars map[string]int32) string {
	for i := int32(1); i < 10; i++ {
		keyword := fmt.Sprintf("texture%d", i)
		if strings.Contains(src, keyword) {
			varNameStart := strings.Index(src, keyword) + 9
			varNameEnd := strings.Index(src[varNameStart:], ";")
			textureVars[src[varNameStart:varNameStart+varNameEnd]] = i
			src = strings.Replace(src, keyword, "uniform sampler2D", 1)
			if strings.Contains(src, keyword) {
				log.ForceErr("LESL: Multiple uses of type '" + keyword + "' not allowed.")
			}
		}
	}
	// fmt.Println(textureVars)
	return src
}

func processReservedVars(src string, reservedVars []string) {
	for _, varName := range reservedVars {
		srcCopy := src
		for strings.Contains(srcCopy, varName) {
			if src[strings.Index(srcCopy, varName)-5:strings.Index(srcCopy, varName)] != "lesl." {
				log.ForceErr("LESL: Variable name '" + varName + "' not allowed in vertex code.")
			}
			srcCopy = srcCopy[strings.Index(srcCopy, varName)+len(varName):]
		}
	}
}

func processLESL(leslPlugins []*LESLPlugin) (string, string, []framework.InstanceDef, map[string]int32) {
	vertVars := vertHeader
	var vertFuncs string
	vertMain := vertMain

	fragVars := fragHeader
	var fragFuncs string
	fragMain := fragMain

	var rawDependencies []*LESLPlugin
	for _, leslPlugin := range leslPlugins {
		rawDependencies = append(rawDependencies, leslPlugin.getDependencies()...)
	}

	textures := make(map[string]int32)

	var dependencies []*LESLPlugin
	keys := make(map[*LESLPlugin]bool)
	for _, dependency := range rawDependencies {
		if _, ok := keys[dependency]; !ok {
			keys[dependency] = true
			dependencies = append(dependencies, dependency)
		}
	}

	for _, dependency := range dependencies {
		vertVars += dependency.vertVars
		vertFuncs += dependency.vertFuncs
		vertMain += dependency.vertMain
		fragVars += dependency.fragVars
		fragFuncs += dependency.fragFuncs
		fragMain += dependency.fragMain
		for key, value := range dependency.textures {
			textures[key] = value
		}
	}

	var instanceDefs []framework.InstanceDef

	inVarLocation := 7
	instanceDefIndex := 16
	instanceIndex := strings.Index(vertVars, "instance")
	for instanceIndex != -1 {
		endIndex := strings.Index(vertVars[instanceIndex:], ";")
		instanceVarLine := strings.Split(vertVars[instanceIndex:][:endIndex], " ")
		var instanceVarSize int
		switch instanceVarLine[1] {
		case "float":
			instanceVarSize = 1
			break
		case "vec2":
			instanceVarSize = 2
			break
		case "vec3":
			instanceVarSize = 3
			break
		case "vec4":
			instanceVarSize = 4
			break
		default:
			log.ForceErr("LESL does not support instance variables of type: " + instanceVarLine[1])
		}

		instanceDefs = append(instanceDefs, framework.InstanceDef{
			Name: instanceVarLine[2], Size: instanceVarSize, Index: instanceDefIndex,
		})
		instanceDefIndex += instanceVarSize

		vertVars = strings.Replace(vertVars, "instance", fmt.Sprint("layout(location =", inVarLocation)+") in", 1)
		inVarLocation++

		instanceIndex = strings.Index(vertVars, "instance")
	}

	// fmt.Println("VERTEX SHADER SRC:\n", vertVars+vertFuncs+vertMain+"}")
	// fmt.Println("FRAGMENT SHADER SRC:\n", fragVars+fragFuncs+fragMain+"}")

	return vertVars + vertFuncs + vertMain + "}", fragVars + fragFuncs + fragMain + "}", instanceDefs, textures
}

type LESLPlugin struct {
	dependencies                  []*LESLPlugin
	vertVars, vertFuncs, vertMain string
	fragVars, fragFuncs, fragMain string
	textures                      map[string]int32
}

func (leslPlugin *LESLPlugin) hasDependency(target *LESLPlugin) bool {
	for _, dependency := range leslPlugin.dependencies {
		if dependency == target || dependency.hasDependency(target) {
			return true
		}
	}
	return false
}

func (leslPlugin *LESLPlugin) getDependencies() []*LESLPlugin {
	var dependencies []*LESLPlugin
	for _, dependency := range leslPlugin.dependencies {
		dependencies = append(dependencies, dependency.getDependencies()...)
	}
	dependencies = append(dependencies, leslPlugin)
	return dependencies
}
