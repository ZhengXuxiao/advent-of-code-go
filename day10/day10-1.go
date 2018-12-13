package main

import (
	"advent-of-code-go/advent-of-code-go/util"

	"gonum.org/v1/gonum/stat"
)

type point struct {
	position_x int
	position_y int
	velocity_x int
	velocity_y int
}

type process struct {
	points     []point
	figure     [][]int8
	max_size_x int
	min_size_x int
	max_size_y int
	min_size_y int
}

func line2Point(line string) point {
	results := util.ParseNum(line)
	p := point{}
	p.position_x = results[0]
	p.position_y = results[1]
	p.velocity_x = results[2]
	p.velocity_y = results[3]
	return p
}

func newProcess(lines chan string) *process {
	pro := process{}
	for line := range lines {
		p := line2Point(line)
		pro.points = append(pro.points, p)
		pro.max_size_x = util.Max(pro.max_size_x, p.position_x)
		pro.max_size_y = util.Max(pro.max_size_y, p.position_y)
		pro.min_size_x = util.Min(pro.min_size_x, p.position_x)
		pro.min_size_y = util.Min(pro.min_size_y, p.position_y)
	}
	return &pro
}

func (p *process) updatePoints(steps int) {
	for i := range p.points {
		p.points[i].position_x += (p.points[i].velocity_x * steps)
		p.points[i].position_y += (p.points[i].velocity_y * steps)
	}
}

func (p *process) getBounds() []int {
	left, right, up, down := p.max_size_x-p.min_size_x, 0, 0, p.max_size_y-p.min_size_y
	for _, point := range p.points {
		x := point.position_x - p.min_size_x
		y := point.position_y - p.min_size_y
		left = util.Min(left, x-1)
		right = util.Max(right, x+1)
		up = util.Max(up, y+1)
		down = util.Min(down, y-1)
	}
	return []int{left, right, up, down}
}

func (p *process) printFigure() {
	size := p.getBounds()
	left, right, up, down := size[0], size[1], size[2], size[3]
	figure := make([][]int, up-down+1)
	for i := down; i <= up; i++ {
		figure[i-down] = make([]int, right-left+1)
	}
	for _, point := range p.points {
		x := point.position_x - p.min_size_x - left
		y := point.position_y - p.min_size_y - down
		figure[y][x] = 1
	}
	for j := 0; j < len(figure); j++ {
		for i := 0; i < len(figure[j]); i++ {
			if figure[j][i] == 0 {
				print("-")
			} else {
				print("#")
			}
		}
		println(" ")
	}

}

func printFigure(input [][]int8, size []int) {
	left, right, up, down := size[0], size[1], size[2], size[3]
	for i := down; i <= up; i++ {
		for j := left; j <= right; j++ {
			if input[i][j] == 0 {
				print("-")
			} else {
				print("#")
			}
		}
		println(" ")
	}
}

func loss2(size []int) int {
	left, right, up, down := size[0], size[1], size[2], size[3]
	return right - left + up - down
}

func loss(input [][]int8, lostfunc func([][]int8) float64) float64 {
	return lostfunc(input)
}

func entropy(input [][]int8) float64 {
	ps := make([]float64, len(input))
	for i, row := range input {
		ps[i] = float64(util.Sum(row)) / float64(len(row))
	}
	return stat.Entropy(ps)
}

func part1() {
	filePath := "./input.txt"
	lines := make(chan string)
	go util.ReadFile(filePath, lines)

	p := newProcess(lines)

	min_entropy := loss2(p.getBounds())
	secs := 0
	for true {
		// 可尝试黄金分割优化
		p.updatePoints(1)
		entropy_tmp := loss2(p.getBounds())
		if min_entropy < entropy_tmp {
			break
		}
		secs += 1
		min_entropy = entropy_tmp
	}
	println(secs)
	p.updatePoints(-1)
	p.printFigure()

}

func main() {
	part1()
}
