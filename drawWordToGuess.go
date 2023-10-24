package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawWordToGuess(word string) {
	// Taille du carré
	width, height := 35, 9

	// Coordonnées du coin supérieur gauche du carré
	x, y := 25, 0

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}

	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(25, 0, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 0, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 8, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 8, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	for i, letter := range word {
		termbox.SetCell(i+40, 4, letter, termbox.ColorDefault, termbox.ColorDefault)
	}

	TbPrint(26, 0, termbox.ColorCyan, termbox.ColorDefault, "\n Mot")
}
