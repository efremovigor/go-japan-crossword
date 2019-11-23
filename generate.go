package main

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"time"
)

const WhileColor = 0
const BlackColor = 1

//Цвета либинта
var colorMap = map[int]color.RGBA{
	BlackColor: colornames.Black,
	WhileColor: colornames.White,
}

//Куб
type Square struct {
	state int
	color color.RGBA
	rect  pixel.Rect
	x     int
	y     int
}

//назначение кода цвета
func (s *Square) getColor() color.RGBA {
	return colorMap[s.state]
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func createMainPlace(size int) {
	minX := pointGameTableX + cubStep
	maxX := minX + cubSize
	minY := pointGameTableY + cubStep
	maxY := minY + cubSize
	for x := 1; x <= size; x++ {
		collectionCubs[x] = make(map[int]*Square)
		for y := 1; y <= size; y++ {
			rect := pixel.R(minX, minY, maxX, maxY)
			square := Square{x: x, y: y, rect: rect}
			square.color = colornames.White
			if collectionCubsIndexX[minX] == nil {
				collectionCubsIndexX[minX] = make(map[float64]*Square)
			}
			collectionCubsIndexX[minX][minY] = &square
			collectionCubs[x][y] = &square
			minX = maxX + cubStep
			maxX = minX + cubSize
		}
		minX = pointGameTableX + cubStep
		maxX = minX + cubSize
		minY = maxY + cubStep
		maxY = minY + cubSize
	}
}
