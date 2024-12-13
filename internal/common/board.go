package common

import (
	"errors"
	"log/slog"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	Upleft
	Upright
	Downleft
	Downright
	Unspecified
)

var (
	// Errors
	ErrOutOfBounds = errors.New("out of bounds")
)

type Position struct {
	x, y int
}

type Dimension struct {
	n, m int // n rows, m columns
}

type Board[T any] struct {
	pos    Position
	grid   [][]T
	visits map[Position]int
}

func (b *Board[T]) Get(p Position) (T, error) {

	var zero T

	d := b.GetDimension()

	if !CheckPos(d, p) {
		return zero, ErrOutOfBounds
	}

	return b.grid[p.y][p.x], nil
}

func (b *Board[T]) Set(p Position, v T) error {

	d := b.GetDimension()

	if !CheckPos(d, p) {
		return ErrOutOfBounds
	}

	b.grid[p.y][p.x] = v

	return nil
}

func (b *Board[T]) Visit(p Position) error {

	d := b.GetDimension()

	if !CheckPos(d, p) {
		return ErrOutOfBounds
	}

	if _, ok := b.visits[p]; !ok {
		b.visits[p] = 1
	} else {
		b.visits[p]++
	}

	b.pos = p

	return nil

}

// GetUnivisted returns a slice of positions that have not been visited
func (b *Board[T]) GetUnvisited() []Position {

	d := b.GetDimension()

	unvisited := make([]Position, 0)

	for i := 0; i < d.m; i++ {
		for j := 0; j < d.n; j++ {
			if _, ok := b.visits[Position{i, j}]; !ok {
				unvisited = append(unvisited, Position{i, j})
			}
		}
	}

	return unvisited

}

func (b *Board[T]) GetDimension() Dimension {

	if len(b.grid) == 0 {
		return Dimension{0, 0}
	}

	return Dimension{len(b.grid), len(b.grid[0])}
}

func (b *Board[T]) GetPosition() Position {
	return b.pos
}

func (d Direction) String() string {
	return [...]string{"up", "down", "left", "right", "upleft", "upright", "downleft", "downright", "unspecified"}[d]
}

func RemoveEmpty(s []string) []string {

	var r []string

	for _, line := range s {
		if len(line) > 0 {
			r = append(r, line)
		}
	}

	return r

}

// ParseRune returns a Board with the given grid of
// runes. Starts at position 0, 0 with no visits
func ParseRune(s []string) Board[rune] {

	b := Board[rune]{
		grid:   make([][]rune, len(RemoveEmpty(s))),
		pos:    Position{0, 0},
		visits: make(map[Position]int),
	}

	for i, line := range s {

		if len(line) == 0 {
			continue
		}

		b.grid[i] = make([]rune, len(line))
		for j, r := range line {
			b.grid[i][j] = r
		}
	}

	slog.Debug("Parsed board", "board size", b.GetDimension())

	return b

}

// GetNeighbours returns the positions of neighbours given
// a grid of N x M and a position p
func GetNeighbours(d Dimension, p Position) map[Direction]Position {

	neighbours := make(map[Direction]Position)

	neighbours[Up] = Position{p.x, p.y - 1}
	neighbours[Down] = Position{p.x, p.y + 1}
	neighbours[Left] = Position{p.x - 1, p.y}
	neighbours[Right] = Position{p.x + 1, p.y}
	neighbours[Upleft] = Position{p.x - 1, p.y - 1}
	neighbours[Upright] = Position{p.x + 1, p.y - 1}
	neighbours[Downleft] = Position{p.x - 1, p.y + 1}
	neighbours[Downright] = Position{p.x + 1, p.y + 1}

	for direction, p := range neighbours {
		if !CheckPos(d, p) {
			delete(neighbours, direction)
		}
	}

	return neighbours

}

// CheckPos returns true if the position is within the bounds
func CheckPos(d Dimension, p Position) bool {

	if p.x < 0 || p.x >= d.m {
		return false
	}

	if p.y < 0 || p.y >= d.n {
		return false
	}

	return true
}
