package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

//Индексы полигона
var collectionCubs = make(map[int]map[int]*Square)
var collectionCubsIndexX = make(map[float64]map[float64]*Square)

var window *pixelgl.Window

var screenX float64 = 1200
var screenY float64 = 800

var pointGameTableX float64 = 200
var pointGameTableY float64 = 30
var cubSize float64 = 30
var cubStep float64 = 2

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
	for !win.Closed() {
		imd := imdraw.New(window)
		window.Clear(colornames.Black)
		for _, value := range collectionCubsIndexX {
			for _, currentCub := range value {
				imd.Color = currentCub.color
				imd.Push(currentCub.rect.Min, currentCub.rect.Max)
				imd.Rectangle(0)
			}
		}
		if window.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Print("мышь x -")
			fmt.Print(win.MousePosition().X)
			fmt.Print("\n")

			fmt.Print("мышь y -")
			fmt.Print(win.MousePosition().Y)
			fmt.Print("\n")

			cub := getCubs(win.MousePosition().X, win.MousePosition().Y)
			fmt.Println(cub)
		}

		imd.Draw(window)
		window.Update()
	}
}

func main() {
	createMainPlace(20)
	pixelgl.Run(run)
}

func getCubs(x float64, y float64) (cub *Square) {
	for indexX, value := range collectionCubsIndexX {
		if indexX < x && indexX+cubSize > x {
			for indexY, currentCub := range value {
				if indexY < y && indexY+cubSize > y {
					cub = currentCub
					cub.color = colornames.Red
				}
			}
		}
	}

	return
}
