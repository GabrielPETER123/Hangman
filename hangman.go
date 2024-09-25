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
	fichier, err := os.Open("words.txt")
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return
	}
	defer fichier.Close()
	scan := bufio.NewScanner(fichier)
	for scan.Scan() {
		mot = scan.Text()
		find_mot = true
	}
	if find_mot == false {
		fmt.Print("No word found\n")
		return
	}
	fmt.Print("Bienvenue dans le jeu du pendu!\nÉcrivez une lettre pour essayer de deviner le mot!\n")
	for attempts > 0 && find_mot == true {
		input := ReadInput()
		if len(input) > 1 {
			fmt.Print("Too much input\n")
		}
		if len(input) < 1 {
			fmt.Print("Not enough input\n")
		} else {
			if VerifyWord(mot, ListInput) {
				fmt.Print("Vous avez déjà trouvé le mot!\n")
				return
			}
			int_input, ListInput = VerifyInput(input, mot, ListInput)
			if int_input == -1 {
				continue
			} else if int_input == 0 {
				bad_guesses++
				PrintHangman(bad_guesses)
				fmt.Printf("Nombre de tentatives restantes: %d\n", attempts-bad_guesses)
			}
			PrintMot(mot, ListInput)
			if strings.Contains(mot, strings.Join(ListInput, "")) {
				fmt.Print("Bien joué!\n")
				break
			} else if bad_guesses == attempts {
				fmt.Print("Pas de chance... Mauvaise réponse !\n")
				break
			}
		}
	}
	return
}

func PrintMot(mot string, ListInput []string) {
	nbr_letter_reveal := (len(mot) / 2) - 1
	for i := 0; i < nbr_letter_reveal; i++ {
		rand_letter := RandomNbr(len(mot))
		if strings.Contains(strings.Join(ListInput, ""), string(mot[rand_letter])) {
			continue
		} else {
			ListInput = append(ListInput, string(mot[rand_letter]))
		}
	}
	for _, letter := range mot {
		if strings.Contains(strings.Join(ListInput, ""), string(letter)) {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
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

func VerifyWord(mot string, ListInput []string) bool {
	var char_word []string
	for _, letter := range mot {
		if char_word == rune 
		char_word = append(char_word, string(letter))
	}

}