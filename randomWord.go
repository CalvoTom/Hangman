package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func RandomWord() {
	f, err := os.Open("thermopylae.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
