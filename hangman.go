package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {

	//Initialisation des variables
	ListInput := []rune{}
	ListToPrint := []rune{}
	bad_guesses := 0
	attempts := 10
	maj := 0
	otherInput := 0
	inputVerify := false
	var LoadSaveInput string
	var saveInput string
	var word string
	var input string
	var lenRand int

	fmt.Print("Welcome to the Hangman game!\n")
	fmt.Print("Did you want to load the last save? (yes/no)\n")
	fmt.Scan(&LoadSaveInput)
	
	if LoadSaveInput == "yes" {
		word, attempts, bad_guesses, ListInput, ListToPrint = loadSave()
		otherInput = 100
	} else {

		//on cherche un mot selon la difficulté choisie
		fmt.Print("New game\n")
		word = FindWord(Difficulty())
	}
	
	//on cherche les lettres du mot
	CharOfWord := CharOfWord(convertStringToListOfRune(word))
	if otherInput == 0 { 
		lenRand = len(CharOfWord) / 2
	
		//Print le mot selon les lettres trouvées
		for i := 0; i < lenRand; i++ {
			ListToPrint = append(ListToPrint, CharOfWord[rand.Intn(len(CharOfWord))])
		}
	}
	PrintWord(CharOfWord, ListToPrint, word)

	fmt.Print("Try to find the word by writing a letter!\n")

	//Boucle du jeu
	for attempts > 0 {
		ListToPrint = compressListToPrint(ListToPrint)
		if Compare(CharOfWord, ListToPrint) {
			fmt.Print("You won\n")
			fmt.Print("The word was: ", word, "\n")
			return
		}
		fmt.Scan(&input)

		RInput := convertStringToListOfRune(input)
		if  LoadSaveInput != "yes" {
			ListInput = append(ListInput, ListToPrint...)
		}
		inputVerify, maj, otherInput = VerifyInput(CharOfWord, RInput, ListInput)
		
		//Sauvegarde de la partie
		if otherInput == 1 {
			fmt.Print("Do you want to save your progress? (yes/no)\n")
			fmt.Scan(&saveInput)
			if saveInput == "yes" {
				saveGame(word, fmt.Sprint(attempts), bad_guesses, ListInput, ListToPrint)
			} 
			if saveInput == "no" {
				fmt.Print("Progress not saved\n")
				otherInput = 0
				break
			} else {
				fmt.Print("Invalid input\n")
			}
		}
		if otherInput == 2 {
			fmt.Print("Do you want to stop the game? (yes/no)\n")
			fmt.Scan(&saveInput)
			if saveInput == "yes" {
				fmt.Print("Game stopped\n")
				break
			}
		} else {
			//Verification de l'input
			if inputVerify == false {
				bad_guesses++
				PrintHangman(bad_guesses)
				fmt.Printf("Bad Guess\n")
			} else {

				//met la lettre en majuscule si elle est en minuscule
				if maj == 1 {
					RInput[0] = RInput[0] + 32
				}
			}
		}
		if otherInput == 0 && inputVerify == true {

			ListInput = append(ListInput, RInput[0])
			ListToPrint = append(ListToPrint, RInput[0])
			fmt.Print("Good Guess\n")	
			attempts--
		}
		PrintWord(CharOfWord, ListToPrint, word)
		fmt.Print("Attempts left: ", attempts, "\n")

		otherInput = 0
		
		//Boucle d'arret de la partie
		if attempts == 0 {
			fmt.Print("You lost, the word was : ", word, "\n")
			break
		}
	}
	return
}

func PrintHangman(bad_guesses int) {
	
	//Variables pour encadrer le print du hangman
	start := 8 * (bad_guesses-1) + 1
	end := start + 7
	read_line := 1
	
	//Ouverture du fichier hangman.txt
	hangman_file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return
	}
	defer hangman_file.Close()
	hangman_scan := bufio.NewScanner(hangman_file)

	//Print le hangman selon le nombre d'erreurs commises
	for hangman_scan.Scan() {
		if lineread := hangman_scan.Text(); read_line >= start && read_line <= end {
			fmt.Println(lineread)
		}
		read_line++
	}
}

