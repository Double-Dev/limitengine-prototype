package framework

type IMesh interface {
	Enable()
	Render(instanceBuffer IInstanceBuffer, instanceDefs []InstanceDef, instanceData []float32, numInstances int32)
	Disable()
	Delete()
}
