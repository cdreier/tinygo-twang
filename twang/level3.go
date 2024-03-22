package twang

type Level3 struct {
	game *Game
}

func NewLevel3(g *Game) *Level3 {
	return &Level3{
		game: g,
	}
}

func (l *Level3) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	spawner := NewEnemySpawn(l.game)
	spawner.enemySpeed = 2
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 40, spawner))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
