package framework

type IAttachment interface {
	AttachToFramebufferColor(framebuffer IFramebuffer)
	ResizeFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	ResizeFramebufferDepth(framebuffer IFramebuffer)
	Delete()
	ID() *uint32
}
