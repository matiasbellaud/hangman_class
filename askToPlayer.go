package hangman

import "fmt"

func Difficulty() string {
	fmt.Print("Choose your difficulty between 1, 2 or 3 : ")
	var difficultyPlayer string
	fmt.Scanln(&difficultyPlayer)
	if difficultyPlayer == "1" || difficultyPlayer == "2" || difficultyPlayer == "3" {
		return difficultyPlayer
	}
	fmt.Println("Your response is not good")
	return Difficulty()
}

func Language() string {
	fmt.Print("Choose your language between French (fr) or English (eng) : ")
	var language string
	fmt.Scanln(&language)
	if language == "fr" || language == "eng" {
		return language
	}
	fmt.Println("Your response is not good")
	return Language()
}

func PlayerAddWord() string {
	var isPlayerAddWord string
	fmt.Print("Do you want to add a word in our game ? o for yes, enter for no  : ")
	fmt.Scanln(&isPlayerAddWord)
	return isPlayerAddWord
}
