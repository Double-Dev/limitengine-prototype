package ui

import (
	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/dependencies/glfw/v3.3/glfw"
)

// Definition for an input action.
const (
	Release = limitengine.Action(glfw.Release)
	Press   = limitengine.Action(glfw.Press)
	Repeat  = limitengine.Action(glfw.Repeat)
)

// Definition for an input joystick.
const (
	Joystick1    = limitengine.Joystick(glfw.Joystick1)
	Joystick2    = limitengine.Joystick(glfw.Joystick2)
	Joystick3    = limitengine.Joystick(glfw.Joystick3)
	Joystick4    = limitengine.Joystick(glfw.Joystick4)
	Joystick5    = limitengine.Joystick(glfw.Joystick5)
	Joystick6    = limitengine.Joystick(glfw.Joystick6)
	Joystick7    = limitengine.Joystick(glfw.Joystick7)
	Joystick8    = limitengine.Joystick(glfw.Joystick8)
	Joystick9    = limitengine.Joystick(glfw.Joystick9)
	Joystick10   = limitengine.Joystick(glfw.Joystick10)
	Joystick11   = limitengine.Joystick(glfw.Joystick11)
	Joystick12   = limitengine.Joystick(glfw.Joystick12)
	Joystick13   = limitengine.Joystick(glfw.Joystick13)
	Joystick14   = limitengine.Joystick(glfw.Joystick14)
	Joystick15   = limitengine.Joystick(glfw.Joystick15)
	Joystick16   = limitengine.Joystick(glfw.Joystick16)
	JoystickLast = limitengine.Joystick(glfw.JoystickLast)
)

