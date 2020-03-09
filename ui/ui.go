package ui

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
)

var (
	log           = limitengine.NewLogger("ui")
	inputControls = make(map[InputEvent]struct {
		float32
		*InputControl
	})
)

func init() {
	if limitengine.Running() {
		// TODO: Implement controller input and ui events.
		limitengine.AddKeyCallback(func(key limitengine.Key, scancode int, action limitengine.Action, mods limitengine.ModKey) {
			moddedKey := GetModdedKey(key, mods)
			for inputEvent, set := range inputControls {
				if inputEvent.Key == key || inputEvent.Key == moddedKey {
					if action == Press {
						set.InputControl.amt = set.float32
					} else if action == Release && set.InputControl.amt == set.float32 {
						set.InputControl.amt = 0.0
					}
				}
			}
		})
		limitengine.AddMouseButtonCallback(func(button limitengine.MouseButton, action limitengine.Action, mods limitengine.ModKey) {
			moddedButton := GetModdedMouseButton(button, mods)
			for inputEvent, set := range inputControls {
				if inputEvent.MouseButton == moddedButton {
					if action == Press || action == Repeat {
						set.InputControl.amt = 1.0 * set.float32
					} else {
						set.InputControl.amt = 0.0
					}
				}
			}
		})
		log.Log("UI online...")
	}
}

type InputEvent struct {
	Joystick    limitengine.Joystick
	Key         limitengine.Key
	MouseButton limitengine.MouseButton
}

type InputControl struct {
	amt float32
}

func (this *InputControl) AddTrigger(inputEvent InputEvent, weight float32) {
	inputControls[inputEvent] = struct {
		float32
		*InputControl
	}{weight, this}
}

func (inputControl *InputControl) Amount() float32 {
	return gmath.Clamp(inputControl.amt, -1.0, 1.0)
}

func ClearKeyControls() {
	inputControls = make(map[InputEvent]struct {
		float32
		*InputControl
	})
}
