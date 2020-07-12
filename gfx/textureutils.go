package gfx

import (
	"sync"

	"github.com/double-dev/limitengine/gmath"
)

var (
	textureBoundsPlugin = CreateLESLPlugin(`
vert{
	vars{
		instance vec4 textureBounds;
		out vec4 fragTextureBounds;
	},
	main{
		fragTextureBounds = textureBounds;
	},
},
frag{
	vars{
		in vec4 fragTextureBounds;
	},
	main{
		lesl.outColor = texture(lesl.texture, vec2(lesl.textureCoord.x * fragTextureBounds.z + fragTextureBounds.x, lesl.textureCoord.y * fragTextureBounds.w + fragTextureBounds.y));
	},
},`)
	textureTintPlugin = CreateLESLPlugin(`
frag{
	vars{
		uniform vec3 tintColor;
		uniform float tintAmount;
	},
	main{
		if (lesl.outColor.a < 0.001) {
			discard;
		} else {
			lesl.outColor = mix(lesl.outColor, vec4(tintColor, lesl.outColor.a), tintAmount);
		}
	},
},`)
)

func TextureBoundsPlugin() *LESLPlugin { return textureBoundsPlugin }
func TextureTintPlugin() *LESLPlugin   { return textureTintPlugin }

type SpriteSheet struct {
	spriteWidth, spriteHeight float32
	padding                   float32
	columns, rows             uint32
}

func NewSpriteSheet(spriteWidth, spriteHeight, padding float32) *SpriteSheet {
	spriteSheet := &SpriteSheet{
		spriteWidth:  spriteWidth,
		spriteHeight: spriteHeight,
		padding:      padding,
		columns:      uint32(1.0 / spriteWidth),
		rows:         uint32(1.0 / spriteHeight),
	}
	return spriteSheet
}

func (spriteSheet *SpriteSheet) Apply(instance *Instance, index uint32) {
	instance.ModifyData(
		"textureBounds",
		spriteSheet.spriteWidth*float32(index%spriteSheet.columns)+spriteSheet.padding,
		spriteSheet.spriteHeight*float32(index/spriteSheet.columns)+spriteSheet.padding,
		spriteSheet.spriteWidth-(2.0*spriteSheet.padding),
		spriteSheet.spriteHeight-(2.0*spriteSheet.padding),
	)
}

func (spriteSheet *SpriteSheet) GetBounds(index uint32) gmath.Vector4 {
	return spriteSheet.GetBoundsFlipped(index, false, false)
}

func (spriteSheet *SpriteSheet) GetBoundsFlipped(index uint32, xAxis, yAxis bool) gmath.Vector4 {
	bounds := gmath.NewVector4(
		spriteSheet.spriteWidth*float32(index%spriteSheet.columns)+spriteSheet.padding,
		spriteSheet.spriteHeight*float32(index/spriteSheet.columns)+spriteSheet.padding,
		spriteSheet.spriteWidth-(2.0*spriteSheet.padding),
		spriteSheet.spriteHeight-(2.0*spriteSheet.padding),
	)
	if xAxis {
		bounds[0] += spriteSheet.spriteWidth - (2.0 * spriteSheet.padding)
		bounds[2] *= -1.0
	}
	if yAxis {
		bounds[1] += spriteSheet.spriteHeight - (2.0 * spriteSheet.padding)
		bounds[3] *= -1.0
	}
	return bounds
}

type TextureAtlas struct {
	atlas map[string]gmath.Vector4
	mutex sync.RWMutex
}

func NewTextureAtlas() *TextureAtlas {
	return &TextureAtlas{
		atlas: make(map[string]gmath.Vector4),
	}
}

func NewTextureAtlasExisting(atlas map[string]gmath.Vector4) *TextureAtlas {
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
	defer textureAtlas.mutex.RUnlock()
	return textureAtlas.atlas[key]
}

func (textureAtlas *TextureAtlas) Apply(instance *Instance, key string) {
	instance.SetData("textureBounds", textureAtlas.Query(key))
}
