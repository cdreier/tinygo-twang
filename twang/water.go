package twang

import (
	"image/color"
)

type Water struct {
	size             int
	index            int
	power            int
	powerCount       int
	colorOffset      int
	colorOffsetCount int
}

func NewWater(size int, position int) *Water {
	return &Water{
		size:  size,
		index: position,
		power: 2,
	}
}

func (w *Water) Update() {

	w.colorOffsetCount++

	if w.colorOffsetCount%10 == 0 {
		w.colorOffset++
		if w.colorOffset == 3 {
			w.colorOffset = 0
		}
	}

}

func (w *Water) Render(index int, colors []color.RGBA) bool {

	if inRange(index, w.index, w.index+w.size) {

		if w.colorOffset == 0 {
			colors[index] = colorWater
		} else {
			colors[index] = colorOff
		}
		return true
	}

	return false
}

func (w *Water) Intersection(g *Game) {
	w.powerCount++
	if w.powerCount%w.power != 0 && inRange(g.Player.index, w.index, w.index+w.size) {
		g.Player.index--
	}
}
