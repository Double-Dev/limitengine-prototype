package gmath

// Vector2 is an alternative to Vector optimized to perform operations on 2
// floats.
// NOTE: While Vector2 is a slice and can have elements appended and deleted
// from it, it was never designed to do so and game devs that do so will do
// at their own risk!
type Vector2 []float32

func NewZeroVector2() Vector2 {
	return Vector2{0.0, 0.0}
}

func NewVector2(x, y float32) Vector2 {
	return Vector2{x, y}
}

func (vector Vector2) Set(x, y float32) Vector2 {
	vector[0] = x
	vector[1] = y
	return vector
}

func (vector Vector2) SetV(other Vector2) Vector2 {
	return vector.Set(other[0], other[1])
}

func (vector Vector2) Add(x, y float32) Vector2 {
	vector[0] += x
	vector[1] += y
	return vector
}

func (vector Vector2) AddV(other Vector2) Vector2 {
	return vector.Add(other[0], other[1])
}

func (vector Vector2) Sub(x, y float32) Vector2 {
	vector[0] -= x
	vector[1] -= y
	return vector
}

func (vector Vector2) SubV(other Vector2) Vector2 {
	return vector.Sub(other[0], other[1])
}

func (vector Vector2) Mul(x, y float32) Vector2 {
	vector[0] *= x
	vector[1] *= y
	return vector
}

func (vector Vector2) MulV(other Vector2) Vector2 {
	return vector.Mul(other[0], other[1])
}

func (vector Vector2) MulSc(scalar float32) Vector2 {
	vector[0] *= scalar
	vector[1] *= scalar
	return vector
}

func (vector Vector2) Div(x, y float32) Vector2 {
	vector[0] /= x
	vector[1] /= y
	return vector
}

func (vector Vector2) DivV(other Vector2) Vector2 {
	return vector.Div(other[0], other[1])
}

func (vector Vector2) DivSc(scalar float32) Vector2 {
	vector[0] /= scalar
	vector[1] /= scalar
	return vector
}

func (vector Vector2) Dot(other Vector2) float32 {
	return vector[0]*other[0] +
		vector[1]*other[1]
}

func (vector Vector2) IsGreater(other Vector2) bool {
	if vector[0] <= other[0] || vector[1] <= other[1] {
		return false
	}
	return true
}

func (vector Vector2) IsLess(other Vector2) bool {
	if vector[0] >= other[0] || vector[1] >= other[1] {
		return false
	}
	return true
}

func (vector Vector2) IsGreaterOrEqual(other Vector2) bool {
	if vector[0] < other[0] || vector[1] < other[1] {
		return false
	}
	return true
}

func (vector Vector2) IsLessOrEqual(other Vector2) bool {
	if vector[0] > other[0] || vector[1] > other[1] {
		return false
	}
	return true
}

func (vector Vector2) LenSq() float32 {
	return vector.Dot(vector)
}

func (vector Vector2) Len() float32 {
	return Sqrt(vector.LenSq())
}

func (vector Vector2) Normalize() Vector2 {
	return vector.DivSc(vector.Len())
}

func (vector Vector2) DstSq(other Vector2) float32 {
	return (vector[0]-other[0])*(vector[0]-other[0]) +
		(vector[1]-other[1])*(vector[1]-other[1])
}

func (vector Vector2) Dst(other Vector2) float32 {
	return Sqrt(vector.DstSq(other))
}

func (vector Vector2) Clone() Vector2 {
	return Vector2{vector[0], vector[1]}
}
