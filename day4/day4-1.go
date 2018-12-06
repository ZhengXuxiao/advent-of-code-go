package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type record struct {
	actionTime time.Time
	action     string
}

type byActionTime []record

func (a byActionTime) Len() int           { return len(a) }
func (a byActionTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byActionTime) Less(i, j int) bool { return a[i].actionTime.Before(a[j].actionTime) }

func getTimeStr(line string) string {
	return line[1:17]
}

func getAction(line string) string {
	return line[19:]
}

func str2time(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04", timeStr)
	return t
}

func lines2record(line string) record {
	var r record
	r.actionTime = str2time(getTimeStr(line))
	r.action = getAction(line)
	return r
}

func hasAction(actionRecord string, action string) bool {
	return strings.HasPrefix(actionRecord, action)
}

func time2index(t time.Time) int {
	if t.Hour() == 23 {
		return t.Minute()
	}
	if t.Hour() == 0 {
		return t.Minute() + 60
	}
	return -1
}

func recordSleepSchedule(schedule []int, beginSleepIndex int, endSleepIndex int) {
	for i := beginSleepIndex; i < endSleepIndex; i++ {
		schedule[i] += 1
	}
}

func getGuardId(action string) string {
	reg := regexp.MustCompile(`[0-9]+`)
	strs := reg.FindAllString(action, -1)
	return strs[0]
}

func getSleepLongestGuardId(scheduleMap map[string][]int) string {
	var sleepMostGuard string
	sleepMostTime := 0
	for k, v := range scheduleMap {
		if sum(v) > sleepMostTime {
			sleepMostTime = sum(v)
			sleepMostGuard = k
		}
	}
	fmt.Println(sleepMostGuard, sleepMostTime)
	return sleepMostGuard
}

func getMostSleepTimeIndex(schedule []int) (int, int) {
	max := schedule[0]
	maxIndex := -1
	for i, v := range schedule {
		if v >= max {
			max = v
			maxIndex = i
		}
	}
	return maxIndex, max
}

func getMostSameTimeGuardId(scheduleMap map[string][]int) (string, int) {
	maxTimes := -1
	maxIndex := -1
	tagetGuardId := "0000"
	for k := range scheduleMap {
		i, v := getMostSleepTimeIndex(scheduleMap[k])
		if v > maxTimes {
			maxTimes = v
			maxIndex = i
			tagetGuardId = k
		}
	}
	return tagetGuardId, maxIndex
}

func sum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
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

	//按行解析内容
	//提取时间与行动内容
	//保存至列表中
	var records []record
	for line := range lines {
		rec := lines2record(line)
		records = append(records, rec)
	}

	//fmt.Println(records)
	//按时间排序
	sort.Sort(byActionTime(records))
	//fmt.Println(records)

	//创建不同守卫的睡眠时间汇总表
	scheduleMap := make(map[string][]int)
	guardId := "000"
	fallSleepFlag := false
	beginSleepIndex := -1
	endSleepIndex := -1
	for _, rec := range records {
		switch {
		case hasAction(rec.action, "falls"):
			fallSleepFlag = true
			beginSleepIndex = time2index(rec.actionTime)
		case hasAction(rec.action, "wakes"):
			endSleepIndex = time2index(rec.actionTime)
			//将对应GuardId的休息数组从beginSleepIndex到endSleepIndex之间+1
			if fallSleepFlag {
				recordSleepSchedule(scheduleMap[guardId], beginSleepIndex, endSleepIndex)
				fallSleepFlag = false
			}
		case hasAction(rec.action, "Guard"):
			guardId = getGuardId(rec.action)
			_, ok := scheduleMap[guardId]
			if !ok {
				schedule := make([]int, 120)
				scheduleMap[guardId] = schedule
			}
		}
	}

	//求和睡眠时间汇总表 得到谁的最多的守卫返回守卫编号
	gID := getSleepLongestGuardId(scheduleMap)

	//从对应时间汇总表中找到最大的值 返回该值
	fmt.Println(scheduleMap[gID])
	fmt.Println(getMostSleepTimeIndex(scheduleMap[gID]))

	gID_int, _ := strconv.Atoi(gID)
	t, _ := getMostSleepTimeIndex(scheduleMap[gID])
	fmt.Println(gID_int * (t - 60))

	a, b := getMostSameTimeGuardId(scheduleMap)
	aint, _ := strconv.Atoi(a)
	fmt.Println(aint * (b - 60))
}

func main() {
	part1()
	// timeStr := getTimeStr("[1518-09-24 00:32] falls asleep")
	// fmt.Println(timeStr)
	// fmt.Println(getAction("[1518-09-24 00:32] falls asleep"))
	// parsedTime := str2time(timeStr)
	// timeStr2 := getTimeStr("[[1518-05-19 00:16] wakes up] falls asleep")
	// fmt.Println(timeStr2)
	// parsedTime2 := str2time(timeStr2)
	// fmt.Println(parsedTime.Unix())
	// fmt.Println(parsedTime2.Unix())
	// fmt.Println(parsedTime2.Before(parsedTime))
}
