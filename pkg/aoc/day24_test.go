package aoc

import "testing"

func TestDay24(t *testing.T) {
	a, _ := Day24("../../data/24.txt")

	if a == 0 {
		t.Fail()
	}
}
