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
	//faire choisir le level de difficulté au joueur
	fmt.Println("Choose your level: ")
	// var then variable level then variable type
	var level string
	// Taking input from user
	fmt.Scanln(&level)
	var dificult string
	if level == 1 {
		dificult = "words.txt"
	} else if level == 2 {
		dificult = "words2.txt"
	} else {
		dificult = "words3.txt"
	}
	fichier, err := os.Open(dificult)
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
			int_input, ListInput = VerifyInput(input, mot, ListInput)
			if int_input == -1 {
				continue
			} else if int_input == 0 {
				bad_guesses++
				PrintHangman(bad_guesses)
				fmt.Printf("Nombre de tentatives restantes: %d\n", attempts-bad_guesses)
			}
			PrintMot(mot, ListInput, attempts)
			if strings.Compare(mot, strings.Join(ListInput, "")) == 0 {
				fmt.Print("Bravo, vous avez trouvé le mot!\n")
				break
			}
			if attempts-bad_guesses == 0 {
				fmt.Print("Vous avez perdu! Le mot était: ", mot, "\n")
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

func PrintWord(CharOfWord []string, ListInput string, attempt int) {
	charprint := []string{}
	if attempt == 10 {
		for i := 0; i < len(CharOfWord)/2-1; i++ {
			charprint = append(charprint, CharOfWord[RandomNbr(len(CharOfWord))])
		}
	}
	for strings.Compare(strings.Join(charprint, ""), ListInput) == 0 {

	}
	for _, c := range CharOfWord {
		if strings.Contains(strings.Join(charprint, ""), c) {
			fmt.Print(c)
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("\n")
}
