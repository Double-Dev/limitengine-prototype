package test

import (
	"image"

	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/tests2d/states"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gio"

	"github.com/pkg/profile"
)

func main() {
	// Profile
	defer profile.Start().Stop()

	// Setup Window
	limitengine.AppView().SetTitle("2D Tests!")
	limitengine.AppView().SetPosition(100, 100)
	limitengine.AppView().SetSize(900, 600)
	limitengine.AppView().SetAspectRatio(3, 2)
	limitengine.AppView().SetIcons([]image.Image{image.Image(gio.LoadPNG("assets/Test.png"))})

	// Load Assets
	assets.LoadAssets()

	// Launch!
	limitengine.Launch(states.NewMenuState())
}
