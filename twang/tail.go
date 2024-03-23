package twang

import "image/color"

type Tail struct {
	color color.RGBA
	size  int
	index int
	speed int
	done  func()
}

func NewTail(c color.RGBA, start int, done func()) *Tail {
	return &Tail{
		color: c,
		size:  10,
		index: start,
		speed: 3,
		done:  done,
	}
}

func (t *Tail) Update() {
	t.index = max(t.index-t.speed, 0)
	if t.index == 0 {
		t.size = max(t.size-t.speed, 0)
	}
	if t.index == 0 && t.size == 0 {
		t.done()
	}
}

func (t *Tail) Render(index int, colors []color.RGBA) bool {

	if inRange(index, t.index, t.index+t.size) {
		colors[index] = t.color
		return true
	}

	return false
}

func (t *Tail) Intersection(p *Game) {}
