package utils2d

import (
	"github.com/double-dev/limitengine/gfx"
)

type SpriteShader struct {
	shader *gfx.Shader
}

func NewSpriteShader(leslPlugins ...*gfx.LESLPlugin) *SpriteShader {
	return &SpriteShader{gfx.NewShader(append([]*gfx.LESLPlugin{gfx.TextureBoundsPlugin(), gfx.TextureTintPlugin()}, leslPlugins...)...)}
}

func (spriteShader *SpriteShader) Shader() *gfx.Shader { return spriteShader.shader }
