package eggtoss

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ENTRY   = "entry screen"
	EXIT    = "exit screen"
	PLAYING = "playing"
)

type Portal struct {
	fontType font.Face
	image    *ebiten.Image
}

func initPortal() *Portal {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Portal{fontType: mplusNormalFont, image: ebiten.NewImage(WINDOW_WIDTH, WINDOW_HEIGHT)}
}

func (p *Portal) Update(screentype string) error {

	// p.screentype = screentype

	return nil
}

func (p *Portal) Draw(g *Game, screen *ebiten.Image) {
	// fmt.Println("g state : ", g.state)
	p.image.DrawImage(bgImage, nil)
	if g.state == ENTRY {

		text.Draw(p.image, "START GAME", p.fontType, WINDOW_WIDTH/2-70, WINDOW_HEIGHT-100, color.White)
		text.Draw(p.image, "Press space to start", p.fontType, WINDOW_WIDTH/2-120, WINDOW_HEIGHT-50, color.White)

	} else if g.state == EXIT {
		text.Draw(p.image, "GAME OVER", p.fontType, WINDOW_WIDTH/2-70, WINDOW_HEIGHT-100, color.White)
		text.Draw(p.image, "Press space to play again", p.fontType, WINDOW_WIDTH/2-125, WINDOW_HEIGHT-50, color.White)
	} else {
		p.image.Clear()
	}

	screen.DrawImage(p.image, nil)

}
