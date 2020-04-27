package gmath

// Vector3 is an alternative to Vector optimized to perform operations on 3
// floats.
// NOTE: While Vector3 is a slice and can have elements appended and deleted
// from it, it was never designed to do so and game devs that do so will do
// at their own risk!
type Vector3 []float32

func NewZeroVector3() Vector3 {
	return Vector3{0.0, 0.0, 0.0}
}

func NewVector3(x, y, z float32) Vector3 {
	return Vector3{x, y, z}
}

func (vector Vector3) Set(x, y, z float32) Vector3 {
	vector[0] = x
	vector[1] = y
	vector[2] = z
	return vector
}

func (vector Vector3) SetV(other Vector3) Vector3 {
	return vector.Set(other[0], other[1], other[2])
}

func (vector Vector3) Add(x, y, z float32) Vector3 {
	vector[0] += x
	vector[1] += y
	vector[2] += z
	return vector
}

func (vector Vector3) AddV(other Vector3) Vector3 {
	return vector.Add(other[0], other[1], other[2])
}

func (vector Vector3) Sub(x, y, z float32) Vector3 {
	vector[0] -= x
	vector[1] -= y
	vector[2] -= z
	return vector
}

func (vector Vector3) SubV(other Vector3) Vector3 {
	return vector.Sub(other[0], other[1], other[2])
}

func (vector Vector3) Mul(x, y, z float32) Vector3 {
	vector[0] *= x
	vector[1] *= y
	vector[2] *= z
	return vector
}

func (vector Vector3) MulV(other Vector3) Vector3 {
	return vector.Mul(other[0], other[1], other[2])
}

func (vector Vector3) MulSc(scalar float32) Vector3 {
	vector[0] *= scalar
	vector[1] *= scalar
	vector[2] *= scalar
	return vector
}

func (vector Vector3) Div(x, y, z float32) Vector3 {
	vector[0] /= x
	vector[1] /= y
	vector[2] /= z
	return vector
}

func (vector Vector3) DivV(other Vector3) Vector3 {
	return vector.Div(other[0], other[1], other[2])
}

func (vector Vector3) DivSc(scalar float32) Vector3 {
	vector[0] /= scalar
	vector[1] /= scalar
	vector[2] /= scalar
	return vector
}

func (vector Vector3) Dot(other Vector3) float32 {
	return vector[0]*other[0] +
		vector[1]*other[1] +
		vector[2]*other[2]
}

func (vector Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		vector[1]*other[2] - vector[2]*other[1],
		vector[2]*other[0] - vector[0]*other[2],
		vector[0]*other[1] - vector[1]*other[0],
	}
}

func (vector Vector3) LenSq() float32 {
	return vector[0]*vector[0] +
		vector[1]*vector[1] +
		vector[2]*vector[2]
}

func (vector Vector3) Len() float32 {
	return Sqrt(vector.LenSq())
}

func (vector Vector3) Normalize() Vector3 {
	return vector.DivSc(vector.Len())
}

func (vector Vector3) DstSq(other Vector3) float32 {
	return (vector[0]-other[0])*(vector[0]-other[0]) +
		(vector[1]-other[1])*(vector[1]-other[1]) +
		(vector[2]-other[2])*(vector[2]-other[2])
}

func (vector Vector3) Dst(other Vector3) float32 {
	return Sqrt(vector.DstSq(other))
}

func (vector Vector3) Clone() Vector3 {
	return Vector3{vector[0], vector[1], vector[2]}
}
