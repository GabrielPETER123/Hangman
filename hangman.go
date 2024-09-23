package main

import (
	"fmt"
	"os"
)

func main() {
	var attempt int
	var wrong int
	var mot []string
	arg := os.Args[1:]
	//stocke la var attempt dans le fichier output pour la ré-import pour la ré-utiliser et reset la var attempt
	fmt.Print(attempt, "\n")
	//définir le mot via un flag (pas encore défini) pour set le mot à trouver
	//définir un flag pour reset (pas encore défini) pour reset le mot à trouver et reset les attempts donc le fichier output
	if len(arg) > 1 {
		fmt.Print("Too much input")
		return
	}
	if len(arg) < 1 {
		fmt.Print("Not enough input")
		return
	}
	if VerifyInput(arg[0], mot) >= -1 && VerifyInput(arg[0]) <= 1 {
		return
	} else {
		fmt.Print("error")
		return
	}
}

func PrintHangman() {
	//pas de forme encore défini
}

func VerifyInput(s string, mot []string) int {
	var ListInput []string
	if len(ListInput) == 0 {
		ListInput = append(ListInput, s)
	}
	if len(ListInput) > 0 {
		for _, r := range ListInput {
			if s == r {
				fmt.Print("Input already used.")
				return -1
			} else {
				for _, r := range Mot {
					ListInput = append(ListInput, s)
					if s == r {
						fmt.Print("Good job!")
						return 1
					} else {
						fmt.Print("Unlucky, try again!")
						return 0
					}
				}
			}
		}
	}
	return 2
}
