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
	state      int
	color      color.RGBA
	rect       pixel.Rect
	x          int
	y          int
	typeSquare int
}

//назначение кода цвета
func (s *Square) getColor() color.RGBA {
	return colorMap[s.state]
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func createMainPlace(picMap map[int]map[int]int) {
	for indexX, value := range picMap {
		collectionCubs[indexX] = make(map[int]*Square)
		for indexY, pixelPic := range value {
			minX := pointGameTableX + cubStep + float64((cubSize+cubStep)*float64(indexX))
			maxX := minX + cubSize
			minY := pointGameTableY + cubStep + float64((cubSize+cubStep)*float64(indexY))
			maxY := minY + cubSize

			rect := pixel.R(minX, minY, maxX, maxY)
			square := Square{x: indexX, y: indexY, rect: rect, typeSquare: pixelPic}
			square.color = colornames.White
			if collectionCubsIndexX[minX] == nil {
				collectionCubsIndexX[minX] = make(map[float64]*Square)
			}
			collectionCubsIndexX[minX][minY] = &square
			collectionCubs[indexX][indexY] = &square
		}
	}
}

func getPicMatrix(size int) (picMap map[int]map[int]int) {
	picMap = make(map[int]map[int]int)
	for i := 0; i < size; i++ {
		picMap[i] = make(map[int]int)
		for j := 0; j < size; j++ {
			picMap[i][j] = r.Intn(2)
		}
	}
	return picMap
}
