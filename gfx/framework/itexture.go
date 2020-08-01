package framework

type ITexture interface {
	Bind()
	NearestFilter(mipmap, antisotrophic bool)
	LinearFilter(mipmap, antisotrophic bool)
	TextureData(image []uint8, width, height int32)

	// Attachment functions:
	AttachToFramebufferColor(framebuffer IFramebuffer)
	ResizeFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	ResizeFramebufferDepth(framebuffer IFramebuffer)
	Delete()
	ID() *uint32
}
