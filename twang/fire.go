package twang

import (
	"image/color"
)

type Fire struct {
	size          int
	index         int
	active        bool
	interval      int
	intervalCount int
	colorOffset   int
}

var activeColors = []color.RGBA{
	colorFireActive1,
	colorFireActive2,
	colorFireActive3,
}

var inactiveColors = []color.RGBA{
	colorFireInactive1,
	colorFireInactive2,
	colorFireInactive3,
}

func NewFire(size int, position int) *Fire {
	return &Fire{
		size:     size,
		index:    position,
		interval: 40,
	}
}

func (f *Fire) Update() {
	if f.intervalCount%f.interval == 0 {
		f.active = !f.active
	}
	f.intervalCount++

	f.colorOffset++
	if f.colorOffset == 3 {
		f.colorOffset = 0
	}
}

func (f *Fire) Render(index int, colors []color.RGBA) bool {

	if inRange(index, f.index, f.index+f.size) {
		i := (index + f.colorOffset) % 3
		if f.active {
			colors[index] = activeColors[i]
		} else {
			colors[index] = inactiveColors[i]
		}
		return true
	}
	return false
}

func (f *Fire) Intersection(g *Game) {
	if f.active && inRange(g.Player.index, f.index, f.index+f.size) {
		g.Retry()
	}
}
