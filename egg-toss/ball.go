package eggtoss

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	x                 float64
	y                 float64
	speed             Speed
	movenment         string
	dir               string
	isThrow           bool
	basket            *Basket
	checkForCollision bool
	fallDown          bool
}

var thrown = 0.0

func (b *Ball) throw() {
	b.isThrow = true
}

func initBall(basket *Basket) *Ball {
	ball := &Ball{basket.x, basket.y, basket.speed, basket.movenment, basket.dir, false, basket, false, false}

	return ball
}

func (b *Ball) Update() error {

	if b.fallDown {
		b.y = b.y + 10
		return nil
	}

	if !b.isThrow {
		b.y = b.basket.y
		b.x = b.basket.x
		return nil
	}

	thrown += 15

	if thrown >= 700 {
		thrown = 0
		b.isThrow = false
		b.checkForCollision = true
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
