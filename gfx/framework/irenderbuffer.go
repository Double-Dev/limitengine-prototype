package framework

type IRenderbuffer interface {
	// Attachment functions:
	AttachToFramebufferColor(framebuffer IFramebuffer, attachment int)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	AttachToFramebufferStencil(framebuffer IFramebuffer)
	AttachToFramebufferDepthStencil(framebuffer IFramebuffer)
}
