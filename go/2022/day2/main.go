package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

var shapeScore = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func getMyShape(s string, res string) string {
	if (s == "A" && res == "X") || (s == "C" && res == "Y") || (s == "B" && res == "Z") {
		return "C"
	}
	if (s == "B" && res == "X") || (s == "A" && res == "Y") || (s == "C" && res == "Z") {
		return "A"
	}
	if (s == "C" && res == "X") || (s == "B" && res == "Y") || (s == "A" && res == "Z") {
		return "B"
	}
	return ""
}

func main() {

	filePath := "input.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	score := 0

	for i, line := range fileLines {
		shapes := strings.Split(line, " ")
		if i < 10 {
			fmt.Printf("%s-%s-%s\n", shapes[0], shapes[1], getMyShape(shapes[0], shapes[1]))
		}
		switch shapes[1] {
		case "X":
			score += shapeScore[getMyShape(shapes[0], shapes[1])]
		case "Y":
			score += 3 + shapeScore[getMyShape(shapes[0], shapes[1])]
		case "Z":
			score += 6 + shapeScore[getMyShape(shapes[0], shapes[1])]
		}
	}

	fmt.Println(score)
}
