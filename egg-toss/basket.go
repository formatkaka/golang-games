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

	SLOW_SPEED   = 1
	MEDIUM_SPEED = 2.5
	FAST_SPEED   = 5
)

var Y_POS = [3]float64{100, DEFAULT_BASKET_POS_Y, 1500}

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
	name      string
	x         float64
	y         float64
	speed     Speed
	movenment string
	dir       string
	down      bool
}

func initBasket(defaultX, defaultY float64, name string) *Basket {
	rand.Seed(time.Now().UnixNano())
	movenment := rand.Intn(2)
	speed := rand.Intn(3)

	basket := &Basket{name, defaultX, defaultY, *speeds[speed], movenments[movenment], LEFT, false}

	return basket
}

func (b *Basket) Update() {
	if b.movenment != STATIC {

		if b.x <= 0 {
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

	fmt.Println(b)
	if b.down && !contains(Y_POS, b.y) {
		b.y = b.y + 5
	} else if b.down && contains(Y_POS, b.y) {
		b.down = false
	}

}

func (b *Basket) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(b.x, b.y)

	screen.DrawImage(basketImg, op)
}

// -600,100,800
func (b *Basket) GoDown() {
	b.down = true
	b.y = b.y + 5
}
