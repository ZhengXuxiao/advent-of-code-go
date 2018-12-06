package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type claimEntity struct {
	claimId string
	area    []int
}

func readFile(filePath string, lines chan string) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		lines <- str
	}
	close(lines)
}

func parseLine(line string) claimEntity {
	var ce claimEntity
	reg := regexp.MustCompile(`[0-9]+`)
	strs := reg.FindAllString(line, -1)
	results := make([]int, 4)
	for i, v := range strs[1:] {
		int_v, _ := strconv.Atoi(v)
		results[i] = int_v
	}
	ce.claimId = strs[0]
	ce.area = results
	return ce
}

func getMaxSize(claimEntities []claimEntity) (int, int) {
	maxRowSize := 0
	maxColSize := 0
	var area []int
	for _, ce := range claimEntities {
		area = ce.area
		if area[0]+area[2] > maxColSize {
			maxColSize = area[0] + area[2]
		}
		if area[1]+area[3] > maxRowSize {
			maxRowSize = area[1] + area[3]
		}
	}
	return maxRowSize, maxColSize
}

func addAreas(area []int, fabric [][]int) {
	for row := area[1]; row < area[1]+area[3]; row++ {
		for col := area[0]; col < area[0]+area[2]; col++ {
			fabric[row][col] += 1
		}
	}
}

func getOverlapClaimSum(fabric [][]int, destClaimTimes int) int {
	overlapClaimSum := 0
	for _, fabricLine := range fabric {
		for _, claimTimes := range fabricLine {
			if claimTimes >= destClaimTimes {
				overlapClaimSum += 1
			}
		}
	}
	return overlapClaimSum
}

func part1() ([]claimEntity, [][]int) {
	filePath := "./input.txt"
	lines := make(chan string)
	go readFile(filePath, lines)

	//收集所有claim
	var claimEntities []claimEntity
	for line := range lines {
		ce := parseLine(line)
		claimEntities = append(claimEntities, ce)
	}

	//找出最大面积
	maxRowSize, maxColSize := getMaxSize(claimEntities)
	fmt.Println(maxRowSize, maxColSize)
	var fabric [][]int
	for i := 0; i < maxRowSize; i++ {
		sl := make([]int, maxColSize)
		fabric = append(fabric, sl)
	}

	for _, ce := range claimEntities {
		addAreas(ce.area, fabric)
	}
	fmt.Println(getOverlapClaimSum(fabric, 2))
	return claimEntities, fabric
}

func main() {
	//_, fabric := part1()
	//fmt.Println(fabric[0][0])
	part2()
}
