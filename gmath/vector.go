package gmath

// Vector is an array of floats with util methods for vector mathematics.
type Vector []float32
type Vector2 [2]float32
type Vector3 [3]float32
type Vector4 [4]float32

// NewZeroVector returns a zero vector with the number of components specified.
func NewZeroVector(size uint) Vector {
	return make([]float32, size)
}

// Set sets each element of this Vector object to the corresponding elements of a float32 vararg.
func (this Vector) Set(other ...float32) Vector {
	for i := 0; i < MinI(len(this), len(other)); i++ {
		this[i] = other[i]
	}
	return this
}

// Add adds a float32 vararg to this Vector object.
func (this Vector) Add(other ...float32) Vector {
	for i := 0; i < MinI(len(this), len(other)); i++ {
		this[i] += other[i]
	}
	return this
}

// AddSc adds a float32 scalar to every element within this Vector object.
func (this Vector) AddSc(scalar float32) Vector {
	for i := 0; i < len(this); i++ {
		this[i] += scalar
	}
	return this
}

// Sub subtracts a float32 vararg from this Vector object.
func (this Vector) Sub(other ...float32) Vector {
	for i := 0; i < MinI(len(this), len(other)); i++ {
		this[i] -= other[i]
	}
	return this
}

// SubSc subtracts a float32 scalar from every element within this Vector object.
func (this Vector) SubSc(scalar float32) Vector {
	for i := 0; i < len(this); i++ {
		this[i] -= scalar
	}
	return this
}

// Mul multiplies this Vector object by a float32 vararg.
func (this Vector) Mul(other ...float32) Vector {
	for i := 0; i < MinI(len(this), len(other)); i++ {
		this[i] *= other[i]
	}
	return this
}

// MulSc multiplies this Vector object by a single float32 scalar.
func (this Vector) MulSc(scalar float32) Vector {
	for i := 0; i < len(this); i++ {
		this[i] *= scalar
	}
	return this
}

// Div divides this Vector object by a float32 vararg.
func (this Vector) Div(other ...float32) Vector {
	for i := 0; i < MinI(len(this), len(other)); i++ {
		this[i] /= other[i]
	}
	return this
}

// Dot returns a float32 result of this Vector's dot product.
func (this Vector) Dot(other Vector) float32 {
	dot := float32(0.0)
	for i := 0; i < MinI(len(this), len(other)); i++ {
		dot += this[i] * other[i]
	}
	return dot
}

// Cross TODO: implement cross product math
func (this Vector) Cross(other Vector) Vector {
	return nil
}

// LenSq returns a float32 result of this Vector's length squared.
func (this Vector) LenSq() float32 {
	l := float32(0.0)
	for i := 0; i < len(this); i++ {
		l += this[i] * this[i]
	}
	return l
}

// Len returns a float32 result of this Vector's length.
func (this Vector) Len() float32 {
	return Sqrt(this.LenSq())
}

// Dst returns a float32 result of this Vector's distance from another Vector.
func (this Vector) Dst(other Vector) float32 {
	return Sqrt(this.Clone().Sub(other...).LenSq())
}

// Clone returns a new Vector with components equal to this Vector.
func (this Vector) Clone() Vector {
	out := make([]float32, len(this))
	copy(out, this)
	return out
}

func (this Vector) ToVector2() Vector2 {
	out := Vector2{}
	for i := 0; i < MinI(len(this), 2); i++ {
		out[i] = this[i]
	}
	return out
}

func (this Vector) ToVector3() Vector3 {
	out := Vector3{}
	for i := 0; i < MinI(len(this), 3); i++ {
		out[i] = this[i]
	}
	return out
}

func (this Vector) ToVector4() Vector4 {
	out := Vector4{}
	for i := 0; i < MinI(len(this), 4); i++ {
		out[i] = this[i]
	}
	return out
}
