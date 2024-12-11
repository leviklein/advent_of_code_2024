package p2

import (
	"embed"
	// "math"
	"slices"
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

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	check(err)
	return n
}

func parseInput(input []string) [][]int {
	var output [][]int
	for _, s := range input {
		var line []int
		result := strings.Split(s, " ")
		for _, n := range result {
			line = append(line, getInt(n))
		}
		output = append(output, line)
	}
	return output
}

func checkLine(l []int) bool {
	init := true
	var upperbound, lowerbound int
	for i, n := range l {
		if i < len(l)-1 {
			diff := n - l[i+1]
			if init {
				switch {
				case diff == 0:
					return false
				case diff > 0:
					upperbound = 3
					lowerbound = 1
				case diff < 0:
					upperbound = -1
					lowerbound = -3
				}
				init = false
			}
			if diff <= upperbound && diff >= lowerbound {
				continue
			} else {
				return false
			}
		}
	}
	return true
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
	var answer int
	parsed := parseInput(input)
	for _, l := range parsed {
		stable := checkLine(l)
		if stable {
			answer += 1
		}
	}
	return answer
}

func part2(input []string) int {
	var answer int
	parsed := parseInput(input)
	for _, l := range parsed {
		stable := checkLine(l)
		if !stable {
			for i := range l {
				s1 := l[:i]
				s2 := l[i+1:]
				testArray := slices.Concat(s1, s2)
				if checkLine(testArray) {
					answer += 1
					break
				}
			}
		} else {
			answer += 1
		}
	}
	return answer
}
