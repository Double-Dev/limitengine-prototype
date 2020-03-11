package gmath

import (
	"fmt"
)

// Matrix contains an array of vectors and has methods to perform matrix mathematics.
type Matrix struct {
	matrix []*Vector
}

// NewIdentityMatrix creates a new Matrix from two uint32s denoting the number of columns and rows.
func NewIdentityMatrix(columns, rows int) *Matrix {
	matrix := &Matrix{}
	for i := 0; i < columns; i++ {
		matrix.matrix = append(matrix.matrix, NewZeroVector(rows))
	}
	matrix.SetIdentity()
	return matrix
}

func NewProjectionMatrix(aspectRatio, nearPlane, farPlane, fov float32) *Matrix {
	matrix := NewIdentityMatrix(4, 4)
	yScale := 1.0 / Tan(ToRadians(fov/2.0))
	xScale := yScale * aspectRatio
	frustumLen := farPlane - nearPlane

	matrix.matrix[0].vector[0] = xScale
	matrix.matrix[1].vector[1] = yScale
	matrix.matrix[2].vector[2] = -((farPlane + nearPlane) / frustumLen)
	matrix.matrix[2].vector[3] = -1.0
	matrix.matrix[3].vector[2] = -((2.0 * nearPlane * farPlane) / frustumLen)
	matrix.matrix[3].vector[3] = 0.0

	return matrix
}

// SetIdentity sets this Matrix equal to the identity Matrix.
func (this *Matrix) SetIdentity() {
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < len(this.matrix[i].vector); j++ {
			if i == j {
				this.matrix[i].vector[j] = 1.0
			} else {
				this.matrix[i].vector[j] = 0.0
			}
		}
	}
}

func (this *Matrix) MulV(vector *Vector) *Vector {
	size := MinI(len(vector.vector), len(this.matrix))
	vOut := NewZeroVector(size)
	for i := 0; i < size; i++ {
		vOut.Add(this.matrix[i].Clone().MulSc(vector.vector[i]).vector...)
	}
	return vOut
}

func (this *Matrix) MulM(other *Matrix) *Matrix {
	mOut := &Matrix{}
	for i := 0; i < MinI(len(this.matrix), len(other.matrix)); i++ {
		mOut.matrix = append(mOut.matrix, this.MulV(other.matrix[i]))
	}
	return mOut
}

func (this *Matrix) Translate(amount *Vector) *Matrix {
	for i := 0; i < len(this.matrix)-1; i++ {
		for j := 0; j < MinI(len(this.matrix)-1, len(amount.vector)-1); j++ {
			this.matrix[len(this.matrix)-1].vector[i] += this.matrix[j].vector[i] * amount.vector[j]
		}
	}
	return this
}

func (this *Matrix) Scale(amount Vector) *Matrix {
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < MinI(len(this.matrix[i].vector), len(amount.vector)); j++ {
			if i == j {
				this.matrix[i].vector[j] *= amount.vector[j]
			}
		}
	}
	return this
}

// TODO: Finish Matrix.go

func (this *Matrix) IsSize(rows, columns int) bool {
	if len(this.matrix) != columns || len(this.matrix[0].vector) != rows {
		return false
	}
	return true
}

func (this *Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < len(this.matrix[i].vector); j++ {
			arr = append(arr, this.matrix[i].vector[j])
		}
	}
	return arr
}

func (this *Matrix) Set(row, col int, val float32) {
	this.matrix[col].vector[row] = val
}

func (this *Matrix) Get(row, col int) float32 {
	return this.matrix[col].vector[row]
}

// Clone returns a new Matrix with components equal to this Matrix.
func (this *Matrix) Clone() *Matrix {
	out := &Matrix{}
	for _, v := range this.matrix {
		out.matrix = append(out.matrix, v.Clone())
	}
	return out
}

func (this *Matrix) String() string {
	s := "\n"
	for i := 0; i < len(this.matrix[0].vector); i++ {
		s += "[" + fmt.Sprintf("%f\t", this.matrix[0].vector[i])
		for j := 1; j < len(this.matrix); j++ {
			s += " " + fmt.Sprintf("%f\t", this.matrix[j].vector[i])
		}
		s += "]\n"
	}
	return s
}
