package limitengine

import (
	"runtime"
	"time"
)

// Version is the current engine version string.
const Version = "0.3.0"

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
	log     = NewLogger("core")
	running bool
	view    View
	state   State

	// TargetUpdatesPerSecond determine the amount of updates the engine will attempt to perform per second.
	TargetUpdatesPerSecond = float32(100.0)

	viewWidth, viewHeight int

	closeCallbacks       []func()
	resizeCallbacks      []func(width, height int)
	sampleCallbacks      []func(samples int)
	joystickCallbacks    []func(joy Joystick, event Action)
	keyCallbacks         []func(key Key, scancode int, action Action, mods ModKey)
	mouseButtonCallbacks []func(button MouseButton, action Action, mod ModKey)
	mouseMotionCallbacks []func(x, y float32)
	mouseScrollCallbacks []func(x, y float32)
	touchMotionCallbacks []func(x, y []float32)
	typingCallbacks      []func(char rune, mods ModKey)
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
		view.delete()
	})
	view.setResizeCallback(func(width, height int) {
		viewWidth = width
		viewHeight = height
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
	view.setTypingCallback(func(char rune, mods ModKey) {
		for _, typingCallback := range typingCallbacks {
			typingCallback(char, mods)
		}
	})
	// TODO: Handle touch input callbacks.

	running = true
	log.Log("Core online...")
}

// Launch begins the thread which runs the engine's update logic until the engine closes.
// This must be called on the main thread.
func Launch(initState State) {
	initState.OnActive()
	state = initState
	go func() {
		currentTime := time.Now().UnixNano()
		for Running() {
			if time.Now().UnixNano()-currentTime > int64((1.0/TargetUpdatesPerSecond)*1000000000.0) {
				lastTime := currentTime
				currentTime = time.Now().UnixNano()
				delta := float32(currentTime-lastTime) / 1000000000.0
				state.Update(delta)
			} else {
				time.Sleep(time.Millisecond)
			}
		}
	}()

	view.show()
	for Running() {
		view.pollEvents()
	}
}

// SetState sets a new application state for the engine to process.
func SetState(newState State) {
	newState.OnActive()
	oldState := state
	state = newState
	oldState.OnInactive()
}

// CurrentState returns the engine's current state.
func CurrentState() State { return state }

// AppView returns the engine's view.
func AppView() View { return view }

// Running returns whether the engine is running.
func Running() bool { return running }

// Width returns the current width (in pixels) of the view.
func Width() int { return viewWidth }

// Height returns the current height (in pixels) of the view.
func Height() int { return viewHeight }

// AspectRatio returns the current aspect ratio of the view.
func AspectRatio() float32 { return float32(viewHeight) / float32(viewWidth) }

// AddCloseCallback adds a close callback function to the application.
func AddCloseCallback(callback func()) {
	closeCallbacks = append(closeCallbacks, callback)
}

// AddResizeCallback adds a resize callback function to the application.
func AddResizeCallback(callback func(width, height int)) {
	resizeCallbacks = append(resizeCallbacks, callback)
}

// AddJoystickCallback adds a joystick callback function to the application.
func AddJoystickCallback(callback func(joy Joystick, event Action)) {
	joystickCallbacks = append(joystickCallbacks, callback)
}

// AddKeyCallback adds a key callback function to the application.
func AddKeyCallback(callback func(key Key, scancode int, action Action, mods ModKey)) {
	keyCallbacks = append(keyCallbacks, callback)
}

// AddMouseButtonCallback adds a mouse button callback function to the application.
func AddMouseButtonCallback(callback func(button MouseButton, action Action, mod ModKey)) {
	mouseButtonCallbacks = append(mouseButtonCallbacks, callback)
}

// AddMouseMotionCallback adds a mouse motion callback function to the application and
// returns the index of the new callbacks entry.
func AddMouseMotionCallback(callback func(x, y float32)) {
	mouseMotionCallbacks = append(mouseMotionCallbacks, callback)
}

// AddMouseScrollCallback adds a mouse scroll callback function to the application.
func AddMouseScrollCallback(callback func(x, y float32)) {
	mouseScrollCallbacks = append(mouseScrollCallbacks, callback)
}

// AddTouchMotionCallback adds a touch motion callback function to the application.
func AddTouchMotionCallback(callback func(x, y []float32)) {
	touchMotionCallbacks = append(touchMotionCallbacks, callback)
}

// AddTypingCallback adds a close typing function to the application.
func AddTypingCallback(callback func(char rune, mods ModKey)) {
	typingCallbacks = append(typingCallbacks, callback)
}
