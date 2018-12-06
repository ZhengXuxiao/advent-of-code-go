package main

import (
	"bufio"
	"os"
)

func readFile(filePath string, lines chan string) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lines <- str
	}
	close(lines)
}

func letterCount(lines chan string, letterCountMaps chan map[string]int) {
	for line := range lines {
		letterCountMap := make(map[string]int)
		for _, letter := range line {
			_, ok := letterCountMap[string(letter)]
			if ok {
				letterCountMap[string(letter)] += 1
			} else {
				letterCountMap[string(letter)] = 1
			}
		}
		letterCountMaps <- letterCountMap
	}
	close(letterCountMaps)
}

func hasCounts(letterCountMap map[string]int, destCount int, hasCountChan chan int) {
	for k, v := range letterCountMap {
		if v == destCount {
			delete(letterCountMap, k)
			hasCountChan <- 1
			return
		}
	}
}

func caculSum(countChan chan int, sumChans chan int) {
	sum := 0
	for count := range countChan {
		sum += count
	}
	sumChans <- sum
}

// func main() {
// 	//读取文件，将文件按行写入channel中
// 	filePath := "./input.txt"
// 	lines := make(chan string)
// 	go readFile(filePath, lines)

// 	//按行转换成letter-count的map，写入channel
// 	letterCountMaps := make(chan map[string]int)
// 	go letterCount(lines, letterCountMaps)

// 	twiceCount := make(chan int)
// 	tripleCount := make(chan int)
// 	go func() {
// 		for letterCountMap := range letterCountMaps {
// 			hasCounts(letterCountMap, 2, twiceCount)
// 			hasCounts(letterCountMap, 3, tripleCount)
// 		}
// 		close(twiceCount)
// 		close(tripleCount)
// 	}()

// 	// //
// 	sumChan := make(chan int)
// 	go caculSum(twiceCount, sumChan)
// 	go caculSum(tripleCount, sumChan)

// 	for sumChan := range sumChan {
// 		fmt.Println(sumChan)
// 	}

// }
