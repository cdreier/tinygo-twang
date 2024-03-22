package main

import (
	"image/color"
	"log"
	"sync"
	"time"
	"tinygoleds/twang"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 100
	scale        = 64
	starsCount   = 1024
	size         = 5
)

type Game struct {
	colors []color.RGBA
	mtx    *sync.Mutex
	twang  *twang.Game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x33, 0x33, 0x33, 0xff})
	g.mtx.Lock()
	defer g.mtx.Unlock()
	if g.colors != nil {
		for i, c := range g.colors {
			led := ebiten.NewImage(size, 10)
			led.Fill(c)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i * (size + 1))), 50)
			screen.DrawImage(led, op)
		}
	}

	// draw a blue rectangle
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) WriteColors(buf []color.RGBA) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	g.colors = buf
	return nil
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("LEDs")

	tg := twang.NewGame(60)

	g := &Game{
		mtx:   &sync.Mutex{},
		twang: tg,
	}

	go func() {
		for {
			keys := inpututil.AppendPressedKeys([]ebiten.Key{})
			handled := false
			for _, k := range keys {
				if k == ebiten.KeySpace {
					tg.Player.Attack(true)
					handled = true
				}
			}
			if !handled {
				tg.Player.Attack(false)
				for _, k := range keys {
					if k == ebiten.KeyLeft {
						tg.Player.Move(-1)
					}
					if k == ebiten.KeyRight {
						tg.Player.Move(1)
					}
				}
			}
			tg.Update()
			tg.Render(g)
			time.Sleep(50 * time.Millisecond)
		}
	}()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
