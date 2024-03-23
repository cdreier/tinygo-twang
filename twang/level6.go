package twang

type Level6 struct {
	game *Game
}

func NewLevel6(g *Game) *Level6 {
	return &Level6{
		game: g,
	}
}

// introducing water, one standing enemy behind the water
func (l *Level6) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	l.game.AddEntity(NewWater(21, 15, 1))
	l.game.AddEntity(NewStandingEnemy(15 + 21 + 5))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
