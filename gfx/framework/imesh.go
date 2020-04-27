package framework

type IMesh interface {
	Enable()
	Render(instanceBuffer IInstanceBuffer, instanceDefs []struct {
		Name  string
		Size  int
		Index int
	}, instanceData []float32, numInstances int32)
	Disable()
	Delete()
}
