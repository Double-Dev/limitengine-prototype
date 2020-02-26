package gmath

import "fmt"

// Matrix is an array of vectors with util methods for matrix mathematics.
type Matrix []Vector

// NewMatrix creates a new Matrix from two ints denoting the number of columns and rows.
func NewMatrix(columns, rows int) Matrix {
	matrix := Matrix(make([]Vector, columns))
	for i := 0; i < columns; i++ {
		matrix[i] = Vector(make([]float32, rows))
	}
	matrix.SetIdentity()
	return matrix
}

// NewMatrixV creates a new Matrix from a Vector array.
func NewMatrixV(columns ...Vector) Matrix {
	return columns
}

func NewProjectionMatrix(aspectRatio, nearPlane, farPlane, fov float32) Matrix {
	matrix := NewMatrix(4, 4)
	yScale := 1.0 / Tan(ToRadians(fov/2.0))
	xScale := yScale * aspectRatio
	frustumLen := farPlane - nearPlane

	matrix[0][0] = xScale
	matrix[1][1] = yScale
	matrix[2][2] = -((farPlane + nearPlane) / frustumLen)
	matrix[3][2] = -1.0
	matrix[2][3] = -((2.0 * nearPlane * farPlane) / frustumLen)
	matrix[3][3] = 0.0

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
	vOut := NewVectorOfSize(uint(size))
	for i := 0; i < size; i++ {
		vOut.AddV(matrix[i].Clone().MulSc(vector[i]))
	}
	return vOut
}

func (matrix Matrix) MulM(other Matrix) Matrix {
	mOut := NewMatrixV()
	for i := 0; i < MinI(len(matrix), len(other)); i++ {
		mOut = append(mOut, mOut.MulV(other[i]))
	}
	return mOut
}

func (matrix Matrix) Translate(amount Vector) Matrix {
	last := matrix[len(matrix)-1]
	for i := 0; i < MinI(len(last), len(amount)); i++ {
		last[i] += amount[i]
	}
	return matrix
}

func (matrix Matrix) Scale(amount Vector) Matrix {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < MinI(len(matrix[i]), len(amount)); j++ {
			if i == j {
				matrix[i][j] *= amount[j]
			}
		}
	}
	return matrix
}

// TODO: Finish Matrix.go

func (matrix Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			arr = append(arr, matrix[j][i])
		}
	}
	return arr
}

// NewVectorV creates a new Vector from a float32 array.
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
