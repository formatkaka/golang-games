package eggtoss

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	STATIC = "static"
	FULL   = "full-movenment"

	SLOW   = "slow"
	MEDIUM = "medium"
	FAST   = "fast"

	SLOW_SPEED   = 2
	MEDIUM_SPEED = 5
	FAST_SPEED   = 10
)

type Speed struct {
	speedType string
	speed     float64
}

var speeds = [3]*Speed{
	{speedType: SLOW, speed: SLOW_SPEED},
	{speedType: MEDIUM, speed: MEDIUM_SPEED},
	{speedType: FAST, speed: FAST_SPEED},
}

var movenments = [2]string{FULL, STATIC}

type Basket struct {
	x         float64
	y         float64
	speed     Speed
	movenment string
	dir       string
}

func initBasket(defaultX, defaultY float64) *Basket {
	rand.Seed(time.Now().UnixNano())
	movenment := rand.Intn(2)
	speed := rand.Intn(3)

	basket := &Basket{defaultX, defaultY, *speeds[speed], movenments[movenment], LEFT}

	fmt.Println(basket)

	return basket
}

func (b *Basket) Update() {
	if b.movenment == STATIC {
		return
	}

	if b.x <= 1 {
		b.dir = RIGHT
	} else if b.x >= WINDOW_WIDTH-IMAGE_WIDTH {
		b.dir = LEFT
	}

	if b.dir == LEFT {
		b.x = b.x - b.speed.speed
	} else {
		b.x = b.x + b.speed.speed
	}

}

func (b *Basket) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(b.x, b.y)

	screen.DrawImage(basketImg, op)
}
