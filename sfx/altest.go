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
	source al.Source
)

func Setup() {
	buffer = al.GenBuffers(1)[0]

	r, err := os.Open(build.Default.GOPATH + "\\src\\github.com\\double-dev\\limitengine\\tests2d\\assets\\jump.ogg")
	if err != nil {
		fmt.Println("WRONG FILE")
	}

	vorbData, err := vorbis.New(r)
	if err != nil {
		fmt.Println("READ ERROR")
	}

	buffer.BufferData(al.FormatMono16, vorbData.Buf(), int32(vorbData.SampleRate))
	vorbData.Close()

	al.SetListenerPosition([3]float32{0.0, 0.0, 0.0})
	al.SetListenerVelocity([3]float32{0.0, 0.0, 0.0})

	source = al.GenSources(1)[0]
	source.SetGain(1.0)
	source.SetPosition([3]float32{0.0, 0.0, 0.0})
	source.SetVelocity([3]float32{0.0, 0.0, 0.0})

	al.PlaySources(source)
}

func PlaySound() {
	al.PlaySources(source)
}

func Delete() {
	al.DeleteBuffers(buffer)
}
