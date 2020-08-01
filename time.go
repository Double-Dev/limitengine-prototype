package limitengine

import "time"

const (
	Nanosecond = float32(1.0 / 1000.0)
	Second     = float32(1.0)
)

// type Timer struct {
// 	duration    float32
// 	currentTime float32
// }

func DelayFunc(function func(), delay float32) {
	go func() {
		time.Sleep(time.Duration(float32(time.Second) * delay))
		function()
	}()
}
