package gfx

// TODO: Think about turning Material into interface to allow easy creation of variations.
type Material struct {
	texture      *Texture
	prefs        uniformLoader
	Transparency bool
}

func CreateMaterial() *Material {
	return &Material{
		texture:      nil,
		prefs:        newUniformLoader(),
		Transparency: false,
	}
}

func CreateTextureMaterial(texture *Texture) *Material {
	return &Material{
		texture:      texture,
		prefs:        newUniformLoader(),
		Transparency: true,
	}
}
