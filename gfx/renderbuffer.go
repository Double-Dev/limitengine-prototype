package gfx

import "github.com/double-dev/limitengine/gfx/framework"

var (
	renderbufferIndex = uint32(1)
	renderbuffers     = make(map[uint32]framework.IRenderbuffer)
)

func deleteRenderbuffers() {
	for _, iRenderbuffer := range renderbuffers {
		iRenderbuffer.Delete()
	}
	renderbuffers = nil
}

type Renderbuffer struct {
	id uint32
}

func NewRenderbuffer(multisample bool) *Renderbuffer {
	renderbuffer := &Renderbuffer{id: renderbufferIndex}
	renderbufferIndex++
	actionQueue = append(actionQueue, func() {
		renderbuffers[renderbuffer.id] = context.NewRenderbuffer(multisample)
	})
	return renderbuffer
}

// Attachment function:
func (renderbuffer *Renderbuffer) frameworkAttachment() framework.IAttachment {
	return renderbuffers[renderbuffer.id]
}

// DeleteRenderbuffer queues a gfx action that deletes the input renderbuffer.
func DeleteRenderbuffer(renderbuffer *Renderbuffer) {
	actionQueue = append(actionQueue, func() {
		iRenderbuffer := renderbuffers[renderbuffer.id]
		iRenderbuffer.Delete()
		delete(renderbuffers, renderbuffer.id)
	})
}
