package gfx

import (
	"github.com/double-dev/limitengine/gmath"
)

type Frame interface {
	Bounds() gmath.Vector4
	endTrigger(time float32) bool
}

type DurationFrame struct {
	bounds   gmath.Vector4
	duration float32
}

func CreateDurationFrame(bounds gmath.Vector4, duration float32) *DurationFrame {
	return &DurationFrame{
		bounds:   bounds,
		duration: duration,
	}
}

func (durationFrame *DurationFrame) Bounds() gmath.Vector4 { return durationFrame.bounds }
func (durationFrame *DurationFrame) endTrigger(time float32) bool {
	return time >= durationFrame.duration
}

type TriggerFrame struct {
	bounds      gmath.Vector4
	triggerFunc func() bool
}

func CreateTriggerFrame(bounds gmath.Vector4, endTrigger func() bool) *TriggerFrame {
	return &TriggerFrame{
		bounds:      bounds,
		triggerFunc: endTrigger,
	}
}

func (triggerFrame *TriggerFrame) Bounds() gmath.Vector4        { return triggerFrame.bounds }
func (triggerFrame *TriggerFrame) endTrigger(time float32) bool { return triggerFrame.triggerFunc() }

type FrameAnimation struct {
	frames []Frame
}

func CreateFrameAnimation(frames ...Frame) *FrameAnimation {
	return &FrameAnimation{
		frames: frames,
	}
}

func (frameAnimation *FrameAnimation) Apply(index int, instance *Instance) {
	instance.SetTextureBoundsV(frameAnimation.frames[index].Bounds())
}

type FrameAnimationPlayer struct {
	currentAnimation *FrameAnimation
	index            int
	time, speed      float32
	playing          bool
	loopNum          int
}

func CreateFrameAnimationPlayer() *FrameAnimationPlayer {
	return &FrameAnimationPlayer{}
}

func (player *FrameAnimationPlayer) Play(animation *FrameAnimation, times int, instance *Instance) {
	if !player.playing {
		player.PlayInterrupt(animation, times, instance)
	}
}

func (player *FrameAnimationPlayer) PlayInterrupt(animation *FrameAnimation, times int, instance *Instance) {
	if player.currentAnimation != animation && len(animation.frames) > 0 {
		player.currentAnimation = animation
		player.playing = true
		player.loopNum = times - 1
		player.time = 0.0
		player.index = 0
		player.currentAnimation.Apply(player.index, instance)
	}
}

func (player *FrameAnimationPlayer) Update(delta float32, instance *Instance) {
	if player.playing {
		player.time += delta * player.speed
		if player.currentAnimation.frames[player.index].endTrigger(player.time) {
			for player.currentAnimation.frames[player.index].endTrigger(player.time) {
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
			}
			player.currentAnimation.Apply(player.index, instance)
		}
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

func (player *FrameAnimationPlayer) SetSpeed(speed float32) {
	player.speed = speed
}

func (player *FrameAnimationPlayer) CurrentAnimation() *FrameAnimation {
	return player.currentAnimation
}
