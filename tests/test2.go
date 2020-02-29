package main

// func main() {
// 	runtime.LockOSThread()
// 	if err := glfw.Init(); err != nil {
// 		panic("Couldn't start glfw!")
// 	}

// 	glfw.WindowHint(glfw.Visible, glfw.False)

// 	running := true

// 	for i := 0; i < 4; i++ {
// 		window, _ := glfw.CreateWindow(400, 400, "window", nil, nil)
// 		window.SetPos(100+450*i, 400)
// 		window.SetCloseCallback(func(w *glfw.Window) {
// 			running = false
// 		})
// 		go func() {
// 			window.MakeContextCurrent()
// 			gl.Init()
// 			glfw.SwapInterval(1)
// 			for running {
// 				gl.ClearColor(0.25*float32(i), 1.0-0.25*float32(i), 0.5, 1.0)
// 				gl.Clear(gl.COLOR_BUFFER_BIT)
// 				window.SwapBuffers()
// 			}
// 		}()
// 		window.Show()
// 	}

// 	for running {
// 		glfw.WaitEvents()
// 	}

// 	glfw.Terminate()
// }
