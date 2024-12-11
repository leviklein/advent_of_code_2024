package template

import (
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
	for _, s := range input {
		// fmt.Println(i, s)
		result := strings.Split(s, "   ")
		if _, err := strconv.Atoi(result[0]); err == nil {
		}
		if _, err := strconv.Atoi(result[1]); err == nil {
		}
	}
	var answer int = 0

	return answer
}

func part2(input []string) int {
	for _, s := range input {
		// fmt.Println(i, s)
		result := strings.Split(s, "   ")
		if _, err := strconv.Atoi(result[0]); err == nil {
		}
		if _, err := strconv.Atoi(result[1]); err == nil {
		}
	}
	var answer int = 0

	return answer
}
