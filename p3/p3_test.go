package p3

import (
	"testing"
)

func TestPart1_TestInput(t *testing.T) {
	got := Solve("part1", "input_test.txt")
	want := 161
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}

func TestPart1_FinalInput(t *testing.T) {
	got := Solve("part1", "input.txt")
	want := 174960292
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}

func TestPart2_TestInput(t *testing.T) {
	got := Solve("part2", "input_test_pt2.txt")
	want := 48
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}
