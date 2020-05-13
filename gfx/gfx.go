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

	// TargetFPS is the amount of frames gfx will attempt to output per second.
	TargetFPS     = float32(100.0)
	AdvanceFrames = 2

	fps        = float32(0.0)
	projMatrix gmath.Matrix4

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
				fmt.Println(width, height)
				for _, camera := range cameras {
					camera.updateProjectionMat(float32(height) / float32(width))
				}
				actionQueue = append(actionQueue, func() { context.Resize(width, height) })
			})

			currentTime := time.Now().UnixNano()
			for limitengine.Running() {
				if len(gfxPipeline) > 0 && time.Now().UnixNano()-currentTime > int64((1.0/TargetFPS)*1000000000.0) {
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
				// iShader.LoadUniformMatrix4fv("vertprojMat", gmath.NewProjectionMatrix2D(limitengine.GetAspectRatio()).ToArray())
				// iShader.LoadUniformMatrix4fv("vertviewMat", gmath.NewIdentityMatrix4().ToArray())
				shader.uniformLoader.loadTo(iShader)

				for material, batch2 := range batch1 {
					material.prefs.loadTo(iShader)
					if material.texture != nil {
						iTexture := textures[material.texture.id]
						iTexture.Bind()
					}
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

		// TODO: Fix transparency sorting for translucent objects of different materials,
		// shaders, meshes, etc., unless that doesn't need to be supported.
		if material.Transparency {
			instances := batch2[mesh]
			i := 0
			for i < len(instances) && instances[i].data["verttransformMat3"][2] > instance.data["verttransformMat3"][2] {
				i++
			}
			instances = append(instances, nil)
			copy(instances[i+1:], instances[i:])
			instances[i] = instance
			batch2[mesh] = instances
		} else {
			batch2[mesh] = append(batch2[mesh], instance)
		}
	})
}
