package gfx

import (
	"sync"

	"github.com/double-dev/limitengine/gmath"
)

type Instance struct {
	data      map[string][]float32
	dataMutex sync.RWMutex
}

func NewInstance() *Instance {
	return &Instance{
		data: map[string][]float32{
			"verttransformMat0": []float32{1.0, 0.0, 0.0, 0.0},
			"verttransformMat1": []float32{0.0, 1.0, 0.0, 0.0},
			"verttransformMat2": []float32{0.0, 0.0, 1.0, 0.0},
			"verttransformMat3": []float32{0.0, 0.0, 0.0, 1.0},
		},
	}
}

func (instance *Instance) SetTransform(transform gmath.Matrix4) {
	instance.dataMutex.Lock()
	instance.data["verttransformMat0"] = transform[0]
	instance.data["verttransformMat1"] = transform[1]
	instance.data["verttransformMat2"] = transform[2]
	instance.data["verttransformMat3"] = transform[3]
	instance.dataMutex.Unlock()
}

func (instance *Instance) SetData(key string, value []float32) {
	instance.dataMutex.Lock()
	instance.data[key] = value
	instance.dataMutex.Unlock()
}

func (instance *Instance) ModifyData(key string, values ...float32) {
	instance.dataMutex.Lock()
	dataLen := gmath.MinI(len(instance.data[key]), len(values))
	for i := 0; i < dataLen; i++ {
		instance.data[key][i] = values[i]
	}
	instance.dataMutex.Unlock()
}

func (instance *Instance) GetData(key string) []float32 {
	instance.dataMutex.RLock()
	result := instance.data[key]
	instance.dataMutex.RUnlock()
	return result
}
