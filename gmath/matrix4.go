package gmath

import (
	"fmt"
)

// Matrix4 contains a slice of vectors and has methods to perform matrix mathematics.
type Matrix4 []Vector4

// NewIdentityMatrix4 news a new Matrix4 from two uint32s denoting the number of columns and rows.
func NewIdentityMatrix4() Matrix4 {
	matrix := Matrix4{NewZeroVector4(), NewZeroVector4(), NewZeroVector4(), NewZeroVector4()}
	matrix[0][0] = 1.0
	matrix[1][1] = 1.0
	matrix[2][2] = 1.0
	matrix[3][3] = 1.0
	return matrix
}

func NewTransformMatrix(translation Vector3, rotation Quaternion, scale Vector3) Matrix4 {
	rMat := NewIdentityMatrix4().SetRotate(rotation)
	tMat := NewIdentityMatrix4().SetTranslate(translation)
	sMat := NewIdentityMatrix4().SetScale(scale)
	return tMat.MulM(rMat).MulM(sMat)
}

func NewViewMatrix(translation Vector3, rotation Quaternion, scale Vector3) Matrix4 {
	rMat := NewIdentityMatrix4().SetRotate(rotation.Clone().Inverse())
	tsMat := NewIdentityMatrix4().SetTranslate(translation.Clone().MulSc(-1.0)).SetScale(scale)
	return rMat.MulM(tsMat)
}

func NewProjectionMatrix2D(aspectRatio float32) Matrix4 {
	matrix := NewIdentityMatrix4()
	matrix[0][0] = aspectRatio
	return matrix
}

func NewProjectionMatrix3D(aspectRatio, nearPlane, farPlane, fov float32) Matrix4 {
	matrix := NewIdentityMatrix4()
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

// SetIdentity sets this Matrix4 equal to the identity Matrix.
func (matrix Matrix4) SetIdentity() {
	matrix[0][0] = 1.0
	matrix[0][1] = 0.0
	matrix[0][2] = 0.0
	matrix[0][3] = 0.0
	matrix[1][0] = 0.0
	matrix[1][1] = 1.0
	matrix[1][2] = 0.0
	matrix[1][3] = 0.0
	matrix[2][0] = 0.0
	matrix[2][1] = 0.0
	matrix[2][2] = 1.0
	matrix[2][3] = 0.0
	matrix[3][0] = 0.0
	matrix[3][1] = 0.0
	matrix[3][2] = 0.0
	matrix[3][3] = 1.0
}

func (matrix Matrix4) MulV(vector Vector4) Vector4 {
	vOut := NewZeroVector4()
	vOut.AddV(matrix[0].Clone().MulSc(vector[0]))
	vOut.AddV(matrix[1].Clone().MulSc(vector[1]))
	vOut.AddV(matrix[2].Clone().MulSc(vector[2]))
	vOut.AddV(matrix[3].Clone().MulSc(vector[3]))
	return vOut
}

func (matrix Matrix4) MulM(other Matrix4) Matrix4 {
	mOut := Matrix4{}
	mOut = append(mOut, matrix.MulV(other[0]))
	mOut = append(mOut, matrix.MulV(other[1]))
	mOut = append(mOut, matrix.MulV(other[2]))
	mOut = append(mOut, matrix.MulV(other[3]))
	return mOut
}

func (matrix Matrix4) SetTranslate(translation Vector3) Matrix4 {
	matrix.SetIdentity()
	matrix[3][0] = translation[0]
	matrix[3][1] = translation[1]
	matrix[3][2] = translation[2]
	matrix[3][3] = 1.0
	return matrix
}

func (matrix Matrix4) SetRotate(rotation Quaternion) Matrix4 {
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

func (matrix Matrix4) SetScale(scale Vector3) Matrix4 {
	matrix[0][0] *= scale[0]
	matrix[1][1] *= scale[1]
	matrix[2][2] *= scale[2]
	return matrix
}

func (matrix Matrix4) ToArray() []float32 {
	arr := []float32{
		matrix[0][0], matrix[0][1], matrix[0][2], matrix[0][3],
		matrix[1][0], matrix[1][1], matrix[1][2], matrix[1][3],
		matrix[2][0], matrix[2][1], matrix[2][2], matrix[2][3],
		matrix[3][0], matrix[3][1], matrix[3][2], matrix[3][3],
	}
	return arr
}

// Clone returns a new Matrix4 with components equal to this Matrix.
func (matrix Matrix4) Clone() Matrix4 {
	out := Matrix4{
		matrix[0].Clone(), matrix[1].Clone(),
		matrix[2].Clone(), matrix[3].Clone(),
	}
	return out
}

func (this Matrix4) String() string {
	return fmt.Sprintf("[%f\t %f\t %f\t %f\t]", this[0][0], this[1][0], this[2][0], this[3][0]) +
		fmt.Sprintf("[%f\t %f\t %f\t %f\t]", this[0][1], this[1][1], this[2][1], this[3][1]) +
		fmt.Sprintf("[%f\t %f\t %f\t %f\t]", this[0][2], this[1][2], this[2][2], this[3][2]) +
		fmt.Sprintf("[%f\t %f\t %f\t %f\t]", this[0][3], this[1][3], this[2][3], this[3][3])
}
