package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var firstArray []string
	var secondArray []string
	var firstNum []int
	var secondNum []int
	var result int

	//Den Text auf zwei String Arrays aufteilen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//Hier wird die line an den Leerzeichen gesplittet, es gibt aber zwei Leerzeichen (da muss man auch erstmal drauf kommen, bei der Fehlersuche ;-))
		twoNumbers := strings.Split(line, " ")
		firstArray = append(firstArray, twoNumbers[0])
		secondArray = append(secondArray, twoNumbers[3])
	}

	//Die Arrayinhalte von String zu Int parsen und in Int Arrays einlesen
	for i := 0; i < len(firstArray); i++ {
		num1, err := strconv.Atoi(firstArray[i])
		num2, err := strconv.Atoi(secondArray[i])

		if err != nil {
			// Fehlerbehandlung, falls die Eingabe keine gültige Zahl ist
			fmt.Println(err)
			return
		}
		firstNum = append(firstNum, num1)
		secondNum = append(secondNum, num2)
	}

	//Arrays sortieren
	sort.Ints(firstNum)
	sort.Ints(secondNum)

	//gleiche Stellen im Array von einander abziehen und als absoluten Betrag addieren
	for i := 0; i < len(firstNum); i++ {
		resultTemp := firstNum[i] - secondNum[i]
		if resultTemp < 0 {
			resultTemp = resultTemp * (-1)
		}
		result = result + resultTemp
	}

	// Resultat ausgeben
	fmt.Println("Endresultat :")
	fmt.Println(result)
}
