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

func (this *Camera) SetProjectionMat(projectionMat gmath.Matrix44) {
	this.prefs.AddMatrix44("projMat", projectionMat)
}

func (this *Camera) SetViewMat(viewMat gmath.Matrix44) {
	this.prefs.AddMatrix44("viewMat", viewMat)
}

func CreateFrameBuffer() {
}

func CreateRenderBuffer() {

}
