package twang

import "image/color"

type Goal struct {
	index int
}

func NewGoal(index int) *Goal {
	return &Goal{
		index: index,
	}
}

func (g *Goal) Update() {}

func (g *Goal) Render(index int, colors []color.RGBA) bool {
	if index == g.index {
		colors[index] = color.RGBA{0xff, 0xff, 0xff, 0xff}
		return true
	}
	return false
}

func (g *Goal) Intersect(p *Game) {
	if g.index == p.Player.index {
		p.NextLevel()
	}
}
