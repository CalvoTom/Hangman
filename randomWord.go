package hangman

import (
	"bufio"
	"math/rand"
	"os"
)

func RandomWord() string {
	var word string
	randomWordLine := rand.Intn(10)

	file, err := os.Open("word.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		if lineNumber == randomWordLine {
			word = scanner.Text()
			break
		}
	}
	return word
}
