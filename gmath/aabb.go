package gmath

type AABB struct {
	Min, Max Vector3
}

func CreateAABB(min, max Vector3) AABB {
	return AABB{
		Min: min,
		Max: max,
	}
}

func (aabb AABB) Contains(min, max Vector3) bool {
	return min.IsGreater(aabb.Min) && max.IsLess(aabb.Max)
}

func (aabb AABB) ContainsAABB(other AABB) bool {
	return aabb.Contains(other.Min, other.Max)
}

func (aabb AABB) ContainsV(vector Vector3) bool {
	return aabb.Contains(vector, vector)
}

func (aabb AABB) Intersects(min, max Vector3) bool {
	return aabb.Min.IsLess(max) && aabb.Max.IsGreater(min)
}

func (aabb AABB) IntersectsAABB(other AABB) bool {
	return aabb.Intersects(other.Min, other.Max)
}
