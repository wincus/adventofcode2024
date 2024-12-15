package common

import (
	"errors"
	"log/slog"
)

type Direction int

const (
	Unspecified Direction = iota
	Up
	Down
	Left
	Right
	Upleft
	Upright
	Downleft
	Downright
)

var (
	// Errors
	ErrOutOfBounds = errors.New("out of bounds")
)

type Position struct {
	X, Y int // X column, Y row
}

type PositionWithDirection struct {
	Position
	Direction
}

type Dimension struct {
	N, M int // N rows, M columns
}

type Board[T any] struct {
	pos       Position
	grid      [][]T
	visits    map[Position]int
	paths     map[PositionWithDirection]int
	overrides map[Position]T
}

func (b *Board[T]) Get(p Position) (T, error) {

	var zero T

	d := b.GetDimension()

	if !CheckPos(d, p) {
		return zero, ErrOutOfBounds
	}

	if v, ok := b.overrides[p]; ok {
		return v, nil
	}

	return b.grid[p.Y][p.X], nil
}

func (b *Board[T]) Set(p Position, v T) error {

	d := b.GetDimension()

	if !CheckPos(d, p) {
		return ErrOutOfBounds
	}

	b.grid[p.Y][p.X] = v

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

// VisitPath visits a position and increments the path count
// for the given direction.
// It will increment the position count.
func (b *Board[T]) VisitPath(p PositionWithDirection) error {

	d := b.GetDimension()

	if !CheckPos(d, p.Position) {
		return ErrOutOfBounds
	}

	if _, ok := b.paths[p]; !ok {
		b.paths[p] = 1
	} else {
		b.paths[p]++
	}

	b.pos = p.Position

	// register the position
	return b.Visit(p.Position)
}

func (b *Board[T]) GetVisits(p Position) int {
	return b.visits[p]
}

func (b *Board[T]) ResetVisits() {
	b.visits = make(map[Position]int)
}

func (b *Board[T]) GetPaths(p PositionWithDirection) int {
	return b.paths[p]
}

func (b *Board[T]) ResetPaths() {
	b.paths = make(map[PositionWithDirection]int)
}

// GetUnivisted returns a slice of positions that have not been visited
// in the given directions
func (b *Board[T]) GetUnvisited() []Position {

	unvisited := make([]Position, 0)

	for i := 0; i < b.GetDimension().N; i++ {
		for j := 0; j < b.GetDimension().M; j++ {
			p := Position{j, i}
			if _, ok := b.visits[p]; !ok {
				unvisited = append(unvisited, p)
			}
		}
	}

	return unvisited

}

func (b *Board[T]) GetVisited() []Position {

	visited := make([]Position, 0)

	for p := range b.visits {
		visited = append(visited, p)
	}

	return visited
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
	return [...]string{"Unspecified", "Up", "Down", "Left", "Right", "Upleft", "Upright", "Downleft", "Downright"}[d]
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

func (b *Board[T]) WithOverrides(overrides map[Position]T) Board[T] {
	b.overrides = overrides
	return *b
}

// ParseRune returns a Board with the given grid of
// runes. Starts at position 0, 0 with no visits
func ParseRune(s []string) Board[rune] {

	b := Board[rune]{
		grid:   make([][]rune, len(RemoveEmpty(s))),
		pos:    Position{0, 0},
		visits: make(map[Position]int),
		paths:  make(map[PositionWithDirection]int),
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

	neighbours[Up] = Position{p.X, p.Y - 1}
	neighbours[Down] = Position{p.X, p.Y + 1}
	neighbours[Left] = Position{p.X - 1, p.Y}
	neighbours[Right] = Position{p.X + 1, p.Y}
	neighbours[Upleft] = Position{p.X - 1, p.Y - 1}
	neighbours[Upright] = Position{p.X + 1, p.Y - 1}
	neighbours[Downleft] = Position{p.X - 1, p.Y + 1}
	neighbours[Downright] = Position{p.X + 1, p.Y + 1}

	for direction, p := range neighbours {
		if !CheckPos(d, p) {
			delete(neighbours, direction)
		}
	}

	return neighbours

}

// CheckPos returns true if the position is within the bounds
func CheckPos(d Dimension, p Position) bool {

	if p.X < 0 || p.X >= d.M {
		return false
	}

	if p.Y < 0 || p.Y >= d.N {
		return false
	}

	return true
}

func (d Direction) TurnRight() Direction {

	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}

	return Unspecified
}

func (d Direction) TurnLeft() Direction {

	switch d {
	case Up:
		return Left
	case Left:
		return Down
	case Down:
		return Right
	case Right:
		return Up
	}

	return Unspecified
}
