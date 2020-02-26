package framework

type ITexture interface {
	Bind()
	Unbind()
	Delete()
}
