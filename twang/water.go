package twang

import (
	"image/color"
)

type Water struct {
	size             int
	index            int
	direction        int
	power            int
	powerCount       int
	colorOffsetCount int
	colors           []color.RGBA
}

func NewWater(size int, position int, direction int) *Water {

	colors := make([]color.RGBA, size)
	for i := 0; i < size; i++ {
		if i%3 == 0 {
			colors[i] = colorWater
		} else {
			colors[i] = colorWaterDark
		}
	}

	return &Water{
		size:      size,
		index:     position,
		power:     2,
		colors:    colors,
		direction: direction,
	}
}

func (w *Water) Update() {
	w.colorOffsetCount++
	if w.colorOffsetCount%5 == 0 {
		// shift colors
		if w.direction > 0 {
			w.colors = append(w.colors[1:], w.colors[0])
		} else {
			w.colors = append([]color.RGBA{w.colors[len(w.colors)-1]}, w.colors[:len(w.colors)-1]...)
		}
	}
}

func (w *Water) Render(index int, colors []color.RGBA) bool {

	if inRange(index, w.index, w.index+w.size-1) {
		colors[index] = w.colors[index-w.index]
		return true
	}

	return false
}

func (w *Water) Intersection(g *Game) {
	w.powerCount++
	if w.powerCount%w.power != 0 && inRange(g.Player.index, w.index, w.index+w.size) {
		g.Player.index -= w.direction
	}
}
