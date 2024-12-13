package day5

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/wincus/adventofcode2024/internal/common"
)

type Rule struct {
	before, after int
}

type Update struct {
	pages []int
}

// Solve returns the solutions for day 5
func Solve(s []string, p common.Part) int {

	rules, updates := split(s)

	r := parseRules(rules)
	u := parseUpdates(updates)

	var orderedUpdates []Update
	var outOfOrder []Update

	for _, update := range u {
		if isInOrder(r, update.pages) {
			orderedUpdates = append(orderedUpdates, update)
		} else {
			outOfOrder = append(outOfOrder, update)
		}
	}

	var total int

	if p == common.Part1 {
		for _, update := range orderedUpdates {
			total += getMiddle(update.pages)
		}

		return total
	}

	if p == common.Part2 {

		for _, update := range fixAll(r, outOfOrder) {
			total += getMiddle(update.pages)
		}

		return total
	}

	return 0
}

func fixAll(rules []Rule, updates []Update) []Update {

	var fixed []Update

	for _, update := range updates {
		fixed = append(fixed, fixUpdate(rules, update))
	}

	return fixed
}

func fixUpdate(rules []Rule, update Update) Update {

	for {

		var tryAgain bool

		for _, rule := range rules {
			if !checkAndFix(rule.before, rule.after, update.pages) {
				tryAgain = true
				break
			}
		}

		if !tryAgain {
			break
		}
	}

	return update
}

func getMiddle(pages []int) int {
	return pages[len(pages)/2]
}

func isInOrder(r []Rule, pages []int) bool {

	for _, rule := range r {
		if !checkAndFix(rule.before, rule.after, pages) {
			return false
		}
	}

	return true

}

func checkAndFix(before, after int, pages []int) bool {

	beforeIndex := -1
	afterIndex := -1

	for i, page := range pages {
		if page == before {
			beforeIndex = i
		}

		if page == after {
			afterIndex = i
		}
	}

	if beforeIndex == -1 || afterIndex == -1 {
		return true
	}

	if beforeIndex < afterIndex {
		return true
	}

	pages[beforeIndex], pages[afterIndex] = pages[afterIndex], pages[beforeIndex]

	return false
}

func split(s []string) ([]string, []string) {

	var prefix []string
	var suffix []string
	var prefixDone bool

	for i, line := range s {

		if line == "" && i != 0 {
			prefixDone = true
			continue
		}

		if prefixDone {
			suffix = append(suffix, line)
		} else {
			prefix = append(prefix, line)
		}
	}

	return prefix, suffix
}

func parseRules(s []string) []Rule {

	var rules []Rule

	for _, line := range s {

		rule := strings.Split(line, "|")

		if len(rule) != 2 {
			slog.Error("Invalid rule", "rule", rule)
			continue
		}

		a, err := strconv.Atoi(rule[0])

		if err != nil {
			slog.Error("Invalid rule", "rule", rule)
			continue
		}

		b, err := strconv.Atoi(rule[1])

		if err != nil {
			slog.Error("Invalid rule", "rule", rule)
			continue
		}

		rules = append(rules, Rule{a, b})
	}

	return rules
}

func parseUpdates(s []string) []Update {

	var updates []Update

	for _, line := range s {

		update := strings.Split(line, ",")

		if len(update) < 2 {
			slog.Error("Invalid update", "update", update)
			continue
		}

		var pages []int

		for _, p := range update {
			page, err := strconv.Atoi(p)

			if err != nil {
				slog.Error("Invalid page", "page", p)
				continue
			}

			pages = append(pages, page)
		}

		updates = append(updates, Update{pages})
	}

	return updates
}
