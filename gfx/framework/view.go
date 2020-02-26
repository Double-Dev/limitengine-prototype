package framework

type View interface {
	OnResume()
	OnPause()

	SetCloseCallback(func())
	SetResizeCallback(func(width, height int))

	SwapBuffers()
	UpdateEvents()

	GetSize() (int, int)
	Delete()
}
