package gmath

// Vector is an array of floats with util methods for vector mathematics.
type Vector []float32

// NewVector creates a new Vector from a float32 vararg.
func NewVector(components ...float32) Vector {
	return components
}

// NewVectorV creates a new Vector from a float32 array.
func NewVectorV(src []float32) Vector {
	dst := make([]float32, len(src))
	copy(dst, src)
	return dst
}

// NewVectorOfSize returns a zero vector with the number of components specified.
func NewVectorOfSize(size uint) Vector {
	return make([]float32, size)
}

// Set sets each element of this Vector object to the corresponding elements of a float32 vararg.
func (vector Vector) Set(other ...float32) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] = other[i]
	}
	return vector
}

// Add adds a float32 vararg to this Vector object.
func (vector Vector) Add(other ...float32) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] += other[i]
	}
	return vector
}

// AddSc adds a float32 scalar to every element within this Vector object.
func (vector Vector) AddSc(scalar float32) Vector {
	for i := 0; i < len(vector); i++ {
		vector[i] += scalar
	}
	return vector
}

// AddV adds a Vector to this Vector object.
func (vector Vector) AddV(other Vector) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] += other[i]
	}
	return vector
}

// Sub subtracts a float32 vararg from this Vector object.
func (vector Vector) Sub(other ...float32) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] -= other[i]
	}
	return vector
}

// SubSc subtracts a float32 scalar from every element within this Vector object.
func (vector Vector) SubSc(scalar float32) Vector {
	for i := 0; i < len(vector); i++ {
		vector[i] -= scalar
	}
	return vector
}

// SubV subtracts a Vector from this Vector object.
func (vector Vector) SubV(other Vector) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] -= other[i]
	}
	return vector
}

// Mul multiplies this Vector object by a float32 vararg.
func (vector Vector) Mul(other ...float32) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] *= other[i]
	}
	return vector
}

// MulSc multiplies this Vector object by a single float32 scalar.
func (vector Vector) MulSc(scalar float32) Vector {
	for i := 0; i < len(vector); i++ {
		vector[i] *= scalar
	}
	return vector
}

// MulV multiplies this Vector object's components by another Vector object's components.
func (vector Vector) MulV(other Vector) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] *= other[i]
	}
	return vector
}

// Div divides this Vector object by a float32 vararg.
func (vector Vector) Div(other ...float32) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] /= other[i]
	}
	return vector
}

// DivV divides this Vector object by a Vector.
func (vector Vector) DivV(other Vector) Vector {
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		vector[i] /= other[i]
	}
	return vector
}

// Dot returns a float32 result of this Vector's dot product.
func (vector Vector) Dot(other Vector) float32 {
	dot := float32(0.0)
	for i := 0; i < MinI(len(vector), len(other)); i++ {
		dot += vector[i] * other[i]
	}
	return dot
}

// Cross TODO: implement cross product math
func (vector Vector) Cross(other Vector) Vector {
	return NewVectorV(vector)
}

// LenSq returns a float32 result of this Vector's length squared.
func (vector Vector) LenSq() float32 {
	l := float32(0.0)
	for i := 0; i < len(vector); i++ {
		l += vector[i] * vector[i]
	}
	return l
}

// Len returns a float32 result of this Vector's length.
func (vector Vector) Len() float32 {
	return Sqrt(vector.LenSq())
}

// Dst returns a float32 result of this Vector's distance from another Vector.
func (vector Vector) Dst(other Vector) float32 {
	return Sqrt(vector.LenSq())
}

// Clone returns a new Vector with components equal to this Vector.
func (vector Vector) Clone() Vector {
	return NewVectorV(vector)
}
