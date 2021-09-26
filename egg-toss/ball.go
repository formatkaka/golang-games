package eggtoss

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x         float64
	y         float64
	speed     Speed
	movenment string
	dir       string
	isThrow   bool
	basket    *Basket
}

var thrown = 0.0

func (b *Ball) throw() {
	b.isThrow = true
}

func initBall(basket *Basket) *Ball {
	ball := &Ball{basket.x, basket.y, basket.speed, basket.movenment, basket.dir, false, basket}

	return ball
}

func (b *Ball) Update(basket1 *Basket, basket2 *Basket) error {
	if !b.isThrow {
		b.y = b.basket.y
		b.x = b.basket.x
		return nil
	}

	thrown += 15

	if thrown >= 700 {
		thrown = 0
		b.isThrow = false

		if b.x > basket2.x-120 && b.x < basket2.x+120 {

			b.basket = basket2
		}
	}

	return nil

}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	if b.isThrow {
		op.GeoM.Translate(b.x, b.y-thrown)
	} else {
		op.GeoM.Translate(b.x, b.y)
	}
	screen.DrawImage(eggImage, op)
}
