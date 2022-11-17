package hangman

import (
	"fmt"
	"os"
)

func AddWord(wordPlayer string) {
	wordPlayerToAdd := "\n" + wordPlayer
	file, _ := os.OpenFile("words.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 1)
	defer file.Close()
	file.WriteString(wordPlayerToAdd)
	fmt.Println("Your word was add in our file !")
}
