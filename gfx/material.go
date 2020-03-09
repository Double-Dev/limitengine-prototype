package gfx

type Material struct {
	texture *Texture
	prefs   uniformLoader
}

func CreateMaterial() *Material {
	return &Material{
		texture: nil,
		prefs:   newUniformLoader(),
	}
}

func CreateTextureMaterial(texture *Texture) *Material {
	return &Material{
		texture: texture,
		prefs:   newUniformLoader(),
	}
}
