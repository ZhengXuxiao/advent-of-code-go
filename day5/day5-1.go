package main

import (
	"fmt"
	"os"
)

type Node struct {
	Value byte
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *Stack) Peek() *Node {
	if s.count == 0 {
		return nil
	}
	return s.nodes[s.count-1]
}

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func compare(a, b byte) bool {
	if Abs(rune(a)-rune(b)) == 32 {
		return true
	}
	return false
}

func readFile(filePath string, contentChan chan []byte) {
	file, _ := os.Open(filePath)
	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	for n > 0 {
		contentChan <- buf[:n]
		n, _ = file.Read(buf)
	}
	close(contentChan)
	// buf, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	// }
	// return string(buf)
}

func react(b byte, s *Stack) {
	if compare(s.Peek().Value, b) {
		s.Pop()
		// beginReact = true
		// bFront = b
	} else {
		s.Push(&Node{b})
	}

}

func isDropUnit(b byte, tagetUnit byte) bool {
	if b == tagetUnit || b == (tagetUnit-32) {
		return true
	}
	return false
}

func part1() {
	filePath := "./input.txt"
	// contentChan := make(chan []byte)
	// go readFile(filePath, contentChan)

	targerStr := "abcdefghijklmnopqrstuvwxyz"
	targerUnits := []byte(targerStr)

	file, _ := os.Open(filePath)
	buf := make([]byte, 102400)
	n, _ := file.Read(buf)
	defer file.Close()
	for _, targerUnit := range targerUnits {
		s := NewStack()
		s.Push(&Node{0})
		for _, b := range buf[:n] {
			if isDropUnit(b, targerUnit) {
				continue
			}
			react(b, s)
			// contentChan <- buf[:n]
		}
		fmt.Println(s.count - 2)
	}

}

func main() {
	// s := NewStack()
	// s.Push(&Node{3})
	// s.Push(&Node{5})
	// s.Push(&Node{7})
	// s.Push(&Node{9})
	// fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())

	//fmt.Println(rune('A') - rune('a'))
	part1()
}
