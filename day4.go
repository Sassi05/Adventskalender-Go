package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var matrix [][]string
	var viceversaMatrix [][]string
	var rows int
	var cols int

	var result int

	//Textdatei einlesen
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	//Text in Matrix umwandeln und Matrix spiegeln (Vertikal), in normaler Matrix zählen
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = result + countSubString(line)
		var row []string
		var viceversaLine string
		var viceversaRow []string
		for _, char := range line {
			row = append(row, string(char))
		}
		viceversaLine = reverse(line)
		for _, char := range viceversaLine {
			viceversaRow = append(viceversaRow, string(char))
		}
		matrix = append(matrix, row)
		viceversaMatrix = append(viceversaMatrix, viceversaRow)
	}

	// Anzahl der Zeilen und Spalten
	rows = len(matrix)
	cols = len(matrix[0])

	//Zeilen und Spalten vertauschen, in dieser Matrix zählen
	for col := 0; col < cols; col++ {
		var verticalString strings.Builder
		for row := 0; row < rows; row++ {
			verticalString.WriteString(matrix[row][col])
		}
		vString := verticalString.String()
		result = result + countSubString(vString)
	}

	//In diagonaler Matrix zählen
	result = result + iterateDiagonal(matrix)
	result = result + iterateDiagonal(viceversaMatrix)

	fmt.Println("Result: ", result)

}

// zum zählen des Substrings
func countSubString(text string) int {
	var sum int
	count := strings.Count(text, "XMAS")
	sum = sum + count
	count = strings.Count(text, "SAMX")
	sum = sum + count
	return sum
}

// Iteration durch die Diagonalen und gleichzeitig zählen
func iterateDiagonal(matrix [][]string) int {
	var sum int
	rows := len(matrix)
	//Diagonale von links nach rechts oben
	j := 0
	for row := 0; row < rows; row++ {
		var diagonalString strings.Builder
		for i := 0; i < rows-j; i++ {
			diagonalString.WriteString(matrix[i][i+j])
		}
		j++
		dString := diagonalString.String()
		sum = sum + countSubString(dString)
	}

	//Diagonale von links nach rechts unten
	k := 1
	for row := 0; row < rows; row++ {
		var diagonalString strings.Builder
		for i := 0; i < rows-k; i++ {
			diagonalString.WriteString(matrix[i+k][i])
		}
		k++
		dString := diagonalString.String()
		sum = sum + countSubString(dString)
	}
	return sum

}

// String umdrehen
func reverse(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
