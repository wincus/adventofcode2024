package day6

import (
	"log/slog"

	"github.com/wincus/adventofcode2024/internal/common"
)

// Solve returns the solutions for day 6
func Solve(s []string, p common.Part) int {

	b := common.ParseRune(s)

	g := findGuard(b)

	if g.Direction == common.Unspecified {
		slog.Error("guard not found")
		return 0
	}

	if p == common.Part1 {

		if loop := walk(b, g); loop {
			return 0
		}

		return len(b.GetVisited())
	}

	if p == common.Part2 {

		var total int

		for y := 0; y < b.GetDimension().N; y++ {
			for x := 0; x < b.GetDimension().M; x++ {

				// skip positions with obstacles
				if v, _ := b.Get(common.Position{X: x, Y: y}); v == '#' {
					continue
				}

				// skip the guard position
				if (common.Position{X: x, Y: y} == g.Position) {
					continue
				}

				override := map[common.Position]rune{
					{X: x, Y: y}: 'O',
				}

				b.ResetPaths() // reset paths

				if loop := walk(b.WithOverrides(override), g); loop {
					total++
				}
			}
		}

		return total
	}

	return 0

}

func findGuard(b common.Board[rune]) common.PositionWithDirection {

	for y := 0; y < b.GetDimension().N; y++ {
		for x := 0; x < b.GetDimension().M; x++ {
			v, _ := b.Get(common.Position{X: x, Y: y})
			if v == '^' {
				return common.PositionWithDirection{
					Position:  common.Position{X: x, Y: y},
					Direction: common.Up,
				}
			}

			if v == 'v' {
				return common.PositionWithDirection{
					Position:  common.Position{X: x, Y: y},
					Direction: common.Down,
				}
			}

			if v == '>' {
				return common.PositionWithDirection{
					Position:  common.Position{X: x, Y: y},
					Direction: common.Right,
				}
			}

			if v == '<' {
				return common.PositionWithDirection{
					Position:  common.Position{X: x, Y: y},
					Direction: common.Left,
				}
			}
		}
	}

	return common.PositionWithDirection{}
}

// walk the board, returns true if the path is a loop
func walk(b common.Board[rune], g common.PositionWithDirection) bool {

	// walk the board
	for {

		// current position and direction
		cp := common.PositionWithDirection{
			Position:  g.Position,
			Direction: g.Direction,
		}

		// get the next position
		err := b.VisitPath(cp)

		// if we are out of bounds, return false
		if err == common.ErrOutOfBounds {
			return false
		}

		// if we have more than one path, return true
		// means we are in a loop
		if b.GetPaths(cp) > 1 {
			return true
		}

		n := common.GetNeighbours(b.GetDimension(), g.Position)

		v, ok := n[g.Direction]

		if !ok {
			break
		}

		r, err := b.Get(v)

		if err != nil {
			break
		}

		switch r {

		case '#':
			g.Direction = g.Direction.TurnRight()
		case 'O':
			g.Direction = g.Direction.TurnRight()

		default:
			g.Position = v
		}
	}

	return false

}