// Definition for an input keyboard key.
const (
	KeyUnknown      = limitengine.Key(glfw.KeyUnknown)
	KeySpace        = limitengine.Key(glfw.KeySpace)
	KeyApostrophe   = limitengine.Key(glfw.KeyApostrophe)
	KeyComma        = limitengine.Key(glfw.KeyComma)
	KeyMinus        = limitengine.Key(glfw.KeyMinus)
	KeyPeriod       = limitengine.Key(glfw.KeyPeriod)
	KeySlash        = limitengine.Key(glfw.KeySlash)
	Key0            = limitengine.Key(glfw.Key0)
	Key1            = limitengine.Key(glfw.Key1)
	Key2            = limitengine.Key(glfw.Key2)
	Key3            = limitengine.Key(glfw.Key3)
	Key4            = limitengine.Key(glfw.Key4)
	Key5            = limitengine.Key(glfw.Key5)
	Key6            = limitengine.Key(glfw.Key6)
	Key7            = limitengine.Key(glfw.Key7)
	Key8            = limitengine.Key(glfw.Key8)
	Key9            = limitengine.Key(glfw.Key9)
	KeySemicolon    = limitengine.Key(glfw.KeySemicolon)
	KeyEqual        = limitengine.Key(glfw.KeyEqual)
	KeyA            = limitengine.Key(glfw.KeyA)
	KeyB            = limitengine.Key(glfw.KeyB)
	KeyC            = limitengine.Key(glfw.KeyC)
	KeyD            = limitengine.Key(glfw.KeyD)
	KeyE            = limitengine.Key(glfw.KeyE)
	KeyF            = limitengine.Key(glfw.KeyF)
	KeyG            = limitengine.Key(glfw.KeyG)
	KeyH            = limitengine.Key(glfw.KeyH)
	KeyI            = limitengine.Key(glfw.KeyI)
	KeyJ            = limitengine.Key(glfw.KeyJ)
	KeyK            = limitengine.Key(glfw.KeyK)
	KeyL            = limitengine.Key(glfw.KeyL)
	KeyM            = limitengine.Key(glfw.KeyM)
	KeyN            = limitengine.Key(glfw.KeyN)
	KeyO            = limitengine.Key(glfw.KeyO)
	KeyP            = limitengine.Key(glfw.KeyP)
	KeyQ            = limitengine.Key(glfw.KeyQ)
	KeyR            = limitengine.Key(glfw.KeyR)
	KeyS            = limitengine.Key(glfw.KeyS)
	KeyT            = limitengine.Key(glfw.KeyT)
	KeyU            = limitengine.Key(glfw.KeyU)
	KeyV            = limitengine.Key(glfw.KeyV)
	KeyW            = limitengine.Key(glfw.KeyW)
	KeyX            = limitengine.Key(glfw.KeyX)
	KeyY            = limitengine.Key(glfw.KeyY)
	KeyZ            = limitengine.Key(glfw.KeyZ)
	KeyLeftBracket  = limitengine.Key(glfw.KeyLeftBracket)
	KeyBackslash    = limitengine.Key(glfw.KeyBackslash)
	KeyRightBracket = limitengine.Key(glfw.KeyRightBracket)
	KeyGraveAccent  = limitengine.Key(glfw.KeyGraveAccent)
	KeyWorld1       = limitengine.Key(glfw.KeyWorld1)
	KeyWorld2       = limitengine.Key(glfw.KeyWorld2)
	KeyEscape       = limitengine.Key(glfw.KeyEscape)
	KeyEnter        = limitengine.Key(glfw.KeyEnter)
	KeyTab          = limitengine.Key(glfw.KeyTab)
	KeyBackspace    = limitengine.Key(glfw.KeyBackspace)
	KeyInsert       = limitengine.Key(glfw.KeyInsert)
	KeyDelete       = limitengine.Key(glfw.KeyDelete)
	KeyRight        = limitengine.Key(glfw.KeyRight)
	KeyLeft         = limitengine.Key(glfw.KeyLeft)
	KeyDown         = limitengine.Key(glfw.KeyDown)
	KeyUp           = limitengine.Key(glfw.KeyUp)
	KeyPageUp       = limitengine.Key(glfw.KeyPageUp)
	KeyPageDown     = limitengine.Key(glfw.KeyPageDown)
	KeyHome         = limitengine.Key(glfw.KeyHome)
	KeyEnd          = limitengine.Key(glfw.KeyEnd)
	KeyCapsLock     = limitengine.Key(glfw.KeyCapsLock)
	KeyScrollLock   = limitengine.Key(glfw.KeyScrollLock)
	KeyNumLock      = limitengine.Key(glfw.KeyNumLock)
	KeyPrintScreen  = limitengine.Key(glfw.KeyPrintScreen)
	KeyPause        = limitengine.Key(glfw.KeyPause)
	KeyF1           = limitengine.Key(glfw.KeyF1)
	KeyF2           = limitengine.Key(glfw.KeyF2)
	KeyF3           = limitengine.Key(glfw.KeyF3)
	KeyF4           = limitengine.Key(glfw.KeyF4)
	KeyF5           = limitengine.Key(glfw.KeyF5)
	KeyF6           = limitengine.Key(glfw.KeyF6)
	KeyF7           = limitengine.Key(glfw.KeyF7)
	KeyF8           = limitengine.Key(glfw.KeyF8)
	KeyF9           = limitengine.Key(glfw.KeyF9)
	KeyF10          = limitengine.Key(glfw.KeyF10)
	KeyF11          = limitengine.Key(glfw.KeyF11)
	KeyF12          = limitengine.Key(glfw.KeyF12)
	KeyF13          = limitengine.Key(glfw.KeyF13)
	KeyF14          = limitengine.Key(glfw.KeyF14)
	KeyF15          = limitengine.Key(glfw.KeyF15)
	KeyF16          = limitengine.Key(glfw.KeyF16)
	KeyF17          = limitengine.Key(glfw.KeyF17)
	KeyF18          = limitengine.Key(glfw.KeyF18)
	KeyF19          = limitengine.Key(glfw.KeyF19)
	KeyF20          = limitengine.Key(glfw.KeyF20)
	KeyF21          = limitengine.Key(glfw.KeyF21)
	KeyF22          = limitengine.Key(glfw.KeyF22)
	KeyF23          = limitengine.Key(glfw.KeyF23)
	KeyF24          = limitengine.Key(glfw.KeyF24)
	KeyF25          = limitengine.Key(glfw.KeyF25)
	KeyKP0          = limitengine.Key(glfw.KeyKP0)
	KeyKP1          = limitengine.Key(glfw.KeyKP1)
	KeyKP2          = limitengine.Key(glfw.KeyKP2)
	KeyKP3          = limitengine.Key(glfw.KeyKP3)
	KeyKP4          = limitengine.Key(glfw.KeyKP4)
	KeyKP5          = limitengine.Key(glfw.KeyKP5)
	KeyKP6          = limitengine.Key(glfw.KeyKP6)
	KeyKP7          = limitengine.Key(glfw.KeyKP7)
	KeyKP8          = limitengine.Key(glfw.KeyKP8)
	KeyKP9          = limitengine.Key(glfw.KeyKP9)
	KeyKPDecimal    = limitengine.Key(glfw.KeyKPDecimal)
	KeyKPDivide     = limitengine.Key(glfw.KeyKPDivide)
	KeyKPMultiply   = limitengine.Key(glfw.KeyKPMultiply)
	KeyKPSubtract   = limitengine.Key(glfw.KeyKPSubtract)
	KeyKPAdd        = limitengine.Key(glfw.KeyKPAdd)
	KeyKPEnter      = limitengine.Key(glfw.KeyKPEnter)
	KeyKPEqual      = limitengine.Key(glfw.KeyKPEqual)
	KeyLeftShift    = limitengine.Key(glfw.KeyLeftShift)
	KeyLeftControl  = limitengine.Key(glfw.KeyLeftControl)
	KeyLeftAlt      = limitengine.Key(glfw.KeyLeftAlt)
	KeyLeftSuper    = limitengine.Key(glfw.KeyLeftSuper)
	KeyRightShift   = limitengine.Key(glfw.KeyRightShift)
	KeyRightControl = limitengine.Key(glfw.KeyRightControl)
	KeyRightAlt     = limitengine.Key(glfw.KeyRightAlt)
	KeyRightSuper   = limitengine.Key(glfw.KeyRightSuper)
	KeyMenu         = limitengine.Key(glfw.KeyMenu)
	KeyLast         = limitengine.Key(glfw.KeyLast)
)

