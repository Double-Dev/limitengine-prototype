package utils2d

import (
	"sync"

	"github.com/double-dev/limitengine/gmath"
)

type TextureAtlas struct {
	atlas map[string]gmath.Vector4
	mutex sync.RWMutex
}

func CreateTextureAtlas() *TextureAtlas {
	return &TextureAtlas{
		atlas: make(map[string]gmath.Vector4),
	}
}

func (textureAtlas *TextureAtlas) Add(key string, bounds gmath.Vector4) {
	textureAtlas.mutex.Lock()
	textureAtlas.atlas[key] = bounds
	textureAtlas.mutex.Unlock()
}

func (textureAtlas *TextureAtlas) Query(key string) gmath.Vector4 {
	textureAtlas.mutex.RLock()
	bounds := textureAtlas.atlas[key]
	textureAtlas.mutex.RUnlock()
	return bounds
}
