package gfx

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	frameBufferIndex = uint32(1)
	frameBuffers     = make(map[uint32]framework.IFrameBuffer)
)

func init() { frameBuffers[0] = nil }

// Camera is a gfx framebuffer.
type Camera struct {
	id    uint32
	prefs uniformLoader
}

// CreateCamera queues a gfx action that creates a new camera.
func CreateCamera() *Camera {
	return &Camera{
		id:    0,
		prefs: newUniformLoader(),
	}
}

func (camera *Camera) SetProjectionMat(projectionMat gmath.Matrix4) {
	camera.prefs.AddMatrix4("vertprojMat", projectionMat)
}

func (camera *Camera) SetViewMat(viewMat gmath.Matrix4) {
	camera.prefs.AddMatrix4("vertviewMat", viewMat)
}

func CreateFrameBuffer() {
}

func CreateRenderBuffer() {

}
