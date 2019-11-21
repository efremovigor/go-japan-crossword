package main

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"time"
)

//Ширина кубов влабиринта
var squareSize = 2

const WhileColor = 0
const BlackColor = 1

//Цвета либинта
var colorMap = map[int]color.RGBA{
	BlackColor: colornames.Black,
	WhileColor: colornames.White,
}

//Индексы полигона
var collection = map[string]Square{}

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

//Генерация матрицы лабиринта
func generate(lengthX int, lengthY int) {
	collection = make(map[string]Square)
	for y := 0; y < lengthY; y++ {
		for x := 0; x < lengthX; x++ {
			curX := x
			curY := y
			rect := pixel.R(float64(curX*squareSize), float64(curY*squareSize), float64(squareSize+curX*squareSize), float64(squareSize+curY*squareSize))
			square := Square{x: curX, y: curY, rect: rect}
			square.state = BlackColor
			collection[getIndex(curX, curY)] = square
		}
	}
}
