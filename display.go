package hangman

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Display() {
	// déclaration des variables
	arg := os.Args
	var toFind string
	if len(arg) == 1 {
		language := Language()
		difficulty := Difficulty()
		toFind = RandomWord(difficulty, language)
	}
	word := RevealRandomLetter(toFind)
	attemps := 10
	var input string
	var listLetter []string
	var newWordUpdated string
	// si le joueur veut reprendre la partie
	attemps, word, toFind = Start(arg, attemps, word, toFind)
	for word != toFind {
		fmt.Println(PrintAscii(word))
		// fait apparaître la liste des lettres deja utilisées
		fmt.Println("The list of letters you have already used :", listLetter)
		fmt.Println()
		// stocke le nouveau mot une fois modifier par l'input
		newWordUpdated, listLetter, input = PrintWord(toFind, word, listLetter, input)
		// si le joueur veut arrêter la partie
		if input == "stop" {
			Stop(attemps, word, toFind)
			ioutil.ReadFile("save.txt")
			break
		}
		// si le nouveau mot est comme le précedent alors
		if word == newWordUpdated {
			attemps -= 1
			fmt.Println()
			fmt.Printf("Not present in the word, %d attemps remaining", attemps)
			fmt.Println()
			// ajoute le dessin du pendu correspondant au nombres d'essai restant
			fmt.Println(DrawingsHangman(9 - attemps))
		} else {
			word = newWordUpdated
			fmt.Println()
		}
		// si on perd
		if attemps == 0 {
			fmt.Println("You loose, the word was :", Color("red", toFind))
			fmt.Println("Try again !")
			os.Remove("save.txt")
			break
		}
		// si on gagne
		if word == toFind {
			fmt.Println("You are right, the word was :", Color("green", toFind))
			fmt.Println("Congratulations !")
			// si le joueur veut ajouter un mot dans les fichiers
			isPlayerAddWord := PlayerAddWord()
			if isPlayerAddWord == "o" {
				fmt.Print("Choose a word to add : ")
				var wordPlayer string
				fmt.Scanln(&wordPlayer)
				if wordPlayer == "stop" || wordPlayer == "" {
					fmt.Println("Successfully stopped")
					break
				} else {
					AddWord(wordPlayer)
				}
			}
			os.Remove("save.txt")
		}
	}
}

func Color(color string, toFind string) string {
	word := PrintAscii(toFind)
	var returnWord string
	// création des variables contenant les couleurs
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Reset = "\033[0m"
	// mise en couleur du mot
	if color == "green" {
		returnWord = (Green + word + Reset)
	} else {
		returnWord = (Red + word + Reset)
	}
	return returnWord
}
