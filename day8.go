package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//Textdatei einlesen
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}

	var linesArray []string
	var result int
	var nodesFound []string

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

	var signArray []rune
	//verschiedene Zeichen aus der Matrix in einem Array sammeln
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '.' {
				if !contains(signArray, matrix[i][j]) {
					signArray = append(signArray, matrix[i][j])
				}
			}
		}
	}

	//für jedes Zeichen die Antinodes zählen
	for k := 0; k < int(len(signArray)); k++ {
		var tempCount int
		tempCount, nodesFound = countAntinodes(matrix, signArray[k], nodesFound)
		//fmt.Println("tempCount", tempCount)

		result = result + tempCount
		//fmt.Println("Zwischenresult", result)
	}

	fmt.Println("Result ", result)
}

func contains(slice []rune, value rune) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func countAntinodes(matrix [][]rune, sign rune, nodeArray []string) (count int, nodes []string) {
	var xArray []int
	var yArray []int
	var counts int
	var antiNode1X int
	var antiNode1Y int
	var antiNode2X int
	var antiNode2Y int

	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == sign {
				xArray = append(xArray, i)
				yArray = append(yArray, j)
			}
		}
	}
	//fmt.Println(xArray)
	//fmt.Println(yArray)

	for k := 0; k < len(xArray)-1; k++ {
		currentX := xArray[k]
		currentY := yArray[k]
		for l := k; l < len(xArray)-1; l++ {
			otherX := xArray[l+1]
			otherY := yArray[l+1]

			//Betrag der Unterschiede errechnen
			x := currentX - otherX
			x = makeAbsolutValue(x)
			y := currentY - otherY
			y = makeAbsolutValue(y)

			//Antinodesstellen finden
			minX := findMin(currentX, otherX)
			maxY := findMax(currentY, otherY)
			maxX := findMax(currentX, otherX)
			minY := findMin(currentY, otherY)
			if currentY >= otherY {
				antiNode1X = minX - x
				antiNode1Y = maxY + y
				antiNode2X = maxX + x
				antiNode2Y = minY - y
			} else {
				antiNode1X = minX - x
				antiNode1Y = minY - y
				antiNode2X = maxX + x
				antiNode2Y = maxY + y
			}
			//Antinodes dürfen nicht doppelt gezählt werden, daher müssen schon vorhandene im Array gesammelt werden
			node1 := strconv.Itoa(antiNode1X) + "/" + strconv.Itoa(antiNode1Y)
			node2 := strconv.Itoa(antiNode2X) + "/" + strconv.Itoa(antiNode2Y)

			//wenn noch keine Antinode an der Stelle ist:
			if !containsString(nodeArray, node1) {
				if antiNode1X > -1 && antiNode1X < rows && antiNode1Y > -1 && antiNode1Y < cols {
					counts = counts + 1
					//neue Antinodes in Array stellen
					nodeArray = append(nodeArray, node1)
				}
			}

			if !containsString(nodeArray, node2) {
				if antiNode2X > -1 && antiNode2X < rows && antiNode2Y > -1 && antiNode2Y < cols {
					counts = counts + 1
					//neue Antinodes in Array stellen
					nodeArray = append(nodeArray, node2)
				}
			}
		}

	}
	return counts, nodeArray
}

func containsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func makeAbsolutValue(value int) (result int) {
	if value < 0 {
		value = -value
	}
	return value
}

func findMin(first int, second int) (result int) {
	if first < second {
		return first
	} else {
		return second
	}
}

func findMax(first int, second int) (result int) {
	if first > second {
		return first
	} else {
		return second
	}
}
