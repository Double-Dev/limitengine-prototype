package framework

type IModel interface {
	Enable()
	Render(instanceBuffer IInstanceBuffer, instanceData []float32, numInstances int32)
	Disable()
	Delete()
}
