package framework

type IAttachment interface {
	AttachToFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	AttachToFramebufferStencil(framebuffer IFramebuffer)
	AttachToFramebufferDepthStencil(framebuffer IFramebuffer)
	ID() uint32
}
