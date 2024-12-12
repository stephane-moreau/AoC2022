package day11

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type monkey struct {
	inspected int
	items     []int
	op        func(int) int
	div       int
	next      []int
}

var (
	input_test = []*monkey{
		{
			items: []int{79, 98},
			op:    func(old int) int { return old * 19 },
			div:   23,
			next:  []int{2, 3},
		},

		{
			items: []int{54, 65, 75, 74},
			op:    func(old int) int { return old + 6 },
			div:   19,
			next:  []int{2, 0},
		},

		{
			items: []int{79, 60, 97},
			op:    func(old int) int { return old * old },
			div:   13,
			next:  []int{1, 3},
		},

		{
			items: []int{74},
			op:    func(old int) int { return old + 3 },
			div:   17,
			next:  []int{0, 1},
		},
	}

	input = []*monkey{
		{
			items: []int{71, 56, 50, 73},
			op:    func(old int) int { return old * 11 },
			div:   13,
			next:  []int{1, 7},
		},
		{
			items: []int{70, 89, 82},
			op:    func(old int) int { return old + 1 },
			div:   7,
			next:  []int{3, 6},
		},
		{
			items: []int{52, 95},
			op:    func(old int) int { return old * old },
			div:   3,
			next:  []int{5, 4},
		},
		{
			items: []int{94, 64, 69, 87, 70},
			op:    func(old int) int { return old + 2 },
			div:   19,
			next:  []int{2, 6},
		},
		{
			items: []int{98, 72, 98, 53, 97, 51},
			op:    func(old int) int { return old + 6 },
			div:   5,
			next:  []int{0, 5},
		},
		{
			items: []int{79},
			op:    func(old int) int { return old + 7 },
			div:   2,
			next:  []int{7, 0},
		},
		{
			items: []int{77, 55, 63, 93, 66, 90, 88, 71},
			op:    func(old int) int { return old * 7 },
			div:   11,
			next:  []int{2, 4},
		},
		{
			items: []int{54, 97, 87, 70, 59, 82, 59},
			op:    func(old int) int { return old + 8 },
			div:   17,
			next:  []int{1, 3},
		},
	}
)

func play(monkeys []*monkey, rounded bool, rounds int) int {
	worrylevel := 1
	for _, m := range monkeys {
		worrylevel *= m.div
	}
	for i := 0; i < rounds; i++ {
		if i == 1 || i == 20 || (i%1000) == 0 {
			_ = i
		}
		for _, m := range monkeys {
			for _, val := range m.items {
				m.inspected++
				v := m.op(val)
				if rounded {
					v = int(float64(v) / 3.0)
				} else {
					v = v % worrylevel
				}
				var next int
				if v%m.div == 0 {
					next = m.next[0]
				} else {
					next = m.next[1]
				}
				n := monkeys[next]
				n.items = append(n.items, v)
			}
			m.items = m.items[:0]
		}
	}
	sort.SliceStable(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}

func TestDay11Part1(t *testing.T) {
	assert.Equal(t, 10605, play(input_test, true, 20))
	assert.Equal(t, 66802, play(input, true, 20))
}

func TestDay11Part2(t *testing.T) {
	assert.Equal(t, 2713310158, play(input_test, false, 10_000))
	assert.Equal(t, 21800916620, play(input, false, 10_000))
}
