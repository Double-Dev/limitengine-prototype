package gfx

import (
	"double-dev/limitengine/core"
	"double-dev/limitengine/gfx/framework"
	"double-dev/limitengine/gfx/gl"
	"double-dev/limitengine/gmath"
	"runtime"
	"time"
)

var (
	view    framework.View
	context framework.Context

	fps        = float32(0.0)
	projMatrix gmath.Matrix

	renderBuffers map[uint32]uint32

	renderBatch = make(map[uint32]map[uint32]uint32)
	actionQueue = []func(){}
	gfxPipeline = [](chan func()){}
)

func init() {
	go func() {
		runtime.LockOSThread()
		view = newGLFWView()
		context, _ = gl.NewGLContext()
		defer view.Delete()

		view.SetCloseCallback(func() {
			core.Running = false
		})
		view.SetResizeCallback(func(width, height int) {
			projMatrix = gmath.NewProjectionMatrix(float32(height)/float32(width), 0.001, 1000.0, 60.0)
			context.Resize(width, height)
		})
		width, height := view.GetSize()
		projMatrix = gmath.NewProjectionMatrix(float32(height)/float32(width), 0.001, 1000.0, 60.0)

		currentTime := time.Now().UnixNano()
		for core.Running {
			if len(gfxPipeline) > 0 {
				lastTime := currentTime
				currentTime = time.Now().UnixNano()
				delta := float32(currentTime-lastTime) / 1000000000.0
				fps = 1.0 / delta

				pipeline := gfxPipeline[0]
				for action := range pipeline {
					action()
				}
				gfxPipeline = gfxPipeline[1:]
				view.SwapBuffers()
			} else {
				time.Sleep(time.Millisecond * 10)
			}
			view.UpdateEvents()
		}
	}()
}

// ClearScreen queues a gfx action that clears the screen based on the input color.
func ClearScreen(r, g, b, a float32) {
	actionQueue = append(actionQueue, func() { context.ClearScreen(r, g, b, a) })
}

// RenderSweep sweeps queued gfx actions onto the render pipeline.
func RenderSweep() {
	if len(gfxPipeline) <= 6 {
		pipeline := make(chan func())
		gfxPipeline = append(gfxPipeline, pipeline)
		queue := actionQueue
		actionQueue = nil
		go func() {
			for _, action := range queue {
				pipeline <- action
			}
			close(pipeline)
		}()
	} else {
		time.Sleep(time.Millisecond * 10)
	}
}

func Render(camera gmath.Matrix, shader *Shader, model *Model, texture *Texture) {
	actionQueue = append(actionQueue, func() {
		iShader := shaders[shader.id]
		iModel := models[model.id]
		iTexture := textures[texture.id]
		iShader.Start()
		iTexture.Bind()
		iShader.LoadUniformMatrix4fv("projMat", projMatrix)
		iShader.LoadUniformMatrix4fv("viewMat", camera)
		iShader.LoadUniform1I("tex", 0)
		iModel.Enable()
		iModel.Render()
		iModel.Disable()
		iShader.Stop()
	})
}

func CreateFrameBuffer() {
}

func CreateRenderBuffer() {

}
