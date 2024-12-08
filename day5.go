package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Textdatei einlesen
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	updateMap := make(map[string][]string)
	var result int

	//Den Text auf zwei String Arrays aufteilen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			twoStrings := strings.Split(line, "|")
			//Aufteilen der Regeln in eine UpdateMap, hierbei, wird die zweite Zahl der Key, die erste der Value, dann kann man hinterher besser prüfen
			//Wenn der Key schon existiert, wird der Value hier angehängt
			if _, exists := updateMap[twoStrings[1]]; exists {
				updateMap[twoStrings[1]] = append(updateMap[twoStrings[1]], twoStrings[0])
			} else {
				// Andernfalls wird ein neues Array mit dem Value erstellt
				updateMap[twoStrings[1]] = []string{twoStrings[0]}
			}
		} else {
			if line != "" {
				regular := true
				proofArray := strings.Split(line, ",")
				middleIndex := len(proofArray) / 2
				middleString := proofArray[middleIndex]
				//proofArray ist die Zeile
				for _, proofs := range proofArray {
					update := updateMap[string(proofs)]
					proofArray = proofArray[1:]
					//prüft die Updateregeln, sobald eine gebrochen wird, springt man aus beiden Schleifen raus
					// man braucht nun nur prüfen, ob im Array hinter der ersten Zahl Werte stehen, die in den Updateregeln sind, da wir umgedreht haben,
					// müssten diese davor stehen, sind also falsch
				outerLoop:
					for _, i := range update {
						for j := 0; j < len(proofArray); j++ {
							if strings.Contains(proofArray[j], i) {
								regular = false
								break outerLoop
							}
						}
					}
				}
				middle, err := strconv.Atoi(middleString)
				if err != nil {
					// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
					fmt.Println(err)
					return
				}
				if regular {
					result = result + middle
				}
			}
		}
	}
	fmt.Println("Result ", result)
}
