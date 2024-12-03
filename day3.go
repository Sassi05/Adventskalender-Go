package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var multiNum1 []int
	var multiNum2 []int
	var result int

	//Textdatei einlesen
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Finde alle Übereinstimmungen im Ausgangstext
		findMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
		matches := findMul.FindAllString(line, -1)

		//Finde alle Zahlen
		number := regexp.MustCompile(`\d+`)
		for _, match := range matches {
			multis := number.FindAllString(match, -1)

			num1, err := strconv.Atoi(multis[0])
			num2, err := strconv.Atoi(multis[1])

			if err != nil {
				// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
				fmt.Println(err)
				return
			}

			//in zwei Int Arrays speichern
			multiNum1 = append(multiNum1, num1)
			multiNum2 = append(multiNum2, num2)

		}

	}

	//Arrays jeweils am gleichen Index multiplizieren und Ergebnisse addieren
	for i := 0; i < len(multiNum1); i++ {
		temp := multiNum1[i] * multiNum2[i]
		result = result + temp
	}

	fmt.Println("Gesamtsumme: ", result)
}
