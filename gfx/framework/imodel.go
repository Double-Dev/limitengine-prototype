package framework

type IModel interface {
	Enable()
	Render()
	Disable()
	Delete()
}
