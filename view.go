package limitengine

// View represents the visual component opened by the host OS to designate the program's run state.
type View interface {
	SetContext()
	ReleaseContext()

	show()
	pollEvents()
	SwapBuffers()

	setCloseCallback(func())
	setResizeCallback(func(width, height int))

	setJoystickCallback(func(joy Joystick, event Action))
	setKeyCallback(func(key Key, scancode int, action Action, mods ModKey))
	setMouseButtonCallback(func(button MouseButton, action Action, mod ModKey))
	setMouseMotionCallback(func(x, y float32))
	setMouseScrollCallback(func(x, y float32))
	setTouchMotionCallback(func(x, y []float32))

	delete()
}

// Action represents an input action.
type Action int

// Joystick represents a joystick.
type Joystick int

// Key represents a keyboard key.
type Key int

// ModKey represents a keyboard modifier key.
type ModKey int

// MouseButton represents a mouse button.
type MouseButton int
