package twang

import (
	"image/color"
)

type Enemy struct {
	index        int
	speed        int // every x updates, player speed is 1, higher is slower
	speedCount   int // every x updates, player speed is 1, higher is slower
	standing     bool
	direction    int
	shouldRemove bool
}

func NewEnemy(start int, direction int) *Enemy {
	return &Enemy{
		index:      start,
		speed:      5,
		speedCount: 0,
		direction:  direction,
	}
}

func NewStandingEnemy(start int) *Enemy {
	return &Enemy{
		index:    start,
		standing: true,
	}
}

type EntitiyAdder interface {
	AddEntity(e Entity)
}

type EnemySpawn struct {
	entityAdder    EntitiyAdder
	enemySpeed     int
	enemyDirection int
}

func NewEnemySpawn(ea EntitiyAdder) *EnemySpawn {
	return &EnemySpawn{
		entityAdder:    ea,
		enemySpeed:     5,
		enemyDirection: -1,
	}
}

func (e *EnemySpawn) Spawn(start int) {
	enemy := NewEnemy(start, -1)
	enemy.speed = e.enemySpeed
	enemy.direction = e.enemyDirection
	e.entityAdder.AddEntity(enemy)
}

func (e *Enemy) Update() {
	if e.standing {
		return
	}
	if e.speedCount%e.speed == 0 {
		e.index = max(e.index+e.direction, 0)
	}
	e.speedCount++
}

func (e *Enemy) Intersection(g *Game) {
	if inRange(e.index, g.Player.index-g.Player.attackRange, g.Player.index+g.Player.attackRange) {
		if g.Player.attack {
			g.RemoveEntity(e)
		} else if e.index <= g.Player.index {
			g.Retry()
		}
	}
}

func (e *Enemy) Render(index int, colors []color.RGBA) bool {
	if index == e.index {
		colors[index] = colorEnemy
		return true
	}
	return false
}

func (e *Enemy) ShouldRemove() bool {
	return e.shouldRemove
}
