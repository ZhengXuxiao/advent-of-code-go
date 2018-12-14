package util

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func PrintHello() {
	print("hello world")
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func ReadFile(filePath string, lines chan string) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lines <- str
	}
	close(lines)
}

func ParseNum(line string) []int {
	reg := regexp.MustCompile(`(\-)?\d+`)
	strs := reg.FindAllString(line, -1)
	results := make([]int, len(strs))
	for i, v := range strs {
		int_v, _ := strconv.Atoi(v)
		results[i] = int_v
	}
	return results
}
