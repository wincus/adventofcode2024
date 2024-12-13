package day4

import (
	"testing"

	"github.com/wincus/adventofcode2024/internal/common"
)

type Test struct {
	input []string
	p     common.Part
	want  int
}

func TestSolver(t *testing.T) {

	tests := []Test{
		{
			input: []string{
				"..X...",
				".SAMX.",
				".A..A.",
				"XMAS.S",
				".X....",
			},
			p:    common.Part1,
			want: 4,
		},
		{
			input: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			p:    common.Part1,
			want: 18,
		},
		{
			input: []string{
				"....XXMAS.",
				".SAMXMS...",
				"...S..A...",
				"..A.A.MS.X",
				"XMASAMX.MM",
				"X.....XA.A",
				"S.S.S.S.SS",
				".A.A.A.A.A",
				"..M.M.M.MM",
				".X.X.XMASX",
			},
			p:    common.Part1,
			want: 18,
		},
		{
			input: []string{
				".M.S......",
				"..A..MSMS.",
				".M.S.MAA..",
				"..A.ASMSM.",
				".M.S.M....",
				"..........",
				"S.S.S.S.S.",
				".A.A.A.A..",
				"M.M.M.M.M.",
				"..........",
			},
			p:    common.Part2,
			want: 9,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}
