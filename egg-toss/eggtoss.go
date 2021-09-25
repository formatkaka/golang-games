package eggtoss

import (
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	PADDING = 20

	WINDOW_WIDTH  = 600
	WINDOW_HEIGHT = 900

	IMAGE_WIDTH = 240

	LEFT  = "going-left"
	RIGHT = "going-right"

	DEFAULT_BASKET_POS_X = (600 - IMAGE_WIDTH) / 2
	DEFAULT_BASKET_POS_Y = 750
)

type Game struct{}

var basketImg, eggImage, bgImage *ebiten.Image
var basket1, basket2 *Basket

func init() {
	var err1, err2 error
	basketImg, _, err1 = ebitenutil.NewImageFromFile("egg-toss/static/basket.png")
	bgImage, _, err2 = ebitenutil.NewImageFromFile("egg-toss/static/bg.jpg")

	basket1 = initBasket(DEFAULT_BASKET_POS_X, DEFAULT_BASKET_POS_Y, "b1")
	basket2 = initBasket(DEFAULT_BASKET_POS_X, 100, "b2")

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
}

func (g *Game) Update() error {

	isSpacebar := inpututil.IsKeyJustReleased(ebiten.KeySpace)

	if isSpacebar {
		// if spacebar released, shift baskets
		basket3 := initBasket(DEFAULT_BASKET_POS_X, -350, "b3")
		basket1.GoDown()
		basket2.GoDown()
		basket3.GoDown()

		time.AfterFunc(1*time.Second, func() {
			basket1 = basket2
			basket2 = basket3
		})
	}

	basket1.Update()
	basket2.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(bgImage, nil)

	basket1.Draw(screen)
	basket2.Draw(screen)
	// basket3.Draw(screen)
}

func (g *Game) Layout(w int, h int) (sW, sH int) {
	return 600, 900
}

func PlayGame() {
	ebiten.SetWindowSize(600, 900)
	ebiten.SetWindowTitle("Eggs")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
