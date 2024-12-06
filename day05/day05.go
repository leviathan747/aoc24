package day05

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day05() {
	input := input.GetInput("./day05/day05_input.txt")
	orderingRules, updates := ParseInput(input)
	validUpdates := []Update{}
	for i := 0; i < len(updates); i++ {
		if UpdateIsValid(updates[i], orderingRules) {
			validUpdates = append(validUpdates, updates[i])
		}
	}
	sum := SumMiddles(validUpdates)
	fmt.Println(sum)
}

type OrderingRule struct {
	prev int
	next int
}

type Update []int

func ParseInput(input string) ([]OrderingRule, []Update) {
	orderingRules := []OrderingRule{}
	updates := []Update{}

	// regular expression for ordering rules
	re := regexp.MustCompile(`([1-9][0-9]*)\|([1-9][0-9]*)`)

	// function

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindStringSubmatch(line)
		if m != nil {
			l, _ := strconv.Atoi(m[1])
			r, _ := strconv.Atoi(m[2])
			orderingRules = append(orderingRules, OrderingRule{l, r})
		} else if strings.ContainsRune(line, ',') {
			f := strings.FieldsFunc(line, func(c rune) bool { return c == ',' })
			update := Update{}
			for i := 0; i < len(f); i++ {
				val, _ := strconv.Atoi(f[i])
				update = append(update, val)
			}
			updates = append(updates, update)
		}
	}

	return orderingRules, updates
}

func UpdateIsValid(update Update, rules []OrderingRule) bool {
	// for each page in the update
	for i := 0; i < len(update); i++ {
		// check "must precede" violations
		mustPrecede := MustPrecede(update[i], rules)
		for j := 0; j < i; j++ {
			if slices.Contains(mustPrecede, update[j]) {
				return false
			}
		}
		// check "must follow" violations
		mustFollow := MustFollow(update[i], rules)
		for k := i + 1; k < len(update); k++ {
			if slices.Contains(mustFollow, update[k]) {
				return false
			}
		}
	}
	return true
}

func MustPrecede(value int, rules []OrderingRule) []int {
	mustPrecede := []int{}
	for i := 0; i < len(rules); i++ {
		if value == rules[i].prev {
			if !slices.Contains(mustPrecede, rules[i].next) {
				mustPrecede = append(mustPrecede, rules[i].next)
			}
		}
	}
	return mustPrecede
}

func MustFollow(value int, rules []OrderingRule) []int {
	mustFollow := []int{}
	for i := 0; i < len(rules); i++ {
		if value == rules[i].next {
			if !slices.Contains(mustFollow, rules[i].prev) {
				mustFollow = append(mustFollow, rules[i].prev)
			}
		}
	}
	return mustFollow
}

func SumMiddles(updates []Update) int {
	sum := 0
	for i := 0; i < len(updates); i++ {
		if len(updates[i]) > 0 {
			sum += updates[i][len(updates[i])/2]
		}
	}
	return sum
}
