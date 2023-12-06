package test

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

func Lissajous(out io.Writer) {
	// 定义一个调色板
	var palette = []color.Color{color.White, color.Black}

	// 定义两个常量
	const (
		WhiteIndex = 0 // 调色板的第一个颜色
		blackIndex = 1 // 调色板的下一个颜色
	)

	const (
		cycles  = 5      //转数
		res     = 0.0001 //
		size    = 100    //
		nframes = 64     //
		delay   = 8      //
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
