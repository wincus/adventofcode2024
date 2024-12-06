package day2

import (
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2024/internal/common"
)

const (
	MIN_LEVEL_SEP       = 1
	MAX_LEVEL_SEP       = 3
	MIN_LEVEL_TOLERANCE = 0
	MAX_LEVEL_TOLERANCE = 1
)

type level []int

// Solve returns the solutions for day 2
func Solve(s []string, p common.Part) int {

	var tolerance int

	levels := parse(s)

	switch p {
	case common.Part1:
		tolerance = MIN_LEVEL_TOLERANCE
	case common.Part2:
		tolerance = MAX_LEVEL_TOLERANCE
	}

	var count int

	for _, l := range levels {

		if l.isSafe(tolerance) {
			count++
		}
	}

	return count

}

func parse(s []string) []level {

	var levels []level

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		levelStr := strings.Split(line, " ")

		level := make([]int, len(levelStr))

		var err error

		for i, l := range levelStr {
			level[i], err = strconv.Atoi(l)

			if err != nil {
				panic(err)
			}
		}

		levels = append(levels, level)
	}

	return levels

}

func (l level) isSafe(tolerance int) bool {

	if l.isAsc(tolerance) || l.isDesc(tolerance) {
		return true
	}

	return false

}

func (l level) isAsc(tolerance int) bool {

	for i := 1; i < len(l); i++ {

		diff := l[i] - l[i-1]

		if diff < 0 || diff < MIN_LEVEL_SEP || diff > MAX_LEVEL_SEP {

			if tolerance == 0 {
				return false
			}

			if l.remove(i - 1).isAsc(tolerance - 1) {
				return true
			}

			if l.remove(i).isAsc(tolerance - 1) {
				return true
			}

			return false
		}
	}

	return true
}

func (l level) isDesc(tolerance int) bool {

	for i := 1; i < len(l); i++ {

		diff := l[i-1] - l[i]

		if diff < 0 || diff < MIN_LEVEL_SEP || diff > MAX_LEVEL_SEP {

			if tolerance == 0 {
				return false
			}

			if l.remove(i - 1).isDesc(tolerance - 1) {
				return true
			}

			if l.remove(i).isDesc(tolerance - 1) {
				return true
			}

			return false
		}
	}

	return true
}

func (l level) remove(i int) level {
	if i < 0 || i >= len(l) {
		panic("index out of bounds")
	}

	newLevel := make(level, len(l)-1)
	copy(newLevel, l[:i])
	copy(newLevel[i:], l[i+1:])

	return newLevel
}
