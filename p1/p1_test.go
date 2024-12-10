package p1

import (
	"testing"
)

func TestPart1(t *testing.T) {
	got := Solve("part1", "input_test.txt")
	if got != 11 {
		t.Errorf("Answer = %d; want 11", got)
	}
}

func TestPart2(t *testing.T) {
	got := Solve("part2", "input_test.txt")
	want := 31
	if got != 31 {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}
