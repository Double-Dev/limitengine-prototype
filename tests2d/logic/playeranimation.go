package logic

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
)

type PlayerAnimationComponent struct {
	direction bool

	Player *gfx.FrameAnimationPlayer

	RightIdleAnim, RightWalkAnim, RightJumpAnim, RightWallAnim *gfx.FrameAnimation
	LeftIdleAnim, LeftWalkAnim, LeftJumpAnim, LeftWallAnim     *gfx.FrameAnimation
}

func NewPlayerAnimationSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			playerAnim := components[0].(*PlayerAnimationComponent)
			render := components[1].(*gfx.RenderComponent)
			control := components[2].(*ControlComponent)
			motion := components[3].(*gmath.MotionComponent)

			if motion.Velocity[0] > 0.1 {
				playerAnim.direction = false
			} else if motion.Velocity[0] < -0.1 {
				playerAnim.direction = true
			}

			if control.canWallJump && !control.canJump {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.RightWallAnim, 1, render.Instance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftWallAnim, 1, render.Instance)
				}
				playerAnim.Player.SetSpeed(3.0)
			} else if gmath.Abs(motion.Velocity[1]) > 0.05 || !control.canJump {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftJumpAnim, 1, render.Instance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightJumpAnim, 1, render.Instance)
				}
			} else if gmath.Abs(motion.Velocity[0]) > 0.05 {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftWalkAnim, 1, render.Instance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightWalkAnim, 1, render.Instance)
				}
				playerAnim.Player.SetSpeed(gmath.Abs(motion.Velocity[0]) * 3.0)
			} else {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftIdleAnim, 1, render.Instance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightIdleAnim, 1, render.Instance)
				}
				playerAnim.Player.SetSpeed(1.0)
			}
			playerAnim.Player.Update(delta, render.Instance)
		}
	}, (*PlayerAnimationComponent)(nil), (*gfx.RenderComponent)(nil), (*ControlComponent)(nil), (*gmath.MotionComponent)(nil))
}
