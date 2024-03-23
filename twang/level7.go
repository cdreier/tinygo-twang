package twang

type Level7 struct {
	game *Game
}

func NewLevel7(g *Game) *Level7 {
	return &Level7{
		game: g,
	}
}

// other direction water, standing enemies
func (l *Level7) Start() {
	l.game.Player.Reset()
	l.game.entities = []Entity{}
	l.game.AddEntity(NewWater(42, 10, -1))
	l.game.AddEntity(NewStandingEnemy(15))
	l.game.AddEntity(NewStandingEnemy(20))
	l.game.AddEntity(NewStandingEnemy(25))
	l.game.AddEntity(NewStandingEnemy(30))
	l.game.AddEntity(NewStandingEnemy(35))
	l.game.AddEntity(NewStandingEnemy(40))
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))

}
