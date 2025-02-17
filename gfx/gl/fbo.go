package gl

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/dependencies/gl/v3.3-core/gl"
	"github.com/double-dev/limitengine/gfx/framework"
)

type fbo struct {
	id              uint32
	colorAttachment framework.IAttachment
	depthAttachment framework.IAttachment

	width, height float32
	samples       int32
}

func newFBO(width, height float32, samples int32) *fbo {
	var id uint32
	gl.GenFramebuffers(1, &id)
	return &fbo{
		id:      id,
		width:   width,
		height:  height,
		samples: samples,
	}
}

func (fbo *fbo) Bind() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, fbo.id)
}

func (fbo *fbo) SetColorAttachment(attachment framework.IAttachment) {
	fbo.Bind()
	attachment.AttachToFramebufferColor(fbo)
	fbo.colorAttachment = attachment
	fbo.unbind()
}

func (fbo *fbo) SetDepthAttachment(attachment framework.IAttachment) {
	fbo.Bind()
	attachment.AttachToFramebufferDepth(fbo)
	fbo.depthAttachment = attachment
	fbo.unbind()
}

func (fbo *fbo) unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (fbo *fbo) Delete() {
	gl.DeleteFramebuffers(1, &fbo.id)
}

func (srcFBO *fbo) BlitToFramebuffer(framebuffer framework.IFramebuffer) {
	if framebuffer != nil {
		targetFBO := framebuffer.(*fbo)
		gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, targetFBO.id)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, srcFBO.id)
		gl.BlitFramebuffer(0, 0, srcFBO.Width(), srcFBO.Height(), 0, 0, framebuffer.Width(), framebuffer.Height(), gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT, gl.NEAREST)
	} else {
		gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
		gl.BindFramebuffer(gl.READ_FRAMEBUFFER, srcFBO.id)
		gl.BlitFramebuffer(0, 0, srcFBO.Width(), srcFBO.Height(), 0, 0, int32(limitengine.Width()), int32(limitengine.Height()), gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT, gl.NEAREST)
	}
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (fbo *fbo) Resize(width, height int32) {
	fbo.Bind()
	if fbo.colorAttachment != nil {
		fbo.colorAttachment.ResizeFramebufferColor(fbo)
	}
	if fbo.depthAttachment != nil {
		fbo.depthAttachment.ResizeFramebufferDepth(fbo)
	}
	gl.Viewport(0, 0, width, height)
	fbo.unbind()
}

func (fbo *fbo) Width() int32 {
	return int32(float32(limitengine.Width()) * fbo.width)
}

func (fbo *fbo) Height() int32 {
	return int32(float32(limitengine.Height()) * fbo.height)
}

func (fbo *fbo) Samples() int32 {
	return fbo.samples // TODO: Provide options for fbo width, height, and samples.
}
