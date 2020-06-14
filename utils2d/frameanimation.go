package utils2d

import (
	"github.com/double-dev/limitengine/gmath"

	"github.com/double-dev/limitengine/gfx"
)

type Frame struct {
	bounds   gmath.Vector4
	duration float32
}

func CreateFrame(bounds gmath.Vector4, duration float32) *Frame {
	return &Frame{
		bounds:   bounds,
		duration: duration,
	}
}

type FrameAnimation struct {
	frames []*Frame
}

func CreateFrameAnimation(frames ...*Frame) *FrameAnimation {
	return &FrameAnimation{
		frames: frames,
	}
}

func (frameAnimation *FrameAnimation) Apply(index int, instance *gfx.Instance) {
	instance.SetTextureBoundsV(frameAnimation.frames[index].bounds)
}

type FrameAnimationPlayer struct {
	currentAnimation *FrameAnimation
	index            int
	time             float32
	playing          bool
	loopNum          int
}

func CreateFrameAnimationPlayer() *FrameAnimationPlayer {
	return &FrameAnimationPlayer{}
}

func (player *FrameAnimationPlayer) Play(animation *FrameAnimation, times int, instance *gfx.Instance) {
	if !player.playing {
		player.PlayInterrupt(animation, times, instance)
	}
}

func (player *FrameAnimationPlayer) PlayInterrupt(animation *FrameAnimation, times int, instance *gfx.Instance) {
	if player.currentAnimation != animation && len(animation.frames) > 0 {
		player.currentAnimation = animation
		player.playing = true
		player.loopNum = times - 1
		player.time = 0.0
		player.index = 0
		player.currentAnimation.Apply(player.index, instance)
	}
}

func (player *FrameAnimationPlayer) Stop() {
	if player.playing {
		player.currentAnimation = nil
		player.playing = false
		player.loopNum = 0
		player.time = 0.0
		player.index = 0
	}
}

func (player *FrameAnimationPlayer) Update(delta float32, instance *gfx.Instance) {
	if player.playing {
		player.time += delta
		if player.time >= player.currentAnimation.frames[player.index].duration {
			player.index++
			player.time = 0.0

			if player.index >= len(player.currentAnimation.frames) {
				if player.loopNum > 0 {
					player.index = 0
					player.loopNum--
				} else {
					player.playing = false
					player.currentAnimation = nil
					return
				}
			}
			player.currentAnimation.Apply(player.index, instance)
		}
	}
}

func (player *FrameAnimationPlayer) CurrentAnimation() *FrameAnimation {
	return player.currentAnimation
}
