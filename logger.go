package limitengine

import (
	"fmt"
	"time"
)

var (
	loggingEnabled = false
	messageQueue   = []string{}
)

func init() {
	if loggingEnabled {
		go func() {
			for Running() {
				if len(messageQueue) > 0 {
					fmt.Println(messageQueue[0])
					messageQueue = messageQueue[1:]
				} else {
					time.Sleep(time.Millisecond * 100)
				}
			}
		}()
	}
}

// TODO: Add cross-platform support.
type logger struct {
	pkg string
}

func NewLogger(pkg string) logger {
	return logger{
		pkg: pkg,
	}
}

func (logger *logger) Log(msg interface{}) {
	if loggingEnabled {
		messageQueue = append(messageQueue, fmt.Sprint(logger.pkg+":", msg))
	}
}

func (logger *logger) ForceErr(msg interface{}) {
	fmt.Print(logger.pkg + ": ")
	panic(msg)
}

func (logger *logger) Err(msg interface{}, err error) {
	fmt.Println(logger.pkg+":", msg)
	panic(err)
}
