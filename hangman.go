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
	var boolInput bool
	ListInput := []string{}
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

	//	attempts := 10

	//lire la sortie standard
	reader := bufio.NewReader(os.Stdin)
	inputLettre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	//	if attempts == 10 {
	//		func(len mot)
	//	}

	//message d'erreur
	if len(inputLettre) > 3 {
		fmt.Print("Too much input")
	}
	if len(inputLettre) < 3 {
		fmt.Print("Not enough input")
	} else {
		boolInput, ListInput = VerifyInput(inputLettre, mot, ListInput)
		if boolInput {
			fmt.Print("Bien joué!")
		} else {
			fmt.Print("Mauvaise réponse")
		}
	}
}

func PrintHangman() {
}

func Intn(n int) int {
	return rand.Intn(n)
}

func VerifyInput(s string, mot string, ListInput []string) (bool, []string) {
	if strings.Contains(strings.Join(ListInput, ""), s) {
		fmt.Println("Lettre déjà utilisée.")
		return false, ListInput
	}
	runeS := []rune(s)[0]
	if strings.ContainsRune(mot, runeS) {
		return true, append(ListInput, s)
	} else {
		return false, append(ListInput, s)
	}
}
