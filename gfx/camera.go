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
	frameBufferIndex = uint32(1)
	frameBuffers     = make(map[uint32]framework.IFramebuffer)
	cameras          = []*Camera{}

	defaultCamera = &Camera{
		id:             0,
		clearColor:     gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType: projNone,
		prefs:          NewUniformLoader(),
	}
)

func init() {
	frameBuffers[0] = nil

	defaultCamera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewIdentityMatrix4(),
	)
	defaultCamera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, defaultCamera)
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
func CreateCamera(colorAttachment, depthAttachment Attachment) *Camera {
	camera := &Camera{
		id:                     frameBufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         projNone,
		prefs:                  NewUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() {
		frameBuffers[camera.id] = context.CreateFramebuffer(camera.colorAttachment.getFrameworkAttachment(), camera.depthStencilAttachment.getFrameworkAttachment(), 1.0, 1.0, 1)
	})
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewIdentityMatrix4(),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

// CreateCamera2D creates a camera initialized with a 2D projection matrix.
// Note: A 2D camera cannot 'see' anything more than 1.0 units away on
// the z-axis.
func CreateCamera2D(colorAttachment, depthAttachment Attachment) *Camera {
	camera := &Camera{
		id:                     frameBufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         proj2D,
		prefs:                  NewUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() {
		frameBuffers[camera.id] = context.CreateFramebuffer(camera.colorAttachment.getFrameworkAttachment(), camera.depthStencilAttachment.getFrameworkAttachment(), 1.0, 1.0, 4)
	})
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix2D(limitengine.AspectRatio()),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func CreateCamera3D(colorAttachment, depthAttachment Attachment, nearPlane, farPlane, fov float32) *Camera {
	camera := &Camera{
		id:                     frameBufferIndex,
		colorAttachment:        colorAttachment,
		depthStencilAttachment: depthAttachment,
		clearColor:             gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType:         proj3D,
		nearPlane:              nearPlane,
		farPlane:               farPlane,
		fov:                    fov,
		prefs:                  NewUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() {
		frameBuffers[camera.id] = context.CreateFramebuffer(camera.colorAttachment.getFrameworkAttachment(), camera.depthStencilAttachment.getFrameworkAttachment(), 1.0, 1.0, 4)
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
	if frameBuffers[camera.id] != nil {
		actionQueue = append(actionQueue, func() { frameBuffers[camera.id].Resize(int32(width), int32(height)) })
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
