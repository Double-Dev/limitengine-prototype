package gmath

type iVector interface {
	SetElement(index int, value float32)
	GetElement(index int) float32
	Dimension() int

	Set(elements ...float32)
	Add(elements ...float32)
	Sub(elements ...float32)
	Mul(elements ...float32)
	MulSc(scalar float32)
	Div(elements ...float32)
	DivSc(scalar float32)

	Dot() float32
	Cross(other iVector) iVector

	LenSq() float32
	Len() float32
	Normalize() float32
	DstSq(other iVector) float32
	Dst(other iVector) float32

	Clone() iVector
}
