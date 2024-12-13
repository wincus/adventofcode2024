package common

import (
	"reflect"
	"testing"
)

func TestBoardGet(t *testing.T) {

	type test struct {
		b    Board[int]
		p    Position
		want int
		err  error
	}

	tests := []test{
		{
			b: Board[int]{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			p:    Position{0, 0},
			want: 1,
			err:  nil,
		},
		{
			b: Board[int]{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			p:    Position{1, 0},
			want: 2,
			err:  nil,
		},
		{
			b: Board[int]{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			p:    Position{0, 1},
			want: 3,
			err:  nil,
		},
		{
			b: Board[int]{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			p:    Position{1, 1},
			want: 4,
			err:  nil,
		},
		{
			b: Board[int]{
				grid: [][]int{
					{1, 2},
					{3, 4},
				},
			},
			p:    Position{2, 0},
			want: 0,
			err:  ErrOutOfBounds,
		},
	}

	for _, test := range tests {
		got, err := test.b.Get(test.p)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}

		if err != test.err {
			t.Errorf("got %v, want %v", err, test.err)
		}
	}
}

func TestParseRune(t *testing.T) {

	type test struct {
		input []string
		grid  [][]rune
	}

	tests := []test{
		{
			input: []string{
				"ABC",
				"DEF",
			},
			grid: [][]rune{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
		},
		{
			input: []string{
				"AB",
			},
			grid: [][]rune{
				{'A', 'B'},
			},
		},
		{
			input: []string{
				"A",
				"B",
			},
			grid: [][]rune{
				{'A'},
				{'B'},
			},
		},
	}

	for _, test := range tests {
		b := ParseRune(test.input)

		if !reflect.DeepEqual(b.grid, test.grid) {
			t.Errorf("got %v, want %v", b.grid, test.grid)
		}
	}

}

func TestGetNeighbours(t *testing.T) {

	type test struct {
		name       string
		dimension  Dimension
		position   Position
		neighbours map[Direction]Position
	}

	tests := []test{
		{
			name:       "0x0",
			dimension:  Dimension{0, 0},
			position:   Position{0, 0},
			neighbours: map[Direction]Position{},
		},
		{
			name:       "1x1",
			dimension:  Dimension{1, 1},
			position:   Position{0, 0},
			neighbours: map[Direction]Position{},
		},
		{
			name:      "1x2",
			dimension: Dimension{1, 2},
			position:  Position{0, 0},
			neighbours: map[Direction]Position{
				Right: {1, 0},
			},
		},
		{
			name:      "2x1",
			dimension: Dimension{2, 1},
			position:  Position{0, 0},
			neighbours: map[Direction]Position{
				Down: {0, 1},
			},
		},
		{
			name:      "2x2 top left",
			dimension: Dimension{2, 2},
			position:  Position{0, 0},
			neighbours: map[Direction]Position{
				Right:     {1, 0},
				Down:      {0, 1},
				Downright: {1, 1},
			},
		},
		{
			name:      "2x2 top right",
			dimension: Dimension{2, 2},
			position:  Position{1, 0},
			neighbours: map[Direction]Position{
				Left:     {0, 0},
				Down:     {1, 1},
				Downleft: {0, 1},
			},
		},
		{
			name:      "2x2 bottom left",
			dimension: Dimension{2, 2},
			position:  Position{0, 1},
			neighbours: map[Direction]Position{
				Right:   {1, 1},
				Up:      {0, 0},
				Upright: {1, 0},
			},
		},
		{
			name:      "2x2 bottom right",
			dimension: Dimension{2, 2},
			position:  Position{1, 1},
			neighbours: map[Direction]Position{
				Left:   {0, 1},
				Up:     {1, 0},
				Upleft: {0, 0},
			},
		},
		{
			name:      "3x3 top left",
			dimension: Dimension{3, 3},
			position:  Position{0, 0},
			neighbours: map[Direction]Position{
				Right:     {1, 0},
				Down:      {0, 1},
				Downright: {1, 1},
			},
		},
		{
			name:      "3x3 top right",
			dimension: Dimension{3, 3},
			position:  Position{2, 0},
			neighbours: map[Direction]Position{
				Left:     {1, 0},
				Down:     {2, 1},
				Downleft: {1, 1},
			},
		},
		{
			name:      "3x3 bottom left",
			dimension: Dimension{3, 3},
			position:  Position{0, 2},
			neighbours: map[Direction]Position{
				Right:   {1, 2},
				Up:      {0, 1},
				Upright: {1, 1},
			},
		},
		{
			name:      "3x3 bottom right",
			dimension: Dimension{3, 3},
			position:  Position{2, 2},
			neighbours: map[Direction]Position{
				Left:   {1, 2},
				Up:     {2, 1},
				Upleft: {1, 1},
			},
		},
		{
			name:      "3x3 middle",
			dimension: Dimension{3, 3},
			position:  Position{1, 1},
			neighbours: map[Direction]Position{
				Left:      {0, 1},
				Right:     {2, 1},
				Up:        {1, 0},
				Down:      {1, 2},
				Upleft:    {0, 0},
				Upright:   {2, 0},
				Downleft:  {0, 2},
				Downright: {2, 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			neighbours := GetNeighbours(test.dimension, test.position)

			if !reflect.DeepEqual(neighbours, test.neighbours) {
				t.Errorf("got %v, want %v", neighbours, test.neighbours)
			}
		})
	}
}
