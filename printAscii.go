package hangman

func PrintAscii(word string) string {
	var index []int
	// parcours le mot pour renvoyer l'équivalent en code ascii
	for _, char := range word {
		if char == '_' {
			index = append(index, int('_')-32) // -32 pour passer de la rune '_' à notre '_' ascii
		} else {
			index = append(index, int(char)-64) // -64 pour passer de la rune à notre ascii majuscule
		}
	}
	return ListAscii(index)
}
