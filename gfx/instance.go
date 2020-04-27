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
			"transformMat0": []float32{1.0, 0.0, 0.0, 0.0},
			"transformMat1": []float32{0.0, 1.0, 0.0, 0.0},
			"transformMat2": []float32{0.0, 0.0, 1.0, 0.0},
			"transformMat3": []float32{0.0, 0.0, 0.0, 1.0},
		},
		dataMutex: sync.RWMutex{},
	}
}

func (instance *Instance) SetTransform(transform gmath.Matrix4) {
	instance.dataMutex.Lock()
	instance.data["transformMat0"] = transform[0]
	instance.data["transformMat1"] = transform[1]
	instance.data["transformMat2"] = transform[2]
	instance.data["transformMat3"] = transform[3]
	instance.dataMutex.Unlock()
}

// type Instance struct {
// 	uniformInts      map[string]int32
// 	uniformMatrix44s map[string]gmath.Matrix
// }

// func NewInstance() *Instance {
// 	return &Instance{
// 		uniformInts:      make(map[string]int32),
// 		uniformMatrix44s: make(map[string]gmath.Matrix),
// 	}
// }

// func (this *Instance) loadTo(iShader framework.IShader) {
// 	gfxMutex.RLock()
// 	for varName, value := range this.uniformInts {
// 		iShader.LoadUniform1I(varName, value)
// 	}
// 	for varName, value := range this.uniformMatrix44s {
// 		iShader.LoadUniformMatrix4fv(varName, value.ToArray())
// 	}
// 	gfxMutex.RUnlock()
// }

// func (this *Instance) AddInt(varName string, val int32) {
// 	gfxMutex.Lock()
// 	this.uniformInts[varName] = val
// 	gfxMutex.Unlock()
// }

// func (this *Instance) AddMatrix44(varName string, val gmath.Matrix) {
// 	if val.IsSize(4, 4) {
// 		gfxMutex.Lock()
// 		this.uniformMatrix44s[varName] = val
// 		gfxMutex.Unlock()
// 	}
// }
