package twang

import (
	"image/color"
)

type Enemy struct {
	index      int
	speed      int // every x updates, player speed is 1, higher is slower
	speedCount int // every x updates, player speed is 1, higher is slower
}

func NewEnemy(start int) *Enemy {
	return &Enemy{
		index:      start,
		speed:      5,
		speedCount: 0,
	}
}

type EntitiyAdder interface {
	AddEntity(e Entity)
}

type EnemySpawn struct {
	entityAdder EntitiyAdder
	enemySpeed  int
}

func NewEnemySpawn(ea EntitiyAdder) *EnemySpawn {
	return &EnemySpawn{
		entityAdder: ea,
		enemySpeed:  5,
	}
}

func (e *EnemySpawn) Spawn(start int) {
	enemy := NewEnemy(start)
	enemy.speed = e.enemySpeed
	e.entityAdder.AddEntity(enemy)
}

func (e *Enemy) Update() {
	if e.speedCount%e.speed == 0 {
		e.index = max(e.index-1, 0)
	}
	e.speedCount++
}

func (e *Enemy) Intersect(g *Game) {
	if inRange(e.index, g.Player.index-g.Player.attackRange, g.Player.index+g.Player.attackRange) {
		if e.index == g.Player.index {
			g.Retry()
		} else if g.Player.attack {
			g.RemoveEntity(e)
		}
	}
}

func (e *Enemy) Render(index int, colors []color.RGBA) bool {
	if index == e.index {
		colors[index] = color.RGBA{0xff, 0x00, 0x00, 0xff}
		return true
	}
	return false
}
