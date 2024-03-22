package twang

import "image/color"

type Fire struct {
	size  int
	index int
}

func NewFire(size int, position int) *Fire {
	return &Fire{
		size:  size,
		index: position,
	}
}

func (f *Fire) Update() {
	// TODO active and passive phaseindex
}

func (f *Fire) Render(index int, colors []color.RGBA) bool {
	if inRange(index, f.index, f.index+f.size) {
		colors[index] = color.RGBA{0xff, 0x00, 0x00, 0xff}
		return true
	}
	return false
}

func (f *Fire) Intersect(g *Game) {
	if inRange(f.index, g.Player.index-g.Player.attackRange, g.Player.index+g.Player.attackRange) {
		g.Retry()
	}
}
