package gl

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx/framework"
	"github.com/go-gl/gl/v3.3-core/gl"
)

type fbo struct {
	id                     uint32
	colorAttachments       []uint32
	depthStencilAttachment uint32

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
	gl.DrawBuffers(int32(len(fbo.colorAttachments)), &fbo.colorAttachments[0])
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, fbo.id)
}

func (fbo *fbo) AddColorAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferColor(fbo, len(fbo.colorAttachments))
	fbo.colorAttachments = append(fbo.colorAttachments, attachment.ID())
	fbo.unbind()
}

func (fbo *fbo) AddDepthAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferDepth(fbo)
	fbo.depthStencilAttachment = attachment.ID()
	fbo.unbind()
}

func (fbo *fbo) AddStencilAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferStencil(fbo)
	fbo.depthStencilAttachment = attachment.ID()
	fbo.unbind()
}

func (fbo *fbo) AddDepthStencilAttachment(attachment framework.IAttachment) {
	fbo.bind()
	attachment.AttachToFramebufferDepthStencil(fbo)
	fbo.depthStencilAttachment = attachment.ID()
	fbo.unbind()
}

func (fbo *fbo) unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}

func (fbo *fbo) UnbindForRender() {
	gl.BindFramebuffer(gl.DRAW_FRAMEBUFFER, 0)
	gl.DrawBuffer(gl.BACK)
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

func (fbo *fbo) Width() int32 {
	return int32(limitengine.Width())
}

func (fbo *fbo) Height() int32 {
	return int32(limitengine.Height())
}

func (fbo *fbo) Samples() int32 {
	return int32(limitengine.Samples())
}
