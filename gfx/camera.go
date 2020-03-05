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
	id       uint32
	position gmath.Vector
	rotation gmath.Vector
	scale    gmath.Vector
}

// CreateCamera queues a gfx action that creates a new camera.
func CreateCamera() *Camera {
	return &Camera{
		id:       0,
		position: gmath.NewVector(0.0, 0.0, 0.0, 1.0),
		rotation: gmath.NewVector(0.0, 0.0, 0.0, 1.0),
		scale:    gmath.NewVector(1.0, 1.0, 1.0, 1.0),
	}
}

func (camera *Camera) Position() gmath.Vector {
	return camera.position
}

func (camera *Camera) Rotation() gmath.Vector {
	return camera.rotation
}

func (camera *Camera) Scale() gmath.Vector {
	return camera.scale
}
