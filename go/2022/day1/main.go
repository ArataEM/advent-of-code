package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
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

	var sums = []uint{}
	var sum = uint(0)
	var max = uint(0)
	for _, line := range fileLines {
		if len(line) > 0 {
			cal, err := strconv.Atoi(line)
			check(err)
			sum += uint(cal)
		} else if len(line) == 0 {
			sums = append(sums, sum)
			if sum > max {
				max = sum
			}
			sum = 0
		}
		// fmt.Println(line)
	}

	sort.Slice(sums, func(i, j int) bool {
		return sums[i] < sums[j]
	})

	fmt.Println(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3])
}
