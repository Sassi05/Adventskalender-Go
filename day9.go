package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Textdatei einlesen
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	var filesTemp []rune
	var spaceTemp []rune
	var filesArray []int
	var spaceArray []int
	var resultArray []int
	var lastIndex int
	var lastIndexValue int
	var noSpace bool

	//Input einlesen in Arrays, dabei werden die files in einem, der space im anderen Array gespeichert
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		for i, char := range lines {
			if i%2 == 0 {
				filesTemp = append(filesTemp, char)
			} else {
				spaceTemp = append(spaceTemp, char)
			}
		}
	}

	// Umwandlung jedes Elements von rune zu int
	for _, r := range filesTemp {
		filesArray = append(filesArray, int(r-'0')) // Umwandlung von rune zu int
	}
	for _, r := range spaceTemp {
		spaceArray = append(spaceArray, int(r-'0')) // Umwandlung von rune zu int
	}

	for i := 0; i < len(filesArray); i++ {
		noSpace = false
		//Anzahl der ID-Zahl, die komprimiert werden soll
		number := filesArray[i]
		for j := 0; j < number; j++ {
			//die ID-Zahl wird entsprechend ihres Wertes zugefügt
			resultArray = append(resultArray, i)
		}
		//auch wenn es noch space gibt, soll nach dem letzten file im fileArray gestoppt werden
		stop := len(filesArray) - i
		if stop == 1 {
		}
		if len(spaceArray) != 0 && stop != 1 {
			//erste Stelle im SpaceArray, Anzahl der Leerstellen, die zur Verfügung stehen
			space := spaceArray[0]
			lastIndex = len(filesArray) - 1
			lastIndexValue = filesArray[lastIndex]
			//in die zur Verfügungen stehenden Leerstellen, werden die letzen Indexstellen des filesArray eingefügt
			for space > lastIndexValue {
				for k := 0; k < lastIndexValue; k++ {
					resultArray = append(resultArray, lastIndex)
				}
				space = space - lastIndexValue
				filesArray = filesArray[:len(filesArray)-1]
				lastIndex = len(filesArray) - 1
				if lastIndex == i { //wenn dies zutrifft, wurde der vorletzte File schon dem Array zugefügt, da er von links kommt, der space könnte aber noch größer sein
					space = 0
				}
				lastIndexValue = filesArray[lastIndex]
			}

			if space == lastIndexValue {
				for l := 0; l < lastIndexValue; l++ {
					resultArray = append(resultArray, lastIndex)
				}
				space = space - lastIndexValue
				filesArray = filesArray[:len(filesArray)-1]
				lastIndex = len(filesArray) - 1
				spaceArray = spaceArray[1:]
				noSpace = true
			}

			if space <= lastIndexValue && !noSpace {
				for m := 0; m < space; m++ {
					resultArray = append(resultArray, lastIndex)
				}
				lastIndexValue = lastIndexValue - space
				space = 0
				filesArray[lastIndex] = lastIndexValue
				spaceArray = spaceArray[1:]
			}

		}
	}

	var result int
	//fmt.Println("Ergebnis", resultArray)
	for i := 0; i < len(resultArray); i++ {
		result = result + (i * resultArray[i])
	}

	fmt.Println("Result ", result)
}
