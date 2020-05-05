package interaction

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type InteractEntity struct {
	entity      limitengine.ECSEntity
	transform   *gmath.TransformComponent
	collider    *ColliderComponent
	physics     *PhysicsComponent
	interactors []Interaction
	interactees []Interaction
}
