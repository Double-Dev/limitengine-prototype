package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
)

type SpriteShader struct {
	renderProgram *gfx.RenderProgram
	uniformLoader gfx.UniformLoader
}

func NewSpriteShader(leslPlugins ...*gfx.LESLPlugin) *SpriteShader {
	return &SpriteShader{
		gfx.NewRenderProgram(append([]*gfx.LESLPlugin{gfx.TextureBoundsPlugin(), gfx.TextureTintPlugin()}, leslPlugins...)...),
		gfx.NewUniformLoader(),
	}
}

func (shader *SpriteShader) RenderProgram() *gfx.RenderProgram { return shader.renderProgram }
func (shader *SpriteShader) UniformLoader() gfx.UniformLoader  { return shader.uniformLoader }
