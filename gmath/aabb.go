package gmath

type AABB struct {
	Min, Max Vector3
}

func NewAABB(min, max Vector3) AABB {
	return AABB{
		Min: min,
		Max: max,
	}
}

func (aabb AABB) Contains(min, max Vector3) bool {
	return min.IsGreater(aabb.Min) && max.IsLess(aabb.Max)
}

func (aabb AABB) Contains2D(min, max Vector3) bool {
	return Vector2(min).IsGreater(Vector2(aabb.Min)) && Vector2(max).IsLess(Vector2(aabb.Max))
}

func (aabb AABB) ContainsAABB(other AABB) bool {
	return aabb.Contains(other.Min, other.Max)
}

func (aabb AABB) ContainsAABB2D(other AABB) bool {
	return aabb.Contains2D(other.Min, other.Max)
}

func (aabb AABB) ContainsV(vector Vector3) bool {
	return aabb.Contains(vector, vector)
}

func (aabb AABB) ContainsV2D(vector Vector3) bool {
	return aabb.Contains2D(vector, vector)
}

func (aabb AABB) Intersects(min, max Vector3) bool {
	return aabb.Min.IsLess(max) && aabb.Max.IsGreater(min)
}

func (aabb AABB) Intersects2D(min, max Vector3) bool {
	return Vector2(aabb.Min).IsLess(Vector2(max)) && Vector2(aabb.Max).IsGreater(Vector2(min))
}

func (aabb AABB) IntersectsAABB(other AABB) bool {
	return aabb.Intersects(other.Min, other.Max)
}

func (aabb AABB) IntersectsAABB2D(other AABB) bool {
	return aabb.Intersects2D(other.Min, other.Max)
}
