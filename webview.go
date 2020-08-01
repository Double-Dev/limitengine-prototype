package limitengine

import (
	"image"

	"github.com/gopherjs/gopherjs/js"
)

type webView struct {
	canvas *js.Object
}

func newWebView() *webView {
	webView := &webView{
		canvas: js.Global.Get("document").Call("querySelector", "#limitengineView"),
	}
	return webView
}

func (webView *webView) SetContext() {}

func (webView *webView) ReleaseContext() {}

func (webView *webView) show() {}

func (*webView) pollEvents() {}

func (webView *webView) SwapBuffers() {}

func (webView *webView) SetPosition(x, y int) {}

func (webView *webView) SetSize(width, height int) {}

func (webView *webView) SetAspectRatio(numer, denom int) {}

func (webView *webView) SetTitle(title string) {}

func (webView *webView) SetIcons(icons []image.Image) {}

func (webView *webView) setCloseCallback(callback func()) {
	// webView.window.SetCloseCallback(func(w *glfw.Window) {
	// 	callback()
	// })
}

func (webView *webView) setResizeCallback(callback func(width, height int)) {
	// webView.window.SetSizeCallback(func(w *glfw.Window, width, height int) {
	// 	callback(width, height)
	// })
}

func (webView *webView) setJoystickCallback(callback func(joy Joystick, event Action)) {
	// TODO: Implement joystick input.
}

func (webView *webView) setKeyCallback(callback func(key Key, scancode int, action Action, mods ModKey)) {
	// webView.window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	// 	callback(Key(key), scancode, Action(action), ModKey(mods))
	// })
}

func (webView *webView) setMouseButtonCallback(callback func(button MouseButton, action Action, mod ModKey)) {
	// webView.window.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	// 	callback(MouseButton(int(button)+1), Action(action), ModKey(mod))
	// })
}

func (webView *webView) setMouseMotionCallback(callback func(x, y float32)) {
	// webView.window.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
	// 	callback(float32(x), float32(y))
	// })
}

func (webView *webView) setMouseScrollCallback(callback func(x, y float32)) {
	// webView.window.SetScrollCallback(func(w *glfw.Window, x, y float64) {
	// 	callback(float32(x), float32(y))
	// })
}

func (webView *webView) setTouchMotionCallback(callback func(x, y []float32)) {
	// TODO: Figure out how to separate this from mouse motion cross-platform.
}

func (webView *webView) setTypingCallback(callback func(char rune, mods ModKey)) {
	// webView.window.SetCharModsCallback(func(w *glfw.Window, char rune, mods glfw.ModifierKey) {
	// 	callback(char, ModKey(mods))
	// })
}

func (webView *webView) delete() {}
