package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var window *pixelgl.Window

var screenX float64 = 1010
var screenY float64 = 770
var cubSize = 30
var cubStep = 2

func run() {
	var win, err = pixelgl.NewWindow(pixelgl.WindowConfig{
		Title:  "Japan Crossword!",
		Bounds: pixel.R(0, 0, screenX, screenY),
		VSync:  true,
	})
	if err != nil {
		panic(err)
	}
	window = win
	for !window.Closed() {
		imd := imdraw.New(window)
		window.Clear(colornames.White)
		rect := pixel.R(50.5, 38.5, 959.5, 731.5)
		imd.Color = colornames.Black
		imd.Push(rect.Min, rect.Max)
		imd.Rectangle(0)
		imd.Color = colornames.White
		minX := screenX/100*5 + float64(cubStep)
		maxX := minX + float64(cubSize)
		minY := screenY/100*5 + float64(cubStep)
		maxY := minY + float64(cubSize)
		for y := 0; y < 20; y++ {
			for x := 0; x < 20; x++ {
				rect := pixel.R(minX, minY, maxX, maxY)
				imd.Push(rect.Min, rect.Max)
				imd.Rectangle(0)
				minX = maxX + float64(cubStep)
				maxX = minX + float64(cubSize)
			}
			minX = screenX/100*5 + float64(cubStep)
			maxX = minX + float64(cubSize)
			minY = maxY + float64(cubStep)
			maxY = minY + float64(cubSize)
		}

		imd.Draw(window)
		window.Update()
	}
}

func main() {
	//generate(lengthX, lengthY)
	pixelgl.Run(run)
}
