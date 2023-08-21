package main

import (
	"errors"
	"math"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

type SineWave struct {
	sampleFactor float64
	phase        float64
}

func (g *SineWave) Stream(samples [][2]float64) (n int, ok bool) {
	for i := range samples {
		v := math.Sin(g.phase * 2.0 * math.Pi)
		samples[i][0] = v
		samples[i][1] = v
		_, g.phase = math.Modf(g.phase + g.sampleFactor)
	}
	return len(samples), true
}

func (*SineWave) Err() error {
	return nil
}

func SineTone(sr beep.SampleRate, freq float64) (beep.Streamer, error) {
	dt := freq / float64(sr)

	if dt >= 1.0/2.0 {
		return nil, errors.New("samplerate must be at least 2 times greater than frequency")
	}
	return &SineWave{dt, 0.1}, nil
}

func SineSetup(sr beep.SampleRate) beep.Streamer {
	sine, err := SineTone(sr, float64(freq))
	if err != nil {
		panic("oh shit")
	}
	return sine
}

func MakeSine() {
	noise := SineSetup(44100)
	speaker.Play(noise)
}
