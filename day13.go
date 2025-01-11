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

	var linesArray []string
	var buttonA []string
	var buttonB []string
	var prize []string

	var buttonAInt []int
	var buttonBInt []int
	var prizeInt []int

	var A int
	var B int
	var result int

	//Input einlesen in Arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		//das Komma hier erleichtert hinterher das Auslesen der Zahlen
		lines += ","
		linesArray = append(linesArray, lines)
	}

	for i := 0; i < len(linesArray); i++ {
		line := linesArray[i]
		if strings.Contains(line, "Button") {
			for j := 0; j < len(line); j++ {
				//durchsuche die Zeile nach dem Pluszeichen, ab dem aktuellen Index
				//um die Zahlen herauszufiltern
				startIndex := strings.Index(line[j:], "+")
				startIndex = j + startIndex
				endIndex := strings.Index(line[startIndex:], ",")
				endIndex = startIndex + endIndex
				//die Zahlen in zwei verschiedenen Arrays speichern
				switch {
				case strings.Contains(line, "A"):
					buttonA = append(buttonA, line[startIndex+1:endIndex])

				case strings.Contains(line, "B"):
					buttonB = append(buttonB, line[startIndex+1:endIndex])

				}
				j = endIndex + 1
			}
		}

		if strings.Contains(line, "Prize") {
			for j := 0; j < len(line); j++ {
				//durchsuche die Zeile nach dem Pluszeichen, ab dem aktuellen Index
				//um die Zahlen herauszufiltern
				startIndex := strings.Index(line[j:], "=")
				startIndex = j + startIndex
				endIndex := strings.Index(line[startIndex:], ",")
				endIndex = startIndex + endIndex
				//im Array speichern
				prize = append(prize, line[startIndex+1:endIndex])
				j = endIndex + 1
			}
		}
	}

	//Strings in Int umwandeln
	buttonAInt = changeStringInInt(buttonA)
	buttonBInt = changeStringInInt(buttonB)
	prizeInt = changeStringInInt(prize)

	//Anzahl der Tokens wird berechnet
	for i := 0; i < len(buttonA); i++ {
		var tempResult int
		B = calculateB(buttonAInt[i], buttonAInt[i+1], buttonBInt[i], buttonBInt[i+1], prizeInt[i], prizeInt[i+1])
		A = calculateA(prizeInt[i], buttonBInt[i], B, buttonAInt[i])
		if (B < 101) && (A < 101) {
			tempResult = A*3 + B
		}
		result = result + tempResult
		i = i + 1
	}

	fmt.Println("result: ", result)

}

func changeStringInInt(StringArray []string) (IntArray []int) {
	for _, str := range StringArray {
		// Versuche, den String in einen Integer umzuwandeln
		num, err := strconv.Atoi(str)
		if err != nil {
			// Fehlerbehandlung, falls der String keine gültige Zahl ist
			fmt.Printf("Fehler beim Umwandeln von '%s': %v\n", str, err)
			continue
		}

		// Füge den Integer-Wert zum intArr-Array hinzu
		IntArray = append(IntArray, num)
	}
	return IntArray
}

// errechnet B und prüft dabei, ob B eine ganze Zahl ist, wenn nicht
// wird 101 ausgegeben, da die Anzahl an Tastendrücken sowieso nur höchstens 100 sein soll
func calculateB(AX int, AY int, BX int, BY int, PX int, PY int) (B int) {
	divisor := (AX * BY) - (AY * BX)
	divident := (PY * AX) - (AY * PX)
	tempB := divident / divisor
	if divident%divisor == 0 {
		// Wert ist eine ganze Zahl
		B = tempB
	} else {
		B = 101
	}
	return B
}

// errechnet A und prüft dabei, ob B eine ganze Zahl ist, wenn nicht
// wird 101 ausgegeben, da die Anzahl an Tastendrücken sowieso nur höchstens 100 sein soll
func calculateA(PX int, BX int, B int, AX int) (A int) {
	divident := PX - (BX * B)
	tempA := divident / AX
	if divident%AX == 0 {
		// Wert ist eine ganze Zahl
		A = tempA
	} else {
		A = 101
	}
	return A
}
