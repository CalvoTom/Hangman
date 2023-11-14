package hangman

import "github.com/nsf/termbox-go"

func DrawButtonsSave() {
	// Taille du Bouton
	width, height := 6, 3

	// Coordonnées du coin supérieur gauche du bouton
	x, y := 25, 29

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(25, 29, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(30, 29, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(25, 31, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(30, 31, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	TbPrint(26, 30, termbox.ColorCyan, termbox.ColorDefault, "SAVE")
}

func DrawButtonsQuit() {
	// Taille du Bouton
	width, height := 6, 3

	// Coordonnées du coin supérieur gauche du bouton
	x, y := 31, 29

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(31, 29, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(36, 29, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(31, 31, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(36, 31, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	TbPrint(32, 30, termbox.ColorCyan, termbox.ColorDefault, "QUIT")
}
