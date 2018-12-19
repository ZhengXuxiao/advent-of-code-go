package main

import (
	"advent-of-code-go/advent-of-code-go/util"
	"sort"
)

type Cart struct {
	position  complex64
	direction complex64
	nextTurn  int
	isCrush   bool
}

func (c *Cart) move() {
	c.position += c.direction
}

func (c *Cart) turn(mark byte) {
	//有没抽象的计算方法？
	if mark == '\\' {
		if real(c.direction) == 0 {
			c.direction *= -1i
		} else {
			c.direction *= 1i
		}
	} else if mark == '/' {
		if real(c.direction) == 0 {
			c.direction *= 1i
		} else {
			c.direction *= -1i
		}
	} else if mark == '+' {
		switch c.nextTurn % 3 {
		case 0:
			c.direction *= 1i
		// case 1:
		// 	c.direction = c.direction
		case 2:
			c.direction *= -1i
		}
		c.nextTurn += 1
	}
}

type CartSort []*Cart

func (c CartSort) Len() int      { return len(c) }
func (c CartSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CartSort) Less(i, j int) bool {
	return real(c[i].position) < real(c[j].position) || real(c[i].position) == real(c[j].position) && imag(c[i].position) < imag(c[j].position)
}

func setup() (map[complex64]byte, []*Cart) {
	carts := make([]*Cart, 0)
	pathMap := make(map[complex64]byte)

	// 读取文件
	lines := util.ReadFile2Lines("./input.txt")

	for i, line := range lines {
		for j, c := range line {
			pos := complex(float32(i), float32(j))
			switch c {
			case '^':
				carts = append(carts, &Cart{pos, -1, 0, false})
			case 'v':
				carts = append(carts, &Cart{pos, 1, 0, false})
			case '<':
				carts = append(carts, &Cart{pos, -1i, 0, false})
			case '>':
				carts = append(carts, &Cart{pos, 1i, 0, false})
			case '/', '\\', '+':
				pathMap[pos] = byte(c)
			}
		}
	}

	return pathMap, carts
}

func part1() {
	pathMap, carts := setup()
	crushNum := 0
	for len(carts)-crushNum > 1 {
		sort.Sort(CartSort(carts))
		// for i, c1 := range carts {
		// 	for j, c2 := range carts {
		// 		if c1.position == c2.position && i != j && !c2.isCrush {
		// 			println(c1.position)
		// 			c1.isCrush = true
		// 			c2.isCrush = true
		// 			c1.position = complex(float32(-1), float32(-1))
		// 			c2.position = complex(float32(-1), float32(-1))
		// 			crushNum += 2
		// 			// c.position = complex(float32(-1), float32(-1))
		// 			// crushNum += 1
		// 		}
		// 	}
		// }
		for _, c := range carts {
			if c.isCrush {
				continue
			}
			c.move()
			mark, ok := pathMap[c.position]
			if ok {
				c.turn(mark)
			}
			for i, c1 := range carts {
				for j, c2 := range carts {
					if c1.position == c2.position && i != j && !c2.isCrush && !c1.isCrush {
						println(c1.position)
						c1.isCrush = true
						c2.isCrush = true
						c1.position = complex(float32(-1), float32(-1))
						c2.position = complex(float32(-1), float32(-1))
						crushNum += 2
						// c.position = complex(float32(-1), float32(-1))
						// crushNum += 1
					}
				}
			}
		}

	}
	println("done")
	for _, c := range carts {
		if c.isCrush == false {
			println(c.position)
		}
	}

}

func main() {
	part1()
}
