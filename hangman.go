package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	var mot string
	if len(arg) > 1 {
		fmt.Print("Too much input")
		return
	}
	if len(arg) < 1 {
		fmt.Print("Not enough input")
		return
	}
	if VerifyInput(arg[0], mot) >= -1 && VerifyInput(arg[0], mot) <= 1 {
		return
	} else {
		fmt.Print("error")
		return
	}
}

func PrintHangman() {
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
