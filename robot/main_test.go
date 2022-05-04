package main

import (
	"testing"
)

func TestMove(t *testing.T) {
	instructions := []Instruction{
		{RIGHT, 3},
		{UP, 1},
		{LEFT, 5},
		{RIGHT, 4},
		{DOWN, 2},
		{UP, 2},
	}
	r := NewRobot()
	for _, i := range instructions {
		r.Move(i)
	}
	// End position
	gotPos := r.Pos
	expectedPos := Position{2, 1}
	if gotPos != expectedPos {
		t.Errorf("End position incorrect. got: %v expected: %v", gotPos, expectedPos)
	}
	// Max right
	got := r.MaxRight
	expected := 3
	if got != expected {
		t.Errorf("Max right position incorrect. got: %v expected: %v", got, expected)
	}
	// Max left
	got = r.MaxLeft
	expected = -2
	if got != expected {
		t.Errorf("Max left position incorrect. got: %v expected: %v", got, expected)
	}
	// Max visited
	maxVisited, count := r.MaxVisited()
	expMaxVisited := Position{2, 1}
	expCount := 2
	if len(maxVisited) != 1 || maxVisited[0] != expMaxVisited {
		t.Errorf("Max visited position incorrect. got: %v expected: %v", maxVisited, expMaxVisited)
	}
	if count != expCount {
		t.Errorf("Max visited count incorrect. got: %v expected: %v", count, expCount)
	}
}

func TestParseInstruction(t *testing.T) {
	type TestCase struct {
		input    string
		expected *Instruction
	}

	cases := []TestCase{
		{"right 2", &Instruction{RIGHT, 2}},
		{"up 3", &Instruction{UP, 3}},
		{"left 1", &Instruction{LEFT, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			got, err := parseInstruction(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if got.Direction != tc.expected.Direction {
				t.Errorf("Direction not correct. got: %d expected: %d", got.Direction, tc.expected.Direction)
			}
			if got.Distance != tc.expected.Distance {
				t.Errorf("Distance not correct. got: %d expected: %d", got.Distance, tc.expected.Distance)
			}
		})
	}
}
