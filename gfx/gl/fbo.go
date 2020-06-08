package gl

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type fbo struct {
	id              uint32
	colorAttachment framework.IAttachment
	depthAttachment framework.IAttachment

	width, height, samples uint32
}

func createFBO() *fbo {
	var id uint32
	gl.GenFramebuffers(1, &id)
	return &fbo{
		id: id,
	}
}

func (fbo *fbo) bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.id)
}

func (fbo *fbo) BindForRender() {
	// gl.DrawBuffers(1, fbo.colorAttachment.ID())
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, fbo.id)
}

func (fbo *fbo) SetColorAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferColor(fbo)
	fbo.colorAttachment = attachment
	fbo.unbind()
}

func (fbo *fbo) SetDepthAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferDepth(fbo)
	fbo.depthAttachment = attachment
	fbo.unbind()
}

func (fbo *fbo) unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (fbo *fbo) UnbindForRender() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	// gl.DrawBuffer(gl.BACK)
}

func (fbo *fbo) Delete() {
	gl.DeleteFramebuffers(1, &fbo.id)
}

func (fbo *fbo) BlitToScreen() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	gl.DrawBuffer(gl.BACK)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, fbo.id)
	gl.ReadBuffer(gl.COLOR_ATTACHMENT0)
	gl.BlitFramebuffer(0, 0, fbo.Width(), fbo.Height(), 0, 0, fbo.Width(), fbo.Height(), gl.COLOR_BUFFER_BIT, gl.NEAREST)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (srcFBO *fbo) BlitToFramebuffer(framebuffer framework.IFramebuffer) {
	targetFBO := framebuffer.(*fbo)
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, targetFBO.id)
	gl.DrawBuffer(gl.BACK)
	gl.BindFramebuffer(gl.READ_FRAMEBUFFER, srcFBO.id)
	gl.ReadBuffer(gl.COLOR_ATTACHMENT0)
	gl.BlitFramebuffer(0, 0, srcFBO.Width(), srcFBO.Height(), 0, 0, srcFBO.Width(), srcFBO.Height(), gl.COLOR_BUFFER_BIT, gl.NEAREST)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (fbo *fbo) Resize(width, height int32) {
	fbo.BindForRender()
	if fbo.colorAttachment != nil {
		fbo.colorAttachment.ResizeFramebufferColor(fbo)
	}
	if fbo.depthAttachment != nil {
		fbo.depthAttachment.ResizeFramebufferDepth(fbo)
	}
	gl.Viewport(0, 0, width, height)
	fbo.UnbindForRender()
}

func (fbo *fbo) Width() int32 {
	return int32(limitengine.Width())
}

func (fbo *fbo) Height() int32 {
	return int32(limitengine.Height())
}

func (fbo *fbo) Samples() int32 {
	return int32(limitengine.Samples())
}
