package framework

type IAttachment interface {
	AttachToFramebufferColor(framebuffer IFramebuffer, attachment int)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	AttachToFramebufferStencil(framebuffer IFramebuffer)
	AttachToFramebufferDepthStencil(framebuffer IFramebuffer)
	ID() uint32
}
