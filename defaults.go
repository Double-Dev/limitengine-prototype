package limitengine

import (
	"go/build"
	"image"
	"image/png"
	"os"
)

// Variables used for default initialization.
const (
	InitWidth       = 600
	InitHeight      = 600
	WindowTitle     = "Hello Limitengine!"
	WindowResizable = true
	BufferSamples   = 8
)

// Variables used for default initialization.
var (
	WindowIcons []image.Image
)

func init() {
	reader, err := os.Open(build.Default.GOPATH + "/src/github.com/double-dev/limitengine/limitengine.png")
	if err != nil {
		log.Log("Could not load limitengine icon: " + err.Error())
	}
	icon, err := png.Decode(reader)
	if err != nil {
		log.Log("Could not decode limitengine icon: " + err.Error())
	}
	WindowIcons = []image.Image{icon}
}
