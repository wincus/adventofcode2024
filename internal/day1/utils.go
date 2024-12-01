package day1

import (
	"log/slog"
	"regexp"
	"sort"
	"strconv"

	"github.com/wincus/adventofcode2024/internal/common"
)

// Solve returns the solutions for day 1
func Solve(s []string, p common.Part) int {

	l1, l2 := parse(s)

	var total int

	for i := 0; i < len(l1); i++ {

		if p == common.Part1 {
			// calculate distance
			distance := dis(l1[i], l2[i])
			total = total + distance
		}

		if p == common.Part2 {
			// calculate similarity
			similarity := sim(l1, l2, i)
			total = total + similarity
		}
	}

	return total
}

func parse(s []string) ([]int, []int) {
	var l1, l2 []int

	re := regexp.MustCompile(`^(\d+)\s+(\d+)$`)

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		res := re.FindStringSubmatch(line)

		if len(res) != 3 {
			slog.Error("Invalid input", "err", "unexpected input", "line", line)
		}

		n1, err := strconv.Atoi(res[1])

		if err != nil {
			slog.Error("Invalid input", "err", err)
		}

		n2, err := strconv.Atoi(res[2])

		if err != nil {
			slog.Error("Invalid input", "err", err)
		}

		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	return l1, l2
}

func dis(a, b int) int {

	if a > b {
		return a - b
	}

	return b - a
}

func sim(l1, l2 []int, j int) int {

	var count int

	for i := 0; i < len(l2); i++ {
		if l1[j] == l2[i] {
			count++
		}
	}

	return count * l1[j]
}
