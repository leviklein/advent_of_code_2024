package common

import (
	"bufio"
	"embed"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(filename string, f embed.FS) []string {
	dat, err := f.Open(filename)
	check(err)
	scanner := bufio.NewScanner(dat)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
