package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/nsf/termbox-go"
)

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func IsAlpha(str string) bool {
	var alphanumeric = regexp.MustCompile("^[[A-zÀ-û]*$")
	return alphanumeric.MatchString(str)
}

func ReplaceAtIndex2(input string, replacement byte, index int) string {
	return strings.Join([]string{input[:index], string(replacement), input[index+1:]}, "")
}

func InitialiseStruc(filename string) (string, *HangManData) {
	var hangman *HangManData
	hangman = new(HangManData)
	RandomWord(hangman, filename)
	arrayTf := []rune(hangman.ToFind)
	for i := 0; i < len(arrayTf); i++ {
		hangman.Word += "_"
	}
	lettersFind := HideWord(hangman)

	return lettersFind, hangman
}

func RandomWord(hangman *HangManData, filename string) {
	randomWordLine := rand.Intn(36-1+1) + 1

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		if lineNumber == randomWordLine {
			hangman.ToFind = scanner.Text()
			break
		}
	}
}

func HideWord(hangman *HangManData) string {
	var lettersFind string
	arrayWord := []rune(hangman.Word)
	arrayTofind := []rune(hangman.ToFind)
	nbLetterReveal := len(arrayTofind)/2 - 1

	for i := 0; i <= nbLetterReveal; i++ {
		randomIndex := rand.Intn(len(arrayWord) - 1)
		arrayWord[randomIndex] = arrayTofind[randomIndex]

		if !strings.Contains(lettersFind, string(arrayTofind[randomIndex])) {
			lettersFind += string(arrayTofind[randomIndex])
		}
		for index, ch := range arrayTofind {
			if ch == arrayTofind[randomIndex] {
				arrayWord[index] = arrayTofind[randomIndex]
			}
		}
	}
	hangman.Word = string(arrayWord)
	return lettersFind
}

func DisplayASCII(filename string, startX, startY int, letter rune) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	liner := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++
		if letter == '_' {
			if liner >= 9 {
				break
			}
			if lineNumber >= 568 {
				drawLine(startX, startY+liner, line)
				liner++
			}

		} else if lineNumber >= 586 {
			if liner >= 9 {
				break
			}
			if lineNumber >= ((int(letter)-97)*9)+586 {
				drawLine(startX, startY+liner, line)
				liner++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func DrawLine(x, y int, line string) {
	for i, ch := range line {
		termbox.SetCell(x+i, y, ch, termbox.ColorDefault, termbox.ColorDefault)
	}
}
