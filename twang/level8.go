package twang

type Level8 struct {
	game *Game
}

func NewLevel8(g *Game) *Level8 {
	return &Level8{
		game: g,
	}
}

func (l *Level8) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	water := NewWater(12, 20)
	water.power = 3
	l.game.AddEntity(water)
	l.game.AddEntity(NewFire(6, 40))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
