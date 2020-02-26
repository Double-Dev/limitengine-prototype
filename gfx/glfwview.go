package gfx

import (
	"image"
	"image/png"
	"os"

	"github.com/vulkan-go/glfw/v3.3/glfw"
)

type glfwView struct {
	window *glfw.Window
}

func newGLFWView() *glfwView {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.Samples, 8)

	reader, _ := os.Open("testIcon.png")
	icon, _ := png.Decode(reader)
	icons := []image.Image{icon}

	window, err := glfw.CreateWindow(400, 400, "Hello World", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetIcon(icons)

	return &glfwView{
		window: window,
	}
}

func (*glfwView) OnResume() {

}

func (*glfwView) OnPause() {

}

func (glfwView *glfwView) SetCloseCallback(callback func()) {
	glfwView.window.SetCloseCallback(func(w *glfw.Window) {
		callback()
	})
}

func (glfwView *glfwView) SetResizeCallback(callback func(width, height int)) {
	glfwView.window.SetSizeCallback(func(w *glfw.Window, width, height int) {
		callback(width, height)
	})
}

func (glfwView *glfwView) SwapBuffers() {
	glfwView.window.SwapBuffers()
}

func (glfwView) UpdateEvents() {
	glfw.PollEvents()
}

func (glfwView *glfwView) GetSize() (int, int) {
	return glfwView.window.GetSize()
}

func (*glfwView) Delete() {
	glfw.Terminate()
}
