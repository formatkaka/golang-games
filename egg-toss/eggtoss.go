package eggtoss

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	PADDING = 20
	LEFT    = "going-left"
	RIGHT   = "going-right"

	DEFAULT_BASKET_POS_X = (600 - 320) / 2
	DEFAULT_BASKET_POS_Y = 900 - 160 - PADDING
)

type Game struct {
	pos float64
	dir string
}

var basketImg, eggImage, bgImage *ebiten.Image

func init() {
	var err error
	basketImg, _, err = ebitenutil.NewImageFromFile("egg-toss/basket.png")

	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {

	if g.pos < -1 {
		g.dir = RIGHT
	} else if g.pos > 1 {
		g.dir = LEFT
	}

	if g.dir == LEFT {
		g.pos = g.pos - 0.03
	} else {
		g.pos = g.pos + 0.03
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	updatedX := DEFAULT_BASKET_POS_X + g.pos*DEFAULT_BASKET_POS_X

	op.GeoM.Translate(updatedX, DEFAULT_BASKET_POS_Y)
	screen.DrawImage(basketImg, op)
}

func (g *Game) Layout(w int, h int) (sW, sH int) {
	return 600, 900
}

func PlayGame() {
	ebiten.SetWindowSize(600, 900)
	ebiten.SetWindowTitle("Eggs")

	if err := ebiten.RunGame(&Game{pos: 0, dir: LEFT}); err != nil {
		log.Fatal(err)
	}
}
