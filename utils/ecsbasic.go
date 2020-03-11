package utils

import (
	"github.com/double-dev/limitengine/gmath"
)

type TransformComponent struct {
	Position *gmath.Vector
	Rotation *gmath.Quaternion
	Scale    *gmath.Vector
}
