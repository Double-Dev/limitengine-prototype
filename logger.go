package limitengine

import "fmt"

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
	fmt.Println(logger.pkg+":", msg)
}

func (logger *logger) ForceErr(msg interface{}) {
	fmt.Print(logger.pkg + ": ")
	panic(msg)
}

func (logger *logger) Err(msg interface{}, err error) {
	fmt.Println(logger.pkg+":", msg)
	panic(err)
}
