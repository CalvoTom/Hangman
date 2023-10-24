package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawHangman(state int) {
	drawHangmanBox()
	// switch {
	// case state == 1:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n\n\n\n\n\n\n\n ========= ")
	// case state == 2:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n|\n|\n|\n|\n|\n ========= ")
	// case state == 3:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n       |\n       |\n       |\n       |\n       |\n ========= ")
	// case state == 4:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n       |\n       |\n       |\n       |\n ========= ")
	// case state == 5:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n       |\n       |\n       |\n ========= ")
	// case state == 6:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n   |   |\n       |\n       |\n ========= ")
	// case state == 7:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n  /|   |\n       |\n       |\n ========= ")
	// case state == 8:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n  /|\\  |\n       |\n       |\n ========= ")
	// case state == 9:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n  /|\\  |\n  /    |\n       |\n ========= ")
	// case state == 10:
	// 	TbPrint(2, 5, termbox.ColorDefault, termbox.ColorDefault, "\n   +---+\n   |   |\n   O   |\n  /|\\  |\n  / \\  |\n       |\n ========= ")
	// default:

	// }
}

func drawHangmanBox() {
	// Taille du carré
	width, height := 25, 27

	// Coordonnées du coin supérieur gauche du carré
	x, y := 0, 0

	for i := 1; i < width-1; i++ {
		termbox.SetCell(x+i, y, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+i, y+height-1, borderVertical, termbox.ColorDefault, termbox.ColorDefault)
	}

	for i := 1; i < height-1; i++ {
		termbox.SetCell(x, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+width-1, y+i, borderHorizontal, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.SetCell(0, 0, borderTopLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 0, borderTopRight, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, 26, borderBotomLeft, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(24, 26, borderBottomRight, termbox.ColorDefault, termbox.ColorDefault)

	TbPrint(1, 0, termbox.ColorCyan, termbox.ColorDefault, "Hangman")
}
