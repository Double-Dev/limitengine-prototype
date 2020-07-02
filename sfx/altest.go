package sfx

import (
	"github.com/double-dev/limitengine/sfx/al"
)

var (
	buffer al.Buffer
)

func Setup() {
	buffer = al.GenBuffers(1)[0]
	// buffer.BufferData()

}

func PlaySound() {

}

func Delete() {
	al.DeleteBuffers(buffer)
}
