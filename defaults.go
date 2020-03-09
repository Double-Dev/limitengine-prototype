package limitengine

import (
	"image"
	"image/png"
	"os"
)

// Variables used for default initialization.
var (
	InitWidth       = 600
	InitHeight      = 600
	WindowTitle     = "Hello World!"
	WindowIcons     []image.Image
	WindowResizable = true
	BufferSamples   = 8
)

// // Variables used to change operating parameters.
// var (
// 	TargetFPS = 60
// )

func init() {
	reader, err := os.Open("../DefaultIcon.png")
	if err != nil {
		log.Err("Error loading default icon.", err)
	}
	icon, err := png.Decode(reader)
	if err != nil {
		log.Err("Error decoding default icon.", err)
	}
	WindowIcons = []image.Image{icon}
}
