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

func NewTransformMatrix(translation Vector, rotation Quaternion, scale Vector) Matrix {
	rMat := NewIdentityMatrix(4, 4).SetRotate(rotation)
	tsMat := NewIdentityMatrix(4, 4).SetTranslate(translation).SetScale(scale)
	return tsMat.MulM(rMat)
}

func NewViewMatrix(translation Vector, rotation Quaternion, scale Vector) Matrix {
	rMat := NewIdentityMatrix(4, 4).SetRotate(rotation.Clone().Inverse())
	tsMat := NewIdentityMatrix(4, 4).SetTranslate(translation.Clone().MulSc(-1.0)).SetScale(scale)
	return rMat.MulM(tsMat)
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

func (this Matrix) MulV(vector Vector) Vector {
	size := MinI(len(vector), len(this))
	vOut := NewZeroVector(size)
	for i := 0; i < size; i++ {
		vOut.Add(this[i].Clone().MulSc(vector[i])...)
	}
	return vOut
}

func (this Matrix) MulM(other Matrix) Matrix {
	mOut := Matrix{}
	for i := 0; i < MinI(len(this), len(other)); i++ {
		mOut = append(mOut, this.MulV(other[i]))
	}
	return mOut
}

func (this Matrix) SetTranslate(translation Vector) Matrix {
	for i := 0; i < MinI(len(this), len(translation)); i++ {
		this[len(this)-1][i] = translation[i]
	}
	return this
}

func (this Matrix) SetRotate(rotation Quaternion) Matrix {
	squares := make([]float32, len(rotation.vector))
	for i := 0; i < len(rotation.vector); i++ {
		squares[i] = rotation.vector[i] * rotation.vector[i]
	}
	this[0][0] = 1 - 2*(squares[1]+squares[2])
	this[1][1] = 1 - 2*(squares[0]+squares[2])
	this[2][2] = 1 - 2*(squares[0]+squares[1])

	this[1][0] = 2.0 * (rotation.vector[0]*rotation.vector[1] - rotation.vector[2]*rotation.vector[3])
	this[0][1] = 2.0 * (rotation.vector[0]*rotation.vector[1] + rotation.vector[2]*rotation.vector[3])

	this[2][0] = 2.0 * (rotation.vector[0]*rotation.vector[2] + rotation.vector[1]*rotation.vector[3])
	this[0][2] = 2.0 * (rotation.vector[0]*rotation.vector[2] - rotation.vector[1]*rotation.vector[3])

	this[2][1] = 2.0 * (rotation.vector[1]*rotation.vector[2] - rotation.vector[0]*rotation.vector[3])
	this[1][2] = 2.0 * (rotation.vector[1]*rotation.vector[2] + rotation.vector[0]*rotation.vector[3])
	// TODO: Fix rotation maths.
	return this
}

func (this Matrix) SetScale(scale Vector) Matrix {
	for i := 0; i < len(this); i++ {
		for j := 0; j < MinI(len(this[i]), len(scale)); j++ {
			if i == j {
				this[i][j] = this[i][j] * scale[j]
			}
		}
	}
	return this
}

// TODO: Finish Matrix.go

func (this Matrix) IsSize(rows, columns int) bool {
	if len(this) != columns || len(this[0]) != rows {
		return false
	}
	return true
}

func (this Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			arr = append(arr, this[i][j])
		}
	}
	return arr
}

func (this Matrix) Set(row, col int, val float32) {
	this[col][row] = val
}

func (this Matrix) Get(row, col int) float32 {
	return this[col][row]
}

// Clone returns a new Matrix with components equal to this Matrix.
func (this Matrix) Clone() Matrix {
	out := Matrix{}
	for _, v := range this {
		out = append(out, v.Clone())
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
