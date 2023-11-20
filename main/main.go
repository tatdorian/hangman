package main

import (
	"fmt"
	"hangman"
	"math/rand"
	"strings"
)

func main() {
	wordsFile := "words.txt"
	words, err := hangman.LoadWords(wordsFile)
	if err != nil {
		fmt.Printf("Error loading words from %s: %v\n", wordsFile, err)
		return
	}

	guessedLetters := []string{}
	word := hangman.RandomWord(words)
	attempts := 10
	display := revealRandomLetters(word)

	PositionsFile := "hangman.txt"
	Positions, err := hangman.LoadHangmanPositions(PositionsFile)
	if err != nil {
		fmt.Printf("Error loading hangman positions: %v\n", err)
		return
	}

	fmt.Println("Good Luck. You have 10 attempts")
	fmt.Println(display)

	for attempts > 0 {
		fmt.Print("Choose a letter or a word: ")
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
				display = hangman.RevealLetter(word, display, letter)
				fmt.Println(display)
				if display == word {
					fmt.Println("Congratulations!")
					return
				}
			} else {
				attempts--
				fmt.Printf("Letter not found, %d attempts remaining:\n", attempts)
				hangman.DisplayHangman(Positions, 10-attempts)
			}
		} else if len(input) >= 2 {
			if input == word {
				fmt.Println("Congratulations! You guessed the word.")
				return
			} else {
				attempts -= 2
				fmt.Printf("Incorrect word, %d attempts remaining:\n", attempts)
				hangman.DisplayHangman(Positions, 10-attempts)
			}
		} else {
			fmt.Println("Please enter a valid letter or word.")
		}
	}

	fmt.Printf("You have run out of attempts. The word was %s.\n", word)
}
func revealRandomLetters(word string) string {
	n := len(word)/2 - 1
	randomIndices := generateRandomIndices(len(word), n)
	display := strings.Repeat("_", len(word))

	for _, idx := range randomIndices {
		display = hangman.RevealLetter(word, display, word[idx:idx+1])
	}

	return display
}

func generateRandomIndices(length, n int) []int {
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
