package p2

import (
	"testing"
)

func TestPart1_TestInput(t *testing.T) {
	got := Solve("part1", "input_test.txt")
	want := 2
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}

func TestPart1_FinalInput(t *testing.T) {
	got := Solve("part1", "input.txt")
	want := 356
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}

func TestPart2_TestInput(t *testing.T) {
	got := Solve("part2", "input_test.txt")
	want := 4
	if got != want {
		t.Errorf("Answer = %d; want %d", got, want)
	}
}
