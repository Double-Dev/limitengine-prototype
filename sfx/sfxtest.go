package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"

	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var streamer beep.StreamSeekCloser

func main() {
	file, err := os.Open("jump.wav")
	if err != nil {
		fmt.Errorf("File read went horribly wrong!")
	}
	streamer, format, err := wav.Decode(file)
	if err != nil {
		fmt.Errorf("File decode went horribly wrong!")
	}
	defer streamer.Close()

	sr := format.SampleRate / 2
	speaker.Init(sr, sr.N(time.Second/10))

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	done := make(chan bool)
	fmt.Println("Playing")
	speaker.Play(beep.Seq(resampled, beep.Callback(func() {
		fmt.Println("Finished")
		done <- true
	})))

	<-done
}