// Definition for an input keyboard modifier key.
const (
	ModShift   = limitengine.ModKey(glfw.ModShift)
	ModControl = limitengine.ModKey(glfw.ModControl)
	ModAlt     = limitengine.ModKey(glfw.ModAlt)
	ModSuper   = limitengine.ModKey(glfw.ModSuper)
)

// Definition for an input mouse button. (Offset is to prevent button with zero value.)
const (
	MouseButton1      = limitengine.MouseButton(glfw.MouseButton1) + 1
	MouseButton2      = limitengine.MouseButton(glfw.MouseButton2) + 1
	MouseButton3      = limitengine.MouseButton(glfw.MouseButton3) + 1
	MouseButton4      = limitengine.MouseButton(glfw.MouseButton4) + 1
	MouseButton5      = limitengine.MouseButton(glfw.MouseButton5) + 1
	MouseButton6      = limitengine.MouseButton(glfw.MouseButton6) + 1
	MouseButton7      = limitengine.MouseButton(glfw.MouseButton7) + 1
	MouseButton8      = limitengine.MouseButton(glfw.MouseButton8) + 1
	MouseButtonLast   = limitengine.MouseButton(glfw.MouseButtonLast) + 1
	MouseButtonLeft   = limitengine.MouseButton(glfw.MouseButtonLeft) + 1
	MouseButtonRight  = limitengine.MouseButton(glfw.MouseButtonRight) + 1
	MouseButtonMiddle = limitengine.MouseButton(glfw.MouseButtonMiddle) + 1
)

func getTotalMods(mods ...limitengine.ModKey) int {
	totalMods := 0
	for _, mod := range mods {
		totalMods = totalMods | int(mod)
	}
	return totalMods
}

// GetModdedKey returns a new key that represents the specified key with the
// specified mods applied.
func GetModdedKey(key limitengine.Key, mods ...limitengine.ModKey) limitengine.Key {
	return limitengine.Key(int(key) | getTotalMods(mods...))
}

// GetModdedMouseButton returns a new mouse button that represents the specified
// mouse button with the specified mods applied.
func GetModdedMouseButton(mouseButton limitengine.MouseButton, mods ...limitengine.ModKey) limitengine.MouseButton {
	return limitengine.MouseButton(int(mouseButton) | getTotalMods(mods...))
}
