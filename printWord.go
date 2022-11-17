package hangman

func PrintWord(word string, newWord string, listLetter []string, inputLetter string) (string, []string, string) {
	sliceWord := []rune(word)
	sliceNewWord := []rune(newWord)
	inputLetterRune := []rune(inputLetter)
	// ajoute à une liste la lettre utilisée
	if len(inputLetterRune) == 1 && inputLetterRune[0] >= 65 && inputLetterRune[0] <= 90 || len(inputLetterRune) == 1 && inputLetterRune[0] >= 97 && inputLetterRune[0] <= 122 {
		listLetter = append(listLetter, inputLetter)
	}
	// si le mot saisie est correct
	if word == inputLetter {
		return word, listLetter, inputLetter
	}
	// changement du mot après la vérification de IsInWord
	if len(inputLetter) == 1 {
		for _, char := range IsInWord(word, inputLetter) {
			sliceNewWord[char] = sliceWord[char]
		}
	}
	return string(sliceNewWord), listLetter, inputLetter
}
