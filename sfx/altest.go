package sfx

import (
	"fmt"
	"go/build"
	"os"

	"github.com/double-dev/limitengine/sfx/al"
	"github.com/double-dev/limitengine/sfx/vorbis"
)

var (
	buffer al.Buffer
)

func Setup() {
	// buffer = al.GenBuffers(1)[0]

	r, err := os.Open(build.Default.GOPATH + "\\src\\github.com\\double-dev\\limitengine\\tests2d\\assets\\jump.ogg")
	if err != nil {
		fmt.Println("WRONG FILE")
	}

	vorbData, err := vorbis.New(r)
	if err != nil {
		fmt.Println("READ ERROR")
	}

	fmt.Println(vorbData.Channels)
	fmt.Println(vorbData.SampleRate)

	// channels := make([]int32, 1)
	// sampleRate := make([]int32, 1)
	// output := make([][]int16)

	// result := vorbis.DecodeFilename("../assets/jump.ogg", channels, sampleRate, output)

	// fmt.Println(output)

	// buffer.BufferData(0, )
}

func PlaySound() {

}

func Delete() {
	al.DeleteBuffers(buffer)
}
