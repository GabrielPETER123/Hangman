package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var mot string
	var int_input int
	ListInput := []string{}
	find_mot := false
	bad_guesses := 0
	attempts := 10
	//	fichier, err := os.Open("words.txt")
	//	if err != nil {
	//		fmt.Printf("The error is: %v", err.Error())
	//		return
	//	}
	//	defer fichier.Close()
	//	scanner := bufio.NewScanner(fichier)
	// compter le nombre de lignes/mots avec le scanner et la mettre dans une variable à utiliser pour la variable aléatoire "nMot"
	// variable aléatoire "nMot" pour choisir un mot
	// variable pour la ligne (qui lit le mot)
	for attempts > 0 || find_mot != true {
		input := ReadInput()
		if len(input) > 1 {
			fmt.Print("Too much input")
		}
		if len(input) < 1 {
			fmt.Print("Not enough input")
		} else {
			attempts--
			int_input, ListInput = VerifyInput(input, mot, ListInput)
			if int_input == 1 {
				fmt.Print("Bien joué!\n")
			} else if int_input == 0 {
				fmt.Print("Mauvaise réponse\n")
				bad_guesses++
				PrintHangman(bad_guesses)
			}
		}
	}
	return
}

func PrintHangman(bad_guesses int) {
	hangman_file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return
	}
	defer hangman_file.Close()
	hangman_state := bufio.NewScanner(hangman_file)
	for hangman_state.Scan() {
	}
}

func RandomNbr(n int) int {
	return rand.Intn(n)
}

func VerifyInput(s string, mot string, ListInput []string) (int, []string) {
	if strings.Contains(strings.Join(ListInput, ""), s) {
		fmt.Print("Lettre déjà utilisée.\n")
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