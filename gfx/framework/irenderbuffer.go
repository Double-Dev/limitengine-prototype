package framework

type IRenderbuffer interface {
	// Attachment functions:
	AttachToFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	AttachToFramebufferStencil(framebuffer IFramebuffer)
	AttachToFramebufferDepthStencil(framebuffer IFramebuffer)
	ID() *uint32
}
