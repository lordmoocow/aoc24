package days

import (
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	rules map[int][]int
	updates [][]int
}

func (d *Day5) Init(input string) {
	orderRuleInput, updateQueueInput, _ := strings.Cut(input, "\n\n")

	orderRules := strings.Split(strings.TrimSpace(orderRuleInput), "\n")
	d.rules = map[int][]int{}
	for _, r := range orderRules {
		k, _ := strconv.Atoi(string(r[:2]))
		v, _ := strconv.Atoi(string(r[3:]))
		d.rules[k] = append(d.rules[k], v)
	}

	updateQueue := strings.Split(strings.TrimSpace(updateQueueInput), "\n")
	d.updates = make([][]int, len(updateQueue))

	for i, update := range updateQueue {
		updateSplit := strings.Split(update, ",")
		d.updates[i] = make([]int, len(updateSplit))
		for j, u := range updateSplit {
			d.updates[i][j], _ = strconv.Atoi(u)
		}
	}
}

func (d *Day5) PartOne() (result int) {
	for _, update := range d.updates {
		if d.checkOrder(update) {
			result += update[len(update)/2]
		}
	}
	return result
}

func (d *Day5) checkOrder(update []int) bool {
	for i, page := range update {
		rules := d.rules[page]
		previousPages := update[:i]
		for _, r := range rules {
			if slices.Contains(previousPages, r){
				return false
			}
		}
	}
	return true
}


func (d *Day5) PartTwo() (result int) {
	return result
}
