package framework

type IRenderbuffer interface {
	// Attachment functions:
	AttachToFramebufferColor(framebuffer IFramebuffer)
	ResizeFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	ResizeFramebufferDepth(framebuffer IFramebuffer)
	Delete()
	ID() *uint32
}
