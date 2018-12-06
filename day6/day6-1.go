package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	name          int
	distance      int
	totalDistance int
	isCentre      bool
}

type coordinates struct {
	x int
	y int
}

type gridMap struct {
	xSize int
	ySize int
	area  [][]point
}

func line2coordinates(line string) coordinates {
	c := new(coordinates)
	reg := regexp.MustCompile(`[0-9]+`)
	strs := reg.FindAllString(line, -1)
	x, _ := strconv.Atoi(strs[0])
	y, _ := strconv.Atoi(strs[1])
	c.x = x
	c.y = y
	return *c
}

func getAreSize(coordinatesList []coordinates) coordinates {
	var mapSize coordinates = coordinates{0, 0}
	for _, c := range coordinatesList {
		if c.x > mapSize.x {
			mapSize.x = c.x
		}
		if c.y > mapSize.y {
			mapSize.y = c.y
		}
	}
	return mapSize
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Manhattan(c1, c2 coordinates) int {
	var distance int = 0

	distance += Abs(c1.x - c2.x)
	distance += Abs(c1.y - c2.y)

	return distance
}

func (m *gridMap) tagMap(c coordinates, name int) {
	for j := 0; j < m.xSize; j++ {
		for i := 0; i < m.ySize; i++ {
			g := m.area[i][j]
			Mdist := Manhattan(coordinates{j, i}, c)
			g.totalDistance += Mdist
			if Mdist < g.distance {
				g.distance = Mdist
				g.name = name
			} else if Mdist == g.distance {
				g.name = -1
			}
			m.area[i][j] = g
		}
	}
}

func (m *gridMap) isInfinite(c coordinates, name int) bool {
	if m.area[0][c.x].name == name || m.area[c.y][0].name == name || m.area[m.ySize-1][c.x].name == name || m.area[c.y][m.xSize-1].name == name {
		return true
	}
	return false
}

func (m *gridMap) getAreaSize(c coordinates, name int) int {
	areaSize := 0
	for i := 0; i < m.ySize; i++ {
		for j := 0; j < m.xSize; j++ {
			if m.area[i][j].name == name {
				areaSize += 1
			}
		}
	}
	return areaSize
}

func (m *gridMap) getRegionSize(safeCoordinatesName []int, safeDist int) int {
	regionSize := 0
	for i := 0; i < m.ySize; i++ {
		for j := 0; j < m.xSize; j++ {
			if m.area[i][j].totalDistance < safeDist {
				regionSize += 1
			}
			// if isIn(m.area[i][j].name, safeCoordinatesName) && m.area[i][j].totalDistance < safeDist {
			// 	regionSize += 1
			// }
			// if m.area[i][j].name == -1 && i != 0 && j != 0 && i != (m.ySize-1) && j != (m.xSize-1) && m.area[i][j].totalDistance < safeDist {
			// 	regionSize += 1
			// }
		}
	}
	return regionSize
}

func isIn(target int, list []int) bool {
	for _, i := range list {
		if i == target {
			return true
		}
	}
	return false
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

func part1() {
	//按行读取文件
	filePath := "./input.txt"
	lines := make(chan string)
	go readFile(filePath, lines)

	//按行转换成坐标
	var allCoordinates []coordinates
	for line := range lines {
		c := line2coordinates(line)
		allCoordinates = append(allCoordinates, c)
	}

	//从坐标中计算地图面积
	mapSize := getAreSize(allCoordinates)
	//fmt.Println(mapSize.x, mapSize.y)

	//创建grid
	gm := new(gridMap)
	gm.xSize = mapSize.x + 1
	gm.ySize = mapSize.y + 1
	initDist := gm.xSize + gm.ySize
	for i := 0; i < gm.ySize; i++ {
		rowArea := make([]point, gm.xSize)
		for j := 0; j < gm.xSize; j++ {
			rowArea[j] = point{-1, initDist, 0, false}
		}
		gm.area = append(gm.area, rowArea)
	}

	//从坐标列表中读取坐标
	//每读取一个坐标后，在地图上进行标记
	for i, c := range allCoordinates {
		gm.tagMap(c, i)
	}
	// for _, row := range gm.area {
	// 	for _, p := range row {
	// 		print(p.name)
	// 	}
	// 	println("")
	// }

	//判断是否infinite区域
	maxArea := 0
	safeCoordinateName := make([]int, 1)
	for i, c := range allCoordinates {
		if gm.isInfinite(c, i) {
			println(i, "is infinite")
		} else {
			iAreaSize := gm.getAreaSize(c, i)
			if iAreaSize > maxArea {
				maxArea = iAreaSize
			}
			safeCoordinateName = append(safeCoordinateName, i)
			//println(i, iAreaSize)
		}
	}
	println("largest area", maxArea)

	//part2
	println(gm.getRegionSize(safeCoordinateName, 10000))
}

func main() {
	part1()
}
