package p3

import (
	"embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/leviklein/advent_of_code_2024/common"
)

var (
	//go:embed input.txt input_test.txt input_test_pt2.txt
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

func findNumbers(input []string) [][]int {
	var output [][]int
	for _, s := range input {
		var line []int
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		result := re.FindAllStringSubmatch(s, -1)
		for _, n := range result {
			line = []int{getInt(n[1]), getInt(n[2])}
			output = append(output, line)
		}
	}
	return output
}

func parseInputPart2(input []string) [][]int {
	// var output [][]int
	var substring string

	big := strings.Join(input, "")
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	substring = re.ReplaceAllString(big, "")
	re = regexp.MustCompile(`don't\(\).*`)
	substring = re.ReplaceAllString(substring, "")

	return findNumbers([]string{substring})

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
	parsed := findNumbers(input)
	for _, l := range parsed {
		answer += l[0] * l[1]
	}
	return answer
}

func part2(input []string) int {
	var answer int
	parsed := parseInputPart2(input)
	for _, l := range parsed {
		answer += l[0] * l[1]
	}
	return answer
}
