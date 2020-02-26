package gfx

var frameBuffers = make(map[FrameBuffer]uint32)

// FrameBuffer is a gfx framebuffer.
type FrameBuffer int
