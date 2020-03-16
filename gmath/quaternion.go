package gmath

// Quaternion is a slice of floats with util methods for quaternion mathematics.
type Quaternion struct {
	vector Vector4
}

// NewIdentityQuaternion returns an identity quaternion with the dimension specified.
func NewIdentityQuaternion() Quaternion {
	return NewQuaternion(0.0, 1.0, 0.0, 0.0)
}

// NewQuaternion returns a quaternion with the axis vector specified.
func NewQuaternion(angle, x, y, z float32) Quaternion {
	return Quaternion{NewZeroVector4()}.Set(angle, x, y, z)
}

// NewQuaternionV returns a quaternion with the axis vector specified.
func NewQuaternionV(angle float32, axis Vector3) Quaternion {
	return NewQuaternion(angle, axis[0], axis[1], axis[2])
}

// Set sets each axis of this Quaternion object to the corresponding axis of a float32 vararg.
func (quaternion Quaternion) Set(angle, x, y, z float32) Quaternion {
	l := Sqrt(x*x + y*y + z*z)
	sin := Sin(angle / 2.0)
	quaternion.vector[0] = sin * x / l
	quaternion.vector[1] = sin * y / l
	quaternion.vector[2] = sin * z / l
	quaternion.vector[3] = Cos(angle / 2.0)
	return quaternion
}

// SetQ sets each axis of this Quaternion object to the corresponding axis of a Quaternion.
func (quaternion Quaternion) SetQ(other Quaternion) Quaternion {

	return quaternion
}

// Mul multiplies this Quaternion object by another quaternion.
func (quaternion Quaternion) Mul(angle, x, y, z float32) Quaternion {
	other := NewQuaternion(angle, x, y, z)
	return quaternion.MulQ(other)
}

// MulQ multiplies this Quaternion object by another quaternion.
func (quaternion Quaternion) MulQ(other Quaternion) Quaternion {
	t0 := (quaternion.vector[2] - quaternion.vector[1]) * (other.vector[1] - other.vector[2])
	t1 := (quaternion.vector[3] + quaternion.vector[0]) * (other.vector[3] + other.vector[0])
	t2 := (quaternion.vector[3] - quaternion.vector[0]) * (other.vector[1] + other.vector[2])
	t3 := (quaternion.vector[2] + quaternion.vector[1]) * (other.vector[3] - other.vector[0])
	t4 := (quaternion.vector[2] - quaternion.vector[0]) * (other.vector[0] - other.vector[1])
	t5 := (quaternion.vector[2] + quaternion.vector[0]) * (other.vector[0] + other.vector[1])
	t6 := (quaternion.vector[3] + quaternion.vector[1]) * (other.vector[3] - other.vector[2])
	t7 := (quaternion.vector[3] - quaternion.vector[1]) * (other.vector[3] + other.vector[2])
	t8 := t5 + t6 + t7
	t9 := 0.5 * (t4 + t8)
	// TODO: Make this work for all dimensions.
	quaternion.Set(
		t1+t9-t8,
		t2+t9-t7,
		t3+t9-t6,
		t0+t9-t5,
	)
	return quaternion
}

// MulSc scales this Quaternion object by a float32.
func (quaternion Quaternion) MulSc(scalar float32) Quaternion {
	quaternion.vector.MulSc(scalar)
	quaternion.vector.Normalize()
	return quaternion
}

// Inverse inverts the quaternion.
func (quaternion Quaternion) Inverse() Quaternion {
	l := quaternion.vector.LenSq()
	quaternion.vector[0] *= -1.0
	quaternion.vector[1] *= -1.0
	quaternion.vector[2] *= -1.0
	quaternion.vector.DivSc(l)
	quaternion.vector.Normalize()
	return quaternion
}

// Clone returns a new Quaternion with components equal to this Quaternion.
func (quaternion Quaternion) Clone() Quaternion {
	return Quaternion{quaternion.vector.Clone()}
}
