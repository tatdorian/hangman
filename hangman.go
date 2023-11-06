package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type HangmanData struct {
	Word        string
	WordDisplay string
	essai       int
	Positions   []string
}

func LoadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, scanner.Err()
}

func RandomWord(words []string) string {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}

func RevealLetter(word, display, letter string) string {
	var builder strings.Builder
	for i := 0; i < len(word); i++ {
		if word[i] == letter[0] || display[i] != '_' {
			builder.WriteByte(word[i])
		} else {
			builder.WriteByte('_')
		}
	}
	return builder.String()
}

func LoadHangmanPositions(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var positions []string
	scanner := bufio.NewScanner(file)
	var position strings.Builder

	for scanner.Scan() {
		if scanner.Text() == "" {
			positions = append(positions, position.String())
			position.Reset()
		} else {
			position.WriteString(scanner.Text() + "\n")
		}
	}
	positions = append(positions, position.String())

	return positions, scanner.Err()
}

func DisplayHangman(positions []string, essai int) {
	if essai < 0 || essai >= len(positions) {
		return
	}
	fmt.Println(positions[essai])
}
