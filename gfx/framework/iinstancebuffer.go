package framework

type IInstanceBuffer interface {
	Bind()
	StoreInstancedData(data []float32)
	Unbind()
}

type InstanceDef struct {
	Name  string
	Size  int
	Index int
}
