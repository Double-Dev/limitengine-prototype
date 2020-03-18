package utils

import (
	"github.com/double-dev/limitengine/gmath"
)

type TransformComponent struct {
	Position gmath.Vector3
	Rotation gmath.Quaternion
	Scale    gmath.Vector3
}
