package hangman

import (
	"github.com/nsf/termbox-go"
)

func DrawVictory() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	size := 0
	for _, letter := range "vicoire" {
		DisplayASCII("standard.txt", 28+size, 2, letter)
		size += 11
	}
	termbox.Flush()
}
