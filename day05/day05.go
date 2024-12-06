package day05

import (
	"bufio"
	"fmt"
	"leviathan747/aoc24/input"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day05() {
	input := input.GetInput("./day05/day05_input.txt")
	orderingRules, updates := ParseInput(input)
	validUpdates := []Update{}
	invalidUpdates := []Update{}
	for i := 0; i < len(updates); i++ {
		if UpdateIsValid(updates[i], orderingRules) {
			validUpdates = append(validUpdates, updates[i])
		} else {
			invalidUpdates = append(invalidUpdates, updates[i])
		}
	}
	sum := SumMiddles(validUpdates)
	fmt.Println(sum)
	FixInvalidUpdates(invalidUpdates, orderingRules)
	sum = SumMiddles(invalidUpdates)
	fmt.Println(sum)
}

type OrderingRules map[int][]int

type Update []int

func ParseInput(input string) (OrderingRules, []Update) {
	orderingRules := OrderingRules{}
	updates := []Update{}

	// regular expression for ordering rules
	re := regexp.MustCompile(`([1-9][0-9]*)\|([1-9][0-9]*)`)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		m := re.FindStringSubmatch(line)
		if m != nil {
			l, _ := strconv.Atoi(m[1])
			r, _ := strconv.Atoi(m[2])
			_, present := orderingRules[l]
			if present {
				orderingRules[l] = append(orderingRules[l], r)
			} else {
				orderingRules[l] = []int{r}
			}
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

func MustPrecede(update Update, i, j int, rules OrderingRules) bool {
	a, b := update[i], update[j]
	aRules := rules[a]

	// get all the numbers that 'a' must precede according to the rules
	aMustPrecede := []int{}
	for k := 0; k < len(update); k++ {
		for l := 0; l < len(aRules); l++ {
			if update[k] == aRules[l] {
				if aRules[l] == b {
					return true
				} else {
					aMustPrecede = append(aMustPrecede, k)
				}
			}
		}
	}
	// if 'b' is not in the list, recursviely check the rules applying to those numbers
	for k := 0; k < len(aMustPrecede); k++ {
		if MustPrecede(update, aMustPrecede[k], j, rules) {
			return true
		}
	}
	// did not find a rule that specified that 'a' must precede 'b'
	return false
}

func UpdateIsValid(update Update, rules OrderingRules) bool {
	ruleMap := map[int]bool{}
	return sort.SliceIsSorted(update, func(i, j int) bool {
		val, present := ruleMap[i*100+j]
		if present {
			return val
		} else {
			val = MustPrecede(update, i, j, rules)
			ruleMap[i*100+j] = val
			return val
		}
	})
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

func FixInvalidUpdates(updates []Update, rules OrderingRules) {
	for i := 0; i < len(updates); i++ {
		update := updates[i]
		for j := 0; j < len(update); j++ { // this is super janky, but running the sort n times fixes things
			ruleMap := map[int]bool{}
			sort.SliceStable(update, func(i, j int) bool {
				val, present := ruleMap[i*100+j]
				if present {
					return val
				} else {
					val = MustPrecede(update, i, j, rules)
					ruleMap[i*100+j] = val
					return val
				}
			})
		}
	}
}
