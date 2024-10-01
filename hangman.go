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

func ReadInput() string { // Lit la réponse de l'utilisateur
	input := bufio.NewScanner(os.Stdin) // Crée un scanner pour lire la réponse de l'utilisateur
	input.Scan() // Lit la réponse de l'utilisateur
	return input.Text() // Retourne la réponse de l'utilisateur
}

func CharOfWord(word []rune) []rune { // Permet de savoir toutes les lettres différentes du mot à trouver	
	CharOfWord := []rune{}
	CharOfWord = append(CharOfWord, word[0]) // Ajoute la première lettre du mot à la liste des lettres du mot
	for i := 1; i < len(word); i++ { // Parcours le mot
		for j := 0; j < len(CharOfWord); j++ { // Parcours la liste des lettres du mot
			if word[i] == CharOfWord[j] { // Si la lettre du mot est déjà dans la liste des lettres du mot alors on passe à la lettre suivante
				break
			}
			if j == len(CharOfWord)-1 { // sinon si on a parcouru toute la liste des lettres du mot
				CharOfWord = append(CharOfWord, word[i]) // On ajoute la lettre du mot à la liste des lettres du mot
			}
		}
	}
	fmt.Print(CharOfWord, "\n")
	return CharOfWord // Retourne la liste des lettres du mot
}

func Compare(CharOfWord, ListInput []rune) bool { // Permet de comparer la liste des lettres du mot avec les lettre trouvées
	SortRune(CharOfWord) // Trie les runes du mot
	SortRune(ListInput) // Trie les runes des lettres trouvées
	if len(CharOfWord) != len(ListInput) { // Si la longeur des deux listes est différente alors le mot n'est pas trouvé
		return false
	} else {
		for i := 0; i < len(CharOfWord); i++ { // Parcours les deux listes
			if CharOfWord[i] != ListInput[i] { // Si une rune est différente alors le mot n'est pas trouvé
				return false
			}
		}
	}
	return true
}

func PrintWord(CharOfWord, ListToPrint []rune, word string) { //Affiche le mot à trouver
	for _, r := range word { // Parcours le mot à trouver
		for index, char := range ListToPrint { // Parcours la liste des lettres à afficher
			if r == char || r == char-32 { // Si la lettre à afficher est la même que la lettre du mot
				fmt.Print(string(r-32), " ")
				break // Si la lettre à afficher est la même que la lettre du mot alors on passe à la lettre suivante du mot
			} else if index == len(ListToPrint)-1 { // sinon si on aucune lettre n'a été trouvée dans la liste des lettres à afficher
				fmt.Print("_ ") // On affiche un underscore
				break // On passe à la lettre suivante du mot
			}
		}
	}
	fmt.Print("\n")
}

func FindWord(dificult string) string { //Trouve un mot aléatoire dans le fichier de mots
	var word string
	file, err := os.Open(dificult) // Ouvre le fichier de mots selon la difficulté choisie
	if err != nil { // Affiche une erreur si le fichier n'est pas trouvé
		fmt.Printf("The error is: %v", err.Error())
		return word
	}
	defer file.Close() // Ferme le fichier à la fin de la fonction
	scan := bufio.NewScanner(file) // Crée un scanner pour lire le fichier
	for scan.Scan() { // Parcours le fichier de mots
		words := []string{}
		for scan.Scan() { // Parcours le fichier de mots
			words = append(words, scan.Text()) // Ajoute les mots du fichier dans une liste de mots
		}
		if len(words) > 0 { // Si la liste de mots n'est pas vide
			word = words[rand.Intn(len(words))] // Prend un mot aléatoire dans la liste de mots
		}
	}
	return word // Retourne le mot aléatoire
}

func Difficulty() string { //Permet de choisir la difficulté
	fmt.Println("Choose your level: ") //Demande à l'utilisateur de choisir un niveau
	var level string
	fmt.Scanln(&level) //récupère le choix de niveau de l'utilisateur
	var dificult string
	switch level { //Permet de choisir le fichier de mots en fonction du niveau choisi
	case "1":
		dificult = "words.txt"
	case "2":
		dificult = "words2.txt"
	case "3":
		dificult = "words3.txt"
	default:
		fmt.Println("Invalid input")
	}
	return dificult //Retourne le fichier de mots correspondant au niveau choisi
}

func TransStringToRune(s string) []rune { //Transforme un string en liste de runes
	r := []rune(s)
	return r
}

func SortRune(Runes []rune) []rune { //Trie les runes dans l'ordre croissant
	sort.Slice(Runes, func(i, j int) bool {	//Transforme les runes en int pour les comparer
		return Runes[i] < Runes[j] 	//Retourne les runes triées
	})
	return Runes //Retourne la liste des runes triées
}
