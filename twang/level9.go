package twang

type Level9 struct {
	game *Game
}

func NewLevel9(g *Game) *Level9 {
	return &Level9{
		game: g,
	}
}

// stronger water, fire rigth behind it
func (l *Level9) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	water := NewWater(18, 15, 1)
	water.power = 3
	l.game.AddEntity(water)
	l.game.AddEntity(NewFire(5, 15+18))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
