package sfx

import (
	"fmt"
	"os"
	"time"

	"github.com/double-dev/limitengine"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

var (
	log = limitengine.NewLogger("sfx")

	playing     = false
	actionQueue = []func(){}
)

func init() {
	if limitengine.Running() {
		go func() {
			for limitengine.Running() {
				if len(actionQueue) > 0 && !playing {
					playing = true
					actionQueue[0]()
					actionQueue = actionQueue[1:]
				} else {
					time.Sleep(time.Millisecond)
				}
			}
		}()
		log.Log("SFX online...")
	}
}

type Sound struct {
	playing    bool
	sampleRate beep.SampleRate
	streamer   beep.StreamSeekCloser
}

func NewSound(path string, speed float32) *Sound {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("File read went horribly wrong!")
		fmt.Println("Error message:", err)
	}
	streamer, format, err := wav.Decode(file)
	if err != nil {
		fmt.Println("File decode went horribly wrong!")
		fmt.Println("Error message:", err)
	}
	sampleRate := format.SampleRate
	streamer.Seek(0)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))
	return &Sound{
		sampleRate: sampleRate,
		streamer:   streamer,
	}
}

func (sound *Sound) Play(speed float32) {
	sound.playing = true
	actionQueue = append(actionQueue, func() {
		sound.streamer.Seek(0)
		// sampleRate := beep.SampleRate(float32(sound.sampleRate) * speed)
		speaker.Play(beep.Seq(sound.streamer, beep.Callback(func() {
			playing = false
			sound.playing = false
		})))
	})
}

func (sound *Sound) PlayOneShot(speed float32) {
	if !sound.playing {
		sound.Play(speed)
	}
}

func (sound *Sound) Playing() bool {
	return sound.playing
}
