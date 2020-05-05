package interaction

import (
	"github.com/double-dev/limitengine/gmath"
)

type Grid2D struct {
	size     float32
	entities map[*InteractEntity][][2]int32
	grid     map[int32]map[int32][]*InteractEntity
}

func NewGrid2D(size float32) *Grid2D {
	return &Grid2D{
		size:     size,
		entities: make(map[*InteractEntity][][2]int32),
		grid:     make(map[int32]map[int32][]*InteractEntity),
	}
}

func (grid *Grid2D) Add(entity *InteractEntity) {
	min := entity.Collider.AABB.Min.Clone().AddV(entity.Transform.Position).DivSc(grid.size)
	max := entity.Collider.AABB.Max.Clone().AddV(entity.Transform.Position).DivSc(grid.size)
	for i := int32(min[0]); i <= int32(max[0]); i++ {
		if _, ok := grid.grid[i]; !ok {
			grid.grid[i] = make(map[int32][]*InteractEntity)
		}
		for j := int32(min[1]); j <= int32(max[1]); j++ {
			grid.grid[i][j] = append(grid.grid[i][j], entity)
			grid.entities[entity] = append(grid.entities[entity], [2]int32{i, j})
		}
	}
}

func (grid *Grid2D) Remove(entity *InteractEntity) {
	for _, info := range grid.entities[entity] {
		cell := grid.grid[info[0]][info[1]]
		for i, cellEntity := range cell {
			if cellEntity == entity {
				cell[i] = cell[len(cell)-1]
				grid.grid[info[0]][info[1]] = cell[:len(cell)-1]
				break
			}
		}
	}
	grid.entities[entity] = nil
}

func (grid *Grid2D) Update(entity *InteractEntity) {
	grid.Remove(entity)
	grid.Add(entity)
}

func (grid *Grid2D) Query(aabb gmath.AABB) []*InteractEntity {
	entityMap := make(map[*InteractEntity]bool)

	min := aabb.Min.Clone().DivSc(grid.size)
	max := aabb.Max.Clone().DivSc(grid.size)
	for i := int32(min[0]); i <= int32(max[0]); i++ {
		if _, ok := grid.grid[i]; ok {
			for j := int32(min[1]); j <= int32(max[1]); j++ {
				for _, entity := range grid.grid[i][j] {
					entityMap[entity] = false
				}
			}
		}
	}

	var entities []*InteractEntity
	for entity := range entityMap {
		checkAABB := gmath.NewAABB(
			entity.Collider.AABB.Min.Clone().AddV(entity.Transform.Position),
			entity.Collider.AABB.Max.Clone().AddV(entity.Transform.Position),
		)
		if aabb.IntersectsAABB2D(checkAABB) || aabb.ContainsAABB2D(checkAABB) || checkAABB.ContainsAABB2D(aabb) {
			entities = append(entities, entity)
		}
	}
	return entities
}
