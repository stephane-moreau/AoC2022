package day13

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type values []any

func (vals values) String() string {
	s := "["
	for i, val := range vals {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprint(val)
	}
	return s + "]"
}

type pair struct {
	left  values
	right values
}

func parseValues(l string) values {
	var res values
	l = strings.TrimSpace(l)
	if l[0] != '[' || l[len(l)-1] != ']' {
		panic("invalid array")
	}
	var sb strings.Builder
	arrStack := values{&res}
	cur := &res
	for i := 1; i < len(l)-1; i++ {
		c := l[i]
		if c >= '0' && c <= '9' {
			sb.WriteByte(c)
		} else if c == ',' && sb.Len() > 0 {
			n, err := strconv.Atoi(strings.TrimSpace(sb.String()))
			if err != nil {
				panic(err)
			}
			sb.Reset()
			*cur = append(*cur, n)
		} else if c == '[' {
			arr := make(values, 0)
			cur = &arr
			arrStack = append(arrStack, cur)
		} else if c == ']' {
			if sb.Len() > 0 {
				n, err := strconv.Atoi(strings.TrimSpace(sb.String()))
				if err != nil {
					panic(err)
				}
				sb.Reset()
				*cur = append(*cur, n)
			}
			prev := arrStack[len(arrStack)-1]
			arrStack = arrStack[:len(arrStack)-1]
			cur = arrStack[len(arrStack)-1].(*values)
			*cur = append(*cur, *prev.(*values))
		}
	}
	if sb.Len() > 0 {
		n, err := strconv.Atoi(strings.TrimSpace(sb.String()))
		if err != nil {
			panic(err)
		}
		sb.Reset()
		*cur = append(*cur, n)
	}

	return res
}

func loadPairs(file string) []pair {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(content), "\n")
	pairs := make([]pair, 0)
	for i := 0; i < len(lines); i += 3 {
		var p pair
		p.left = parseValues(lines[i])
		p.right = parseValues(lines[i+1])
		pairs = append(pairs, p)
	}
	return pairs
}

func compare(left, right any) int {
	li, isLeftInt := left.(int)
	ri, isRightInt := right.(int)
	lAr, isLeftArr := left.(values)
	rAr, isRightArr := right.(values)
	if isLeftInt && isRightInt {
		if li < ri {
			return -1
		}
		if li == ri {
			return 0
		}
		return 1
	}
	if isLeftInt && isRightArr {
		return compare(values{li}, rAr)
	}
	if isLeftArr && isRightInt {
		return compare(lAr, values{ri})
	}
	for i := range lAr {
		if i < len(rAr) {
			cmpItem := compare(lAr[i], rAr[i])
			if cmpItem != 0 {
				return cmpItem
			}
		}
		if i == len(rAr) {
			return 1
		}
	}
	if len(rAr) > len(lAr) {
		return -1
	}
	return 0
}

func TestDay13Part1(t *testing.T) {
	ps := loadPairs("input_test.txt")
	res := 0
	for ndx, p := range ps {
		if compare(p.left, p.right) <= 0 {
			res += ndx + 1
		} else {
			fmt.Printf("ndx %d: %v - %v\n", ndx, p.left, p.right)
		}
	}
	assert.Equal(t, 13, res)

	ps = loadPairs("input.txt")
	res = 0
	for ndx, p := range ps {
		if compare(p.left, p.right) <= 0 {
			res += ndx + 1
		} else {
			fmt.Printf("ndx %d: %v - %v\n", ndx, p.left, p.right)
		}
	}
	assert.Equal(t, 13, res)
}

func loadFile(file string) []values {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	lines := strings.Split(string(content), "\n")
	vals := make([]values, 0)
	for i := 0; i < len(lines); i++ {
		l := strings.TrimSpace(lines[i])
		if l == "" {
			continue
		}
		v := parseValues(l)
		vals = append(vals, v)
	}
	return vals
}

func TestDay13Part2(t *testing.T) {
	vals := loadFile("input_test.txt")
	v6 := values{values{6}}
	v2 := values{values{2}}
	vals = append(vals, v6, v2)
	sort.SliceStable(vals, func(i, j int) bool {
		return compare(vals[i], vals[j]) < 0
	})
	var n6, n2 int
	for n, v := range vals {
		if compare(v, v6) == 0 {
			n6 = n + 1
		}
		if compare(v, v2) == 0 {
			n2 = n + 1
		}
	}
	assert.Equal(t, 140, n6*n2)

	vals = loadFile("input.txt")
	vals = append(vals, v6, v2)
	sort.SliceStable(vals, func(i, j int) bool {
		return compare(vals[i], vals[j]) < 0
	})
	for n, v := range vals {
		if compare(v, v6) == 0 {
			n6 = n + 1
		}
		if compare(v, v2) == 0 {
			n2 = n + 1
		}
	}
	assert.Equal(t, 140, n6*n2)
}
