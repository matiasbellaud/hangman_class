package hangman

func IsInWord(word string, inputLetter string) []int {
	var res []int
	sliceInput := []rune(inputLetter)
	letterInput := sliceInput[0]
	// on parcours le mot en verifiant si notre lettre est dedans et on resort une liste des index correspondant
	for index, char := range word {
		if letterInput == char {
			res = append(res, index)
		}
	}
	return res
}
