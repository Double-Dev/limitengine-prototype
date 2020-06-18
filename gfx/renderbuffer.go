package gfx

import "github.com/double-dev/limitengine/gfx/framework"

var (
	renderbufferIndex = uint32(1)
	renderbuffers     = make(map[uint32]framework.IRenderbuffer)
)

type Renderbuffer struct {
	id uint32
}

func NewRenderbuffer(multisample bool) *Renderbuffer {
	renderbuffer := &Renderbuffer{
		id: renderbufferIndex,
	}
	renderbufferIndex++
	actionQueue = append(actionQueue, func() {
		renderbuffers[renderbuffer.id] = context.NewRenderbuffer(multisample)
	})
	return renderbuffer
}

// Attachment function:
func (renderbuffer *Renderbuffer) getFrameworkAttachment() framework.IAttachment {
	return renderbuffers[renderbuffer.id]
}
