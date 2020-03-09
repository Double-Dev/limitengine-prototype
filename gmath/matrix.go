package gmath

import (
	"fmt"
)

// Matrix is an array of vectors with util methods for matrix mathematics.
type Matrix []Vector
type Matrix22 [2]Vector2
type Matrix33 [3]Vector3
type Matrix44 [4]Vector4

// NewIdentityMatrix creates a new Matrix from two ints denoting the number of columns and rows.
func NewIdentityMatrix(columns, rows int) Matrix {
	matrix := Matrix{}
	for i := 0; i < columns; i++ {
		matrix = append(matrix, make([]float32, rows))
	}
	matrix.SetIdentity()
	return matrix
}

func NewProjectionMatrix(aspectRatio, nearPlane, farPlane, fov float32) Matrix {
	matrix := NewIdentityMatrix(4, 4)
	yScale := 1.0 / Tan(ToRadians(fov/2.0))
	xScale := yScale * aspectRatio
	frustumLen := farPlane - nearPlane

	matrix[0][0] = xScale
	matrix[1][1] = yScale
	matrix[2][2] = -((farPlane + nearPlane) / frustumLen)
	matrix[2][3] = -1.0
	matrix[3][2] = -((2.0 * nearPlane * farPlane) / frustumLen)
	matrix[3][3] = 0.0

	return matrix
}

// SetIdentity sets this Matrix equal to the identity Matrix.
func (this Matrix) SetIdentity() {
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			if i == j {
				this[i][j] = 1.0
			} else {
				this[i][j] = 0.0
			}
		}
	}
}

func (this Matrix) MulV(vector []float32) Vector {
	size := MinI(len(vector), len(this))
	vOut := NewZeroVector(uint(size))
	for i := 0; i < size; i++ {
		vOut.Add(this[i].Clone().MulSc(vector[i])...)
	}
	return vOut
}

func (this Matrix) MulM(other [][]float32) Matrix {
	mOut := Matrix{}
	for i := 0; i < MinI(len(this), len(other)); i++ {
		mOut = append(mOut, this.MulV(other[i]))
	}
	return mOut
}

func (this Matrix) Translate(amount Vector) Matrix {
	for i := 0; i < len(this)-1; i++ {
		for j := 0; j < MinI(len(this)-1, len(this)-1); j++ {
			this[len(this)-1][i] += this[j][i] * amount[j]
		}
	}
	return this
}

func (this Matrix) Scale(amount Vector) Matrix {
	for i := 0; i < len(this); i++ {
		for j := 0; j < MinI(len(this[i]), len(amount)); j++ {
			if i == j {
				this[i][j] *= amount[j]
			}
		}
	}
	return this
}

// TODO: Finish Matrix.go

func (this Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			arr = append(arr, this[i][j])
		}
	}
	return arr
}

// Clone returns a new Matrix with components equal to this Matrix.
func (this Matrix) Clone() Matrix {
	var out Matrix
	for _, v := range this {
		out = append(out, v.Clone())
	}
	return out
}

func (this Matrix) ToMatrix22() Matrix22 {
	out := Matrix22{}
	for i, v := range this {
		out[i] = v.ToVector2()
	}
	return out
}
func (this Matrix) ToMatrix33() Matrix33 {
	out := Matrix33{}
	for i, v := range this {
		out[i] = v.ToVector3()
	}
	return out
}
func (this Matrix) ToMatrix44() Matrix44 {
	out := Matrix44{}
	for i, v := range this {
		out[i] = v.ToVector4()
	}
	return out
}
func (this Matrix) String() string {
	s := "\n"
	for i := 0; i < len(this[0]); i++ {
		s += "[" + fmt.Sprintf("%f\t", this[0][i])
		for j := 1; j < len(this); j++ {
			s += " " + fmt.Sprintf("%f\t", this[j][i])
		}
		s += "]\n"
	}
	return s
}

func (this Matrix22) ToMatrix() Matrix {
	out := Matrix{}
	for _, v := range this {
		out = append(out, NewZeroVector(2).Set(v[:]...))
	}
	return out
}
func (this Matrix22) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			arr = append(arr, this[i][j])
		}
	}
	return arr
}
func (this Matrix22) String() string { return this.ToMatrix().String() }

func (this Matrix33) ToMatrix() Matrix {
	out := Matrix{}
	for _, v := range this {
		out = append(out, NewZeroVector(3).Set(v[:]...))
	}
	return out
}
func (this Matrix33) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			arr = append(arr, this[i][j])
		}
	}
	return arr
}
func (this Matrix33) String() string { return this.ToMatrix().String() }

func (this Matrix44) ToMatrix() Matrix {
	out := Matrix{}
	for _, v := range this {
		out = append(out, NewZeroVector(4).Set(v[:]...))
	}
	return out
}
func (this Matrix44) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			arr = append(arr, this[i][j])
		}
	}
	return arr
}
func (this Matrix44) String() string { return this.ToMatrix().String() }
