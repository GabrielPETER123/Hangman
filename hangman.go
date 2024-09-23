package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var mot string
	fichier, err := os.Open("words.txt")
	if err != nil {
		fmt.Printf("The error is: %v", err.Error())
		return
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	// variable aléatoire pour choisir un mot
	// variable pour la ligne (qui lit le mot)

	attempts := 10

	//lire la sortie standard
	standardOutput := bufio.NewScanner(os.Stdin)
	for standardOutput.Scan() {
		line := standardOutput.Text()
	}
	if err := standardOutput.Err(); err != nil {
		fmt.Println("Error reading from standard output:", err)
	}

	if attempts == 10 {
		func(len mot)
	}

	//message d'erreur
	if len(line) > 1 {
		fmt.Print("Too much input")
	}
	if len(line) < 1 {
		fmt.Print("Not enough input")
	}
	if VerifyInput(line, mot) >= -1 && VerifyInput(line, mot) <= 1 {
	} else {
		fmt.Print("error")
	}
}

func PrintHangman() {
}

func Intn(n int) int {
	return rand.Intn(n)
}

func VerifyInput(s string, mot string) int {
	var ListInput []string
	if len(ListInput) == 0 {
		ListInput = append(ListInput, s)
	}
	if len(ListInput) > 0 {
		for _, r := range ListInput {
			if s == r {
				fmt.Print("lettre déjà utilisée.")
				return -1
			} else {
				for _, r := range mot {
					ListInput = append(ListInput, s)
					if s == r {
						fmt.Print("Bien joué!")
						return 1
					} else {
						fmt.Print("Mauvaise lettre.")
						return 0
					}
				}
			}
		}
	}
	return 2
}
