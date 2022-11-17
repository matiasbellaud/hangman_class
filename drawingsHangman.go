package hangman

import (
	"io/ioutil"
	"strings"
)

func DrawingsHangman(index int) string {
	data, _ := ioutil.ReadFile("hangman.txt")
	strDrawings := string(data)
	// création d'une liste contenant chaque positions du pendu
	drawings := strings.Split(strDrawings, ",")
	return drawings[index]
}
