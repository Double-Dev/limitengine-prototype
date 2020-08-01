package interaction

import "github.com/double-dev/limitengine/gmath"

type SpacialStructure interface {
	Add(entity *InteractEntity)
	Remove(entity *InteractEntity)
	Update(entity *InteractEntity)
	Query(aabb gmath.AABB) []*InteractEntity
}
