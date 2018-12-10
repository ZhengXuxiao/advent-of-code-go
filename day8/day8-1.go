package main

import (
	"os"
	"strconv"
	"strings"
)

type node struct {
	metadataLen int
	childCount  int
	metadata    []int
	childNode   []*node
}

func getNode(s []string) (*node, []string, int) {
	var n node
	n.childCount, _ = strconv.Atoi(s[0])
	n.metadataLen, _ = strconv.Atoi(s[1])
	n.metadata = make([]int, n.metadataLen)
	n.childNode = make([]*node, n.childCount)
	metedaSum := 0
	s = s[2:]
	var cn *node = new(node)
	var ms int
	for i := 0; i < n.childCount; i++ {
		cn, s, ms = getNode(s)
		// println(cn)
		n.childNode[i] = cn
		metedaSum += ms
	}
	for j := 0; j < n.metadataLen; j++ {
		n.metadata[j], _ = strconv.Atoi(s[j])
		mtmp, _ := strconv.Atoi(s[j])
		metedaSum += mtmp
		//print(s[j], " ")
	}
	return &n, s[n.metadataLen:], metedaSum
}

func readFile(filePath string) string {
	file, _ := os.Open(filePath)
	buf := make([]byte, 102400)
	n, _ := file.Read(buf)
	defer file.Close()
	return string(buf[:n])
}

func part1() *node {
	//读取文件
	//需要去除换行符
	filePath := "./input.txt"
	fileContent := readFile(filePath)

	s := strings.Split(fileContent, " ")
	n, s, ms := getNode(s)
	println("rest str len:", len(s))
	println(n.childCount)
	println("part1", ms)
	// println(n.childNode[0].childCount)
	return n
}

func main() {
	part2()
}
