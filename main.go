package main

import (
	"fmt"
	"strings"
	"bufio"
	"math/rand"
	"os"
)

func main() {
	var badGuesses int
	var word string
	var hiddenWord string
	var guess string
	var guesses []string
	var hiddenRunes []rune
	var wordRunes []rune
	var wrongMessage string

	//** Clear the console */
	fmt.Print("\033[H\033[2J")

	//** Difficulty Choice */
	fmt.Println("Choose a difficulty level:")
	fmt.Println("1. Easy")
	fmt.Println("2. Medium")
	fmt.Println("3. Hard")
	word = chooseDifficulty()
	wordRunes = []rune(word)

	//** Clear the console */
	fmt.Print("\033[H\033[2J")

	//** Hide the word */
	hiddenWord, hiddenRunes = hideWord(word)

	//** The runes not hidden in the word */
	wordRunes = constructWordRunes(wordRunes, hiddenRunes)

	//** Attemp loop */
	for i := 0; i < 10; i++ {
		//** Clear the console */
		fmt.Print("\033[H\033[2J")
		
		//** Print informations in the console */
		fmt.Println("The hidden word is:", hiddenWord)
		fmt.Println("You have", 10-i, "attempts left.")
		printHangman(badGuesses)
		if len(wrongMessage) > 0 {		
			fmt.Println(wrongMessage)
		}
		fmt.Println("--------------------------------------------------------")
		fmt.Println("Try to guess a letter:")

		fmt.Scan(&guess)
		
		//** Verify the input is a single letter */
		if !verifySingleLetter(guess) {
			wrongMessage = "Invalid input. Please enter a single letter."
			i--
			continue
		}

		guess = strings.ToLower(guess)

		//** Verify the guess has not already be tried */
		if verifyAlreadyTried(guess, guesses) {
			wrongMessage = "You already tried the letter " + guess + ". Please try a different letter."
			i--
			continue
		}
		guesses = append(guesses, guess)

		//** Verify if the guess is not a revealed rune */ 
		if verifyRevealedRune(guess, wordRunes) {
			wrongMessage = "Your guess " + guess + ", is already revealed in the word."
			i--
			continue
		}

		//** Check if the guess is in the hidden runes */
		var found bool
		found, hiddenRunes, hiddenWord = verifyHiddenRunes(guess, hiddenRunes, hiddenWord, word)
		if !found {
			badGuesses++
			if badGuesses == 10 {
				fmt.Println("You lost! The word was:", word)
				return
			}
			wrongMessage = "You guessed the letter " + guess + ", but it is not in the word."

		}
		if len(hiddenRunes) == 0 {
			fmt.Println("Congratulations! You guessed the word:", word)
			return
		}
	}
}

//** Print the hangman based on the number of bad guesses */
func printHangman(badGuesses int) {
	if badGuesses == 0 {
		return
	}

	start := 8 * (badGuesses - 1)
	end := start + 7

	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Error reading hangman.txt:", err)
		return
	}
	defer file.Close()

	hangman := bufio.NewScanner(file)

	var lines []string
	for hangman.Scan() {
		lines = append(lines, hangman.Text())
	}
	for i := start; i <= end && i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

//** Choose the difficulty level and return the word */
func chooseDifficulty() string {
	var word string
	var difficulty string
	fmt.Println("Enter 1, 2, or 3: ")
	fmt.Scan(&difficulty)

	for difficulty != "1" && difficulty != "2" && difficulty != "3" {
		fmt.Println("Invalid choice. Please choose a valid difficulty level.")
		fmt.Println("Enter 1, 2, or 3: ")
		fmt.Scan(&difficulty)
	}
	switch difficulty {
	case "2":
		fmt.Println("You chose Medium difficulty.")
		word = getWord("words2.txt")
		break
	case "3":
		fmt.Println("You chose Hard difficulty.")
		word = getWord("words3.txt")
		break
	default:
		fmt.Println("You chose Easy difficulty.")
		word = getWord("words.txt")
		break
	}

	return word
}

