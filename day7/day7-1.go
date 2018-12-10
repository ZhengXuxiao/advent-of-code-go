package main

import (
	"bufio"
	"os"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type step struct {
	name         rune
	preStepCount int
	nextStepList []*step
	beginTime    int
	endTime      int
}

func creteStep(name rune) step {
	var s step
	s.name = name
	s.preStepCount = 0
	s.nextStepList = make([]*step, 0)
	return s
}

func (s *step) finish() {
	s.endTime = s.beginTime + int(rune(s.name)) - 64 + 60
	for _, nextStep := range s.nextStepList {
		nextStep.preStepCount -= 1
		nextStep.beginTime = max(nextStep.beginTime, s.endTime)
		//nextStep.endTime = nextStep.beginTime + int(rune(nextStep.name)) - 64
	}
}

func (s *step) addStep(stepPoint *step) {
	stepPoint.preStepCount += 1
	s.nextStepList = append(s.nextStepList, stepPoint)
}

var stepMap map[rune]*step

type StepSlice []rune

func (ss *StepSlice) getAlphabetically() rune {
	if len(*ss) >= 1 {
		AlphabeticallyRune := rune('Z')
		for i, v := range *ss {
			if v < AlphabeticallyRune {
				AlphabeticallyRune = v
				(*ss)[0], (*ss)[i] = (*ss)[i], (*ss)[0]
			}
		}
		(*ss) = (*ss)[1:]
		return AlphabeticallyRune
	}
	return ' '
}

// func (ss StepSlice) Len() int      { return len(ss) }
// func (ss StepSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
// func (ss StepSlice) Less(i, j int) bool {

// }
func line2Step(line string) (preStep, nextStep rune) {
	return rune(line[5]), rune(line[36])
}

func lines2StepMap(lines chan string) map[rune]*step {
	stepMap = make(map[rune]*step)
	for line := range lines {
		pre, next := line2Step(line)
		_, ok := stepMap[next]
		if !ok {
			s := creteStep(next)
			stepMap[next] = &s
		}
		// stepMap[next].preStepCount += 1
		_, ok = stepMap[pre]
		if !ok {
			s := creteStep(pre)
			stepMap[pre] = &s
		}
		stepMap[pre].addStep(stepMap[next])
		// stepMap[pre].nextStepList = append(stepMap[pre].nextStepList, stepMap[next])
	}
	return stepMap
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
	//读取文件
	filePath := "./input.txt"
	lines := make(chan string)
	go readFile(filePath, lines)

	//按行读取信息
	//创建map[byte]*step
	stepMap := lines2StepMap(lines)

	//遍历map
	//将preStepCount==0的step保存至输出列表
	//输出alphabetically的元素并删除
	//并将map对应的step删除
	//ss = append(ss, rune('Z')+1)
	for len(stepMap) > 0 {
		var ss StepSlice
		for _, v := range stepMap {
			if v.preStepCount == 0 {
				ss = append(ss, v.name)
			}
		}
		a := ss.getAlphabetically()
		print(string(a))
		stepMap[a].finish()
		delete(stepMap, a)
	}
}

func main() {
	part2()
	//println(rune('A'))
}
