package main

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/cjslep/noise"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func main() {
	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))
	defer speaker.Close()

	wnd := g.NewMasterWindow("Tone Generator", 800, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}

var (
	freq float32
	val  float64
)

func DoNoise() {
	perlinGenerator := noise.NewPerlin(69420)
	val = perlinGenerator.Noise(0.5, 0.5)
}

func loop() {
	g.SingleWindow().Layout(
		g.Row(
			g.Label("Frequency"),
			g.ArrowButton(g.DirectionLeft),
			g.Event().OnClick(g.MouseButtonLeft, func() { freq -= 10 }),
			g.SliderFloat(&freq, 20, 20000),
			g.ArrowButton(g.DirectionRight),
			g.Event().OnClick(g.MouseButtonLeft, func() { freq += 10 }),
			g.InputFloat(&freq),
		),
		g.Row(
			g.Button("Sine"),
			g.Event().OnClick(g.MouseButtonLeft, func() { go MakeSine() }),
		),
		g.Row(
			g.Button("Stop"),
			g.Event().OnClick(g.MouseButtonLeft, func() { speaker.Clear() }),
		),
		g.Spacing(),
		g.Row(
			g.Button("White Noise"),
			g.Event().OnClick(g.MouseButtonLeft, DoNoise),
		),
		g.Row(),
	)
}
