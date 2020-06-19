package states

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/tests2d/logic"
)

type MainState struct {
	ecs *limitengine.ECS

	gfxListener      gfx.GFXListener
	interactionWorld *interaction.World

	renderSystem          *limitengine.ECSSystem
	motionSystem          *limitengine.ECSSystem
	controlSystem         *limitengine.ECSSystem
	playerAnimationSystem *limitengine.ECSSystem
}

func NewMainState() *MainState {
	mainState := &MainState{
		ecs: limitengine.NewECS(),
	}

	// Post-Processing Render
	mainState.ecs.NewEntity(
		&gfx.RenderComponent{
			Camera:   gfx.DefaultCamera(),
			Shader:   assets.PostShader,
			Material: assets.PostMaterial,
			Mesh:     gfx.SpriteMesh(),
			Instance: gfx.NewInstance(),
		},
	)

	// Entities
	logic.NewPlayerEntity(mainState.ecs)

	textInstance := gfx.NewInstance()
	charBounds := assets.CalibriFont.GetChar("E").Bounds()
	textInstance.SetTextureBoundsV(charBounds)
	mainState.ecs.NewEntity(
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
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(-1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Right Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(1.5, 0.0, -0.45), gmath.NewVector3(0.1, 1.0, 1.0))
	// Top Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(0.0, 1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))
	// Bottom Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(0.0, -1.0, -0.4), gmath.NewVector3(1.5, 0.1, 1.0))
	// Seperator Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(-0.6, -0.3, -0.4), gmath.NewVector3(0.1, 0.7, 1.0))

	// Systems
	mainState.interactionWorld = interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)
	mainState.interactionWorld.AddInteraction(&logic.ControlInteraction{})

	mainState.gfxListener = gfx.NewGFXListener()

	mainState.renderSystem = gfx.NewRenderSystem()
	mainState.motionSystem = gmath.NewMotionSystem(1.0)
	mainState.controlSystem = logic.NewControlSystem()
	mainState.playerAnimationSystem = logic.NewPlayerAnimationSystem()

	return mainState
}

func (mainState *MainState) OnActive() {
	mainState.ecs.AddECSListener(mainState.interactionWorld)
	mainState.ecs.AddECSListener(mainState.gfxListener)
	mainState.ecs.AddECSSystem(mainState.renderSystem)
	mainState.ecs.AddECSSystem(mainState.motionSystem)
	mainState.ecs.AddECSSystem(mainState.controlSystem)
	mainState.ecs.AddECSSystem(mainState.playerAnimationSystem)
}

func (mainState *MainState) Update(delta float32) {
	mainState.renderSystem.Update(delta)
	mainState.motionSystem.Update(delta)
	mainState.controlSystem.Update(delta)
	mainState.playerAnimationSystem.Update(delta)
}

func (mainState *MainState) OnInactive() {
	mainState.ecs.RemoveECSListener(mainState.interactionWorld)
	mainState.ecs.RemoveECSListener(mainState.gfxListener)
	mainState.ecs.RemoveECSSystem(mainState.renderSystem)
	mainState.ecs.RemoveECSSystem(mainState.motionSystem)
	mainState.ecs.RemoveECSSystem(mainState.controlSystem)
	mainState.ecs.RemoveECSSystem(mainState.playerAnimationSystem)
}
