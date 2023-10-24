package hangman

import (
	"log"
	"strings"

	"github.com/nsf/termbox-go"
)

var borderTopLeft rune = 0x250C
var borderTopRight rune = 0x2510
var borderBotomLeft rune = 0x2514
var borderBottomRight rune = 0x2518
var borderHorizontal rune = 0x2502
var borderVertical rune = 0x2500

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func fillbg(bg termbox.Attribute) {
	w, h := termbox.Size()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			termbox.SetBg(x, y, bg)
		}
	}
}

func Drawer() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	hangman := HangManData{"__c", "abc", 0, hangmanArray}
	var lettersTried string
	var lettersFind string

loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DrawWordToGuess(string(hangman.Word))
		DrawLettersTried(lettersTried)
		DrawLettersFind(lettersFind)
		DrawHangman(hangman.Attempts)
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			} else if ev.Ch != 0 {
				guess := string(ev.Ch)
				if strings.Contains(string(hangman.ToFind), guess) {
					for i, letter := range hangman.ToFind {
						if string(letter) == guess {
							hangman.Word[i] = letter
							if strings.Contains(string(lettersFind), guess) {
								continue
							} else {
								lettersFind += guess
								lettersTried += guess
							}
						}
					}
				} else {
					if strings.Contains(string(lettersTried), guess) {
						continue
					} else {
						lettersTried += guess
						hangman.Attempts++
					}
				}
			}
		}
	}
}
