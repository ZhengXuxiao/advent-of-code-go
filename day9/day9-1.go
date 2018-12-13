package main

import "advent-of-code-go/advent-of-code-go/util"

type Marble struct {
	owner      int
	score      int
	preMarble  *Marble
	nextMarble *Marble
}

func (m *Marble) addMarble(newMarble *Marble) *Marble {
	newMarble.nextMarble = m.nextMarble.nextMarble
	m.nextMarble.nextMarble.preMarble = newMarble
	newMarble.preMarble = m.nextMarble
	m.nextMarble.nextMarble = newMarble
	// m = newMarble
	return newMarble
}

func (m *Marble) removeMarble() (int, *Marble) {
	delMarble := m.preMarble.preMarble.preMarble.preMarble.preMarble.preMarble.preMarble
	delMarble.nextMarble.preMarble = delMarble.preMarble
	delMarble.preMarble.nextMarble = delMarble.nextMarble
	// m = delMarble.nextMarble
	return delMarble.score, delMarble.nextMarble
}

func (m *Marble) printMarbles() {
	print(m.score, " ")
	m_tmp := m.nextMarble
	for m_tmp != m {
		print(m_tmp.score, " ")
		m_tmp = m_tmp.nextMarble
	}
	println()
}

func NewMarble(socre int) *Marble {
	m := Marble{}
	m.score = socre
	m.nextMarble = &m
	m.preMarble = &m
	return &m
}

func initPlayers(playCount int) map[int]int {
	playerScore := make(map[int]int)
	for i := 0; i < playCount; i++ {
		playerScore[i] = 0
	}
	return playerScore
}

func part1(maxScore int, players int) {
	m := NewMarble(0)
	// rootMarble := m

	playerScore := initPlayers(players)

	p := 0
	s := 0
	removedScore := 0
	for true {
		s += 1
		if s > maxScore {
			break
		}
		if s%23 == 0 {
			removedScore, m = m.removeMarble()
			playerScore[p%players] += (removedScore + s)
		} else {
			m_tmp := NewMarble(s)
			m = m.addMarble(m_tmp)
		}
		p += 1
		// println(m.score)
		// rootMarble.printMarbles()
	}

	finalMax := 0
	for i := 0; i < players; i++ {
		finalMax = util.Max(playerScore[i], finalMax)
		//println(i, playerScore[i])
	}

	println(finalMax)

}

func main() {
	// util.PrintHello()
	println()
	part1(71938*100, 462)
}
