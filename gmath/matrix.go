package gmath

import (
	"fmt"
)

// Matrix contains a slice of vectors and has methods to perform matrix mathematics.
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

func NewTransformMatrix(translation *Vector, rotation *Quaternion, scale *Vector) *Matrix {
	rMat := NewIdentityMatrix(4, 4).SetRotate(rotation)
	tsMat := NewIdentityMatrix(4, 4).SetTranslate(translation).SetScale(scale)
	return tsMat.MulM(rMat)
}

func NewViewMatrix(translation *Vector, rotation *Quaternion, scale *Vector) *Matrix {
	rMat := NewIdentityMatrix(4, 4).SetRotate(rotation.Clone().Inverse())
	tsMat := NewIdentityMatrix(4, 4).SetTranslate(translation.Clone().MulSc(-1.0)).SetScale(scale)
	return rMat.MulM(tsMat)
}

func NewProjectionMatrix(aspectRatio, nearPlane, farPlane, fov float32) *Matrix {
	matrix := NewIdentityMatrix(4, 4)
	yScale := 1.0 / Tan(ToRadians(fov/2.0))
	xScale := yScale * aspectRatio
	frustumLen := farPlane - nearPlane

	matrix.matrix[0].SetElement(0, xScale)
	matrix.matrix[1].SetElement(1, yScale)
	matrix.matrix[2].SetElement(2, -((farPlane + nearPlane) / frustumLen))
	matrix.matrix[2].SetElement(3, -1.0)
	matrix.matrix[3].SetElement(2, -((2.0 * nearPlane * farPlane) / frustumLen))
	matrix.matrix[3].SetElement(3, 0.0)

	return matrix
}

// SetIdentity sets this Matrix equal to the identity Matrix.
func (this *Matrix) SetIdentity() {
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < this.matrix[i].Dimension(); j++ {
			if i == j {
				this.matrix[i].SetElement(j, 1.0)
			} else {
				this.matrix[i].SetElement(j, 0.0)
			}
		}
	}
}

func (this *Matrix) MulV(vector *Vector) *Vector {
	size := MinI(vector.Dimension(), len(this.matrix))
	vOut := NewZeroVector(size)
	for i := 0; i < size; i++ {
		vOut.Add(this.matrix[i].Clone().MulSc(vector.GetElement(i)).vector...)
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

func (this *Matrix) SetTranslate(translation *Vector) *Matrix {
	for i := 0; i < MinI(len(this.matrix), len(translation.vector)); i++ {
		this.matrix[len(this.matrix)-1].SetElement(i, translation.GetElement(i))
	}
	return this
}

func (this *Matrix) SetRotate(rotation *Quaternion) *Matrix {
	squares := make([]float32, rotation.quaternion.Dimension())
	for i := 0; i < rotation.quaternion.Dimension(); i++ {
		squares[i] = rotation.GetElement(i) * rotation.GetElement(i)
	}
	this.matrix[0].SetElement(0, 1-2*(squares[1]+squares[2]))
	this.matrix[1].SetElement(1, 1-2*(squares[0]+squares[2]))
	this.matrix[2].SetElement(2, 1-2*(squares[0]+squares[1]))

	this.matrix[1].SetElement(0, 2.0*(rotation.GetElement(0)*rotation.GetElement(1)-rotation.GetElement(2)*rotation.GetElement(3)))
	this.matrix[0].SetElement(1, 2.0*(rotation.GetElement(0)*rotation.GetElement(1)+rotation.GetElement(2)*rotation.GetElement(3)))

	this.matrix[2].SetElement(0, 2.0*(rotation.GetElement(0)*rotation.GetElement(2)+rotation.GetElement(1)*rotation.GetElement(3)))
	this.matrix[0].SetElement(2, 2.0*(rotation.GetElement(0)*rotation.GetElement(2)-rotation.GetElement(1)*rotation.GetElement(3)))

	this.matrix[2].SetElement(1, 2.0*(rotation.GetElement(1)*rotation.GetElement(2)-rotation.GetElement(0)*rotation.GetElement(3)))
	this.matrix[1].SetElement(2, 2.0*(rotation.GetElement(1)*rotation.GetElement(2)+rotation.GetElement(0)*rotation.GetElement(3)))
	// TODO: Fix rotation maths.
	return this
}

func (this *Matrix) SetScale(scale *Vector) *Matrix {
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < MinI(this.matrix[i].Dimension(), scale.Dimension()); j++ {
			if i == j {
				this.matrix[i].SetElement(j, this.matrix[i].GetElement(j)*scale.GetElement(j))
			}
		}
	}
	return this
}

// TODO: Finish Matrix.go

func (this *Matrix) IsSize(rows, columns int) bool {
	if len(this.matrix) != columns || this.matrix[0].Dimension() != rows {
		return false
	}
	return true
}

func (this *Matrix) ToArray() []float32 {
	arr := []float32{}
	for i := 0; i < len(this.matrix); i++ {
		for j := 0; j < this.matrix[i].Dimension(); j++ {
			arr = append(arr, this.matrix[i].GetElement(j))
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
	for i := 0; i < this.matrix[0].Dimension(); i++ {
		s += "[" + fmt.Sprintf("%f\t", this.matrix[0].GetElement(i))
		for j := 1; j < len(this.matrix); j++ {
			s += " " + fmt.Sprintf("%f\t", this.matrix[j].GetElement(i))
		}
		s += "]\n"
	}
	return s
}
