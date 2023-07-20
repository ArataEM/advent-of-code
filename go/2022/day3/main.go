package main

import (
	"bufio"
	"fmt"
	"os"
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

func byteToPriority(r byte) int {
	if 'a' <= r && r <= 'z' {
		return int((r-'a')%26 + 1)
	}
	if 'A' <= r && r <= 'Z' {
		return int((r-'A')%26 + 27)
	}
	return 0
}

func getPriority(line string) int {
	firstPart, secondPart := line[0:len(line)/2], line[(len(line)/2):]
	for i := range firstPart {
		for j := range secondPart {
			if firstPart[i] == secondPart[j] {
				return byteToPriority(firstPart[i])
			}
		}
	}

	return 0
}

func getCommonItem(l1 string, l2 string, l3 string) byte {
	for i := range l1 {
		for j := range l2 {
			if l1[i] == l2[j] {
				for k := range l3 {
					if l1[i] == l3[k] {
						return l1[i]
					}
				}
			}
		}
	}
	return 0
}

func main() {

	fileLines := readFile("input.txt")

	sum := 0
	for _, line := range fileLines {
		sum += getPriority(line)
	}

	fmt.Println(sum)

	sum = 0
	for i := 0; i < len(fileLines); i += 3 {
		sum += byteToPriority(getCommonItem(fileLines[i], fileLines[i+1], fileLines[i+2]))
	}

	fmt.Println(sum)

}
