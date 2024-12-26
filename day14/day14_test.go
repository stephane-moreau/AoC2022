package day14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const CRLF = "\r\n"

type Map struct {
	m        [][]byte
	min, max point
}

func (m *Map) Get(x, y int) byte {
	return m.m[y-m.min.y][x-m.min.x]
}
func (m *Map) Set(x, y int, c byte) {
	m.m[y-m.min.y][x-m.min.x] = c
}

type point struct {
	x int
	y int
}

var debug = true

func (m Map) dump() {
	if !debug {
		return
	}
	fmt.Println("")
	var sb strings.Builder
	for y := 0; y < m.max.y-m.min.y; y++ {
		for x := 0; x < m.max.x-m.min.x; x++ {
			c := m.m[y][x]
			if c == 0 {
				sb.WriteRune('.')
			} else {
				sb.WriteRune(rune(c))
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

func (m *Map) addFloor() {
	m.m[m.max.y] = make([]byte, m.max.x-m.min.x)
	for x := m.min.x; x < m.max.x; x++ {
		m.Set(x, m.max.y, '#')
	}
	m.max.y++
}

func (m Map) dropSand(d point) bool {
	for ; d.y < m.max.y-1; d.y++ {
		if m.Get(d.x, d.y+1) != 0 {
			if m.Get(d.x-1, d.y+1) == 0 {
				d.x--
				continue
			}
			if d.x+1 == m.max.x {
				return false
			}
			if m.Get(d.x+1, d.y+1) == 0 {
				d.x++
				continue
			}
			break
		}
	}
	if m.Get(d.x, d.y) != 0 || d.y == m.max.y-1 {
		return false
	}
	m.Set(d.x, d.y, 'o')
	return true
}

func loadInput(file string) (*Map, error) {
	content, err := os.ReadFile(file)
	var ptMax point
	ptMin := point{-1, 0}
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), CRLF)
	tracks := make([][]point, 0)
	for _, l := range lines {
		coords := strings.Split(l, "->")
		points := make([]point, len(coords))
		for i, c := range coords {
			xy := strings.Split(strings.TrimSpace(c), ",")
			points[i].x, err = strconv.Atoi(xy[0])
			if err != nil {
				return nil, err
			}
			if points[i].x >= ptMax.x {
				ptMax.x = points[i].x + 1
			}
			if points[i].x < ptMin.x || ptMin.x == -1 {
				ptMin.x = points[i].x - 1
			}
			points[i].y, err = strconv.Atoi(xy[1])
			if err != nil {
				return nil, err
			}
			if points[i].y >= ptMax.y {
				ptMax.y = points[i].y + 1
			}
		}
		tracks = append(tracks, points)
	}
	m := Map{
		m:   make([][]byte, ptMax.y+2-ptMin.y),
		min: point{ptMin.x - ptMax.y - 2, 0},
		max: point{ptMax.x + ptMax.y + 2, ptMax.y + 1},
	}
	for y := 0; y < m.max.y-m.min.y; y++ {
		m.m[y] = make([]byte, m.max.x-m.min.x)
	}
	for _, t := range tracks {
		for i := 0; i < len(t)-1; i++ {
			p := t[i]
			n := t[i+1]
			if p.x == n.x {
				for y := min(p.y, n.y); y <= max(p.y, n.y); y++ {
					m.Set(p.x, y, '#')
				}
			}
			if p.y == n.y {
				for x := min(p.x, n.x); x <= max(p.x, n.x); x++ {
					m.Set(x, p.y, '#')
				}
			}
		}
	}
	return &m, nil
}

func TestDay13Part1(t *testing.T) {
	m, err := loadInput("input_test.txt")
	require.NoError(t, err)
	i := 0
	for m.dropSand(point{500, 0}) {
		m.dump()
		i++
	}
	assert.Equal(t, 24, i)
	debug = false
	m, err = loadInput("input.txt")
	require.NoError(t, err)
	i = 0
	for m.dropSand(point{500, 0}) {
		m.dump()
		i++
	}
	assert.Equal(t, 913, i)
}

func TestDay13Part2(t *testing.T) {
	m, err := loadInput("input_test.txt")
	m.addFloor()
	require.NoError(t, err)
	i := 0
	for m.dropSand(point{500, 0}) {
		m.dump()
		i++
	}
	assert.Equal(t, 93, i)
	debug = false
	m, err = loadInput("input.txt")
	m.addFloor()
	require.NoError(t, err)
	i = 0
	for m.dropSand(point{500, 0}) {
		m.dump()
		i++
	}
	assert.Equal(t, 30762, i)
}
