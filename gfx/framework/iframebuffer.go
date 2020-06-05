package framework

type IFramebuffer interface {
	Bind()
	AddColorAttachment(attachment IAttachment)
	AddDepthAttachment(attachment IAttachment)
	AddStencilAttachment(attachment IAttachment)
	AddDepthStencilAttachment(attachment IAttachment)
	Unbind()
	Delete()

	BlitToScreen()

	Width() int32
	Height() int32
	Samples() int32
}
