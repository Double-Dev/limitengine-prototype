package framework

type ITexture interface {
	Bind()
	NearestFilter(mipmap, antisotrophic bool)
	LinearFilter(mipmap, antisotrophic bool)
	TextureData(image []uint8, width, height int32)
	Unbind()
	Delete()

	// Attachment functions:
	AttachToFramebufferColor(framebuffer IFramebuffer)
	AttachToFramebufferDepth(framebuffer IFramebuffer)
	AttachToFramebufferStencil(framebuffer IFramebuffer)
	AttachToFramebufferDepthStencil(framebuffer IFramebuffer)
	ID() *uint32
}
