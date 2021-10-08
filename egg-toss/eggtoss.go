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

type Game struct {
	state string
}

var basketImg, eggImage, bgImage *ebiten.Image
var basket1, basket2 *Basket
var ball *Ball
var portal *Portal

func init() {
	var err1, err2, err3 error
	basketImg, _, err1 = ebitenutil.NewImageFromFile("egg-toss/static/basket.png")
	bgImage, _, err2 = ebitenutil.NewImageFromFile("egg-toss/static/bg.jpg")
	eggImage, _, err3 = ebitenutil.NewImageFromFile("egg-toss/static/panipuri.png")

	basket1 = initBasket(DEFAULT_BASKET_POS_X, DEFAULT_BASKET_POS_Y, "b1")
	basket2 = initBasket(DEFAULT_BASKET_POS_X, 100, "b2")

	ball = initBall(basket1)

	portal = initPortal()

	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatal(err1, err2, err3)
	}
}

func (g *Game) Update() error {

	isSpacebar := inpututil.IsKeyJustReleased(ebiten.KeySpace)

	if isSpacebar {

		if g.state == ENTRY {
			g.state = PLAYING
		} else if g.state == PLAYING {
			ball.throw()
		} else {
			g.state = PLAYING
			ball.fallDown = false
		}

	}

	basket1.Update()
	basket2.Update()
	ball.Update()

	if ball.checkForCollision {
		if ball.x > basket2.x-120 && ball.x < basket2.x+120 {
			// Successfull collision
			ball.basket = basket2
			basket1.GoDown()
			basket2.GoDown()
			basketNew := initBasket(DEFAULT_BASKET_POS_X, -200, "b3")
			basketNew.GoDown()

			time.AfterFunc(1*time.Second, func() {
				basket1 = basket2
				basket2 = basketNew
			})
		} else {
			// No collision. Prepare to end the game
			ball.fallDown = true

		}

		ball.checkForCollision = false
	}

	if ball.fallDown && ball.y >= WINDOW_HEIGHT {
		g.state = EXIT

		// log.Fatal("End the game")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	// if g.state == ENTRY || g.state == EXIT {
	portal.Draw(g, screen)
	// }

	if g.state == PLAYING {
		screen.DrawImage(bgImage, nil)
		basket1.Draw(screen)
		basket2.Draw(screen)
		ball.Draw(screen)
	}
	// basket3.Draw(screen)
}

func (g *Game) Layout(w int, h int) (sW, sH int) {
	return 600, 900
}

func PlayGame() {
	ebiten.SetWindowSize(600, 900)
	ebiten.SetWindowTitle("Eggs")

	if err := ebiten.RunGame(&Game{state: ENTRY}); err != nil {
		log.Fatal(err)
	}
}
