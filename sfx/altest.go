package sfx

import (
	"fmt"

	"github.com/double-dev/limitengine/sfx/al"
	"github.com/double-dev/limitengine/sfx/vorbis"
)

var (
	buffer al.Buffer
)

func Setup() {
	buffer = al.GenBuffers(1)[0]

	channels := make([]int32, 1)
	sampleRate := make([]int32, 1)
	output := make([][]int16)

	result := vorbis.DecodeFilename("../assets/jump.ogg", channels, sampleRate, output)

	fmt.Println(output)

	// buffer.BufferData(0, )
}

func PlaySound() {

}

func Delete() {
	al.DeleteBuffers(buffer)
}
