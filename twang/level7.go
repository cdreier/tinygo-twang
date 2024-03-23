package twang

type Level7 struct {
	game *Game
}

func NewLevel7(g *Game) *Level7 {
	return &Level7{
		game: g,
	}
}

func (l *Level7) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewWater(21, 20))
	spawner := NewEnemySpawn(l.game)
	spawner.enemySpeed = 2
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 100, spawner))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
