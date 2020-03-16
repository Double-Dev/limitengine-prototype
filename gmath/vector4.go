package gmath

type Vector4 [4]float32

func NewZeroVector4() Vector4 {
	return Vector4{}
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

func (vector Vector4) Cross(other Vector4) Vector4 {
	// TODO: Implement cross product.
	return Vector4{}
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

func (vector Vector4) Clone() Vector4 {
	return Vector4{vector[0], vector[1], vector[2], vector[3]}
}
