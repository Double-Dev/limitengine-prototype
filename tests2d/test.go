package test2d

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gio"
)

func main() {
	testLesl := gio.LoadAsString("testShader.lesl")
	gfx.ProcessLESL(testLesl)
	limitengine.Launch()
}
