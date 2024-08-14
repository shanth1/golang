package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

const (
	whitelndex = 0
	blacklndex = 1
)

type Options struct {
	Cycles int
}

func lissajous(out io.Writer, params *Options) {
	const (
		defaultCycles = 10    // Количество полных колебаний x
		res           = 0.001 // Угловое разрешение
		size          = 100   // Канва изображения охватывает [size..+size]
		nframes       = 64    // Количество кадров анимации
		delay         = 8     // Задержка между кадрами (единица - 10мс)
	)

	seed := int64(42)
	randomGen := rand.New(rand.NewSource(seed))

	cycles := float64(defaultCycles)
	if params != nil && params.Cycles > 0 {
		cycles = float64(params.Cycles)
	}

	freq := randomGen.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blacklndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Примечание: игнорируем ошибки
}
