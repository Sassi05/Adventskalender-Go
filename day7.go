package main

import (
	"bufio"
	"fmt"
	"math"
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

	var resultStringArray []string
	var tempArray []string
	var operationArray [][]int
	var resultArray []int
	var trueResult int

	//Input einlesen in Arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		linesArray := strings.Split(lines, ": ")
		resultStringArray = append(resultStringArray, linesArray[0])
		tempArray = append(tempArray, linesArray[1])
	}

	//Die Arrayinhalte des tempArray von String zu Int parsen und in Int Matrix einlesen
	for _, temps := range tempArray {

		splits := strings.Fields(temps)
		var row []int
		for _, split := range splits {

			num, err := strconv.Atoi(split)

			if err != nil {
				// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
				fmt.Println(err)
				return
			}
			row = append(row, num)
		}
		operationArray = append(operationArray, row)
	}
	for _, results := range resultStringArray {

		num, err := strconv.Atoi(results)

		if err != nil {
			// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
			fmt.Println(err)
			return
		}
		resultArray = append(resultArray, num)
	}

	//Array ausgeben bei Bedarf
	//fmt.Println(resultArray)

	//fmt.Println(operationArray)

	for i := 0; i < len(resultArray); i++ {
		count := len(operationArray[i]) - 1
		//es soll eine Position mehr da sein, als Operatoren
		positions := count + 1
		var binarySt string
		var operator []string
		//es gibt 2^Anzahl Operatoren Möglichkeiten
		possibilities := math.Pow(2, float64(count))
		for i := 0; i < int(possibilities); i++ {
			//Binary String mit dynamischer Breite (positions)
			//(Binärzahl, dynamische Breite, Zahl, die als Binärzahl dargestellt wird)
			binarySt = fmt.Sprintf("%0*b", positions, i)
			for _, char := range binarySt {
				operator = append(operator, string(char))
			}
		}

		result := 0
		result = sum(int(possibilities), operator, resultArray[i], operationArray[i]...)
		trueResult = trueResult + result
	}
	fmt.Println("Ergebnis ", trueResult)

}

// nimmt ein dynamisches Array (nums), das erwartete Ergebnis (result), die Anzahl an Möglichkeiten (possibilities)
// und ein Operatorarray mit Nullen und Einsen entgegen, dabei steht die 0 für Plus und die Eins für Multiplizieren
// Es werden alle Möglichkeiten durchgetestet, stimmt das Ergebnis mit dem erwarteten Ergebnis überein
// wird abgebrochen und das Ergebnis zurückgegeben
func sum(possibilities int, operator []string, result int, nums ...int) (trueResult int) {
	total := 0
	for j := 0; j < possibilities; j++ {
		for i, num := range nums {
			switch operator[i] {
			case "0":
				total += num

			case "1":
				total *= num
			}
		}
		operator = operator[len(nums):]

		if total == result {
			//fmt.Println("Ergebnis stimmt: ", total)
			return total
		} else {
			total = 0
		}
	}
	return total
}
