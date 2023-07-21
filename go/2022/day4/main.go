package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	return fileLines
}

func contains(s []string) int {
	x := strings.Split(s[0], "-")
	y := strings.Split(s[1], "-")

	x0, _ := strconv.Atoi(x[0])
	y0, _ := strconv.Atoi(y[0])
	y1, _ := strconv.Atoi(y[1])
	x1, _ := strconv.Atoi(x[1])
	if ((x0 <= y0) && (x1 >= y1)) || ((x0 >= y0) && (x1 <= y1)) {
		return 1
	}
	return 0
}

func overlap(s []string) int {
	x := strings.Split(s[0], "-")
	y := strings.Split(s[1], "-")

	x0, _ := strconv.Atoi(x[0])
	y0, _ := strconv.Atoi(y[0])
	y1, _ := strconv.Atoi(y[1])
	x1, _ := strconv.Atoi(x[1])
	if ((x0 <= y0) && (x1 >= y1)) ||
		((x0 >= y0) && (x1 <= y1)) ||
		((x0 <= y0) && (x1 >= y0)) ||
		((x0 <= y1) && (x1 >= y1)) {
		return 1
	}
	return 0
}

func main() {

	fileLines := readFile("input.txt")

	sum := 0
	for _, line := range fileLines {
		assign := strings.Split(line, ",")
		sum += contains(assign)
	}

	fmt.Println(sum)

	sum = 0
	for _, line := range fileLines {
		assign := strings.Split(line, ",")
		sum += overlap(assign)
	}

	fmt.Println(sum)
}
