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

	var linesArray []string

	//Input einlesen in Arrays
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		linesArray = append(linesArray, lines)
	}

	//Matrix erstellen
	cols := len(linesArray)
	matrix := make([][]rune, cols)
	for i, line := range linesArray {
		matrix[i] = []rune(line)
	}

	// Matrix ausgeben bei Bedarf
	/*for _, row := range matrix {
		fmt.Println(string(row))
	}*/

	//rund um den Ausgangspunkt nach Wegen suchen
	countPath := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			//die Null in Rune ist 48
			if matrix[i][j] == 48 {
				//Hier wird eine temporäre Matrix erstellt, damit man die gefundenen Neunen streichen kann
				var tempMatrix [][]rune
				for i := 0; i < len(matrix); i++ {
					tempMatrix = make([][]rune, len(matrix))
					for j := range matrix {
						tempMatrix[j] = append([]rune(nil), matrix[j]...) // Erstellt eine Kopie der Zeile
					}
				}
				// & countPath = Adresse von countPath
				countPath = search(48, i, j, tempMatrix, &countPath)
			}
		}

	}
	fmt.Println("Result: ", countPath)
}

// rekursive Funktion um nach Wegen zu suchen
func search(k int, i int, j int, matrix [][]rune, count *int) (pathes int) {
	k = k + 1
	if k == 58 {
		//Pointer auf den Zähler
		*count++
		return *count
	}
	runeSignUp := getRune(i-1, j, matrix)
	if runeSignUp == rune(k) {
		//57 ist die Neun in Rune
		if k == 57 {
			matrix[i-1][j] = 'X'
		}
		search(k, i-1, j, matrix, count)
	}
	runeSignDown := getRune(i+1, j, matrix)
	if runeSignDown == rune(k) {
		if k == 57 {
			matrix[i+1][j] = 'X'
		}
		search(k, i+1, j, matrix, count)
	}
	runeSignLeft := getRune(i, j-1, matrix)
	if runeSignLeft == rune(k) {
		if k == 57 {
			matrix[i][j-1] = 'X'
		}
		search(k, i, j-1, matrix, count)
	}
	runeSignRight := getRune(i, j+1, matrix)
	if runeSignRight == rune(k) {
		if k == 57 {
			matrix[i][j+1] = 'X'
		}
		search(k, i, j+1, matrix, count)
	}
	return *count
}

func getRune(x int, y int, matrix [][]rune) rune {
	switch {
	case x < 0:
		return ' '

	case x > len(matrix[0])-1:
		return ' '

	case y < 0:
		return ' '

	case y > len(matrix)-1:
		return ' '
	}
	return matrix[x][y]
}
