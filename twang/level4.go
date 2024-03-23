package twang

type Level4 struct {
	game *Game
}

func NewLevel4(g *Game) *Level4 {
	return &Level4{
		game: g,
	}
}

func (l *Level4) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewFire(10, 25))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
