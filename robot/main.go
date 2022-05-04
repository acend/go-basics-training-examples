package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

type Robot struct {
	Pos       Position
	MaxRight  int
	MaxLeft   int
	Positions map[Position]int
}

func NewRobot() *Robot {
	return &Robot{
		Positions: map[Position]int{},
	}
}

func (r *Robot) Move(i Instruction) {
	r.Pos.Move(i)
	if r.Pos.X > r.MaxRight {
		r.MaxRight = r.Pos.X
	}
	if r.Pos.X < r.MaxLeft {
		r.MaxLeft = r.Pos.X
	}
	r.addPos(r.Pos)
}

func (r *Robot) MaxVisited() ([]Position, int) {
	maxCount := 0
	positions := []Position{}

	for pos, count := range r.Positions {
		if count > maxCount {
			maxCount = count
			positions = []Position{pos}
		} else if count == maxCount {
			positions = append(positions, pos)
		}
	}

	return positions, maxCount
}

func (r *Robot) addPos(p Position) {
	_, ok := r.Positions[p]
	if !ok {
		r.Positions[p] = 1
		return
	}

	r.Positions[p] += 1
}

type Position struct {
	X int
	Y int
}

func (p *Position) Move(i Instruction) {
	switch i.Direction {
	case UP:
		p.Y += i.Distance
	case RIGHT:
		p.X += i.Distance
	case DOWN:
		p.Y -= i.Distance
	case LEFT:
		p.X -= i.Distance
	default:
		panic("invalid direction")
	}
}

type Instruction struct {
	Direction int
	Distance  int
}

func parseInstruction(input string) (*Instruction, error) {
	fields := strings.Fields(input)
	if len(fields) != 2 {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	direction, err := strToDirection(fields[0])
	if err != nil {
		return nil, err
	}
	distance, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	}

	return &Instruction{
		Direction: direction,
		Distance:  distance,
	}, nil
}

func readInstructions(fileName string) ([]Instruction, error) {
	rawData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	instructions := []Instruction{}
	rawStrData := string(rawData)
	rawStrData = strings.TrimSpace(rawStrData)
	for i, line := range strings.Split(rawStrData, "\n") {
		instruction, err := parseInstruction(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line %d '%s': %w", i+1, line, err)
		}
		instructions = append(instructions, *instruction)
	}
	return instructions, nil
}

func strToDirection(name string) (int, error) {
	switch name {
	case "up":
		return UP, nil
	case "right":
		return RIGHT, nil
	case "down":
		return DOWN, nil
	case "left":
		return LEFT, nil
	default:
		return -1, fmt.Errorf("invalid instruction '%s'", name)
	}
}

func run() error {
	flag.Parse()
	inputFile := flag.Arg(0)
	if inputFile == "" {
		inputFile = "input.txt"
	}

	instructions, err := readInstructions(inputFile)
	if err != nil {
		return err
	}

	r := NewRobot()
	for _, i := range instructions {
		r.Move(i)
	}
	fmt.Println("end position", r.Pos)
	fmt.Println("most right", r.MaxRight)
	fmt.Println("most left", r.MaxLeft)

	maxVisited, count := r.MaxVisited()
	fmt.Printf("most visited %v (%d times)\n", maxVisited, count)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
