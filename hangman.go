package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var word string
	var int_input int
	ListInput := []string{}
	bad_guesses := 0
	attempts := 10
	fmt.Print("Welcome to the Hangman game!\n")
	word, find_word := FindWord(Difficulty())
	if find_word == false {
		return
	}
	fmt.Print("Try to find the word by writing a letter!\n")
	CharOfWord := CharOfWord(word)
	for attempts > 0 && find_word == true {
		input := ReadInput()
		if len(input) > 1 {
			fmt.Print("Too much input\n")
		}
		if len(input) < 1 {
			fmt.Print("Not enough input\n")
		} else {
			int_input, ListInput = VerifyInput(input, word, ListInput)
			if int_input == -1 {
				continue
			} else if int_input == 0 {
				bad_guesses++
				PrintHangman(bad_guesses)
				fmt.Printf("You have %d attempts left.\n", attempts-bad_guesses)
			}
			CharOfWord = PrintWord(CharOfWord, word, attempts, int_input, input)
			if CompareWordAndListInput(CharOfWord, ListInput) {
				fmt.Print("You won\n")
				return
			}
			if attempts-bad_guesses == 0 {
				fmt.Print("You lost, the word was : ", word, "\n")
				break
			}
			attempts--
		}
	}
	return
}

func PrintHangman(bad_guesses int) {
	start := 1 * bad_guesses
	end := 7*bad_guesses + start
	read_line := 1
	hangman_file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return
	}
	defer hangman_file.Close()
	hangman_scan := bufio.NewScanner(hangman_file)
	for hangman_scan.Scan() {
		if lineread := hangman_scan.Text(); read_line >= start && read_line <= end {
			fmt.Println(lineread)
		}
		read_line++
	}
}

func VerifyInput(s string, mot string, ListInput []string) (int, []string) {
	if strings.Contains(strings.Join(ListInput, ""), s) {
		fmt.Print("Already used.\n")
		return -1, ListInput
	}
	runeS := []rune(s)[0]
	if strings.ContainsRune(mot, runeS) {
		return 1, append(ListInput, s)
	} else {
		return 0, append(ListInput, s)
	}
}

func ReadInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func CharOfWord(mot string) []string {
	listChar := []string{}
	for _, char := range mot {
		for _, r := range listChar {
			for _, charoflist := range r {
				if char == charoflist {
					break
				} else {
					listChar = append(listChar, string(char))
				}
			}
		}
	}
	return listChar
}

func CompareWordAndListInput(WordChar, ListInput []string) bool {
	InputRight := 0
	if len(WordChar) != len(ListInput) {
		return false
	}
	for i := 0; i < len(WordChar)-1; i++ {
		if WordChar[i] != ListInput[i] {
			return false
		} else {
			InputRight++
		}
	}
	if InputRight == len(WordChar) {
		return true
	}
	return false
}

func PrintWord(CharOfWord []string, word string, attempts, int_input int, input string) []string {
	if int_input == 1 {
		CharOfWord = append(CharOfWord, input)
	}
	for _, char := range word {
		if strings.Contains(strings.Join(CharOfWord, ""), string(char)) {
			fmt.Print(string(char))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
	return CharOfWord
}

func FindWord(dificult string) (string, bool) {
	find_word := false
	var word string
	file, err := os.Open(dificult)
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return word, false
	}
	defer file.Close()
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		words := []string{}
		for scan.Scan() {
			words = append(words, scan.Text())
		}
		if len(words) > 0 {
			word = words[rand.Intn(len(words))]
			find_word = true
		}
		find_word = true
	}
	if find_word == false {
		fmt.Print("No word found\n")
		return word, false
	}
	return word, true
}

func Difficulty() string {
	fmt.Println("Choose your level: ")
	var level string
	fmt.Scanln(&level)
	var dificult string
	switch level {
	case "1":
		dificult = "words.txt"
	case "2":
		dificult = "words2.txt"
	case "3":
		dificult = "words3.txt"
	default:
		fmt.Println("Invalid input")
	}
	return dificult
}
