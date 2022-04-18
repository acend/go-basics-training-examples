package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Direction int

func (d Direction) String() string {
	switch d {
	case UP:
		return "up"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	default:
		panic("invalid direction")
	}
}

type Move struct {
	Direction Direction
	Steps     int
}

func (m *Move) String() string {
	return fmt.Sprintf("%s %d", m.Direction, m.Steps)
}

type Generator struct {
	rand          *rand.Rand
	totalCommands int
	maxSteps      int
}

func (g *Generator) Generate() []Move {
	moves := []Move{}
	prev := Direction(0)
	for i := 0; i < g.totalCommands; i++ {
		direction := Direction(rand.Intn(4))
		if direction == prev {
			direction = (direction + 1) % 4
		}
		prev = direction
		steps := g.rand.Intn(5) + 1
		moves = append(moves, Move{
			Direction: direction,
			Steps:     steps,
		})
	}
	return moves
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Robot struct {
	position  Position
	maxUp     int
	maxDown   int
	maxLeft   int
	maxRight  int
	positions map[Position]int
}

func (r *Robot) Summary() string {
	out := &bytes.Buffer{}

	fmt.Fprintf(out, "end position: %v\n", r.position)

	fmt.Fprintf(out, "max Y: %d\n", r.maxUp)
	fmt.Fprintf(out, "min Y: %d\n", r.maxDown)
	fmt.Fprintf(out, "max X: %d\n", r.maxRight)
	fmt.Fprintf(out, "min X: %d\n", r.maxLeft)

	pos, count := r.MaxPositions()
	fmt.Fprintf(out, "most visited positions (%d times): %v", count, pos)

	return out.String()
}

func (r *Robot) Run(moves []Move) {
	for _, move := range moves {
		r.Move(move)
		r.addPosition(r.position)
	}
}

func (r *Robot) addPosition(p Position) {
	if r.positions == nil {
		r.positions = make(map[Position]int)
	}
	_, ok := r.positions[p]
	if ok {
		r.positions[p] += 1
	} else {
		r.positions[p] = 1
	}
}

// MaxPositions returns the most often visited position and the count
func (r *Robot) MaxPositions() ([]Position, int) {
	maxPos := []Position{}
	maxCount := 0
	for pos, count := range r.positions {
		if count > maxCount {
			maxPos = nil
			maxCount = count
			maxPos = append(maxPos, pos)
		} else if count == maxCount {
			maxPos = append(maxPos, pos)
		}
	}
	return maxPos, maxCount
}

func (r *Robot) Move(m Move) {
	switch m.Direction {
	case UP:
		r.position.Y += m.Steps
		r.maxUp = max(r.maxUp, r.position.Y)
	case RIGHT:
		r.position.X += m.Steps
		r.maxRight = max(r.maxRight, r.position.X)
	case DOWN:
		r.position.Y -= m.Steps
		r.maxDown = min(r.maxDown, r.position.Y)
	case LEFT:
		r.position.X -= m.Steps
		r.maxLeft = min(r.maxLeft, r.position.X)
	default:
		panic("invalid direction")
	}

}

type Position struct {
	X int
	Y int
}

func readMoves(file string) ([]Move, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fd)
	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		//line = strings.TrimSpace(line)
		command, rawSteps, found := strings.Cut(line, " ")
		if !found {
			return nil, fmt.Errorf("invalid line: '%s'", line)
		}

		var direction Direction
		switch command {
		case "up":
			direction = UP
		case "right":
			direction = RIGHT
		case "down":
			direction = DOWN
		case "left":
			direction = LEFT
		default:
			return nil, fmt.Errorf("invalid command '%s' in line: '%s'", command, line)
		}
		steps, err := strconv.Atoi(rawSteps)
		if err != nil {
			return nil, fmt.Errorf("invalid steps '%s' in line: '%s'", rawSteps, line)
		}
		moves = append(moves, Move{direction, steps})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}

func main() {
	seed := flag.Int64("seed", 7, "seed value")
	flag.Parse()
	action := flag.Arg(0)

	switch action {
	case "generate":
		numberOfCommands, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		g := &Generator{
			rand:          rand.New(rand.NewSource(*seed)),
			totalCommands: numberOfCommands,
			maxSteps:      5,
		}
		moves := g.Generate()
		for _, move := range moves {
			fmt.Fprintf(os.Stdout, "%s\n", move.String())
		}

	case "solve":
		moves, err := readMoves(flag.Arg(1))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		r := &Robot{}
		r.Run(moves)
		fmt.Println(r.Summary())

	default:
		fmt.Println("unknown action:", action)
		os.Exit(1)
	}
}
