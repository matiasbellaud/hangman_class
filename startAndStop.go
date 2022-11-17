package hangman

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

func Start(arg []string, attemps int, word string, toFind string) (int, string, string) {
	// fonction qui vient regarder si l'on veut reprendre la partie d'avant
	if len(arg) == 2 {
		if arg[1] == "save.txt" {
			//décode des infos dans un fichier txt
			attemps, word, toFind = UnMarshall()
			return attemps, word, toFind
		}
	}
	return attemps, word, toFind
}

func Stop(attempsInt int, word string, toFind string) {
	// si on veut arrêter la partie
	//encode des informations dans un fichier txt
	attemps := strconv.Itoa(attempsInt)
	Marshall(attemps, word, toFind)
}

type Args struct {
	Attemps string `json:"attemps"`
	Word    string `json:"word"`
	ToFind  string `json:"ToFind"`
}

func Marshall(attemps string, word string, toFind string) {
	// on remet à zéro le fichier save.txt
	os.Remove("save.txt")
	os.Create("save.txt")
	args := &Args{
		Attemps: attemps,
		Word:    word,
		ToFind:  toFind,
	}
	// encode les informations
	data, _ := json.Marshal(args)
	file, _ := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY, 1)
	defer file.Close()
	// ajoute les informations encodées dans save.txt
	file.WriteString(string(data))
}

func UnMarshall() (int, string, string) {
	var args Args
	// on vient lire le fichier save.txt
	data, _ := ioutil.ReadFile("save.txt")
	// on décode les informations
	json.Unmarshal(data, &args)
	attempsInt, _ := strconv.Atoi(args.Attemps)
	return attempsInt, args.Word, args.ToFind
}
