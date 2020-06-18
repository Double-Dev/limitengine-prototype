package states

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/tests2d/logic"
)

func NewMainState() *limitengine.State {
	state := limitengine.NewState(nil)

	// Post-Processing Render
	state.NewEntity(
		&gfx.RenderComponent{
			Camera:   gfx.DefaultCamera(),
			Shader:   assets.PostShader,
			Material: assets.PostMaterial,
			Mesh:     gfx.SpriteMesh(),
			Instance: gfx.NewInstance(),
		},
	)

	// Entities
	logic.NewPlayerEntity(state)

	textInstance := gfx.NewInstance()
	charBounds := assets.CalibriFont.GetChar("E").Bounds()
	textInstance.SetTextureBoundsV(charBounds)
	state.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.0, -0.6),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(charBounds[2], charBounds[3], 1.0),
		},
		&gfx.RenderComponent{
			Camera:   assets.SceneCamera,
			Shader:   assets.SceneShader,
			Material: assets.CalibriMaterial,
			Mesh:     gfx.SpriteMesh(),
			Instance: textInstance,
		},
	)

	// Left Wall
	logic.NewLevelWallEntity(state, gmath.NewVector3(-1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Right Wall
	logic.NewLevelWallEntity(state, gmath.NewVector3(1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Top Wall
	logic.NewLevelWallEntity(state, gmath.NewVector3(0.0, 1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))
	// Bottom Wall
	logic.NewLevelWallEntity(state, gmath.NewVector3(0.0, -1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))
	// Seperator Wall
	logic.NewLevelWallEntity(state, gmath.NewVector3(-0.6, -0.3, -0.4), gmath.NewVector3(0.1, 0.7, 1.0))

	// Systems
	interactionWorld := interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)
	interactionWorld.AddInteraction(&logic.ControlInteraction{})
	state.AddListener(interactionWorld)

	gfxListener := gfx.NewGFXListener()
	state.AddListener(gfxListener)

	state.AddSystem(gfx.NewRenderSystem())
	state.AddSystem(gmath.NewMotionSystem(1.0))
	state.AddSystem(logic.NewControlSystem())
	state.AddSystem(logic.NewPlayerAnimationSystem())

	return state
}