//Vérifie si l'input est valide
func VerifyInput(CharOfWord, RInput, ListInput []rune) (bool, int, int) {
	//On cherche un input de longueur seulement égale à 1
	var input string
		
	if len(RInput) > 1 {
		if string(RInput) == "save" {
			return true, 0, 1
		}
		if string(RInput) == "stop" { 
			return true, 0, 2
		}else {
			fmt.Print("Too much input\n")
		}
		fmt.Scan(&input)
		VerifyInput(CharOfWord, convertStringToListOfRune(input), ListInput)
	}
	if len(RInput) < 1 {
		fmt.Print("Not enough input\n")
		fmt.Scan(&input)
		VerifyInput(CharOfWord, convertStringToListOfRune(input), ListInput)
	}

	for _, char := range ListInput {
		if char == RInput[0] || char == RInput[0]-32 {
			fmt.Print("Already inputed\n")
			fmt.Scan(&input)
			VerifyInput(CharOfWord, convertStringToListOfRune(input), ListInput)
		}
	}

	//Erreur si le charactère n'est pas une lettre de l'alphabet
	if !((RInput[0] >= 65 && RInput[0] <= 90) || (RInput[0] >= 97 && RInput[0] <= 122)) {
		fmt.Print("Invalid input\n")
		fmt.Scan(&input)
		VerifyInput(CharOfWord, convertStringToListOfRune(input), ListInput)
	}

	//transforme en majuscule si nécessaire
	
	for _, char := range CharOfWord {
		if char == RInput[0] {
			return true, 0, 0
		}
	}

	for _, char := range CharOfWord {
		if char == RInput[0]+32 {
			return true, 1, 0
		}
	}

	return false, 0, 0
}

func CharOfWord(word []rune) []rune {
	CharOfWord := []rune{}
	CharOfWord = append(CharOfWord, word[0])
	
	//Parcourt le mot pour trouver les lettres différentes
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
	return CharOfWord
}

func Compare(CharOfWord, ListInput []rune) bool {
	//Trie les runes de chaque liste
	SortRune(CharOfWord)
	SortRune(ListInput)
	
	//Compare les deux listes
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
	//Print le mot selon les lettres trouvées

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
	
	//Ouverture du fichier de mots selon la difficulté choisie
	file, err := os.Open(dificult)
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return word
	}

	defer file.Close()
	scan := bufio.NewScanner(file)
	
	//scan les mots du fichier et en choisi un aléatoirement
	for scan.Scan() {
		words := []string{}
		
		//Met tous les mots du fichier dans une liste
		for scan.Scan() {
			words = append(words, scan.Text())
		}

		//Choisi un mot aléatoire
		if len(words) > 0 {
			word = words[rand.Intn(len(words))]
		}
	}
	return word
}

func Difficulty() string {

	fmt.Println("Choose your level: ")
	var level string
	fmt.Scan(&level)
	var dificult string

	//choix des diffirenetes difficultés
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

func SortRune(Runes []rune) []rune {
	//Trie les runes via l'import "sort"
	sort.Slice(Runes, func(i, j int) bool {
		return Runes[i] < Runes[j]
	})
	return Runes
}

func convertStringToInt(s string) int {
	var n int
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n
}

func convertStringToListOfRune(s string) []rune {
	var list []rune
	for _, c := range s {
		list = append(list, c)
	}
	return list
}

func loadSave() (string, int, int, []rune, []rune) {

	//Initialisation des variables
	var word string
	var attempts string
	var bad_guesses int
	var ListInput []rune
	var ListToPrint []rune

	//Ouverture du fichier save.txt
	file, err := os.Open("save.txt")
		if err != nil {
			fmt.Println("Error opening save file")
		}
		defer file.Close()
		saveScanner := bufio.NewScanner(file)
		
		//Lecture du fichier save.txt
		for saveScanner.Scan() {
			word = saveScanner.Text()
			saveScanner.Scan()
			attempts = saveScanner.Text()
			saveScanner.Scan()
			bad_guesses = convertStringToInt(saveScanner.Text())
			saveScanner.Scan()
			ListInput = convertStringToListOfRune(saveScanner.Text())
			saveScanner.Scan()
			ListToPrint = convertStringToListOfRune(saveScanner.Text())
		}
	fmt.Println("Save loaded")
	return word, convertStringToInt(attempts), bad_guesses, ListInput, ListToPrint
}

func saveGame(word string, attempts string, bad_guesses int, ListInput []rune, ListToPrint []rune) {
	
	//Suppression du fichier save.txt
	err := os.Remove("save.txt")
		if err != nil {
			fmt.Println("Error deleting save file")
		}

	//Création du fichier save.txt
	file, err := os.Create("save.txt")
		if err != nil {
			fmt.Println("Error creating save file")
			return
		}
		defer file.Close()
	
	//Écriture des variables dans le fichier save.txt
	file.WriteString(word + "\n")
	file.WriteString(attempts + "\n")
	file.WriteString(fmt.Sprint(bad_guesses) + "\n")
	file.WriteString(string(ListInput) + "\n")
	file.WriteString(string(ListToPrint) + "\n")
	fmt.Println("Game saved")
}

//cette fonction sert seulement pour la fonction load car la liste de rune dans la save se répète (je ne sais pas pourquoi)
func compressListToPrint(ListToPrint []rune) []rune {
	newList := []rune{}

	//Compresser la liste de rune, c'est à dire enlever les doublons
	for i := 0; i < len(ListToPrint); i++ {
		for j := 0; j < len(newList); j++ {
			if ListToPrint[i] == newList[j] {
				break
			}
			if j == len(newList)-1 {
				newList = append(newList, ListToPrint[i])
			}
		}
		if len(newList) == 0 {
			newList = append(newList, ListToPrint[i])
		}
	}
	return newList
}