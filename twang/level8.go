package twang

type Level8 struct {
	game *Game
}

func NewLevel8(g *Game) *Level8 {
	return &Level8{
		game: g,
	}
}

// water with moving enemies towards the player
func (l *Level8) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	l.game.AddEntity(NewWater(21, 20, 1))
	spawner := NewEnemySpawn(l.game)
	spawner.enemySpeed = 2
	l.game.AddEntity(NewSpawn(len(l.game.colors)-1, 100, spawner))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
}
