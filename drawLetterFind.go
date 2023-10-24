package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawLettersFind(letters string) {
	// Taille du carré
	width, height := 35, 9

	// Coordonnées du coin supérieur gauche du carré
	x, y := 25, 9

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(25, 9, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 9, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 17, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(59, 17, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	for i, letter := range letters {
		termbox.SetCell(i+27, 13, letter, termbox.ColorDefault, termbox.ColorDefault)
	}

	TbPrint(26, 9, termbox.ColorCyan, termbox.ColorDefault, "Lettre trouvé \n\n\n\n")
}
