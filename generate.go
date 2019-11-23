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
