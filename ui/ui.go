package ui

import (
	"github.com/double-dev/limitengine"
)

var (
	log = limitengine.NewLogger("ui")
)

func init() {
	if limitengine.Running() {
		limitengine.AddMouseMotionCallback(func(x, y float32) {
			// fmt.Println("MouseX:", x, "\t MouseY:", y)
		})
		// TODO: Sort out input.
		log.Log("UI online...")
	}
}
