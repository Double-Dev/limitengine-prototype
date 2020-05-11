package interaction

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type InteractEntity struct {
	Entity                    limitengine.ECSEntity
	Transform                 *gmath.TransformComponent
	Motion                    *gmath.MotionComponent
	Collider                  *ColliderComponent
	previousCollidingEntities map[*InteractEntity]gmath.Vector3
	collidingEntities         map[*InteractEntity]gmath.Vector3
	interactors               []Interaction
	interactees               []Interaction
}
