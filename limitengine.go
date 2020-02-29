package limitengine

import (
	"runtime"
)

// Version is the current engine version string.
const Version = "0.0.1"

// Const variables that store supported platform names.
const (
	ANDROID   = "android"
	DARWIN    = "darwin"
	DRAGONFLY = "dragonfly"
	FREEBSD   = "freebsd"
	ILLUMOS   = "illumos"
	JS        = "js"
	LINUX     = "linux"
	NETBSD    = "netbsd"
	OPENBSD   = "openbsd"
	PLAN9     = "plan9"
	SOLARIS   = "solaris"
	WINDOWS   = "windows"
)

var (
	log                  = NewLogger("core")
	running              bool
	view                 View
	closeCallbacks       []func()
	resizeCallbacks      []func(width, height int)
	joystickCallbacks    []func(joy Joystick, event Action)
	keyCallbacks         []func(key Key, scancode int, action Action, mods ModKey)
	mouseButtonCallbacks []func(button MouseButton, action Action, mod ModKey)
	mouseMotionCallbacks []func(x, y float32)
	mouseScrollCallbacks []func(x, y float32)
	touchMotionCallbacks []func(x, y []float32)
)

func init() {
	runtime.LockOSThread()
	switch runtime.GOOS {
	case ANDROID:
		log.ForceErr("Android is not yet supported.")
	case DARWIN, DRAGONFLY, FREEBSD, LINUX, WINDOWS:
		view = newGLFWView()
	case JS, NETBSD:
		log.ForceErr("Browser applications are not yet supported.")
	default:
		log.ForceErr("The host OS is not yet supported.")
	}
	view.setCloseCallback(func() {
		running = false
		for _, closeCallback := range closeCallbacks {
			closeCallback()
		}
	})
	view.setResizeCallback(func(width, height int) {
		for _, resizeCallback := range resizeCallbacks {
			resizeCallback(width, height)
		}
	})
	// TODO: Handle joystick callbacks.
	view.setKeyCallback(func(key Key, scancode int, action Action, mods ModKey) {
		for _, keyCallback := range keyCallbacks {
			keyCallback(key, scancode, action, mods)
		}
	})
	view.setMouseButtonCallback(func(button MouseButton, action Action, mods ModKey) {
		for _, mouseButtonCallback := range mouseButtonCallbacks {
			mouseButtonCallback(button, action, mods)
		}
	})
	view.setMouseMotionCallback(func(x, y float32) {
		for _, mouseMotionCallback := range mouseMotionCallbacks {
			mouseMotionCallback(x, y)
		}
	})
	view.setMouseScrollCallback(func(x, y float32) {
		for _, mouseScrollCallback := range mouseScrollCallbacks {
			mouseScrollCallback(x, y)
		}
	})
	// TODO: Handle touch input callbacks.
	running = true
	log.Log("Core online...")
}

// Launch runs the core's Run() func until the engine closes and must be called on the main thread.
func Launch() {
	view.show()
	for Running() {
		view.pollEvents()
	}
}

// AppView returns the engine's view.
func AppView() View {
	return view
}

// Running returns whether the engine is running.
func Running() bool {
	return running
}

func AddCloseCallback(callback func()) { closeCallbacks = append(closeCallbacks, callback) }
func AddResizeCallback(callback func(width, height int)) {
	resizeCallbacks = append(resizeCallbacks, callback)
}
func AddJoystickCallback(callback func(joy Joystick, event Action)) {
	joystickCallbacks = append(joystickCallbacks, callback)
}
func AddKeyCallback(callback func(key Key, scancode int, action Action, mods ModKey)) {
	keyCallbacks = append(keyCallbacks, callback)
}
func AddMouseButtonCallback(callback func(button MouseButton, action Action, mod ModKey)) {
	mouseButtonCallbacks = append(mouseButtonCallbacks, callback)
}
func AddMouseMotionCallback(callback func(x, y float32)) {
	mouseMotionCallbacks = append(mouseMotionCallbacks, callback)
}
func AddMouseScrollCallback(callback func(x, y float32)) {
	mouseScrollCallbacks = append(mouseScrollCallbacks, callback)
}
func AddTouchMotionCallback(callback func(x, y []float32)) {
	touchMotionCallbacks = append(touchMotionCallbacks, callback)
}
