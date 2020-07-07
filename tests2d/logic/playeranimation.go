package logic

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gfx"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/utils2d"
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
			sprite := components[1].(*utils2d.SpriteComponent)
			control := components[2].(*ControlComponent)
			motion := components[3].(*gmath.MotionComponent)

			spriteInstance := sprite.Renderable.Instance

			if motion.Velocity[0] > 0.1 {
				playerAnim.direction = false
			} else if motion.Velocity[0] < -0.1 {
				playerAnim.direction = true
			}

			if control.canWallJump && !control.canJump {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.RightWallAnim, 1, spriteInstance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftWallAnim, 1, spriteInstance)
				}
				playerAnim.Player.SetSpeed(3.0)
			} else if gmath.Abs(motion.Velocity[1]) > 0.05 || !control.canJump {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftJumpAnim, 1, spriteInstance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightJumpAnim, 1, spriteInstance)
				}
			} else if gmath.Abs(motion.Velocity[0]) > 0.05 {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftWalkAnim, 1, spriteInstance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightWalkAnim, 1, spriteInstance)
				}
				playerAnim.Player.SetSpeed(gmath.Abs(motion.Velocity[0]) * 3.0)
			} else {
				if playerAnim.direction {
					playerAnim.Player.PlayInterrupt(playerAnim.LeftIdleAnim, 1, spriteInstance)
				} else {
					playerAnim.Player.PlayInterrupt(playerAnim.RightIdleAnim, 1, spriteInstance)
				}
				playerAnim.Player.SetSpeed(1.0)
			}
			playerAnim.Player.Update(delta, spriteInstance)
		}
	}, (*PlayerAnimationComponent)(nil), (*utils2d.SpriteComponent)(nil), (*ControlComponent)(nil), (*gmath.MotionComponent)(nil))
}
