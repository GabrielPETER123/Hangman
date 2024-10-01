package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {
	ListInput := []rune{}
	ListToPrint := []rune{}
	bad_guesses := 0
	attempts := 10
	maj := 0
	inputVerify := false
	fmt.Print("Welcome to the Hangman game!\n")
	word := FindWord(Difficulty())
	fmt.Print(word, "\n")
	CharOfWord := CharOfWord(TransStringToRune(word))
	lenRand := len(CharOfWord) / 2
	for i := 0; i < lenRand; i++ {
		ListToPrint = append(ListToPrint, CharOfWord[rand.Intn(len(CharOfWord))])
	}
	PrintWord(CharOfWord, ListToPrint, word)
	fmt.Print("Try to find the word by writing a letter!\n")

	for attempts > 0 {
		fmt.Print(CharOfWord, "\n")
		fmt.Print(ListToPrint, "\n")
		if Compare(CharOfWord, ListToPrint) {
			fmt.Print("You won\n")
			fmt.Print("The word was: ", word, "\n")
			return
		}
		input := ReadInput()
		RInput := TransStringToRune(input)
		ListInput = append(ListInput, ListToPrint...)
		inputVerify, maj = VerifyInput(CharOfWord, RInput, ListInput)
		if inputVerify == false {
			bad_guesses++
			PrintHangman(bad_guesses)
			fmt.Printf("Bad Guess\n")
		} else {
			fmt.Print(ListToPrint, "\n")
			if maj == 1 {
				RInput[0] = RInput[0] + 32
			}
			ListInput = append(ListInput, RInput[0])
			ListToPrint = append(ListToPrint, RInput[0])
			fmt.Print("Good Guess\n")
		}
		PrintWord(CharOfWord, ListToPrint, word)
		fmt.Print("Attempts left: ", attempts, "\n")
		attempts--
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

func VerifyInput(CharOfWord, RInput, ListInput []rune) (bool, int) {
	if len(RInput) > 1 {
		fmt.Print("Too much input\n")
		VerifyInput(CharOfWord, TransStringToRune(ReadInput()), ListInput)
	}
	if len(RInput) < 1 {
		fmt.Print("Not enough input\n")
		VerifyInput(CharOfWord, TransStringToRune(ReadInput()), ListInput)
	}
	for _, char := range ListInput {
		if char == RInput[0] || char == RInput[0]-32 {
			fmt.Print("Already inputed\n")
			VerifyInput(CharOfWord, TransStringToRune(ReadInput()), ListInput)
		}
	}
	if !((RInput[0] >= 65 && RInput[0] <= 90) || (RInput[0] >= 97 && RInput[0] <= 122)) {
		fmt.Print("Invalid input\n")
		VerifyInput(CharOfWord, TransStringToRune(ReadInput()), ListInput)
	}
	for _, char := range CharOfWord {
		if char == RInput[0] {
			return true, 0
		}
	}
	for _, char := range CharOfWord {
		if char == RInput[0]+32 {
			return true, 1
		}
	}
	return false, 0
}

func ReadInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func CharOfWord(word []rune) []rune {
	CharOfWord := []rune{}
	CharOfWord = append(CharOfWord, word[0])
	for i := 1; i < len(word); i++ {
		for j := 0; j < len(CharOfWord); j++ {
			if word[i] == CharOfWord[j] {
				break
			}
			if j == len(CharOfWord)-1 {
				CharOfWord = append(CharOfWord, word[i])
			}
		}
	}
	fmt.Print(CharOfWord, "\n")
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

func PrintWord(CharOfWord, ListToPrint []rune, word string) {
	for _, r := range word {
		for index, char := range ListToPrint {
			if r == char || r == char-32 {
				fmt.Print(string(r-32), " ")
				break
			} else if index == len(ListToPrint)-1 {
				fmt.Print("_ ")
				break
			}
		}
	}
	fmt.Print("\n")
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
