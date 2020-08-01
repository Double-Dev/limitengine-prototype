package limitengine

import (
	"image"
	"runtime"

	"github.com/double-dev/limitengine/dependencies/glfw/v3.3/glfw"
)

type glfwView struct {
	window *glfw.Window
}

func newGLFWView() *glfwView {
	if err := glfw.Init(); err != nil {
		log.Err("Error initializing GLFW:", err)
	}
	glfw.WindowHint(glfw.Visible, glfw.False)
	if windowResizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	if runtime.GOOS == DARWIN {
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	}
	window, err := glfw.CreateWindow(initWidth, initHeight, windowTitle, nil, nil)
	if err != nil {
		log.Err("Error creating window.", err)
	}
	glfwView := &glfwView{
		window: window,
	}
	return glfwView
}

func (glfwView *glfwView) SetContext() {
	glfwView.window.MakeContextCurrent()
	glfw.SwapInterval(1)
}

func (glfwView *glfwView) ReleaseContext() {
	glfw.DetachCurrentContext()
}

func (glfwView *glfwView) show() {
	glfwView.window.Show()
}

func (*glfwView) pollEvents() {
	glfw.WaitEvents()
}

func (glfwView *glfwView) SwapBuffers() {
	glfwView.window.SwapBuffers()
}

func (glfwView *glfwView) SetPosition(x, y int) {
	glfwView.window.SetPos(x, y)
}

func (glfwView *glfwView) SetSize(width, height int) {
	glfwView.window.SetSize(width, height)
}

func (glfwView *glfwView) SetAspectRatio(numer, denom int) {
	glfwView.window.SetAspectRatio(numer, denom)
}

func (glfwView *glfwView) SetTitle(title string) {
	glfwView.window.SetTitle(title)
}

func (glfwView *glfwView) SetIcons(icons []image.Image) {
	glfwView.window.SetIcon(icons)
}

func (glfwView *glfwView) setCloseCallback(callback func()) {
	glfwView.window.SetCloseCallback(func(w *glfw.Window) {
		callback()
	})
}

func (glfwView *glfwView) setResizeCallback(callback func(width, height int)) {
	glfwView.window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		callback(width, height)
	})
}

func (glfwView *glfwView) setJoystickCallback(callback func(joy Joystick, event Action)) {
	// TODO: Implement joystick input.
}

func (glfwView *glfwView) setKeyCallback(callback func(key Key, scancode int, action Action, mods ModKey)) {
	glfwView.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		callback(Key(key), scancode, Action(action), ModKey(mods))
	})
}

func (glfwView *glfwView) setMouseButtonCallback(callback func(button MouseButton, action Action, mod ModKey)) {
	glfwView.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		callback(MouseButton(int(button)+1), Action(action), ModKey(mod))
	})
}

func (glfwView *glfwView) setMouseMotionCallback(callback func(x, y float32)) {
	glfwView.window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		callback(float32(x), float32(y))
	})
}

func (glfwView *glfwView) setMouseScrollCallback(callback func(x, y float32)) {
	glfwView.window.SetScrollCallback(func(w *glfw.Window, x, y float64) {
		callback(float32(x), float32(y))
	})
}

func (glfwView *glfwView) setTouchMotionCallback(callback func(x, y []float32)) {
	// TODO: Figure out how to separate this from mouse motion cross-platform.
}

func (glfwView *glfwView) setTypingCallback(callback func(char rune, mods ModKey)) {
	glfwView.window.SetCharModsCallback(func(w *glfw.Window, char rune, mods glfw.ModifierKey) {
		callback(char, ModKey(mods))
	})
}

func (glfwView *glfwView) delete() {
	glfwView.window.SetShouldClose(true)
	glfw.Terminate()
}
