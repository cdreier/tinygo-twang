package twang

import (
	"image/color"
)

type Spawn struct {
	interval      int // spawn every X render cycles
	intervalCount int // spawn every X render cycles
	position      int
	spawn         Spawnable
}

type Spawnable interface {
	Spawn(start int)
}

func NewSpawn(position, interval int, cb Spawnable) *Spawn {
	return &Spawn{
		interval:      interval,
		intervalCount: 0,
		position:      position,
		spawn:         cb,
	}
}

func (s *Spawn) Update() {
	if s.intervalCount%s.interval == 0 {
		s.spawn.Spawn(s.position)
	}
	s.intervalCount++
}

func (s *Spawn) Render(index int, colors []color.RGBA) bool {
	return false
}

func (s *Spawn) Intersect(p *Game) {}
