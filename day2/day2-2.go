package main

import (
	"fmt"
)

func lineLoop(lines chan string, matchResults chan string) {
	linesSlices := make([]string, 100)
	for line1 := range lines {
		for _, line2 := range linesSlices {
			lineMatch(line2, line1, matchResults)
		}
		linesSlices = append(linesSlices, line1)
	}
	close(matchResults)
}

func lineMatch(line1 string, line2 string, matchResults chan string) {
	matchResult := ""
	for i := range line1 {
		if line1[i] == line2[i] {
			matchResult += string(line1[i])
		}
	}
	//fmt.Println(matchResult)
	matchResults <- matchResult
}

func part2() {
	//读取文件，将文件按行写入channel中
	filePath := "./input.txt"
	lines := make(chan string)
	go readFile(filePath, lines)

	//
	matchResults := make(chan string)
	go lineLoop(lines, matchResults)

	//
	maxMatchCount := 0
	var maxMatchResult string
	for matchResult := range matchResults {
		if len(matchResult) > maxMatchCount {
			maxMatchCount = len(matchResult)
			maxMatchResult = matchResult
		}
	}
	fmt.Println(maxMatchResult)
	fmt.Println("done!")
}

func main() {
	part2()
}
