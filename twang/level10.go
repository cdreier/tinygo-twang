package twang

type Level10 struct {
	game *Game
}

func NewLevel10(g *Game) *Level10 {
	return &Level10{
		game: g,
	}
}

// stronger water, fire rigth behind it
func (l *Level10) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	water1 := NewWater(18, 10, 1)
	water1.power = 3
	l.game.AddEntity(water1)

	water2 := NewWater(18, 28, -1)
	water2.power = 3
	l.game.AddEntity(water2)

	l.game.AddEntity(NewFire(5, 28+18))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
