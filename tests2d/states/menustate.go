package states

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/ui"
)

type MenuState struct {
	ecs *limitengine.ECS

	control *ui.InputControl

	textListener *gfx.GFXListener
	textSystem   *limitengine.ECSSystem
}

func NewMenuState() *MenuState {
	menuState := &MenuState{
		ecs: limitengine.NewECS(),
	}

	menuState.ecs.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.2, 0.0),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(1.0, 1.0, 1.0),
		},
		gfx.NewTextComponent(
			2, gfx.DefaultCamera(), assets.TextShader,
			gfx.NewFont(assets.SegoeFont, gmath.NewVector3(0.25, 0.75, 0.25), 0.5, 0.1, gmath.NewVector3(1.0, 1.0, 1.0), 0.4, 0.5),
			"Blob: The Game", 1.0,
		),
	)
	menuState.ecs.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, -0.2, 0.0),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.5, 0.5, 1.0),
		},
		gfx.NewTextComponent(
			2, gfx.DefaultCamera(), assets.TextShader,
			gfx.NewFont(assets.SegoeFont, gmath.NewVector3(0.25, 0.75, 0.25), 0.5, 0.1, gmath.NewZeroVector3(), 0.4, 0.5),
			"Press enter to start...", 1.0,
		),
	)

	menuState.control = &ui.InputControl{}
	menuState.control.AddTrigger(ui.InputEvent{Key: ui.KeyEnter}, 1.0)

	menuState.textListener = gfx.NewTextListener()
	menuState.textSystem = gfx.NewTextSystem()

	return menuState
}

func (menuState *MenuState) OnActive() {
	menuState.ecs.AddECSListener(menuState.textListener)
	menuState.ecs.AddECSSystem(menuState.textSystem)
}

func (menuState *MenuState) Update(delta float32) {
	menuState.textSystem.Update(delta)

	if menuState.control.Amount() > 0.0 {
		limitengine.SetState(NewMainState())
	}

	gfx.Sweep()
}

func (menuState *MenuState) OnInactive() {
	// menuState.ecs.RemoveECSListener(menuState.textListener)
	// menuState.ecs.RemoveECSSystem(menuState.textSystem)
}
