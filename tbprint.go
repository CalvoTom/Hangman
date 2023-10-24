package hangman

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func TbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for i := 0; i < len(text); i++ {
		if rune(text[i]) == 10 {
			y += 1
		} else {
			termbox.SetCell(x, y, rune(text[i]), fg, bg)
			x += runewidth.RuneWidth(rune(text[i]))
		}
	}
}
