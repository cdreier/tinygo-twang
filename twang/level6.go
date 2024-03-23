package twang

type Level6 struct {
	game *Game
}

func NewLevel6(g *Game) *Level6 {
	return &Level6{
		game: g,
	}
}

func (l *Level6) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewWater(36, 15))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
