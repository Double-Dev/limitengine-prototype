package gl

import (
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type rbo struct {
	id uint32
}

func createRBO() *rbo {
	var id uint32
	gl.GenRenderbuffers(1, &id)
	return &rbo{
		id: id,
	}
}

func (rbo *rbo) bind() {
	gl.BindRenderbuffer(gl.RENDERBUFFER, rbo.id)
}

// Attachment functions
func (rbo *rbo) AttachToFramebufferColor(framebuffer framework.IFramebuffer) {
	rbo.bind()
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, framebuffer.Samples(), gl.RGB, framebuffer.Width(), framebuffer.Height())
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, uint32(gl.COLOR_ATTACHMENT0), gl.RENDERBUFFER, rbo.id)
}

func (rbo *rbo) AttachToFramebufferDepth(framebuffer framework.IFramebuffer) {
	rbo.bind()
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, framebuffer.Samples(), gl.DEPTH_COMPONENT, framebuffer.Width(), framebuffer.Height())
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, gl.RENDERBUFFER, rbo.id)
}

func (rbo *rbo) AttachToFramebufferStencil(framebuffer framework.IFramebuffer) {
	rbo.bind()
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, framebuffer.Samples(), gl.STENCIL_INDEX, framebuffer.Width(), framebuffer.Height())
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo.id)
}

func (rbo *rbo) AttachToFramebufferDepthStencil(framebuffer framework.IFramebuffer) {
	rbo.bind()
	gl.RenderbufferStorageMultisample(gl.RENDERBUFFER, framebuffer.Samples(), gl.DEPTH_STENCIL, framebuffer.Width(), framebuffer.Height())
	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.DEPTH_STENCIL_ATTACHMENT, gl.RENDERBUFFER, rbo.id)
}

func (rbo *rbo) ID() uint32 {
	return rbo.id
}
