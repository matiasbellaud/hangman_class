package hangman

import (
	"io/ioutil"
	"strings"
)

func ListAscii(index []int) string {
	var splitLines []string
	var sliceAscii []string
	var resultAscii string
	// on ouvre le fichier des caractères ascii et on les stocke dans une variable string
	data, _ := ioutil.ReadFile("standard.txt")
	letters := string(data)
	// on viens créer une liste avec tous les caractères
	splitLetters := strings.Split(letters, ",,")
	// A = 33
	// Z = 58
	// _ = 63
	// on crée une liste de taille 8
	for i := 0; i < 8; i++ {
		sliceAscii = append(sliceAscii, "")
	}
	// on ajoute chaque première ligne de chaque lettre à une case de la liste créée avant
	for _, char := range index {
		splitLines = strings.Split(splitLetters[char], "\n")
		for i := 0; i < 8; i++ {
			sliceAscii[i] += splitLines[i]
		}
	}
	// créer une variable string qui contiendra la liste
	for i := 0; i < 8; i++ {
		resultAscii += sliceAscii[i] + "\n"
	}
	return resultAscii
}
