package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawLettersTried(letters string) {
	// Taille du carré
	width, height := 35, 9

	// Coordonnées du coin supérieur gauche du carré
	x, y := 25, 18

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(25, 18, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 18, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 26, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 26, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	for i, letter := range letters {
		termbox.SetCell(i+27, 22, letter, termbox.ColorDefault, termbox.ColorDefault)
	}

	TbPrint(26, 18, termbox.ColorCyan, termbox.ColorDefault, "Lettre testé")
}
