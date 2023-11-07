package hangman

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/nsf/termbox-go"
)

const (
	borderTopLeft     = 0x250C
	borderTopRight    = 0x2510
	borderBotomLeft   = 0x2514
	borderBottomRight = 0x2518
	borderHorizontal  = 0x2502
	borderVertical    = 0x2500
)

type HangManData struct {
	Word     string // Word composed of '_', ex: H_ll_
	ToFind   string // Final word chosen by the program at the beginning. It is the word to find
	Attempts int    // Number of attempts left
}

func Drawer() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	file := "words.txt"
	police := "standard.txt"

	flag.String("letterFile", "default", "File of word to start with")
	flag.Parse()
	if len(os.Args[1:]) == 2 {
		police = os.Args[2]
	}

	lettersFind, hangman := InitialiseStruc(file)
	var lettersTried string

loop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		DrawWordToGuess(hangman, police)
		DrawAttempts(hangman)
		DrawLettersTried(lettersTried)
		DrawLettersFind(lettersFind)
		DrawHangman(hangman)
		termbox.Flush()
		if hangman.Word == hangman.ToFind {
			break loop
		}
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || hangman.Attempts >= 10 {
				break loop
			} else if IsAlpha(string(ev.Ch)) {
				guess := string(ev.Ch)
				if strings.Contains(string(hangman.ToFind), guess) {
					for i, letter := range hangman.ToFind {
						if string(letter) == guess {
							hangman.Word = ReplaceAtIndex(hangman.Word, letter, i)
							if strings.Contains(string(lettersFind), guess) {
								continue
							} else {
								lettersFind += guess
								lettersTried += guess
							}
						}
					}
				} else {
					hangman.Attempts += 1
					if strings.Contains(string(lettersTried), guess) {
						continue
					} else {
						lettersTried += guess
					}
				}
			}
		}
	}
}
