package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	game_func "github.com/tatdorian/hangman-classic/game_func"
	"github.com/tatdorian/hangman-classic/presentations"
)

func Game() {
	args := os.Args[1:]
	if len(args) == 1 {
		// normal game case
		word := strings.ToUpper(game_func.RandomWord(os.Args[1]))
		randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
		var hiddenLetters []int   // stocking the indices of the hidden letters
		var revealedLetters []int // stocking the indices of the revealed letters
		wordRune := []rune(word)
		for i := 0; i < len(word)/2-1; i++ {
			randTemp := randSource.Intn(len(wordRune))
			if wordRune[randTemp] != 0 {
				revealedLetters = append(revealedLetters, randTemp)
				wordRune[randTemp] = 0
			} else {
				i--
			}
		}
		for j := 0; j < len(wordRune); j++ {
			if wordRune[j] != 0 {
				hiddenLetters = append(hiddenLetters, j)
			}
		}
		for _, i := range revealedLetters {
			wordRune[i] = rune(word[i])
		}
		for _, i := range hiddenLetters {
			wordRune[i] = '_'
		}
		attempts := 10
		var lettersSuggested []string
		fmt.Printf("Good Luck, you have 10 attempts.\n%s\n", string(wordRune))
		for attempts > 0 {
			if string(wordRune) == word {
				fmt.Printf("%s\n", "Congrats !")
				return
			}
			var letter string
			_, err := fmt.Scan(&letter)
			if err != nil {
				fmt.Println(err)
			}
			letter = strings.ToUpper(letter)
			if len(letter) > 1 {
				if letter == "STOP" {
					game_func.SaveGame(&game_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: ""})
					fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
					return
				}
				if letter == word {
					fmt.Printf("%s\n", "Congrats !")
					return
				} else {
					attempts -= 2
					fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
					presentations.PrintMan(attempts)
				}
			} else {
				if game_func.ContainsTable(lettersSuggested, letter) {
					fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
				} else {
					if game_func.ContainsString(word, letter) {
						lettersSuggested = append(lettersSuggested, letter)
						indexes := game_func.LetterInWorld(word, letter)
						for _, i := range indexes {
							wordRune[i] = rune(word[i])
						}
						fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
					} else {
						attempts--
						fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
						presentations.PrintMan(attempts)
					}
				}
			}
		}
		fmt.Printf("You Loose !\nThe word to find was %s", word)
	} else if len(args) == 2 && args[0] == "--startWith" && args[1] == "save.txt" {
		// save game case
		data, err := os.ReadFile("save.txt")
		if err != nil {
			fmt.Println(err)
		}
		var GameStats *game_func.NewSave
		if err2 := json.Unmarshal(data, &GameStats); err != nil {
			fmt.Println(err2)
		}
		word, wordRune, attempts, lettersSuggested, LetterFile := GameStats.Word, GameStats.WordRune, GameStats.Attempts, GameStats.LettersSuggested, GameStats.LetterFile
		if LetterFile == "" {
			fmt.Printf("Welcome Back, you have %v attempts remaining.\n%s\n", attempts, string(wordRune))
			for attempts > 0 {
				if string(wordRune) == word {
					fmt.Printf("%s\n", "Congrats !")
					return
				}
				var letter string
				_, err := fmt.Scan(&letter)
				if err != nil {
					fmt.Println(err)
				}
				letter = strings.ToUpper(letter)
				if len(letter) > 1 {
					if letter == "STOP" {
						game_func.SaveGame(&game_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: LetterFile})
						fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
						return
					}
					if letter == word {
						fmt.Printf("%s\n", "Congrats !")
						return
					} else {
						attempts--
						fmt.Printf("Wrong ! You have %d attempts left.\n%s\n", attempts, string(wordRune))
						presentations.PrintMan(attempts)
					}
				} else {
					if game_func.ContainsTable(lettersSuggested, letter) {
						fmt.Printf("You already suggested this letter !\n%s\n", string(wordRune))
					} else {
						if game_func.ContainsString(word, letter) {
							lettersSuggested = append(lettersSuggested, letter)
							indexes := game_func.LetterInWorld(word, letter)
							for _, i := range indexes {
								wordRune[i] = rune(word[i])
							}
							fmt.Printf("Choose: %s\n%s\n", letter, string(wordRune))
						} else {
							attempts--
							fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
							presentations.PrintMan(attempts)
						}
					}
				}
			}
			fmt.Printf("You Loose !\nThe word to find was %s", word)
		} else {
			// letter file case
			fmt.Printf("Welcome Back, you have %v attempts remaining.\n", attempts)
			for attempts > 0 {
				if string(wordRune) == word {
					return
				}
				var letter string
				_, err := fmt.Scan(&letter)
				if err != nil {
					fmt.Println(err)
				}
				letter = strings.ToUpper(letter)
				if len(letter) > 1 {
					if letter == "STOP" {
						game_func.SaveGame(&game_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: LetterFile})
						fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
						return
					}
					if letter == word {
						return
					} else {
						attempts--
						fmt.Printf("Wrong ! You have %d attempts left.\n", attempts)
						presentations.PrintMan(attempts)
					}
				} else {
					if game_func.ContainsTable(lettersSuggested, letter) {
						fmt.Printf("You already suggested this letter !\n")
					} else {
						if game_func.ContainsString(word, letter) {
							lettersSuggested = append(lettersSuggested, letter)
							indexes := game_func.LetterInWorld(word, letter)
							for _, i := range indexes {
								wordRune[i] = rune(word[i])
							}
							fmt.Printf("Choose: %s\n", letter)
						} else {
							attempts--
							fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
							presentations.PrintMan(attempts)
						}
					}
				}
			}
			fmt.Printf("You Loose !\nThe word to find was \n")
		}
	} else if len(args) == 3 && args[1] == "--letterFile" {
		// letter file case
		word := strings.ToUpper(game_func.RandomWord(os.Args[1]))
		randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
		var hiddenLetters []int   // stocking the indices of the hidden letters
		var revealedLetters []int // stocking the indices of the revealed letters
		wordRune := []rune(word)
		for i := 0; i < len(word)/2-1; i++ {
			randTemp := randSource.Intn(len(wordRune))
			if wordRune[randTemp] != 0 {
				revealedLetters = append(revealedLetters, randTemp)
				wordRune[randTemp] = 0
			} else {
				i--
			}
		}
		for j := 0; j < len(wordRune); j++ {
			if wordRune[j] != 0 {
				hiddenLetters = append(hiddenLetters, j)
			}
		}
		for _, i := range revealedLetters {
			wordRune[i] = rune(word[i])
		}
		for _, i := range hiddenLetters {
			wordRune[i] = '_'
		}
		attempts := 10
		var lettersSuggested []string
		fmt.Printf("Good Luck, you have 10 attempts.\n")
		for attempts > 0 {
			if string(wordRune) == word {
				return
			}
			var letter string
			_, err := fmt.Scan(&letter)
			if err != nil {
				fmt.Println(err)
			}
			letter = strings.ToUpper(letter)
			if len(letter) > 1 {
				if letter == "STOP" {
					game_func.SaveGame(&game_func.NewSave{Word: word, WordRune: wordRune, Attempts: attempts, LettersSuggested: lettersSuggested, LetterFile: args[2]})
					fmt.Printf("Choose: %s\n\nGame Saved in save.txt.\n", letter)
					return
				}
				if letter == word {
					return
				} else {
					attempts -= 2
					fmt.Printf("Wrong ! You have %d attempts left.\n", attempts)
					presentations.PrintMan(attempts)
				}
			} else {
				if game_func.ContainsTable(lettersSuggested, letter) {
					fmt.Printf("You already suggested this letter !\n")
				} else {
					if game_func.ContainsString(word, letter) {
						lettersSuggested = append(lettersSuggested, letter)
						indexes := game_func.LetterInWorld(word, letter)
						for _, i := range indexes {
							wordRune[i] = rune(word[i])
						}
						fmt.Printf("Choose: %s\n", letter)
					} else {
						attempts--
						fmt.Printf("Choose: %s\nNot present in the Word, %v attempts remaining\n", letter, attempts)
						presentations.PrintMan(attempts)
					}
				}
			}
		}
		fmt.Printf("You Loose !\nThe word to find was\n")
	} else {
		fmt.Println("Impossible...")
	}
}

func main() {
	Game()
}
