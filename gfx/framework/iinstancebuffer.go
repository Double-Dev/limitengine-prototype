package framework

type IInstanceBuffer interface {
	Bind()
	StoreInstancedData(data []float32)
	Unbind()
}
