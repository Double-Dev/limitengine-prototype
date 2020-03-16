package gmath

import "fmt"

// Vector is a slice of floats with util methods for vector mathematics.
type Vector struct {
	vector []float32
}

// NewZeroVector returns a zero vector with the number of components specified.
func NewZeroVector(size int) *Vector {
	return &Vector{make([]float32, size)}
}

// NewVector returns a vector with the components specified.
func NewVector(components ...float32) *Vector {
	return NewZeroVector(len(components)).Set(components...)
}

// SetElement sets the specified element of this Vector object to the float32 value.
func (vector *Vector) SetElement(index int, value float32) {
	vector.vector[index] = value
}

// GetElement returns the specified element of this Vector object.
func (vector *Vector) GetElement(index int) float32 {
	return vector.vector[index]
}

// Dimension returns the dimension of this Vector object.
func (vector *Vector) Dimension() int {
	return len(vector.vector)
}

// Set sets each element of this Vector object to the corresponding elements of a float32 vararg.
func (vector *Vector) Set(other ...float32) *Vector {
	for i := 0; i < MinI(len(vector.vector), len(other)); i++ {
		vector.vector[i] = other[i]
	}
	return vector
}

// SetV sets each element of this Vector object to the corresponding elements of a Vector.
func (vector *Vector) SetV(other *Vector) *Vector {
	return vector.Set(other.vector...)
}

// Add adds a float32 vararg to this Vector object.
func (vector *Vector) Add(other ...float32) *Vector {
	for i := 0; i < MinI(len(vector.vector), len(other)); i++ {
		vector.vector[i] += other[i]
	}
	return vector
}

// AddV adds a Vector vararg to this Vector object.
func (vector *Vector) AddV(other *Vector) *Vector {
	return vector.Add(other.vector...)
}

// AddSc adds a float32 scalar to every element within vector Vector object.
func (vector *Vector) AddSc(scalar float32) *Vector {
	for i := 0; i < len(vector.vector); i++ {
		vector.vector[i] += scalar
	}
	return vector
}

// Sub subtracts a float32 vararg from this Vector object.
func (vector *Vector) Sub(other ...float32) *Vector {
	for i := 0; i < MinI(len(vector.vector), len(other)); i++ {
		vector.vector[i] -= other[i]
	}
	return vector
}

// SubV subtracts a vector from this Vector object.
func (vector *Vector) SubV(other *Vector) *Vector {
	return vector.Sub(other.vector...)
}

// SubSc subtracts a float32 scalar from every element within vector Vector object.
func (vector *Vector) SubSc(scalar float32) *Vector {
	for i := 0; i < len(vector.vector); i++ {
		vector.vector[i] -= scalar
	}
	return vector
}

// Mul multiplies this Vector object by a float32 vararg.
func (vector *Vector) Mul(other ...float32) *Vector {
	for i := 0; i < MinI(len(vector.vector), len(other)); i++ {
		vector.vector[i] *= other[i]
	}
	return vector
}

// MulV multiplies this Vector object by a Vector.
func (vector *Vector) MulV(other *Vector) *Vector {
	return vector.Mul(other.vector...)
}

// MulSc multiplies this Vector object by a single float32 scalar.
func (vector *Vector) MulSc(scalar float32) *Vector {
	for i := 0; i < len(vector.vector); i++ {
		vector.vector[i] *= scalar
	}
	return vector
}

// Div divides this Vector object by a float32 vararg.
func (vector *Vector) Div(other ...float32) *Vector {
	for i := 0; i < MinI(len(vector.vector), len(other)); i++ {
		vector.vector[i] /= other[i]
	}
	return vector
}

// DivV divides this Vector object by a Vector.
func (vector *Vector) DivV(other *Vector) *Vector {
	return vector.Div(other.vector...)
}

// DivSc divides this Vector object by a single float32 scalar.
func (vector *Vector) DivSc(scalar float32) *Vector {
	for i := 0; i < len(vector.vector); i++ {
		vector.vector[i] /= scalar
	}
	return vector
}

// Dot returns a float32 result of this Vector's dot product.
func (vector *Vector) Dot(other *Vector) float32 {
	dot := float32(0.0)
	for i := 0; i < MinI(len(vector.vector), len(other.vector)); i++ {
		dot += vector.vector[i] * other.vector[i]
	}
	return dot
}

// Cross TODO: implement cross product math
func (vector *Vector) Cross(other *Vector) *Vector {
	return nil
}

// LenSq returns a float32 result of this Vector's length squared.
func (vector *Vector) LenSq() float32 {
	l := float32(0.0)
	for i := 0; i < len(vector.vector); i++ {
		l += vector.vector[i] * vector.vector[i]
	}
	return l
}

// Len returns a float32 result of this Vector's length.
func (vector *Vector) Len() float32 {
	return Sqrt(vector.LenSq())
}

// Normalize normalizes this Vector.
func (vector *Vector) Normalize() *Vector {
	return vector.DivSc(vector.Len())
}

// Dst returns a float32 result of this Vector's distance from another Vector.
func (vector *Vector) Dst(other *Vector) float32 {
	return Sqrt(vector.Clone().SubV(other).LenSq())
}

// Clone returns a new Vector with components equal to this Vector.
func (vector *Vector) Clone() *Vector {
	out := make([]float32, len(vector.vector))
	copy(out, vector.vector)
	return &Vector{out}
}

func (vector *Vector) String() string {
	return fmt.Sprint(vector.vector)
}
