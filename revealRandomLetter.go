package hangman

import (
	"math/rand"
	"time"
)

func RevealRandomLetter(word string) string {
	rand.Seed(time.Now().UnixNano()) // crée une graine qui permettra de réinitialiser l'aléatoire à chaque fois
	var listIndex []int
	resultWord := ""
	nbrLetter := len(word)/2 - 1
	var randomNbr int
	isTheFirstNbr := true
	// utilisation d'un tableau qui viens garder chaque nombre aleatoire
	for len(listIndex) <= nbrLetter-1 {
		randomNbr = rand.Intn(len(word))
		for _, number := range listIndex { // si il y a déjà le nombre dans le tableau il ne l'ajoutera pas
			if randomNbr == number {
				isTheFirstNbr = false
			}
		}
		if isTheFirstNbr {
			listIndex = append(listIndex, randomNbr)
		}
		isTheFirstNbr = true
	}
	isIndex := false
	for index, letter := range word {
		for _, number := range listIndex {
			// si l'index se trouve dans le tableau de nombre alors il ajoute la lettre au mot
			if index == number {
				isIndex = true
				resultWord += string(letter)
			}
		} // si l'index ne se trouve pas dans le tableau de nombre alors il ajoute un "_" au mot
		if !isIndex {
			resultWord += "_"
		}
		isIndex = false
	}
	return resultWord
}
