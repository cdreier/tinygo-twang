package twang

type Level3 struct {
	game *Game
}

func NewLevel3(g *Game) *Level3 {
	return &Level3{
		game: g,
	}
}

// enemies moving quickly
func (l *Level3) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	spawner := NewEnemySpawn(l.game)
	spawner.enemySpeed = 2

	spawner2 := NewEnemySpawn(l.game)
	spawner2.enemySpeed = 2
	spawner2.enemyDirection = 1
	l.game.AddEntity(NewSpawn(len(l.game.colors)/2, 40, spawner))
	l.game.AddEntity(NewSpawn(len(l.game.colors)/2, 40, spawner2))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
