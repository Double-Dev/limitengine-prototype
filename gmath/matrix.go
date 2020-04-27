package gmath

import (
	"fmt"
)

// Matrix contains a slice of vectors and has methods to perform matrix mathematics.
type Matrix []Vector

// NewIdentityMatrix creates a new Matrix from two uint32s denoting the number of columns and rows.
func NewIdentityMatrix(columns, rows int) Matrix {
	matrix := Matrix{}
	for i := 0; i < columns; i++ {
		matrix = append(matrix, NewZeroVector(rows))
	}
	matrix.SetIdentity()
	return matrix
}

// SetIdentity sets this Matrix equal to the identity Matrix.
func (matrix Matrix) SetIdentity() {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = 1.0
			} else {
				matrix[i][j] = 0.0
			}
		}
	}
}

func (matrix Matrix) MulV(vector Vector) Vector {
	size := MinI(len(vector), len(matrix))
	vOut := NewZeroVector(size)
	for i := 0; i < size; i++ {
		vOut.Add(matrix[i].Clone().MulSc(vector[i])...)
	}
	return vOut
}

func (matrix Matrix) MulM(other Matrix) Matrix {
	mOut := Matrix{}
	for i := 0; i < MinI(len(matrix), len(other)); i++ {
		mOut = append(mOut, matrix.MulV(other[i]))
	}
	return mOut
}

func (matrix Matrix) SetTranslate(translation Vector) Matrix {
	for i := 0; i < MinI(len(matrix), len(translation)); i++ {
		matrix[len(matrix)-1][i] = translation[i]
	}
	return matrix
}

func (matrix Matrix) SetRotate(rotation Quaternion) Matrix {
	squares := make([]float32, len(rotation.vector))
	for i := 0; i < len(rotation.vector); i++ {
		squares[i] = rotation.vector[i] * rotation.vector[i]
	}
	matrix[0][0] = 1 - 2*(squares[1]+squares[2])
	matrix[1][1] = 1 - 2*(squares[0]+squares[2])
	matrix[2][2] = 1 - 2*(squares[0]+squares[1])

	matrix[1][0] = 2.0 * (rotation.vector[0]*rotation.vector[1] - rotation.vector[2]*rotation.vector[3])
	matrix[0][1] = 2.0 * (rotation.vector[0]*rotation.vector[1] + rotation.vector[2]*rotation.vector[3])

	matrix[2][0] = 2.0 * (rotation.vector[0]*rotation.vector[2] + rotation.vector[1]*rotation.vector[3])
	matrix[0][2] = 2.0 * (rotation.vector[0]*rotation.vector[2] - rotation.vector[1]*rotation.vector[3])

	matrix[2][1] = 2.0 * (rotation.vector[1]*rotation.vector[2] - rotation.vector[0]*rotation.vector[3])
	matrix[1][2] = 2.0 * (rotation.vector[1]*rotation.vector[2] + rotation.vector[0]*rotation.vector[3])
	// TODO: Fix rotation maths.
	return matrix
}

func (matrix Matrix) SetScale(scale Vector) Matrix {
	for i := 0; i < MinI(MinI(len(matrix), len(matrix[i])), len(scale)); i++ {
		matrix[i][i] = matrix[i][i] * scale[i]
	}
	return matrix
}

// TODO: Finish Matrix.go

func (matrix Matrix) IsSize(rows, columns int) bool {
	if len(matrix) != columns || len(matrix[0]) != rows {
		return false
	}
	return true
}

func (matrix Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[i][j])
		}
	}
	return arr
}

// Clone returns a new Matrix with components equal to matrix Matrix.
func (matrix Matrix) Clone() Matrix {
	out := Matrix{}
	for _, v := range matrix {
		out = append(out, v.Clone())
	}
	return out
}

func (matrix Matrix) String() string {
	s := "\n"
	for i := 0; i < len(matrix[0]); i++ {
		s += "[" + fmt.Sprintf("%f\t", matrix[0][i])
		for j := 1; j < len(matrix); j++ {
			s += " " + fmt.Sprintf("%f\t", matrix[j][i])
		}
		s += "]\n"
	}
	return s
}
