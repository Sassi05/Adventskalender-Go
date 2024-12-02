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
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	var levels []string
	var levelsNum []int
	var count int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//Heute mal Textzeile mit Fields splitten
		levels = strings.Fields(line)

		//Die Arrayinhalte von String zu Int parsen und in Int Array einlesen
		for i := 0; i < len(levels); i++ {
			num, err := strconv.Atoi(levels[i])

			if err != nil {
				// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
				fmt.Println(err)
				return
			}

			levelsNum = append(levelsNum, num)

		}
		//Einzelne Ebenen auf "Sicherheit" prüfen
		result := proofLevels(levelsNum)
		count = count + result
		//Array wieder leeren
		levelsNum = []int{}
	}

	// Resultat ausgeben
	fmt.Println("Endresultat :")
	fmt.Println(count)

}

// Funktion, um die Ebenen auf "Sicherheit" zu überprüfen
func proofLevels(array []int) int {
	var safe bool = false
	//bei absteigender Reihenfolge
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			safe = true
		} else {
			safe = false
			break
		}
	}
	if safe {
		for i := 0; i < len(array)-1; i++ {
			if array[i]-array[i+1] == 1 || array[i]-array[i+1] == 2 || array[i]-array[i+1] == 3 {
				safe = true
			} else {
				safe = false
				break
			}
		}
	}
	//bei aufsteigender Reihenfolge
	if !safe {
		for i := 0; i < len(array)-1; i++ {
			if array[i] < array[i+1] {
				safe = true
			} else {
				safe = false
				break
			}
		}
		if safe {
			for i := 0; i < len(array)-1; i++ {
				if array[i]-array[i+1] == -1 || array[i]-array[i+1] == -2 || array[i]-array[i+1] == -3 {
					safe = true
				} else {
					safe = false
					break
				}
			}
		}
	}
	//Wenn die Ebene sicher ist, wird sie gezählt, sonst nicht
	var count int
	if safe {
		count = 1
	} else {
		count = 0
	}
	return count
}
