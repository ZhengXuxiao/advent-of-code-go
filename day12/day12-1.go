package main

import (
	"advent-of-code-go/advent-of-code-go/util"
	"strings"
)

type pots struct {
	generation     []byte
	nextGeneration []byte
	beginPotIndex  int
	beginIndex     int
	endIndex       int
	spreadRuleMap  map[string]byte
	prodecesSum    int
}

func (p *pots) initGeneration(initState string, initIndex int) {
	p.generation = []byte(strings.Join([]string{"....", initState, "...."}, ""))

	p.beginIndex = 4
	p.beginPotIndex = p.beginIndex + p.beginPotIndex - initIndex
	p.endIndex = len(p.generation) - 5
	// p.addProdecesSum()
}

func (p *pots) getSpreadRule(lines []string) {
	p.spreadRuleMap = make(map[string]byte)
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ".") {
			k, v := getSpreadRuleKeyValue(line)
			p.spreadRuleMap[k] = v
		}
	}

}

func (p *pots) getNextGeneration() {
	p.nextGeneration = make([]byte, len(p.generation))
	beginIndexTmp, endIndexTmp := len(p.generation), 0
	for i := p.beginIndex - 2; i <= p.endIndex+2; i++ {
		v, ok := p.spreadRuleMap[string(p.generation[i-2:i+3])]
		if ok {
			p.nextGeneration[i] = v
			if p.nextGeneration[i] == byte('#') && i < beginIndexTmp {
				beginIndexTmp = i
			}
			if p.nextGeneration[i] == byte('#') && i > endIndexTmp {
				endIndexTmp = i
			}
		} else {
			p.nextGeneration[i] = '.'
		}

	}
	p.initGeneration(string(p.nextGeneration[beginIndexTmp:endIndexTmp+1]), beginIndexTmp)
	// p.generation = []byte(strings.Join([]string{"....", string(p.nextGeneration[p.beginIndex : p.endIndex+1]), "...."}, ""))
	// p.beginIndex = 2
	// p.endIndex = len(p.generation) - 3
}

func (p *pots) addProdecesSum() {
	increment := 0
	for i, c := range p.generation {
		if c == '#' {
			increment += i - p.beginPotIndex
		}
	}
	print(increment - p.prodecesSum)
	print(" ")
	p.prodecesSum = increment
}

func getSpreadRuleKeyValue(line string) (string, byte) {
	key := line[:5]
	value := line[9]
	return key, value
}

func part1() {
	// 读取文件
	lines := util.ReadFile2Lines("./input.txt")

	p := pots{}
	p.prodecesSum = 0
	p.beginPotIndex = 4
	p.initGeneration("##...#......##......#.####.##.#..#..####.#.######.##..#.####...##....#.#.####.####.#..#.######.##...", 4)

	//获取转换状态
	p.getSpreadRule(lines)
	println(string(p.generation))
	//细胞自动机？
	for i := 1; i <= 150; i++ {
		//根据转换规律生成下一代
		p.getNextGeneration()
		p.addProdecesSum()
		// println(string(p.generation))
	}

	println(p.prodecesSum)

}

func main() {
	part1()

	println()
	print(8845 + 51*(50000000000-150))
}
