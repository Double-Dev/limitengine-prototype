package main

import (
	"github.com/double-dev/limitengine/sfx"
)

func main() {
	// Profile
	// defer profile.Start().Stop()

	// // Setup Window
	// limitengine.AppView().SetTitle("2D Tests!")
	// limitengine.AppView().SetPosition(100, 100)
	// limitengine.AppView().SetSize(900, 600)
	// limitengine.AppView().SetAspectRatio(3, 2)
	// limitengine.AppView().SetIcons([]image.Image{image.Image(gio.LoadPNG("assets/Test.png"))})

	// Load Assets
	// assets.LoadAssets()

	sfx.Setup()

	// Launch!
	// limitengine.Launch(states.NewMainState())
}
