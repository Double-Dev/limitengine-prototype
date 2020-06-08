package framework

type IFramebuffer interface {
	BindForRender()
	SetColorAttachment(attachment IAttachment)
	SetDepthAttachment(attachment IAttachment)
	UnbindForRender()
	Delete()

	BlitToScreen()
	BlitToFramebuffer(framebuffer IFramebuffer)

	Resize(width, height int32)
	Width() int32
	Height() int32
	Samples() int32
}
