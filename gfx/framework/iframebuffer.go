package framework

type IFramebuffer interface {
	BindForRender()
	AddColorAttachment(attachment IAttachment)
	AddDepthAttachment(attachment IAttachment)
	AddStencilAttachment(attachment IAttachment)
	AddDepthStencilAttachment(attachment IAttachment)
	UnbindForRender()
	Delete()

	BlitToScreen()
	BlitToFramebuffer(framebuffer IFramebuffer)

	Width() int32
	Height() int32
	Samples() int32
}
