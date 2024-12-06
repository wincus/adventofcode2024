package day2

import (
	"reflect"
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
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			p:    common.Part1,
			want: 2,
		},
		{
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			p:    common.Part2,
			want: 4,
		},
	}

	for _, test := range tests {
		got := Solve(test.input, test.p)

		if got != test.want {
			t.Errorf("got %v, want %v for part %v", got, test.want, test.p)
		}
	}
}

func TestRemoveLevel(t *testing.T) {

	tests := []struct {
		input  level
		remove int
		want   level
	}{
		{
			input:  level{0, 1, 2, 3, 4},
			remove: 0,
			want:   level{1, 2, 3, 4},
		},
		{
			input:  level{0, 1, 2, 3, 4},
			remove: 1,
			want:   level{0, 2, 3, 4},
		},
		{
			input:  level{0, 1, 2, 3, 4},
			remove: 2,
			want:   level{0, 1, 3, 4},
		},
		{
			input:  level{0, 1, 2, 3, 4},
			remove: 3,
			want:   level{0, 1, 2, 4},
		},
		{
			input:  level{0, 1, 2, 3, 4},
			remove: 4,
			want:   level{0, 1, 2, 3},
		},
	}

	for _, test := range tests {
		got := test.input.remove(test.remove)

		if len(got) != len(test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
