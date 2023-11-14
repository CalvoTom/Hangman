package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawWordToGuess(hangman *HangManData) {
	// Taille du carré
	width, height := 100, 13

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
	termbox.SetCell(124, 0, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 12, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(124, 12, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)
	if hangman.IsASCII {
		size := 0
		for _, letter := range hangman.Word {
			DisplayASCII(hangman.Police, 28+size, 2, letter)
			size += 11
		}
	} else {
		TbPrint(27, 5, termbox.ColorDefault, termbox.ColorDefault, hangman.Word)
	}
	TbPrint(26, 0, termbox.ColorCyan, termbox.ColorDefault, "Mot")
}
