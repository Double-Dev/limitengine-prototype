package limitengine

import (
	"fmt"
	"time"
)

var messageQueue []string

func init() {
	go func() {
		if len(messageQueue) > 0 {
			for _, message := range messageQueue {
				fmt.Println(message)
			}
		} else {
			time.Sleep(time.Millisecond * 100)
		}
	}()
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
	messageQueue = append(messageQueue, fmt.Sprintf(logger.pkg+":", msg))
}

func (logger *logger) ForceErr(msg interface{}) {
	fmt.Print(logger.pkg + ": ")
	panic(msg)
}

func (logger *logger) Err(msg interface{}, err error) {
	fmt.Println(logger.pkg+":", msg)
	panic(err)
}
