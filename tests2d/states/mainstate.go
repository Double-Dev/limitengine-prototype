package states

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/tests2d/logic"
	"github.com/double-dev/limitengine/utils2d"
)

type MainState struct {
	ecs *limitengine.ECS

	renderListener   *gfx.GFXListener
	spriteListener   *gfx.GFXListener
	textListener     *gfx.GFXListener
	interactionWorld *interaction.World

	renderSystem *limitengine.ECSSystem
	spriteSystem *limitengine.ECSSystem
	textSystem   *limitengine.ECSSystem

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
		gfx.NewRenderComponent(10, gfx.DefaultCamera(), assets.PostShader, assets.PostMaterial, gfx.SpriteMesh(), gfx.NewInstance()),
	)

	// Entities
	logic.NewPlayerEntity(mainState.ecs)

	mainState.ecs.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.0, 0.0),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
		},
		gfx.NewTextComponent(
			2, assets.SceneCamera, assets.TextShader,
			gfx.NewFont(assets.SegoeFont, gmath.NewVector3(0.75, 0.25, 0.75), 0.5, 0.1, gmath.NewZeroVector3(), 0.4, 0.5),
			"Hello World!", 1.0,
		),
	)
	mainState.ecs.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, -0.25, 0.0),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
		},
		gfx.NewTextComponent(
			2, assets.SceneCamera, assets.TextShader,
			gfx.NewFont(assets.SegoeFont, gmath.NewVector3(0.75, 0.25, 0.75), 0.5, 0.1, gmath.NewZeroVector3(), 0.4, 0.5),
			"<Start Game>", 1.0,
		),
	)

	// Left Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(-1.5, 0.0, 0.0), gmath.NewVector3(0.1, 1.0, 1.0))
	// Right Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(1.5, 0.0, 0.0), gmath.NewVector3(0.1, 1.0, 1.0))
	// Top Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(0.0, 1.0, 0.0), gmath.NewVector3(1.5, 0.1, 1.0))
	// Bottom Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(0.0, -1.0, 0.0), gmath.NewVector3(1.5, 0.1, 1.0))
	// Seperator Wall
	logic.NewLevelWallEntity(mainState.ecs, gmath.NewVector3(-0.6, -0.3, 0.0), gmath.NewVector3(0.1, 0.7, 1.0))

	// Systems
	mainState.interactionWorld = interaction.NewWorld(interaction.NewGrid2D(0.5), 60.0)
	mainState.interactionWorld.AddInteraction(&logic.ControlInteraction{})

	mainState.renderListener = gfx.NewRenderListener()
	mainState.spriteListener = utils2d.NewSpriteListener()
	mainState.textListener = gfx.NewTextListener()

	mainState.renderSystem = gfx.NewRenderSystem()
	mainState.spriteSystem = utils2d.NewSpriteSystem()
	mainState.textSystem = gfx.NewTextSystem()
	mainState.motionSystem = gmath.NewMotionSystem(1.0)
	mainState.controlSystem = logic.NewControlSystem()
	mainState.playerAnimationSystem = logic.NewPlayerAnimationSystem()

	return mainState
}

func (mainState *MainState) OnActive() {
	mainState.ecs.AddECSListener(mainState.interactionWorld)
	mainState.ecs.AddECSListener(mainState.renderListener)
	mainState.ecs.AddECSListener(mainState.spriteListener)
	mainState.ecs.AddECSListener(mainState.textListener)

	mainState.ecs.AddECSSystem(mainState.renderSystem)
	mainState.ecs.AddECSSystem(mainState.spriteSystem)
	mainState.ecs.AddECSSystem(mainState.textSystem)
	mainState.ecs.AddECSSystem(mainState.motionSystem)
	mainState.ecs.AddECSSystem(mainState.controlSystem)
	mainState.ecs.AddECSSystem(mainState.playerAnimationSystem)
}

func (mainState *MainState) Update(delta float32) {
	mainState.renderSystem.Update(delta)
	mainState.spriteSystem.Update(delta)
	mainState.textSystem.Update(delta)
	mainState.motionSystem.Update(delta)
	mainState.controlSystem.Update(delta)
	mainState.playerAnimationSystem.Update(delta)

	gfx.Sweep()
}

func (mainState *MainState) OnInactive() {
	mainState.ecs.RemoveECSListener(mainState.interactionWorld)
	mainState.ecs.RemoveECSListener(mainState.renderListener)
	mainState.ecs.RemoveECSListener(mainState.spriteListener)
	mainState.ecs.RemoveECSListener(mainState.textListener)

	mainState.ecs.RemoveECSSystem(mainState.renderSystem)
	mainState.ecs.RemoveECSSystem(mainState.spriteSystem)
	mainState.ecs.RemoveECSSystem(mainState.textSystem)
	mainState.ecs.RemoveECSSystem(mainState.motionSystem)
	mainState.ecs.RemoveECSSystem(mainState.controlSystem)
	mainState.ecs.RemoveECSSystem(mainState.playerAnimationSystem)
}
