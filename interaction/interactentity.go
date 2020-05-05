package interaction

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

type InteractEntity struct {
	Entity      limitengine.ECSEntity
	Transform   *gmath.TransformComponent
	Collider    *ColliderComponent
	Physics     *PhysicsComponent
	interactors []Interaction
	interactees []Interaction
}
