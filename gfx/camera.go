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
		prefs:          newUniformLoader(),
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
	colorAttachments       []Attachment
	depthStencilAttachment Attachment

	clearColor gmath.Vector4

	projectionType           string
	nearPlane, farPlane, fov float32
	prefs                    uniformLoader
}

func DefaultCamera() *Camera {
	return defaultCamera
}

func CreateCamera() *Camera {
	camera := &Camera{
		id:             frameBufferIndex,
		clearColor:     gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType: projNone,
		prefs:          newUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() { frameBuffers[camera.id] = context.CreateFramebuffer() })
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
func CreateCamera2D() *Camera {
	camera := &Camera{
		id:             frameBufferIndex,
		clearColor:     gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType: proj2D,
		prefs:          newUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() { frameBuffers[camera.id] = context.CreateFramebuffer() })
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix2D(limitengine.AspectRatio()),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func CreateCamera3D(nearPlane, farPlane, fov float32) *Camera {
	camera := &Camera{
		id:             frameBufferIndex,
		clearColor:     gmath.NewVector4(0.0, 0.0, 0.0, 1.0),
		projectionType: proj3D,
		nearPlane:      nearPlane,
		farPlane:       farPlane,
		fov:            fov,
		prefs:          newUniformLoader(),
	}
	frameBufferIndex++
	actionQueue = append(actionQueue, func() { frameBuffers[camera.id] = context.CreateFramebuffer() })
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

func (camera *Camera) SetClearColor(r, g, b, a float32) {
	camera.clearColor.Set(r, g, b, a)
}

func (camera *Camera) AddColorAttachment(attachment Attachment) {
	if len(camera.colorAttachments) <= 32 {
		actionQueue = append(actionQueue, func() { frameBuffers[camera.id].AddColorAttachment(attachment.getFrameworkAttachment()) })
		camera.colorAttachments = append(camera.colorAttachments, attachment)
	}
}

func (camera *Camera) AddDepthAttachment(attachment Attachment) {
	if camera.depthStencilAttachment != nil {
		actionQueue = append(actionQueue, func() { frameBuffers[camera.id].AddDepthAttachment(attachment.getFrameworkAttachment()) })
		camera.depthStencilAttachment = attachment
	}
}

func (camera *Camera) AddStencilAttachment(attachment Attachment) {
	if camera.depthStencilAttachment != nil {
		actionQueue = append(actionQueue, func() { frameBuffers[camera.id].AddStencilAttachment(attachment.getFrameworkAttachment()) })
		camera.depthStencilAttachment = attachment
	}
}

func (camera *Camera) AddDepthStencilAttachment(attachment Attachment) {
	if camera.depthStencilAttachment != nil {
		actionQueue = append(actionQueue, func() { frameBuffers[camera.id].AddDepthStencilAttachment(attachment.getFrameworkAttachment()) })
		camera.depthStencilAttachment = attachment
	}
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

func (camera *Camera) SetViewMat(viewMat gmath.Matrix4) {
	camera.prefs.AddMatrix4("vertviewMat", viewMat)
}
