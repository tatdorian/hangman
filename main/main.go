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
		fmt.Printf("Erreur pour charger le mot %s: %v\n", wordsFile, err)
		return
	}

	word := hangman.RandomWord(words)
	essai := 10
	display := revealRandomLetters(word)

	PositionsFile := "hangman.txt"
	Positions, err := hangman.LoadHangmanPositions(PositionsFile)
	if err != nil {
		fmt.Printf("Erreur pour afficher les positions : %v\n", err)
		return
	}

	fmt.Println("Bonne Chance. Vous avez 10 essais")
	fmt.Println(display)

	for essai > 0 {
		fmt.Print("")
		fmt.Print("Choose a letter: ")
		var letter string
		fmt.Scanln(&letter)
		if len(letter) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}
		if strings.Contains(word, letter) {
			display = hangman.RevealLetter(word, display, letter)
			fmt.Println(display)
			if display == word {
				fmt.Println("Bravo !")
				return
			}
		} else {
			fmt.Printf("Not present in the word, %d attempts remaining\n", essai)
			hangman.DisplayHangman(Positions, 10-essai)
			essai--
		}
	}
	fmt.Printf("You've run out of attempts. The word was %s.\n", word)
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
