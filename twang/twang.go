package twang

import (
	"image/color"
)

type Game struct {
	colors           []color.RGBA
	Player           *Player
	entities         []Entity
	currentLevel     int
	levels           []Level
	animationRunning bool
}

type Entity interface {
	Update()
	Render(index int, colors []color.RGBA) bool
	Intersection(p *Game)
}

type Level interface {
	Start()
}

func NewGame(length int) *Game {

	colors := make([]color.RGBA, length)

	g := &Game{
		colors:       colors,
		entities:     []Entity{},
		Player:       NewPlayer(length),
		currentLevel: 0,
	}

	return g
}

func (g *Game) LoadLevels(levels []Level) {
	g.levels = levels
	g.StartLevel()
}

func (g *Game) StartLevel() {
	g.levels[g.currentLevel].Start()
}

type Renderer interface {
	WriteColors(buf []color.RGBA) (err error)
}

func (g *Game) Render(r Renderer) {

	// normal render loop only when interaction is allowed
	for i := range g.colors {
		playerHandled := false
		if !g.animationRunning {
			playerHandled = g.Player.Render(i, g.colors)
		}

		entityHandled := false
		if !playerHandled {
			for _, e := range g.entities {
				// do not overwrite handle if one handled
				if tmp := e.Render(i, g.colors); tmp && !entityHandled {
					entityHandled = true
				}
			}
		}
		if !playerHandled && !entityHandled {
			g.colors[i] = colorOff
		}
	}
	r.WriteColors(g.colors)
}

func (g *Game) Update() {
	for _, e := range g.entities {
		e.Update()
		e.Intersection(g)
	}
}

func (g *Game) AddEntity(e Entity) {
	g.entities = append(g.entities, e)
}

func (g *Game) RemoveEntity(e Entity) {
	for i, entity := range g.entities {
		if entity == e {
			g.entities = append(g.entities[:i], g.entities[i+1:]...)
			return
		}
	}
}

func (g *Game) Retry() {
	g.entities = []Entity{}
	g.animationRunning = true
	retryTail := NewTail(colorEnemy, g.Player.index, func() {
		g.animationRunning = false
		g.StartLevel()
	})
	g.AddEntity(retryTail)
}

func (g *Game) NextLevel() {
	g.entities = []Entity{}
	g.animationRunning = true
	successTail := NewTail(colorPlayer, g.Player.index, func() {
		g.animationRunning = false
		if g.currentLevel+1 < len(g.levels) {
			g.currentLevel++
		} else {
			g.currentLevel = 0
		}
		g.StartLevel()
	})
	g.AddEntity(successTail)
}
