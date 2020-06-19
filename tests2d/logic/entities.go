package logic

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/tests2d/assets"
	"github.com/double-dev/limitengine/ui"
)

func NewPlayerEntity(ecs *limitengine.ECS) limitengine.ECSEntity {
	xAxis := &ui.InputControl{}
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyA}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyLeft}, -1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyD}, 1.0)
	xAxis.AddTrigger(ui.InputEvent{Key: ui.KeyRight}, 1.0)
	yAxis := &ui.InputControl{}
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyW}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyUp}, 1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyS}, -1.0)
	yAxis.AddTrigger(ui.InputEvent{Key: ui.KeyDown}, -1.0)
	playerMotion := &gmath.MotionComponent{
		Velocity:        gmath.NewVector3(0.1, 0.0, 0.0),
		Acceleration:    gmath.NewZeroVector3(),
		AngVelocity:     gmath.NewQuaternion(0.1, 0.0, 1.0, 0.0),
		AngAcceleration: gmath.NewIdentityQuaternion(),
	}
	return ecs.NewEntity(
		&gmath.TransformComponent{
			Position: gmath.NewVector3(0.0, 0.0, -0.3),
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    gmath.NewVector3(0.075, 0.075, 1.0),
		},
		playerMotion,
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-0.075, -0.075, 0.0), gmath.NewVector3(0.075, 0.075, 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   assets.SceneCamera,
			Shader:   assets.SceneShader,
			Material: assets.PlayerMaterial,
			Mesh:     gfx.SpriteMesh(),
			Instance: gfx.NewInstance(),
		},
		&PlayerAnimationComponent{
			Player:        gfx.NewFrameAnimationPlayer(),
			RightIdleAnim: assets.PlayerRightIdle,
			RightWalkAnim: assets.PlayerRightWalk,
			RightJumpAnim: gfx.NewFrameAnimation(
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(8), func() bool {
					return !(playerMotion.Velocity[1] >= 1.9)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(9), func() bool {
					return !(playerMotion.Velocity[1] < 1.9 && playerMotion.Velocity[1] > 0.0)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(10), func() bool {
					return !(playerMotion.Velocity[1] <= 0.0 && playerMotion.Velocity[1] > -1.0)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(11), func() bool {
					return !(playerMotion.Velocity[1] <= -1.0)
				}),
			),
			RightWallAnim: assets.PlayerRightWallSlide,
			LeftIdleAnim:  assets.PlayerLeftIdle,
			LeftWalkAnim:  assets.PlayerLeftWalk,
			LeftJumpAnim: gfx.NewFrameAnimation(
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(12), func() bool {
					return !(playerMotion.Velocity[1] >= 1.9)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(13), func() bool {
					return !(playerMotion.Velocity[1] < 1.9 && playerMotion.Velocity[1] > 0.0)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(14), func() bool {
					return !(playerMotion.Velocity[1] <= 0.0 && playerMotion.Velocity[1] > -1.0)
				}),
				gfx.NewTriggerFrame(assets.PlayerSpriteSheet.GetBounds(15), func() bool {
					return !(playerMotion.Velocity[1] <= -1.0)
				}),
			),
			LeftWallAnim: assets.PlayerLeftWallSlide,
		},
		&ControlComponent{
			XAxis: xAxis,
			YAxis: yAxis,
		},
	)
}

func NewLevelWallEntity(ecs *limitengine.ECS, position, scale gmath.Vector3) {
	ecs.NewEntity(
		&gmath.TransformComponent{
			Position: position,
			Rotation: gmath.NewIdentityQuaternion(),
			Scale:    scale,
		},
		&interaction.ColliderComponent{
			AABB: gmath.NewAABB(gmath.NewVector3(-1.0*scale[0], -1.0*scale[1], 0.0), gmath.NewVector3(scale[0], scale[1], 0.0)),
		},
		&gfx.RenderComponent{
			Camera:   assets.SceneCamera,
			Shader:   assets.SceneShader,
			Material: assets.LevelMaterial,
			Mesh:     gfx.SpriteMesh(),
			Instance: gfx.NewInstance(),
		},
	)
}