func getWord(file string) string {
	wordFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error reading word file:", err)
		return ""
	}
	defer wordFile.Close()

	scanner := bufio.NewScanner(wordFile)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	randomIndex := rand.Intn(len(words))
	var word string
	word = words[randomIndex]

	fmt.Println("The word is:", word)
	return word
}

//** Create a hidden word from the original word and the hidden runes from it */
func hideWord(word string) (string, []rune) {
	var hiddenWord string
	var runeList = getRunes(word)
	var hiddenRunes []rune
	numberOfHiddenLetters := len(runeList) / 2

	//** Create a list of hidden runes */
	for i := 0; i < numberOfHiddenLetters; i++ {
		randomIndex := rand.Intn(len(runeList) - 1)
		if i == 0 {
			hiddenRunes = append(hiddenRunes, runeList[randomIndex])
		} else {
			//** Check if the rune is already in the hidden runes */
			for i := 0; i < len(hiddenRunes); i++ {
				if runeList[randomIndex] == hiddenRunes[i] {
					randomIndex = rand.Intn(len(runeList) - 1)
					i = -1
				}
			}
			hiddenRunes = append(hiddenRunes, runeList[randomIndex])
		}
	}
	hiddenWord = createHiddenWord(word, hiddenRunes)

	return hiddenWord, hiddenRunes
}

func createHiddenWord(word string, hiddenRunes []rune) string {
	wordRunes := []rune(word)

	//** Replace hidden runes with underscores */
	for i, r := range wordRunes {
		for _, hiddenRune := range hiddenRunes {
			if r == hiddenRune {
				wordRunes[i] = '_'
			}
		}
	}

	return string(wordRunes)
}

//** Get the runes from the word and remove duplicates */
func getRunes(word string) []rune {
	var runeList []rune
	var wordRunes = []rune(word)

	//** Make sure the hidden runes are not repeated */
	uniqueRunes := make(map[rune]bool)
	var dedupedRunes []rune
	for _, r := range wordRunes {
		if !uniqueRunes[r] {
			uniqueRunes[r] = true
			dedupedRunes = append(dedupedRunes, r)
		}
	}
	//** Transfer map to array */
	for _, r := range dedupedRunes {
		runeList = append(runeList, r)
	}
	
	return runeList
}

//** Construct the runes list that are not hidden in the word */
func constructWordRunes(wordRunes []rune, hiddenRunes []rune) []rune {
	var runes []rune
	for _, r := range wordRunes {
		found := false
		for _, hiddenRune := range hiddenRunes {
			if r == hiddenRune {
				found = true
				break
			}
		}
		if !found {
			runes = append(runes, r)
		}
	}

	return runes
}

//** Remove a rune from the hidden runes list */
func removeRune(hiddenRunes []rune, r rune) []rune {
	for i, hiddenRune := range hiddenRunes {
		if hiddenRune == r {
			return append(hiddenRunes[:i], hiddenRunes[i+1:]...)
		}
	}
	return hiddenRunes
}

//** Verify functions for the guess */
func verifySingleLetter(guess string) bool {
	if len(guess) != 1 {
		return false
	}
	if (guess < "a" || guess > "z") && (guess < "A" || guess > "Z") {
		return false
	}
	return true
}

func verifyAlreadyTried(guess string, guesses []string) bool {
	for _, r := range guesses {
		if guess[0] == r[0] {
			return true
		}
	}
	return false
}

func verifyRevealedRune(guess string, wordRunes []rune) bool {
	for _, r := range wordRunes {
		if guess[0] == byte(r) {
			return true
		}
	}
	return false
}

func verifyHiddenRunes(guess string, hiddenRunes []rune, hiddenWord string, word string) (bool, []rune, string) {
	for _, r := range hiddenRunes {
		if guess[0] == byte(r) {
			hiddenRunes = removeRune(hiddenRunes, rune(guess[0]))
			hiddenWord = createHiddenWord(word, hiddenRunes)
			fmt.Println("Good guess! The letter", guess, "is in the word.")
			return true, hiddenRunes, hiddenWord
		}
	}
	return false, hiddenRunes, hiddenWord
}
