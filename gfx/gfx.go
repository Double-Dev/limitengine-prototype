package gfx

import (
	"fmt"
	"runtime"
	"time"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gfx/gl"
	"github.com/double-dev/limitengine/gmath"
)

var (
	log     = limitengine.NewLogger("gfx")
	context framework.Context

	fps        = float32(0.0)
	projMatrix gmath.Matrix

	renderBuffers map[uint32]uint32

	renderBatch = make(map[uint32]map[uint32]uint32)
	actionQueue = []func(){}
	gfxPipeline = [](chan func()){}
)

func init() {
	if limitengine.Running() {
		go func() {
			runtime.LockOSThread()
			view := limitengine.AppView()
			view.SetContext()
			var err error
			context, err = gl.NewGLContext()
			if err != nil {
				log.Err("Context could not be initialized.", err)
			}

			limitengine.AddResizeCallback(func(width, height int) {
				projMatrix = gmath.NewProjectionMatrix(float32(height)/float32(width), 0.001, 1000.0, 60.0)
				fmt.Println(width, height)
				actionQueue = append(actionQueue, func() { context.Resize(width, height) })
			})
			projMatrix = gmath.NewProjectionMatrix(float32(limitengine.InitHeight)/float32(limitengine.InitWidth), 0.001, 1000.0, 60.0)

			// currentTime := time.Now().UnixNano()
			for limitengine.Running() {
				if len(gfxPipeline) > 0 {
					// lastTime := currentTime
					// currentTime = time.Now().UnixNano()
					// delta := float32(currentTime-lastTime) / 1000000000.0
					// fps = 1.0 / delta
					// fmt.Println(fps)

					pipeline := gfxPipeline[0]
					for action := range pipeline {
						action()
					}
					gfxPipeline = gfxPipeline[1:]
					view.SwapBuffers()
				} else {
					time.Sleep(time.Millisecond * 10)
				}
			}
		}()
		log.Log("GFX online...")
	}
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
		// iShader.LoadUniform1I("tex", 0)
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
