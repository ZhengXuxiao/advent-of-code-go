package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string, integers chan int) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Later I want to create a buffer of lines, not just line-by-line here ...
		b, _ := strconv.Atoi(scanner.Text())
		// fmt.Println(b)
		integers <- b
	}
	close(integers)
}

func getFreqList(integers chan int, freqs chan int) {
	sum := 0
	for v := range integers {
		sum += v
		freqs <- sum
	}
	close(freqs)
}

func getDuplicate(filePath string, m map[int]bool, basicFreq int) (int, bool) {
	integers := make(chan int)
	//freqs := make(chan int)
	go readFile(filePath, integers)
	//go getFreqList(integers, freqs)
	//var i int
	for i := range integers {
		basicFreq += i
		_, ok := m[basicFreq]
		if ok == false {
			m[basicFreq] = true
			//fmt.Println("first", i)
		} else {
			return basicFreq, true
		}
	}
	return basicFreq, false
}

func main() {
	filePath := "./input.txt"
	m := make(map[int]bool)
	basicFreq := 0
	basicFreq, done := getDuplicate(filePath, m, basicFreq)
	for done == false {
		basicFreq, done = getDuplicate(filePath, m, basicFreq)
	}

	// for !done {
	// 	getDuplicate(filePath, m, basicFreq)
	// 	fmt.Println(basicFreq)
	// }
	fmt.Println(basicFreq, done)

	// An artificial input source.  Normally this is a file passed on the command line.
	// const input = "Foo\n(555) 123-3456\nBar\nBaz\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456\n(555) 123-3456"
	// numberOfTelephoneNumbers := telephoneNumbersInFile(input)
	// fmt.Println(numberOfTelephoneNumbers)
}
