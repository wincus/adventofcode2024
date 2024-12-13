package day4

import (
	"github.com/wincus/adventofcode2024/internal/common"
)

var XMAS = [4]rune{'X', 'M', 'A', 'S'}

// Solve returns the solutions for day 4
func Solve(s []string, p common.Part) int {

	b := common.ParseRune(s)

	var total int

	for {
		v := b.GetUnvisited()

		if len(v) == 0 {
			break
		}

		b.Visit(v[0])

		if p == common.Part1 {
			total += countXMAS(b, v[0], 0, 0, common.Unspecified)
		}

		if p == common.Part2 {
			if isMAS(b) {
				total++
			}
		}
	}

	return total

}

func countXMAS(board common.Board[rune], position common.Position, iteration int, count int, from common.Direction) int {

	if iteration > len(XMAS)-1 {
		return count
	}

	v, _ := board.Get(position)

	if v != XMAS[iteration] {
		return 0
	}

	if iteration == len(XMAS)-1 {
		return 1
	}

	var sum int

	n := common.GetNeighbours(board.GetDimension(), position)

	for direction, neighbour := range n {
		if from == common.Unspecified || from == direction {
			sum += countXMAS(board, neighbour, iteration+1, count, direction)
		}
	}

	return sum
}

func isMAS(board common.Board[rune]) bool {

	v, err := board.Get(board.GetPosition())

	if err != nil {
		return false
	}

	if v != XMAS[2] {
		return false
	}

	n := common.GetNeighbours(board.GetDimension(), board.GetPosition())

	count := make(map[rune]int)
	simetry := make(map[common.Direction]rune)

	for _, direction := range []common.Direction{common.Upleft, common.Downleft, common.Upright, common.Downright} {
		if neighbour, ok := n[direction]; ok {
			if v, err := board.Get(neighbour); err == nil {
				count[v]++
				simetry[direction] = v
			}
		}
	}

	if count[XMAS[1]] != 2 || count[XMAS[3]] != 2 {
		return false
	}

	if simetry[common.Upleft] == simetry[common.Downright] {
		return false
	}

	return true
}
