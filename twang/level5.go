package twang

type Level5 struct {
	game *Game
}

func NewLevel5(g *Game) *Level5 {
	return &Level5{
		game: g,
	}
}

func (l *Level5) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 100, NewEnemySpawn(l.game)))
	l.game.AddEntity(NewFire(5, 10))
	secondFire := NewFire(5, 40)
	secondFire.interval = 60
	l.game.AddEntity(secondFire)
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
