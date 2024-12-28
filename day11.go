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

	//Input einlesen in Arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		linesArray = strings.Split(lines, " ")
	}

	for i := 0; i < 25; i++ {
		var stonesArray []string
		for j := 0; j < len(linesArray); j++ {
			switch {
			case linesArray[j] == "0":
				linesArray[j] = "1"

			case evenNumberOfDigits(linesArray[j]) == true:
				first, second := splitStone(linesArray[j])
				stonesArray = append(stonesArray, first)
				linesArray[j] = second

			default:
				// Umwandlung von String zu int
				num, err := strconv.Atoi(linesArray[j])
				if err != nil {
					fmt.Println("Fehler beim Umwandeln:", err)
				} else {
					num = num * 2024
				}
				linesArray[j] = strconv.Itoa(num)
			}
			stonesArray = append(stonesArray, linesArray[j])
		}
		//ersetzt linesArray durch stonesArray
		linesArray = append(linesArray[:0], stonesArray...)

		//fmt.Println("Stones ", stonesArray)

	}
	fmt.Println("Result ", len(linesArray))

}

func evenNumberOfDigits(digits string) (answer bool) {
	var even bool
	if len([]rune(digits))%2 == 0 {
		even = true
	} else {
		even = false
	}
	return even
}

func splitStone(digits string) (first string, second string) {
	runes := []rune(digits)
	num := len(runes) / 2
	firstrunes := runes[:num]
	secondrunes := runes[num:]
	firstrunes = removeLeadingZeros(firstrunes)
	secondrunes = removeLeadingZeros(secondrunes)
	stone1 := string(firstrunes)
	stone2 := string(secondrunes)
	return stone1, stone2
}

func removeLeadingZeros(array []rune) (newArray []rune) {
	for i := 0; i < len(array); i++ {
		if array[i] != 48 || len(array) == 1 {
			break
		} else {
			array = array[1:]
			i--
		}
	}
	return array
}
