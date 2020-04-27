package gfx

import (
	"fmt"
	"runtime"
	"sync"
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
	projMatrix gmath.Matrix4

	AdvanceFrames = 2

	gfxMutex      = sync.RWMutex{}
	renderBatches = []map[*Camera]map[*Shader]map[*Material]map[*Mesh][]*Instance{}
	actionQueue   = []func(){}
	gfxPipeline   = [](chan func()){}
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
				projMatrix = gmath.NewProjectionMatrix3D(float32(height)/float32(width), 0.001, 1000.0, 60.0)
				fmt.Println(width, height)
				actionQueue = append(actionQueue, func() { context.Resize(width, height) })
			})
			projMatrix = gmath.NewProjectionMatrix3D(float32(limitengine.InitHeight)/float32(limitengine.InitWidth), 0.001, 1000.0, 60.0)

			currentTime := time.Now().UnixNano()
			for limitengine.Running() {
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
			}
		}()
		log.Log("GFX online...")
	}
}

// ClearScreen queues a gfx action that clears the screen based on the input color.
func ClearScreen(r, g, b, a float32) {
	actionQueue = append(actionQueue, func() { context.ClearScreen(r, g, b, a) })
}

// Sweep sweeps queued gfx actions onto the render pipeline and renders all gfx objects added to the batch.
func Sweep() {
	for len(gfxPipeline)-1 > AdvanceFrames {
		time.Sleep(time.Millisecond * 10)
	}
	actionQueue = append(actionQueue, func() {
		// Batching System TODO: Improve with instanced rendering.
		for camera, batch0 := range renderBatches[0] {
			// iFrameBuffer := frameBuffers[camera.id]
			// fmt.Println(iFrameBuffer)
			// Enable framebuffer
			for shader, batch1 := range batch0 {
				iShader := shaders[shader.id]
				iShader.Start()
				camera.prefs.loadTo(iShader)
				iShader.LoadUniformMatrix4fv("projMat", projMatrix.ToArray())
				// iShader.LoadUniformMatrix4fv("viewMat", gmath.NewIdentityMatrix(4, 4).Translate(camera.position).ToMatrix44())

				for material, batch2 := range batch1 {
					material.prefs.loadTo(iShader)
					iTexture := textures[material.texture.id]
					iTexture.Bind()
					for mesh, instances := range batch2 {
						mesh.prefs.loadTo(iShader)
						iMesh := meshes[mesh.id]
						iMesh.Enable()

						data := []float32{}

						for _, instance := range instances {
							instanceDefs := shader.GetInstanceDefs()
							for _, instanceDef := range instanceDefs {
								instance.dataMutex.RLock()
								data = append(data, instance.data[instanceDef.Name][0:instanceDef.Size]...)
								instance.dataMutex.RUnlock()
							}
						}
						iMesh.Render(shader.instanceBuffer, shader.GetInstanceDefs(), data, int32(len(instances)))

						iMesh.Disable()
					}
					iTexture.Unbind()
				}
				iShader.Stop()
			}
			// Disable framebuffer
		}
		renderBatches = renderBatches[1:]
	})
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
}

func Render(camera *Camera, shader *Shader, material *Material, mesh *Mesh, instance *Instance) {
	actionQueue = append(actionQueue, func() {
		if len(renderBatches) == 0 {
			renderBatches = append(renderBatches, make(map[*Camera]map[*Shader]map[*Material]map[*Mesh][]*Instance))
		}
		renderBatch := renderBatches[len(renderBatches)-1]
		batch0 := renderBatch[camera]
		if batch0 == nil {
			batch0 = make(map[*Shader]map[*Material]map[*Mesh][]*Instance)
			renderBatch[camera] = batch0
		}
		batch1 := batch0[shader]
		if batch1 == nil {
			batch1 = make(map[*Material]map[*Mesh][]*Instance)
			batch0[shader] = batch1
		}
		batch2 := batch1[material]
		if batch2 == nil {
			batch2 = make(map[*Mesh][]*Instance)
			batch1[material] = batch2
		}
		batch2[mesh] = append(batch2[mesh], instance)

		// iShader := shaders[shader.id]
		// var iModel framework.IModel
		// if model != nil {
		// 	iModel = models[model.id]
		// } else {
		// 	iModel = models[0]
		// }
		// iTexture := textures[texture.id]
		// iShader.Start()
		// iTexture.Bind()
		// iModel.Enable()

		// iShader.LoadUniformMatrix4fv("projMat", projMatrix.ToMatrix44())
		// vMat := gmath.NewIdentityMatrix(4, 4).Translate(camera.Position())

		// iShader.LoadUniformMatrix4fv("viewMat", vMat.ToMatrix44())
		// iShader.LoadUniformMatrix4fv("transformMat", transform.ToMatrix44())
		// iShader.LoadUniform1I("tex", 0)

		// iModel.Render()
		// iModel.Disable()
		// iShader.Stop()
	})
}
