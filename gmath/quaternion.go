package gmath

import "fmt"

// Quaternion is a slice of floats with util methods for quaternion mathematics.
type Quaternion struct {
	quaternion []float32
}

// NewZeroQuaternion returns a zero quaternion with the dimension specified.
func NewZeroQuaternion(dimension int) *Quaternion {
	return &Quaternion{make([]float32, dimension+1)}
}

// NewQuaternion returns a quaternion with the dimen specified.
func NewQuaternion(axis ...float32) *Quaternion {
	return NewZeroQuaternion(len(axis)).Set(axis...)
}

// SetAxis sets the specified axis of this Quaternion object to the float32 value.
func (quaternion *Quaternion) SetAxis(index int, value float32) {
	quaternion.quaternion[index] = value
}

// GetAxis returns the specified axis of this Quaternion object.
func (quaternion *Quaternion) GetAxis(index int) float32 {
	return quaternion.quaternion[index]
}

// Set sets each axis of this Quaternion object to the corresponding axis of a float32 vararg.
func (quaternion *Quaternion) Set(other ...float32) *Quaternion {
	for i := 0; i < MinI(len(quaternion.quaternion), len(other)); i++ {
		quaternion.quaternion[i] = other[i]
	}
	return quaternion
}

// SetQ sets each axis of this Quaternion object to the corresponding axis of a Quaternion.
func (quaternion *Quaternion) SetQ(other *Quaternion) *Quaternion {
	return quaternion.Set(other.quaternion...)
}

// Add adds a float32 vararg to this Quaternion object.
func (quaternion *Quaternion) Add(other ...float32) *Quaternion {
	for i := 0; i < MinI(len(quaternion.quaternion), len(other)); i++ {
		quaternion.quaternion[i] += other[i]
	}
	return quaternion
}

// AddQ adds a Quaternion vararg to this Quaternion object.
func (quaternion *Quaternion) AddQ(other *Quaternion) *Quaternion {
	return quaternion.Add(other.quaternion...)
}

// AddSc adds a float32 scalar to every axis within quaternion Quaternion object.
func (quaternion *Quaternion) AddSc(scalar float32) *Quaternion {
	for i := 0; i < len(quaternion.quaternion); i++ {
		quaternion.quaternion[i] += scalar
	}
	return quaternion
}

// Sub subtracts a float32 vararg from this Quaternion object.
func (quaternion *Quaternion) Sub(other ...float32) *Quaternion {
	for i := 0; i < MinI(len(quaternion.quaternion), len(other)); i++ {
		quaternion.quaternion[i] -= other[i]
	}
	return quaternion
}

// SubQ subtracts a quaternion from this Quaternion object.
func (quaternion *Quaternion) SubQ(other *Quaternion) *Quaternion {
	return quaternion.Sub(other.quaternion...)
}

// SubSc subtracts a float32 scalar from every axis within quaternion Quaternion object.
func (quaternion *Quaternion) SubSc(scalar float32) *Quaternion {
	for i := 0; i < len(quaternion.quaternion); i++ {
		quaternion.quaternion[i] -= scalar
	}
	return quaternion
}

// Mul multiplies this Quaternion object by a float32 vararg.
func (quaternion *Quaternion) Mul(other ...float32) *Quaternion {
	for i := 0; i < MinI(len(quaternion.quaternion), len(other)); i++ {
		quaternion.quaternion[i] *= other[i]
	}
	return quaternion
}

// MulQ multiplies this Quaternion object by a Quaternion.
func (quaternion *Quaternion) MulQ(other *Quaternion) *Quaternion {
	return quaternion.Mul(other.quaternion...)
}

// MulSc multiplies this Quaternion object by a single float32 scalar.
func (quaternion *Quaternion) MulSc(scalar float32) *Quaternion {
	for i := 0; i < len(quaternion.quaternion); i++ {
		quaternion.quaternion[i] *= scalar
	}
	return quaternion
}

// Div divides this Quaternion object by a float32 vararg.
func (quaternion *Quaternion) Div(other ...float32) *Quaternion {
	for i := 0; i < MinI(len(quaternion.quaternion), len(other)); i++ {
		quaternion.quaternion[i] /= other[i]
	}
	return quaternion
}

// DivQ divides this Quaternion object by a Quaternion.
func (quaternion *Quaternion) DivQ(other *Quaternion) *Quaternion {
	return quaternion.Div(other.quaternion...)
}

// DivSc divides this Quaternion object by a single float32 scalar.
func (quaternion *Quaternion) DivSc(scalar float32) *Quaternion {
	for i := 0; i < len(quaternion.quaternion); i++ {
		quaternion.quaternion[i] /= scalar
	}
	return quaternion
}

// Dot returns a float32 result of this Quaternion's dot product.
func (quaternion *Quaternion) Dot(other *Quaternion) float32 {
	dot := float32(0.0)
	for i := 0; i < MinI(len(quaternion.quaternion), len(other.quaternion)); i++ {
		dot += quaternion.quaternion[i] * other.quaternion[i]
	}
	return dot
}

// LenSq returns a float32 result of this Quaternion's length squared.
func (quaternion *Quaternion) LenSq() float32 {
	l := float32(0.0)
	for i := 0; i < len(quaternion.quaternion); i++ {
		l += quaternion.quaternion[i] * quaternion.quaternion[i]
	}
	return l
}

// Len returns a float32 result of this Quaternion's length.
func (quaternion *Quaternion) Len() float32 {
	return Sqrt(quaternion.LenSq())
}

// Normalize normalizes this Quaternion.
func (quaternion *Quaternion) Normalize() *Quaternion {
	return quaternion.DivSc(quaternion.Len())
}

// Clone returns a new Quaternion with components equal to this Quaternion.
func (quaternion *Quaternion) Clone() *Quaternion {
	out := make([]float32, len(quaternion.quaternion))
	copy(out, quaternion.quaternion)
	return &Quaternion{out}
}

func (quaternion *Quaternion) String() string {
	return fmt.Sprint(quaternion.quaternion)
}
