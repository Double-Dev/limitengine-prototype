package gmath

// Vector4 is an alternative to Vector optimized to perform operations on 4
// floats.
// NOTE: While Vector4 is a slice and can have elements appended and deleted
// from it, it was never designed to do so and game devs that do so will do
// at their own risk!
type Vector4 []float32

func NewZeroVector4() Vector4 {
	return Vector4{0.0, 0.0, 0.0, 0.0}
}

func NewVector4(x, y, z, w float32) Vector4 {
	return Vector4{x, y, z, w}
}

func (vector Vector4) Set(x, y, z, w float32) Vector4 {
	vector[0] = x
	vector[1] = y
	vector[2] = z
	vector[3] = w
	return vector
}

func (vector Vector4) SetV(other Vector4) Vector4 {
	return vector.Set(other[0], other[1], other[2], other[3])
}

func (vector Vector4) Add(x, y, z, w float32) Vector4 {
	vector[0] += x
	vector[1] += y
	vector[2] += z
	vector[3] += w
	return vector
}

func (vector Vector4) AddV(other Vector4) Vector4 {
	return vector.Add(other[0], other[1], other[2], other[3])
}

func (vector Vector4) Sub(x, y, z, w float32) Vector4 {
	vector[0] -= x
	vector[1] -= y
	vector[2] -= z
	vector[3] -= w
	return vector
}

func (vector Vector4) SubV(other Vector4) Vector4 {
	return vector.Sub(other[0], other[1], other[2], other[3])
}

func (vector Vector4) Mul(x, y, z, w float32) Vector4 {
	vector[0] *= x
	vector[1] *= y
	vector[2] *= z
	vector[3] *= w
	return vector
}

func (vector Vector4) MulV(other Vector4) Vector4 {
	return vector.Mul(other[0], other[1], other[2], other[3])
}

func (vector Vector4) MulSc(scalar float32) Vector4 {
	vector[0] *= scalar
	vector[1] *= scalar
	vector[2] *= scalar
	vector[3] *= scalar
	return vector
}

func (vector Vector4) Div(x, y, z, w float32) Vector4 {
	vector[0] /= x
	vector[1] /= y
	vector[2] /= z
	vector[3] /= w
	return vector
}

func (vector Vector4) DivV(other Vector4) Vector4 {
	return vector.Div(other[0], other[1], other[2], other[3])
}

func (vector Vector4) DivSc(scalar float32) Vector4 {
	vector[0] /= scalar
	vector[1] /= scalar
	vector[2] /= scalar
	vector[3] /= scalar
	return vector
}

func (vector Vector4) Dot(other Vector4) float32 {
	return vector[0]*other[0] +
		vector[1]*other[1] +
		vector[2]*other[2] +
		vector[3]*other[3]
}

func (vector Vector4) IsGreater(other Vector4) bool {
	if vector[0] <= other[0] || vector[1] <= other[1] || vector[2] <= other[2] || vector[3] <= other[3] {
		return false
	}
	return true
}

func (vector Vector4) IsLess(other Vector4) bool {
	if vector[0] >= other[0] || vector[1] >= other[1] || vector[2] >= other[2] || vector[3] >= other[3] {
		return false
	}
	return true
}

func (vector Vector4) IsGreaterOrEqual(other Vector4) bool {
	if vector[0] < other[0] || vector[1] < other[1] || vector[2] < other[2] || vector[3] < other[3] {
		return false
	}
	return true
}

func (vector Vector4) IsLessOrEqual(other Vector4) bool {
	if vector[0] > other[0] || vector[1] > other[1] || vector[2] > other[2] || vector[3] > other[3] {
		return false
	}
	return true
}

func (vector Vector4) LenSq() float32 {
	return vector[0]*vector[0] +
		vector[1]*vector[1] +
		vector[2]*vector[2] +
		vector[3]*vector[3]
}

func (vector Vector4) Len() float32 {
	return Sqrt(vector.LenSq())
}

func (vector Vector4) Normalize() Vector4 {
	return vector.DivSc(vector.Len())
}

func (vector Vector4) DstSq(other Vector4) float32 {
	return (vector[0]-other[0])*(vector[0]-other[0]) +
		(vector[1]-other[1])*(vector[1]-other[1]) +
		(vector[2]-other[2])*(vector[2]-other[2]) +
		(vector[3]-other[3])*(vector[3]-other[3])
}

func (vector Vector4) Dst(other Vector4) float32 {
	return Sqrt(vector.DstSq(other))
}

func (vector Vector4) ToVector3() Vector3 {
	return Vector3{vector[0], vector[1], vector[2]}
}

func (vector Vector4) Clone() Vector4 {
	return Vector4{vector[0], vector[1], vector[2], vector[3]}
}
