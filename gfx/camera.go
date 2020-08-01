package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

const (
	projNone = "None"
	proj2D   = "2D"
	proj3D   = "3D"
)

var (
	framebufferIndex = uint32(1)
	framebuffers     = make(map[uint32]framework.IFramebuffer)
	cameras          = []*Camera{}

	defaultCamera = &Camera{
		id:             0,
		clearColor:     gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType: projNone,
		prefs:          NewUniformLoader(),
	}
)

func init() {
	framebuffers[0] = nil
	defaultCamera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewIdentityMatrix4(),
	)
	defaultCamera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, defaultCamera)
}

func deleteFramebuffers() {
	framebufferIndex = 0
	for _, iFramebuffer := range framebuffers {
		if iFramebuffer != nil {
			iFramebuffer.Delete()
		}
	}
	framebuffers = nil
	cameras = nil
	defaultCamera = nil
}

// Camera is a gfx framebuffer.
type Camera struct {
	id                     uint32
	colorAttachment        Attachment
	depthStencilAttachment Attachment
	blitCameras            []*Camera

	clearColor gmath.Vector4

	projectionType           string
	nearPlane, farPlane, fov float32
	prefs                    UniformLoader
}

func DefaultCamera() *Camera {
	return defaultCamera
}

// TODO: Think about turning Camera into interface to allow easy creation of variations.
func NewCamera(colorAttachment, depthAttachment Attachment) *Camera {
	camera := &Camera{
		id:                     framebufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         projNone,
		prefs:                  NewUniformLoader(),
	}
	framebufferIndex++
	actionQueue = append(actionQueue, func() {
		framebuffers[camera.id] = context.NewFramebuffer(camera.colorAttachment.frameworkAttachment(), camera.depthStencilAttachment.frameworkAttachment(), 1.0, 1.0, 1)
	})
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewIdentityMatrix4(),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

// NewCamera2D news a camera initialized with a 2D projection matrix.
// Note: A 2D camera cannot 'see' anything more than 1.0 units away on
// the z-axis.
func NewCamera2D(colorAttachment, depthAttachment Attachment) *Camera {
	camera := &Camera{
		id:                     framebufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         proj2D,
		prefs:                  NewUniformLoader(),
	}
	framebufferIndex++
	actionQueue = append(actionQueue, func() {
		framebuffers[camera.id] = context.NewFramebuffer(camera.colorAttachment.frameworkAttachment(), camera.depthStencilAttachment.frameworkAttachment(), 1.0, 1.0, 4)
	})
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix2D(limitengine.AspectRatio()),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func NewCamera3D(colorAttachment, depthAttachment Attachment, nearPlane, farPlane, fov float32) *Camera {
	camera := &Camera{
		id:                     framebufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         proj3D,
		nearPlane:              nearPlane,
		farPlane:               farPlane,
		fov:                    fov,
		prefs:                  NewUniformLoader(),
	}
	framebufferIndex++
	actionQueue = append(actionQueue, func() {
		framebuffers[camera.id] = context.NewFramebuffer(camera.colorAttachment.frameworkAttachment(), camera.depthStencilAttachment.frameworkAttachment(), 1.0, 1.0, 4)
	})
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix3D(
			limitengine.AspectRatio(),
			nearPlane,
			farPlane,
			fov,
		),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func (camera *Camera) HasBlitCamera(targetBlitCamera *Camera) bool {
	for _, blitCamera := range camera.blitCameras {
		if blitCamera == targetBlitCamera || blitCamera.HasBlitCamera(targetBlitCamera) {
			return true
		}
	}
	return false
}

func (camera *Camera) AddBlitCamera(blitCamera *Camera) {
	if !camera.HasBlitCamera(blitCamera) {
		camera.blitCameras = append(camera.blitCameras, blitCamera)
	} // TODO: Print error otherwise.
}

func (camera *Camera) RemoveBlitCamera(targetBlitCamera *Camera) {
	for i, blitCamera := range camera.blitCameras {
		if blitCamera == targetBlitCamera {
			camera.blitCameras[i] = camera.blitCameras[len(camera.blitCameras)-1]
			camera.blitCameras = camera.blitCameras[:len(camera.blitCameras)-1]
		}
	}
}

func (camera *Camera) SetClearColor(r, g, b, a float32) {
	camera.clearColor.Set(r, g, b, a)
}

func (camera *Camera) SetViewMat(viewMat gmath.Matrix4) {
	camera.prefs.AddMatrix4("vertviewMat", viewMat)
}

func (camera *Camera) resize(width, height int) {
	if framebuffers[camera.id] != nil {
		actionQueue = append(actionQueue, func() { framebuffers[camera.id].Resize(int32(width), int32(height)) })
	}
	aspectRatio := float32(height) / float32(width)
	switch camera.projectionType {
	case proj2D:
		camera.prefs.AddMatrix4(
			"vertprojMat",
			gmath.NewProjectionMatrix2D(aspectRatio),
		)
		break
	case proj3D:
		camera.prefs.AddMatrix4(
			"vertprojMat",
			gmath.NewProjectionMatrix3D(aspectRatio, camera.nearPlane, camera.farPlane, camera.fov),
		)
		break
	}
}

// DeleteCamera queues a gfx action that deletes the input camera.
func DeleteCamera(camera *Camera) {
	actionQueue = append(actionQueue, func() {
		iFramebuffer := framebuffers[camera.id]
		iFramebuffer.Delete()
		delete(framebuffers, camera.id)
	})
}
