package twang

type Level1 struct {
	game *Game
}

func NewLevel1(g *Game) *Level1 {
	return &Level1{
		game: g,
	}
}

// 3 standing enemies
func (l *Level1) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
	l.game.AddEntity(NewStandingEnemy(15))
	l.game.AddEntity(NewStandingEnemy(30))
	l.game.AddEntity(NewStandingEnemy(45))

}
