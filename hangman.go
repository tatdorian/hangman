package hangmanweb

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
func Main() {
	wordsFile := "words.txt"
	words, err := LoadWords(wordsFile)
	if err != nil {
		fmt.Printf("Error loading words from %s: %v\n", wordsFile, err)
		return
	}

	guessedLetters := []string{}
	word := RandomWord(words)
	attempts := 10
	display := revealLetters(word)

	PositionsFile := "hangman.txt"
	Positions, err := LoadHangmanPositions(PositionsFile)
	if err != nil {
		fmt.Printf("Error loading hangman positions: %v\n", err)
		return
	}

	fmt.Println("Good Luck. You have 10 attempts")
	fmt.Println(display)

	for attempts > 0 {
		fmt.Print("\nChoose a letter or a word: ")
		var input string
		fmt.Scanln(&input)

		if len(input) == 1 {
			letter := input
			if contains(guessedLetters, letter) {
				fmt.Printf("You have already guessed this letter '%s'.\n", letter)
				continue
			}
			guessedLetters = append(guessedLetters, letter)

			if strings.Contains(word, letter) {
				display = RevealLetter(word, display, letter)
				fmt.Println(display)
				if display == word {
					fmt.Println("\n Congratulations!")
					return
				}
			} else {
				attempts--
				fmt.Printf("Letter not found, %d attempts remaining:\n", attempts)
				DisplayHangman(Positions, 10-attempts)
				fmt.Printf(display)
			}
		} else if len(input) >= 2 {
			if input == word {
				fmt.Println("\n Congratulations! You guessed the word.")
				return
			} else {
				attempts -= 2
				fmt.Printf("Incorrect word, %d attempts remaining:\n", attempts)
				DisplayHangman(Positions, 10-attempts)
				fmt.Printf(display)
			}
		} else {
			fmt.Println("Please enter a valid letter or word.")
		}
	}

	fmt.Printf("\n You have run out of attempts. The word was %s.\n", word)
}
func revealLetters(word string) string {
	n := len(word)/2 - 1
	randomIndices := generateIndices(len(word), n)
	display := strings.Repeat("_", len(word))

	for _, idx := range randomIndices {
		display = RevealLetter(word, display, word[idx:idx+1])
	}

	return display
}

func generateIndices(length, n int) []int {
	if n > length {
		n = length
	}
	indices := rand.Perm(length)[:n]
	return indices
}
func contains(slice []string, letter string) bool {
	for _, l := range slice {
		if l == letter {
			return true
		}
	}
	return false
}
