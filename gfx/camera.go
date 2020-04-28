package gfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/double-dev/limitengine/gmath"
)

var (
	frameBufferIndex = uint32(1)
	frameBuffers     = make(map[uint32]framework.IFrameBuffer)
	cameras          = []*Camera{}
)

func init() { frameBuffers[0] = nil }

// Camera is a gfx framebuffer.
type Camera struct {
	id                       uint32
	perspective3D            bool
	nearPlane, farPlane, fov float32
	prefs                    uniformLoader
}

// CreateCamera2D creates a camera initialized with a 2D projection matrix.
// Note: A 2D camera cannot 'see' anything more than 1.0 units away on
// the z-axis.
func CreateCamera2D() *Camera {
	camera := &Camera{
		id:            0,
		perspective3D: false,
		prefs:         newUniformLoader(),
	}
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix2D(limitengine.GetAspectRatio()),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func CreateCamera3D(nearPlane, farPlane, fov float32) *Camera {
	camera := &Camera{
		id:            0,
		perspective3D: true,
		nearPlane:     nearPlane,
		farPlane:      farPlane,
		fov:           fov,
		prefs:         newUniformLoader(),
	}
	camera.prefs.AddMatrix4(
		"vertprojMat",
		gmath.NewProjectionMatrix3D(
			limitengine.GetAspectRatio(),
			nearPlane,
			farPlane,
			fov,
		),
	)
	camera.SetViewMat(gmath.NewIdentityMatrix4())
	cameras = append(cameras, camera)
	return camera
}

func (camera *Camera) updateProjectionMat(aspectRatio float32) {
	if camera.perspective3D {
		camera.prefs.AddMatrix4(
			"vertprojMat",
			gmath.NewProjectionMatrix3D(aspectRatio, camera.nearPlane, camera.farPlane, camera.fov),
		)
	} else {
		camera.prefs.AddMatrix4(
			"vertprojMat",
			gmath.NewProjectionMatrix2D(aspectRatio),
		)
	}
}

func (camera *Camera) SetViewMat(viewMat gmath.Matrix4) {
	camera.prefs.AddMatrix4("vertviewMat", viewMat)
}

func CreateFrameBuffer() {
}

func CreateRenderBuffer() {

}
