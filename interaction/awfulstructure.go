package interaction

import (
	"github.com/double-dev/limitengine/gmath"
)

type AwfulStructure struct {
	entities []*InteractEntity
}

func NewAwfulStructure() *AwfulStructure {
	return &AwfulStructure{}
}

func (awfulStructure *AwfulStructure) Add(entity *InteractEntity) {
	awfulStructure.entities = append(awfulStructure.entities, entity)
}

func (awfulStructure *AwfulStructure) Remove(entity *InteractEntity) {
	for i, structureEntity := range awfulStructure.entities {
		if structureEntity == entity {
			awfulStructure.entities[i] = awfulStructure.entities[len(awfulStructure.entities)-1]
			awfulStructure.entities = awfulStructure.entities[:len(awfulStructure.entities)-1]
			break
		}
	}
}

func (awfulStructure *AwfulStructure) Update(entity *InteractEntity) {
}

func (awfulStructure *AwfulStructure) Query(aabb gmath.AABB) []*InteractEntity {
	var query []*InteractEntity
	for _, entity := range awfulStructure.entities {
		checkAABB := gmath.NewAABB(
			entity.collider.AABB.Min.Clone().AddV(entity.transform.Position),
			entity.collider.AABB.Max.Clone().AddV(entity.transform.Position),
		)
		if aabb.IntersectsAABB2D(checkAABB) ||
			aabb.ContainsAABB2D(checkAABB) ||
			checkAABB.ContainsAABB2D(aabb) {
			query = append(query, entity)
		}
	}
	return query
}
