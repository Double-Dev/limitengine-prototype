package sfx

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/sfx/framework"
)

var (
	log     = limitengine.NewLogger("gfx")
	context framework.Context

	actionQueue = []func(){}
	sfxPipeline = [](chan func()){}
)

func init() {
	if limitengine.Running() {
		go func() {
			// var err error
			// context, err = gl.NewGLContext()
			// if err != nil {
			// 	log.Err("Context could not be initialized.", err)
			// }

		}()
		log.Log("SFX online...")
	}
}
