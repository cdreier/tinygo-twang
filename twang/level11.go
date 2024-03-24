package twang

type Level11 struct {
	game *Game
}

func NewLevel11(g *Game) *Level11 {
	return &Level11{
		game: g,
	}
}

// water and 2 enemy spawns with fast enemies
func (l *Level11) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}

	water1 := NewWater(18, 10, 1)
	water1.power = 3
	l.game.AddEntity(water1)

	spawner := NewEnemySpawn(l.game)
	spawner.enemySpeed = 1
	l.game.AddEntity(NewSpawn(len(l.game.colors), 40, spawner))
	spawner2 := NewEnemySpawn(l.game)
	spawner2.enemySpeed = 1
	l.game.AddEntity(NewSpawn(len(l.game.colors), 60, spawner2))

	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
