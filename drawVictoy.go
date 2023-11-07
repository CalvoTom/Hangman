package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawVictory() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	TbPrint(15, 5, termbox.ColorDefault, termbox.ColorDefault, "Félicitation !")
	termbox.Flush()
}
