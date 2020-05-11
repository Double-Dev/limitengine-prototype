package interaction

import (
	"github.com/double-dev/limitengine/gmath"
)

type ColliderComponent struct {
	IsTrigger bool

	AABB gmath.AABB
}
