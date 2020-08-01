package gfx

import "github.com/double-dev/limitengine/gfx/framework"

type Attachment interface {
	frameworkAttachment() framework.IAttachment
}
