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

	// TargetFPS is the amount of frames gfx will attempt to output per second.
	TargetFPS = float32(100.0)
	// AdvanceFrames is the number of frames for which the GFX pipeline will accept commands for in advance.
	AdvanceFrames = 2

	queuedFrames = 0

	fps        = float32(0.0)
	projMatrix gmath.Matrix4

	renderBatch = make(map[*Camera]map[Shader]map[Material]map[*Mesh][]*Instance)
	actionQueue = []func(){}
	gfxPipeline = [](chan func()){}
)

func init() {
	if limitengine.Running() {
		go func() {
			// GFX Init
			runtime.LockOSThread()
			view := limitengine.AppView()
			view.SetContext()
			var err error
			context, err = gl.NewContext()
			if err != nil {
				log.Err("Context could not be initialized.", err)
			}
			limitengine.AddResizeCallback(func(width, height int) {
				fmt.Println(width, height)
				actionQueue = append(actionQueue, func() { context.Resize(width, height) })
				for _, camera := range cameras {
					camera.resize(width, height)
				}
			})
			// GFX Loop
			currentTime := time.Now().UnixNano()
			for limitengine.Running() {
				if len(gfxPipeline) > 0 && time.Now().UnixNano()-currentTime > int64((1.0/TargetFPS)*1000000000.0) {
					lastTime := currentTime
					currentTime = time.Now().UnixNano()
					delta := float32(currentTime-lastTime) / 1000000000.0
					fps = 1.0 / delta
					// log.Log(fps)

					pipeline := gfxPipeline[0]
					for action := range pipeline { // TODO: Make this not send unexpected signal.
						action()
					}
					gfxPipeline = gfxPipeline[1:]
					view.SwapBuffers()
				} else {
					time.Sleep(time.Millisecond)
				}
			}
			// GFX Cleanup
			renderBatch = nil
			actionQueue = nil
			gfxPipeline = nil
			deleteFramebuffers()
			deleteShaders()
			deleteRenderbuffers()
			deleteTextures()
			deleteMeshes()
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
	for queuedFrames > AdvanceFrames {
		time.Sleep(time.Millisecond)
	}
	queuedFrames++
	actionQueue = append(actionQueue, func() {
		queuedFrames--
		for camera, batch0 := range renderBatch {
			iFrameBuffer := framebuffers[camera.id]
			if iFrameBuffer != nil {
				iFrameBuffer.Bind()
			} else {
				context.UnbindFramebuffers()
			}
			context.ClearScreen(camera.clearColor[0], camera.clearColor[1], camera.clearColor[2], camera.clearColor[3])
			for shader, batch1 := range batch0 {
				renderProgram := shader.RenderProgram()
				iShader := shaders[renderProgram.id]
				iShader.Start()
				camera.prefs.loadTo(iShader)
				shader.UniformLoader().loadTo(iShader)

				for material, batch2 := range batch1 {
					material.Prefs().loadTo(iShader)
					iTexture := textures[material.Texture().id]
					if iTexture != nil {
						iTexture.Bind()
					} else {
						context.UnbindTextures()
					}
					for mesh, instances := range batch2 {
						mesh.prefs.loadTo(iShader)
						context.DepthTest(mesh.DepthTest)
						context.BackCulling(mesh.BackCulling)
						context.WriteDepth(mesh.WriteDepth)
						iMesh := meshes[mesh.id]
						iMesh.Enable()

						data := []float32{} // TODO: Don't new a new slice every frame.

						for _, instance := range instances {
							instanceDefs := renderProgram.InstanceDefs()
							for _, instanceDef := range instanceDefs {
								instance.dataMutex.RLock()
								data = append(data, instance.data[instanceDef.Name][0:instanceDef.Size]...)
								instance.dataMutex.RUnlock()
							}
						}
						// TODO: Look into optimizing GPU overhead from instanced rendering.
						iMesh.Render(renderProgram.instanceBuffer, renderProgram.InstanceDefs(), data, int32(len(instances)))

						iMesh.Disable()
					}
				}
				iShader.Stop()
			}
			for _, blitCamera := range camera.blitCameras {
				iFrameBuffer.BlitToFramebuffer(framebuffers[blitCamera.id])
			}
		}

	})
	pipeline := make(chan func())
	gfxPipeline = append(gfxPipeline, pipeline)
	queue := actionQueue
	actionQueue = nil
	// go func() {
	for _, action := range queue {
		pipeline <- action // TODO: Make this not send unexpected signal.
	}
	close(pipeline)
	// }()
}

type Renderable struct {
	Camera   *Camera
	Shader   Shader
	Material Material
	Mesh     *Mesh
	Instance *Instance
}

func AddRenderable(renderable *Renderable) {
	actionQueue = append(actionQueue, func() {
		batch0 := renderBatch[renderable.Camera]
		if batch0 == nil {
			batch0 = make(map[Shader]map[Material]map[*Mesh][]*Instance)
			renderBatch[renderable.Camera] = batch0
		}
		batch1 := batch0[renderable.Shader]
		if batch1 == nil {
			batch1 = make(map[Material]map[*Mesh][]*Instance)
			batch0[renderable.Shader] = batch1
		}
		batch2 := batch1[renderable.Material]
		if batch2 == nil {
			batch2 = make(map[*Mesh][]*Instance)
			batch1[renderable.Material] = batch2
		}

		// TODO: Fix transparency sorting for translucent objects of different materials,
		// shaders, meshes, etc., unless that doesn't need to be supported.
		if renderable.Material.Transparency() {
			instances := batch2[renderable.Mesh]
			i := 0
			for i < len(instances) && instances[i].GetData("verttransformMat3")[2] > renderable.Instance.GetData("verttransformMat3")[2] {
				i++
			}
			instances = append(instances, nil)
			copy(instances[i+1:], instances[i:])
			instances[i] = renderable.Instance
			batch2[renderable.Mesh] = instances
		} else {
			batch2[renderable.Mesh] = append(batch2[renderable.Mesh], renderable.Instance)
		}
	})
}

func RemoveRenderable(renderable *Renderable) {
	actionQueue = append(actionQueue, func() {
		batch0 := renderBatch[renderable.Camera]
		if batch0 == nil {
			batch0 = make(map[Shader]map[Material]map[*Mesh][]*Instance)
			renderBatch[renderable.Camera] = batch0
		}
		batch1 := batch0[renderable.Shader]
		if batch1 == nil {
			batch1 = make(map[Material]map[*Mesh][]*Instance)
			batch0[renderable.Shader] = batch1
		}
		batch2 := batch1[renderable.Material]
		if batch2 == nil {
			batch2 = make(map[*Mesh][]*Instance)
			batch1[renderable.Material] = batch2
		}

		instances := batch2[renderable.Mesh]
		for i, batchInstance := range instances {
			if batchInstance == renderable.Instance {
				copy(instances[i:], instances[i+1:])
				instances[len(instances)-1] = nil
				instances = instances[:len(instances)-1]
				break
			}
		}
		batch2[renderable.Mesh] = instances
	})
}

func ClearRenderables() {
	actionQueue = append(actionQueue, func() {
		renderBatch = make(map[*Camera]map[Shader]map[Material]map[*Mesh][]*Instance)
	})
}
