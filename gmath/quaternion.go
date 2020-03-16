package gmath

import "fmt"

// Quaternion is a slice of floats with util methods for quaternion mathematics.
type Quaternion struct {
	quaternion *Vector
}

// NewIdentityQuaternion returns an identity quaternion with the dimension specified.
func NewIdentityQuaternion(dimension int) *Quaternion {
	axis := make([]float32, dimension)
	axis[0] = 1.0
	return NewQuaternion(0.0, axis...)
}

// NewQuaternion returns a quaternion with the axis vector specified.
func NewQuaternion(angle float32, axis ...float32) *Quaternion {
	return (&Quaternion{NewZeroVector(len(axis) + 1)}).Set(angle, axis...)
}

// SetElement sets the specified element of this Vector object to the float32 value.
func (quaternion *Quaternion) SetElement(index int, value float32) {
	quaternion.quaternion.SetElement(index, value)
}

// GetElement returns the specified element of this Vector object.
func (quaternion *Quaternion) GetElement(index int) float32 {
	return quaternion.quaternion.GetElement(index)
}

// Set sets each axis of this Quaternion object to the corresponding axis of a float32 vararg.
func (quaternion *Quaternion) Set(angle float32, axis ...float32) *Quaternion {
	normAxis := NewVector(axis...)
	normAxis.Normalize()
	sin := Sin(angle)
	for i := 0; i < MinI(quaternion.quaternion.Dimension()-1, len(normAxis.vector)); i++ {
		quaternion.quaternion.SetElement(i, sin*normAxis.GetElement(i))
	}
	quaternion.quaternion.SetElement(quaternion.quaternion.Dimension()-1, Cos(angle))
	return quaternion
}

// SetQ sets each axis of this Quaternion object to the corresponding axis of a Quaternion.
func (quaternion *Quaternion) SetQ(other *Quaternion) *Quaternion {
	for i := 0; i < MinI(quaternion.quaternion.Dimension(), other.quaternion.Dimension()); i++ {
		quaternion.quaternion.SetElement(i, other.quaternion.GetElement(i))
	}
	return quaternion
}

// Mul multiplies this Quaternion object by another quaternion.
func (quaternion *Quaternion) Mul(angle float32, axis ...float32) *Quaternion {
	other := NewQuaternion(angle, axis...)
	return quaternion.MulQ(other)
}

// MulQ multiplies this Quaternion object by another quaternion.
func (quaternion *Quaternion) MulQ(other *Quaternion) *Quaternion {
	t0 := (quaternion.quaternion.GetElement(2) - quaternion.quaternion.GetElement(1)) * (other.quaternion.GetElement(1) - other.quaternion.GetElement(2))
	t1 := (quaternion.quaternion.GetElement(3) + quaternion.quaternion.GetElement(0)) * (other.quaternion.GetElement(3) + other.quaternion.GetElement(0))
	t2 := (quaternion.quaternion.GetElement(3) - quaternion.quaternion.GetElement(0)) * (other.quaternion.GetElement(1) + other.quaternion.GetElement(2))
	t3 := (quaternion.quaternion.GetElement(2) + quaternion.quaternion.GetElement(1)) * (other.quaternion.GetElement(3) - other.quaternion.GetElement(0))
	t4 := (quaternion.quaternion.GetElement(2) - quaternion.quaternion.GetElement(0)) * (other.quaternion.GetElement(0) - other.quaternion.GetElement(1))
	t5 := (quaternion.quaternion.GetElement(2) + quaternion.quaternion.GetElement(0)) * (other.quaternion.GetElement(0) + other.quaternion.GetElement(1))
	t6 := (quaternion.quaternion.GetElement(3) + quaternion.quaternion.GetElement(1)) * (other.quaternion.GetElement(3) - other.quaternion.GetElement(2))
	t7 := (quaternion.quaternion.GetElement(3) - quaternion.quaternion.GetElement(1)) * (other.quaternion.GetElement(3) + other.quaternion.GetElement(2))
	t8 := t5 + t6 + t7
	t9 := 0.5 * (t4 + t8)
	// TODO: Make this work for all dimensions.
	quaternion.quaternion.Set(
		t1+t9-t8,
		t2+t9-t7,
		t3+t9-t6,
		t0+t9-t5,
	)
	return quaternion
}

// MulSc scales this Quaternion object by a float32.
func (quaternion *Quaternion) MulSc(scalar float32) *Quaternion {
	quaternion.quaternion.MulSc(scalar)
	quaternion.quaternion.Normalize()
	return quaternion
}

// // Sub subtracts a float32 vararg from this Quaternion object.
// func (quaternion *Quaternion) Sub(angle float32, axis ...float32) *Quaternion {
// 	other := NewQuaternion(angle, axis...)
// 	return quaternion.SubQ(other)
// }

// // SubQ subtracts a quaternion from this Quaternion object.
// func (quaternion *Quaternion) SubQ(other *Quaternion) *Quaternion {
// 	quaternion.quaternion.DivV(other.quaternion)
// 	return quaternion
// }

// Inverse inverts the quaternion.
func (quaternion *Quaternion) Inverse() *Quaternion {
	l := quaternion.quaternion.LenSq()
	for i := 0; i < quaternion.quaternion.Dimension()-1; i++ {
		quaternion.quaternion.SetElement(i, -1.0*quaternion.quaternion.GetElement(i))
	}
	quaternion.quaternion.DivSc(l)
	quaternion.quaternion.Normalize()
	return quaternion
}

// Clone returns a new Quaternion with components equal to this Quaternion.
func (quaternion *Quaternion) Clone() *Quaternion {
	return &Quaternion{quaternion.quaternion.Clone()}
}

func (quaternion *Quaternion) String() string {
	return fmt.Sprint(quaternion.quaternion)
}
