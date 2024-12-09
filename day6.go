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
	var steps int

	//Input einlesen in Array
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

	//Bewegen
	startX, startY := findStart(matrix)
	steps = movingForward(matrix, startX, startY, steps)
	fmt.Println("Result: ", steps)

	// Matrix ausgeben bei Bedarf
	/*for _, row := range matrix {
		fmt.Println(string(row))
	}*/

}

// Startposition finden
func findStart(matrix [][]rune) (int, int) {
	var x int
	var y int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '^' {
				x = i
				y = j
			}
		}
	}
	return x, y
}

// nach oben
func movingForward(matrix [][]rune, x int, y int, steps int) (newSteps int) {
	moving := "up"
	for {
		switch moving {
		case "up":
			nextX := -1
			for i := x; i > -1; i-- {
				if matrix[i][y] == '#' {
					nextX = i + 1
					break
				}
				//X werden nicht doppelt gezählt
				if matrix[i][y] != 'X' {
					steps = steps + 1
				}
				matrix[i][y] = 'X'

			}
			steps = steps - 1 //zählt einen Step zu viel
			//zum Schluß, darf der Step nicht abgezogen werden
			if nextX == -1 {
				return steps + 1
			}
			x = nextX
			matrix[x][y] = '>' // Neue Position markieren
			moving = "right"

		case "right":
			nextY := -1
			for i := y; i < len(matrix[x]); i++ {
				if matrix[x][i] == '#' {
					nextY = i - 1
					break
				}
				//X werden nicht doppelt gezählt
				if matrix[x][i] != 'X' {
					steps = steps + 1
				}
				matrix[x][i] = 'X'
			}
			steps = steps - 1 //zählt einen Step zu viel
			//zum Schluß, darf der Step nicht abgezogen werden
			if nextY == -1 {
				return steps + 1
			}
			y = nextY
			matrix[x][y] = '▼' // Neue Position markieren
			moving = "down"

		case "down":
			nextX := -1
			for i := x; i < len(matrix); i++ {
				if matrix[i][y] == '#' {
					nextX = i - 1
					break
				}
				//X werden nicht doppelt gezählt
				if matrix[i][y] != 'X' {
					steps = steps + 1
				}
				matrix[i][y] = 'X'
			}
			steps = steps - 1 //zählt einen Step zu viel
			//zum Schluß, darf der Step nicht abgezogen werden
			if nextX == -1 {
				return steps + 1
			}
			x = nextX
			matrix[x][y] = '<' // Neue Position markieren
			moving = "left"

		case "left":
			nextY := -1
			for i := y; i > -1; i-- {
				if matrix[x][i] == '#' {
					nextY = i + 1
					break
				}
				//X werden nicht doppelt gezählt
				if matrix[x][i] != 'X' {
					steps = steps + 1
				}
				matrix[x][i] = 'X'

			}
			steps = steps - 1 //zählt einen Step zu viel
			//zum Schluß, darf der Step nicht abgezogen werden
			if nextY == -1 {
				return steps + 1
			}
			y = nextY
			matrix[x][y] = '^' // Neue Position markieren
			moving = "up"

		}

	}
}
