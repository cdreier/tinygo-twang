package twang

import (
	"image/color"
)

type Player struct {
	index       int
	world       int
	attack      bool
	attackRange int
}

// TODO make player and interface
func NewPlayer(world int) *Player {
	return &Player{
		world:       world,
		attackRange: 2,
	}
}

func (p *Player) Move(dir int) {
	if p.index+dir < 0 {
		p.index = 0
	} else if p.index+dir >= p.world {
		p.index = p.world - 1
	} else {
		p.index += dir
	}
}

func (p *Player) Attack(onoff bool) {
	p.attack = onoff
}

func (p *Player) Reset() {
	p.index = 0
	p.attackRange = 2
}

func (p *Player) Render(index int, colors []color.RGBA) bool {

	if p.attack && inRange(index, p.index-p.attackRange, p.index+p.attackRange) {
		colors[minMax(index, 0, p.world-1)] = colorPlayerAttack
		return true
	}

	if index == p.index {
		colors[index] = colorPlayer
		return true
	}
	return false
}
