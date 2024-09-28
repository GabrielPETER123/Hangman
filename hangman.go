package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {
	AlreadyGuessed := false
	var ListInput []rune
	bad_guesses := 0
	attempts := 10
	fmt.Print("Welcome to the Hangman game!\n")
	word := FindWord(Difficulty())
	CharOfWord := CharOfWord(TransStringToRune(word))
	fmt.Print(PrintWord(CharOfWord, ListInput, word), "\n")
	fmt.Print("Try to find the word by writing a letter!\n")
	for attempts > 0 {
		input := ReadInput()
		RInput := TransStringToRune(input)
		if len(RInput) > 1 {
			fmt.Print("Too much input\n")
		}
		if len(RInput) < 1 {
			fmt.Print("Not enough input\n")
		}
		if !(VerifyInput(CharOfWord, RInput)) && len(RInput) == 1 {
			bad_guesses++
			PrintHangman(bad_guesses)
			fmt.Printf("Bad Guess\n")
		} else if len(RInput) == 1 {
			for _, char := range ListInput {
				if char == RInput[0] {
					fmt.Printf("Already guessed\n")
					AlreadyGuessed = true
					break
				}
			}
			if !AlreadyGuessed {
				ListInput = append(ListInput, RInput[0])
				fmt.Printf("Good Guess\n")
				AlreadyGuessed = false
			}
		}
		fmt.Print(PrintWord(CharOfWord, ListInput, word), "\n")
		fmt.Print("Attempts left: ", attempts, "\n")
		attempts--
		if Compare(CharOfWord, ListInput) {
			fmt.Print("You won\n")
			fmt.Print("The word was: ", word, "\n")
			return
		}
		if attempts == 0 {
			fmt.Print("You lost, the word was : ", word, "\n")
			break
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

func VerifyInput(CharOfWord, Input []rune) bool {
	for _, char := range CharOfWord {
		if char == Input[0] {
			return false
		}
	}
	return true
}

func ReadInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func CharOfWord(word []rune) []rune {
	var CharOfWord []rune
	CharOfWord = append(CharOfWord, word[0])
	for i := 1; i < len(word); i++ {
		for _, char := range CharOfWord {
			if char != word[i] {
				CharOfWord = append(CharOfWord, word[i])
			}
		}
	}
	return CharOfWord
}

func Compare(CharOfWord, ListInput []rune) bool {
	SortRune(CharOfWord)
	SortRune(ListInput)
	if len(CharOfWord) != len(ListInput) {
		return false
	} else {
		for i := 0; i < len(CharOfWord); i++ {
			if CharOfWord[i] != ListInput[i] {
				return false
			}
		}
	}
	return true
}

func PrintWord(CharOfWord, ListInput []rune, word string) string {
	var wordToPrint string
	ListToPrint := []rune{}
	for i := 0; i < (len(CharOfWord)/2)-1; i++ {
		ListToPrint = append(ListToPrint, CharOfWord[rand.Intn(len(CharOfWord))])
	}
	for _, r := range ListToPrint {
		for _, char := range ListInput {
			if r != char {
				ListToPrint = append(ListToPrint, char)
			}
		}
	}
	for _, r := range word {
		for _, char := range ListToPrint {
			if r == char {
				wordToPrint += string(r)
			} else {
				wordToPrint += "_"
			}
		}
	}
	return wordToPrint
}

func FindWord(dificult string) string {
	var word string
	file, err := os.Open(dificult)
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return word
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
		}
	}
	return word
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

func TransStringToRune(s string) []rune {
	r := []rune(s)
	return r
}

func SortRune(Runes []rune) []rune {
	sort.Slice(Runes, func(i, j int) bool {
		return Runes[i] < Runes[j]
	})
	return Runes
}
