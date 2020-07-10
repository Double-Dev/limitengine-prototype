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
	TargetFPS = float32(100.0)
	// AdvanceFrames is the number of frames for which the GFX pipeline will accept commands for in advance.
	AdvanceFrames = 2

	queuedFrames = 0

	fps        = float32(0.0)
	projMatrix gmath.Matrix4

	renderOrder   []int32
	renderCameras []*Camera
	renderBatch   = make(map[int32]map[*Camera]map[Shader]map[Material]map[*Mesh]*instanceData)
	renderMutex   = sync.RWMutex{}
	actionQueue   = []func(){}
	gfxPipeline   = [](chan func()){}
)

type instanceData struct {
	instances []*Instance
	data      [][]float32
}

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
			renderOrder = nil
			renderCameras = nil
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
		for _, camera := range renderCameras {
			iFrameBuffer := framebuffers[camera.id]
			if iFrameBuffer != nil {
				iFrameBuffer.Bind()
			} else {
				context.UnbindFramebuffers()
			}
			context.ClearScreen(camera.clearColor[0], camera.clearColor[1], camera.clearColor[2], camera.clearColor[3])
		}
		for _, layer := range renderOrder {
			for camera, batch0 := range renderBatch[layer] {
				iFrameBuffer := framebuffers[camera.id]
				if iFrameBuffer != nil {
					iFrameBuffer.Bind()
				} else {
					context.UnbindFramebuffers()
				}
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
						for mesh, batch3 := range batch2 {
							mesh.prefs.loadTo(iShader)
							context.DepthTest(mesh.DepthTest)
							context.BackCulling(mesh.BackCulling)
							context.WriteDepth(mesh.WriteDepth)
							iMesh := meshes[mesh.id]
							iMesh.Enable()

							for i := 0; i < (len(batch3.instances)/context.GetMaxInstances())+1; i++ {
								for j, instance := range batch3.instances[i*context.GetMaxInstances() : gmath.MinI((i+1)*context.GetMaxInstances(), len(batch3.instances))] {
									if instance.dataModified {
										instance.dataMutex.RLock()
										for k, instanceDef := range renderProgram.InstanceDefs() {
											for l, value := range instance.data[instanceDef.Name][0:instanceDef.Size] {
												batch3.data[i][j*renderProgram.InstanceSize()+k*instanceDef.Size+l] = value
											}
										}
										instance.dataMutex.RUnlock()
										instance.dataModified = false
									}
								}
								// TODO: Look into optimizing GPU overhead from instanced rendering.
								if len(batch3.data[i]) > 0 {
									iMesh.Render(
										renderProgram.instanceBuffer,
										renderProgram.InstanceDefs(),
										batch3.data[i],
										int32(len(batch3.data[i])/renderProgram.InstanceSize()),
									)
								}
							}
							iMesh.Disable()
						}
					}
					iShader.Stop()
				}
				for _, blitCamera := range camera.blitCameras {
					iFrameBuffer.BlitToFramebuffer(framebuffers[blitCamera.id])
				}
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
	Layer    int32
	Camera   *Camera
	Shader   Shader
	Material Material
	Mesh     *Mesh
	Instance *Instance
}

func AddRenderable(renderable *Renderable) {
	actionQueue = append(actionQueue, func() {
		batchLayer := renderBatch[renderable.Layer]
		if batchLayer == nil {
			batchLayer = make(map[*Camera]map[Shader]map[Material]map[*Mesh]*instanceData)
			renderMutex.Lock()
			renderBatch[renderable.Layer] = batchLayer
			renderMutex.Unlock()
			i := 0
			for i < len(renderOrder) && renderOrder[i] < renderable.Layer {
				i++
			}
			renderOrder = append(renderOrder, -1)
			copy(renderOrder[i+1:], renderOrder[i:])
			renderOrder[i] = renderable.Layer
		}
		batch0 := batchLayer[renderable.Camera]
		if batch0 == nil {
			batch0 = make(map[Shader]map[Material]map[*Mesh]*instanceData)
			renderMutex.Lock()
			batchLayer[renderable.Camera] = batch0
			renderMutex.Unlock()
			renderCameras = append(renderCameras, renderable.Camera)
		}
		batch1 := batch0[renderable.Shader]
		if batch1 == nil {
			batch1 = make(map[Material]map[*Mesh]*instanceData)
			renderMutex.Lock()
			batch0[renderable.Shader] = batch1
			renderMutex.Unlock()
		}
		batch2 := batch1[renderable.Material]
		if batch2 == nil {
			batch2 = make(map[*Mesh]*instanceData)
			renderMutex.Lock()
			batch1[renderable.Material] = batch2
			renderMutex.Unlock()
		}
		batch3 := batch2[renderable.Mesh]
		if batch3 == nil {
			batch3 = &instanceData{}
			renderMutex.Lock()
			batch2[renderable.Mesh] = batch3
			renderMutex.Unlock()
		}
		batch3.instances = append(batch3.instances, renderable.Instance)
		data := []float32{}
		instanceDefs := renderable.Shader.RenderProgram().InstanceDefs()
		renderable.Instance.dataMutex.RLock()
		for _, instanceDef := range instanceDefs {
			data = append(data, renderable.Instance.data[instanceDef.Name][0:instanceDef.Size]...)
		}
		renderable.Instance.dataMutex.RUnlock()
		index := len(batch3.instances) / context.GetMaxInstances()
		if len(batch3.data) <= index {
			batch3.data = append(batch3.data, data)
		} else {
			batch3.data[index] = append(batch3.data[index], data...)
		}
	})
}

func RemoveRenderable(renderable *Renderable) {
	actionQueue = append(actionQueue, func() {
		batchLayer := renderBatch[renderable.Layer]
		if batchLayer == nil {
			return
		}
		batch0 := batchLayer[renderable.Camera]
		if batch0 == nil {
			return
		}
		batch1 := batch0[renderable.Shader]
		if batch1 == nil {
			return
		}
		batch2 := batch1[renderable.Material]
		if batch2 == nil {
			return
		}
		batch3 := batch2[renderable.Mesh]
		if batch3 == nil {
			return
		}

		// TODO: If anything starts behaving weirdly in the GFX, investigate here first.
		for i, instance := range batch3.instances {
			if instance == renderable.Instance {
				copy(batch3.instances[i:], batch3.instances[i+1:])
				batch3.instances[len(batch3.instances)-1] = nil
				batch3.instances = batch3.instances[:len(batch3.instances)-1]
				index := i / context.GetMaxInstances()
				instanceSize := renderable.Shader.RenderProgram().InstanceSize()
				copy(batch3.data[index][i*instanceSize:], batch3.data[index][(i+1)*instanceSize:])
				for j := len(batch3.data[index]) - instanceSize; j < len(batch3.data[index]); j++ {
					batch3.data[index][j] = 0.0
				}
				batch3.data[index] = batch3.data[index][:len(batch3.data[index])-instanceSize]
				break
			}
		}
	})
}

func ClearRenderables() {
	actionQueue = append(actionQueue, func() {
		renderOrder = nil
		renderCameras = nil
		renderBatch = make(map[int32]map[*Camera]map[Shader]map[Material]map[*Mesh]*instanceData)
	})
}
