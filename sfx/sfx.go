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

	mixer       = new(beep.Mixer)
	sounds      []*Sound
	playing     = false
	actionQueue = []func(){}
)

func init() {
	if limitengine.Running() {
		go func() {
			for limitengine.Running() {
				if len(actionQueue) > 0 {
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

func Stop() {
	actionQueue = append(actionQueue, func() {
		if playing {
			speaker.Clear()
			playing = false
			for _, sound := range sounds {
				sound.playing = false
			}
		}
	})
}

type Sound struct {
	playing    bool
	sampleRate beep.SampleRate
	streamer   beep.StreamSeekCloser
}

func NewSound(path string) *Sound {
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
	sound := &Sound{
		sampleRate: sampleRate,
		streamer:   streamer,
	}
	sounds = append(sounds, sound)
	return sound
}

func (sound *Sound) Play(speed float32) {
	sound.playing = true
	actionQueue = append(actionQueue, func() {
		sound.streamer.Seek(0)
		// sampleRate := beep.SampleRate(float32(sound.sampleRate) * speed)
		mixer.Add(beep.Seq(sound.streamer, beep.Callback(func() {
			sound.playing = false
		})))
		if !playing {
			playing = true
			speaker.Play(beep.Seq(mixer, beep.Callback(func() {
				playing = false
			})))
		}
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
