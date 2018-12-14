package main

import "advent-of-code-go/advent-of-code-go/util"

type fuelGrid struct {
	grid [][]int
}

func (fG *fuelGrid) initGrid(serialNumber int) {
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			rackId := j + 1 + 10
			powerLv := (rackId*(i+1)+serialNumber)*rackId%1000/100 - 5
			fG.grid[i][j] = powerLv
		}
	}
}

func (fG *fuelGrid) getMaxSqure(squreSize int) (int, int, int) {
	maxTotalPower := 0
	coordinate_x, coordinate_y := 0, 0
	for i := 0; i < 300-squreSize; i++ {
		for j := 0; j < 300-squreSize; j++ {
			squreTotalPower := 0
			for k := 0; k < squreSize; k++ {
				squreTotalPower += util.Sum(fG.grid[i+k][j : j+squreSize])
			}
			if squreTotalPower > maxTotalPower {
				coordinate_x, coordinate_y = j+1, i+1
				maxTotalPower = squreTotalPower
			}
		}
	}
	return coordinate_x, coordinate_y, maxTotalPower
}

//可用积分图优化——Summed-area table
func (fG *fuelGrid) getBestSqure() (int, int, int) {
	best_x, best_y, best_SqureSize, best_TotalPower := 0, 0, 0, 0
	for i := 0; i < 300; i++ {
		x, y, totalPower := fG.getMaxSqure(i)
		if totalPower > best_TotalPower {
			best_x, best_y, best_TotalPower = x, y, totalPower
			best_SqureSize = i
		}
	}
	return best_x, best_y, best_SqureSize
}

func newFuelGrid() *fuelGrid {
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}
	fG := fuelGrid{grid}
	return &fG
}

func part1() {
	fg := newFuelGrid()

	//初始化计算各个cell的能量等级
	fg.initGrid(3999)
	println(fg.grid[4][2])
	x, y, _ := fg.getMaxSqure(3)
	println(x, y)

	//part2
	a, b, c := fg.getBestSqure()
	println(a, b, c)
}

func main() {
	part1()
}
