package hangman

import "fmt"

func InputPlayer(listLetter []string, word string) string {
	var verification bool
	letter := ""
	fmt.Print("Choose: ")
	fmt.Scanln(&letter) // prend une string en input
	verification, letter = VerifInput(letter, word)
	// si verification est faux alors un message d'erreur apparaît
	if !verification {
		fmt.Println("Your character is not available.")
		return InputPlayer(listLetter, word)
	}
	for _, letterUsed := range listLetter {
		if letterUsed == letter {
			fmt.Println("Your character was already used.")
			return InputPlayer(listLetter, word)
		}
	}
	return (letter)
}
func VerifInput(letter, word string) (bool, string) {
	// Si le mot saisie est bon, renvoit true et le mot
	if letter == word {
		return true, letter
	}
	if letter == "stop" {
		return true, letter
	}
	letterRune := []rune(letter)
	// verifie la taille de l'input
	if len(letterRune) != 1 {
		return false, letter
	}
	// verifie si c'est pas une lettre
	if letterRune[0] < 65 || letterRune[0] > 90 && letterRune[0] < 97 || letterRune[0] > 123 {
		return false, string(letterRune)
	}
	// passe les minuscule en majuscule
	if letterRune[0] >= 65 && letterRune[0] <= 90 || letter[0] >= 97 && letter[0] <= 122 {
		if letterRune[0] >= 65 && letterRune[0] <= 90 {
			letterRune[0] += 32
		}
		return true, string(letterRune)
	}
	return true, string(letterRune)
}
