package template

import (
	"container/heap"
	"embed"
	// "fmt"
	// "math"
	"strconv"
	"strings"

	"github.com/leviklein/advent_of_code_2024/common"
)

var (
	//go:embed input.txt input_test.txt
	f embed.FS
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func Solve(part, filename string) any {
	input := common.GetInput(filename, f)
	switch part {
	case "part1":
		return part1(input)
	case "part2":
		return part2(input)
	default:
		return -1
	}
}

func part1(input []string) int {
	b1 := &IntHeap{}
	b2 := &IntHeap{}
	heap.Init(b1)
	heap.Init(b2)
	for _, s := range input {
		// fmt.Println(i, s)
		result := strings.Split(s, "   ")
		if s, err := strconv.Atoi(result[0]); err == nil {
			heap.Push(b1, s)
		}
		if s, err := strconv.Atoi(result[1]); err == nil {
			heap.Push(b2, s)
		}
	}
	var answer int = 0
	for b1.Len() > 0 {
		n1, n2 := heap.Pop(b1).(int), heap.Pop(b2).(int)
		answer += abs(n1 - n2)
	}
	return answer
}

func part2(input []string) int {
	b1 := &IntHeap{}
	b2 := &IntHeap{}
	heap.Init(b1)
	heap.Init(b2)
	for _, s := range input {
		result := strings.Split(s, "   ")
		if s, err := strconv.Atoi(result[0]); err == nil {
			heap.Push(b1, s)
		}
		if s, err := strconv.Atoi(result[1]); err == nil {
			heap.Push(b2, s)
		}
	}

	map2 := make(map[int]int)
	for b2.Len() > 0 {
		n := heap.Pop(b2).(int)
		map2[n] += 1
	}

	var answer int
	for b1.Len() > 0 {
		n1 := heap.Pop(b1).(int)
		answer += (n1 * map2[n1])
	}
	return answer
}
