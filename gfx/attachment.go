package gfx

import "github.com/double-dev/limitengine/gfx/framework"

type Attachment interface {
	getFrameworkAttachment() framework.IAttachment
}
