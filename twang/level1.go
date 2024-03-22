package twang

type Level1 struct {
	game *Game
}

func NewLevel1(g *Game) *Level1 {
	return &Level1{
		game: g,
	}
}

func (l *Level1) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 100, NewEnemySpawn(l.game)))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))

	// l.game.AddEntity(NewFire(5, 20))
}
