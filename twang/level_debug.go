package twang

type LevelDebug struct {
	game *Game
}

func NewLevelDebug(g *Game) *LevelDebug {
	return &LevelDebug{
		game: g,
	}
}

func (l *LevelDebug) Start() {
	l.game.Player.index = 0
	l.game.entities = []Entity{}
	l.game.AddEntity(NewGoal(len(l.game.colors) - 1))
	l.game.AddEntity(NewWater(10, 20))
	// l.game.AddEntity(NewFire(10, 20))
}
