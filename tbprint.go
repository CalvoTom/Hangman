package hangman

import (
	"github.com/nsf/termbox-go"
)

func TbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
