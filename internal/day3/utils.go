package day3

import (
	"log/slog"
	"regexp"
	"strconv"

	"github.com/wincus/adventofcode2024/internal/common"
)

type OP int

const (
	UNKNOWN OP = iota
	MUL
	DO
	DONT
)

type instructions struct {
	expressions []expr
}

type expr struct {
	op OP
	a  int
	b  int
}

// Solve returns the solutions for day 3
func Solve(s []string, p common.Part) int {

	var total int

	var disabled bool

	for _, line := range s {

		if len(line) == 0 {
			continue
		}

		i := parse(line)

		for _, e := range i.expressions {

			switch e.op {
			case DO:
				disabled = false
			case DONT:
				disabled = true
			case MUL:

				// part 1 evaluates all expressions
				if p == common.Part1 {
					total += e.eval()
				}

				// part 2 only evaluates expressions if the instructions are enabled
				if p == common.Part2 && !disabled {
					total += e.eval()
				}
			}
		}
	}

	return total
}

func parse(s string) instructions {

	// instructions are enabled by default
	var i instructions

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := re.FindAll([]byte(s), -1)

	for _, match := range matches {

		expr := getExpression(string(match))

		if expr.op != UNKNOWN {
			i.expressions = append(i.expressions, expr)
		}
	}

	return i
}

func getExpression(s string) expr {

	var re *regexp.Regexp

	re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	if re.MatchString(s) {

		matches := re.FindStringSubmatch(s)

		a, err := strconv.Atoi(matches[1])

		if err != nil {
			slog.Warn("Error converting string to int")
			return expr{
				op: UNKNOWN,
			}
		}

		b, err := strconv.Atoi(matches[2])

		if err != nil {
			slog.Warn("Error converting string to int")
			return expr{
				op: UNKNOWN,
			}
		}

		return expr{
			op: MUL,
			a:  a,
			b:  b,
		}
	}

	re = regexp.MustCompile(`do\(\)`)

	if re.MatchString(s) {
		return expr{
			op: DO,
		}
	}

	re = regexp.MustCompile(`don't\(\)`)
	if re.MatchString(s) {
		return expr{
			op: DONT,
		}
	}

	return expr{
		op: UNKNOWN,
	}

}

func (e expr) eval() int {

	if e.op == MUL {
		return e.a * e.b
	}

	return 0

}
