package twang

import (
	"image/color"
)

type Water struct {
	size             int
	index            int
	power            int
	powerCount       int
	colorOffsetCount int
	colors           []color.RGBA
}

func NewWater(size int, position int) *Water {

	colors := make([]color.RGBA, size)
	for i := 0; i < size; i++ {
		if i%3 == 0 {
			colors[i] = colorWater
		} else {
			colors[i] = colorWaterDark
		}
	}

	return &Water{
		size:   size,
		index:  position,
		power:  2,
		colors: colors,
	}
}

func (w *Water) Update() {
	w.colorOffsetCount++
	if w.colorOffsetCount%4 == 0 {
		// shift colors
		var tmp color.RGBA
		tmp, w.colors = w.colors[0], w.colors[1:]
		w.colors = append(w.colors, tmp)
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
		g.Player.index--
	}
}
