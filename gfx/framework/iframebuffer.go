package framework

type IFramebuffer interface {
	Bind()
	SetColorAttachment(attachment IAttachment)
	SetDepthAttachment(attachment IAttachment)
	Delete()

	BlitToFramebuffer(framebuffer IFramebuffer)

	Resize(width, height int32)
	Width() int32
	Height() int32
	Samples() int32
}
