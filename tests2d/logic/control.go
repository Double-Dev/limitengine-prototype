package logic

import (
	"fmt"
	"reflect"

	"github.com/double-dev/limitengine"
	"github.com/double-dev/limitengine/gmath"
	"github.com/double-dev/limitengine/interaction"
	"github.com/double-dev/limitengine/ui"
)

type ControlComponent struct {
	XAxis, YAxis              *ui.InputControl
	canJump                   bool
	canWallJump, wallJumpLeft bool
	gravityEnabled            bool
}

func NewControlSystem() *limitengine.ECSSystem {
	return limitengine.NewSystem(func(delta float32, entities [][]limitengine.ECSComponent) {
		for _, components := range entities {
			control := components[0].(*ControlComponent)
			motion := components[1].(*gmath.MotionComponent)

			speed := float32(3.0)
			maxSpeed := float32(1.0)
			if control.XAxis.Amount() > 0.01 {
				motion.Acceleration[0] = speed
			} else if control.XAxis.Amount() < -0.01 {
				motion.Acceleration[0] = -speed
			} else {
				// TODO: Implement friction in collision calculations to avoid doing this.
				motion.Acceleration[0] = 0.0
				if !control.gravityEnabled {
					motion.Velocity[0] *= 0.95
				}
			}

			if gmath.Abs(motion.Velocity[0]) > maxSpeed {
				motion.Velocity[0] = maxSpeed * gmath.Sign(motion.Velocity[0])
			}

			if control.YAxis.Amount() > 0.01 {
				if control.canJump {
					control.gravityEnabled = true
					control.canJump = false
					motion.Velocity[1] = 1.0
				} else if control.canWallJump {
					control.gravityEnabled = true
					control.canWallJump = false
					motion.Velocity[1] = 1.0
					if control.wallJumpLeft {
						motion.Velocity[0] = -2.0
					} else {
						motion.Velocity[0] = 2.0
					}
				}
			}

			if !control.canJump && !control.canWallJump {
				control.gravityEnabled = true
			}

			if control.gravityEnabled {
				motion.Acceleration[1] = -2.0
				if motion.Velocity[1] < 0.0 {
					motion.Acceleration[1] = -3.0
				}
			} else {
				motion.Acceleration[1] = -0.25
			}
		}
	}, (*ControlComponent)(nil), (*gmath.MotionComponent)(nil))
}

type ControlInteraction struct{}

func (interation *ControlInteraction) StartInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3, penetration float32) {
	fmt.Println("BEGIN INTERACT")
	control := interactor.Entity.GetComponent((*ControlComponent)(nil)).(*ControlComponent)
	if !interactee.Collider.IsTrigger {
		if normal[1] < -0.5 {
			control.canJump = true
			control.gravityEnabled = false
		} else if gmath.Abs(normal[0]) > 0.9 {
			interactor.Motion.Velocity[1] = -0.1
			control.canWallJump = true
			control.gravityEnabled = false
			if normal[0] < 0.0 {
				control.wallJumpLeft = false
			} else {
				control.wallJumpLeft = true
			}
		}
	}
}

func (interation *ControlInteraction) EndInteract(delta float32, interactor, interactee interaction.InteractEntity, normal gmath.Vector3) {
	fmt.Println("end interact")
	control := interactor.Entity.GetComponent((*ControlComponent)(nil)).(*ControlComponent)
	if !interactee.Collider.IsTrigger {
		if normal[1] < -0.5 {
			control.canJump = false
			control.gravityEnabled = true
		} else if gmath.Abs(normal[0]) > 0.9 {
			control.canWallJump = false
		}
	}
}

func (interation *ControlInteraction) GetInteractorComponents() []reflect.Type {
	return []reflect.Type{
		reflect.TypeOf((*ControlComponent)(nil)),
		reflect.TypeOf((*gmath.MotionComponent)(nil)),
	}
}

func (interation *ControlInteraction) GetInteracteeComponents() []reflect.Type {
	return []reflect.Type{}
}
