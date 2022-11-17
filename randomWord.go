package hangman

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func RandomWord(difficulty string, language string) string {
	rand.Seed(time.Now().UnixNano())
	var data []byte
	if language == "eng" {
		data, _ = ioutil.ReadFile("words_eng.txt")
	} else if language == "fr" {
		data, _ = ioutil.ReadFile("words.txt")
	}
	words := string(data)
	var word []rune
	var wordString string
	// création d'une liste contenant tous les mots
	splitWords := strings.Split(words, "\n")
	length := len(splitWords)
	// choix du mot grâce à l'aléatoire
	randomInt := rand.Intn(length)
	wordRune := []rune(splitWords[randomInt])
	// changement de mot selon la difficulté
	wordIsGood := false
	if difficulty == "0" || difficulty == "1" && len(wordRune) <= 4 || difficulty == "2" && len(wordRune) >= 5 && len(wordRune) <= 7 || difficulty == "3" && len(wordRune) >= 8 {
		wordIsGood = true
	} else {
		return RandomWord(difficulty, language)
	}
	if wordIsGood {
		// règle le problème du caractère inconnu qui apparaîssait à la fin du mot
		for i := 0; i < len(wordRune); i++ {
			if wordRune[i] >= 97 && wordRune[i] <= 122 {
				word = append(word, wordRune[i])
			}
		}
		for _, letter := range word {
			wordString += string(letter)
		}
	}
	return wordString
}
