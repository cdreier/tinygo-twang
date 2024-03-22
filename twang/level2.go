package twang

type Level2 struct {
	game *Game
}

func NewLevel2(g *Game) *Level2 {
	return &Level2{
		game: g,
	}
}

func (l *Level2) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 50, NewEnemySpawn(l.game)))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))

	// l.game.AddEntity(NewFire(5, 20))
}
