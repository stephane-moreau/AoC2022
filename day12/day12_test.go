package day12

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Map [][]byte

func loadFile(file string) (Map, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	museumMap := make(Map, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.ReplaceAll(strings.TrimSpace(lines[i]), ".", "\x00")
		museumMap[i] = []byte(line)
	}
	return museumMap, nil
}

type point struct {
	x int
	y int
}

func (p point) move(mv point) point {
	if mv.x == 0 && mv.y == 0 {
		panic("wrong move")
	}
	return point{p.x + mv.x, p.y + mv.y}
}

func findPos(hikeMap Map, search byte) point {
	yMax := len(hikeMap)
	xMax := len(hikeMap[0])
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if hikeMap[y][x] == search {
				return point{x, y}
			}
		}
	}
	return point{-1, -1}
}

const (
	North = 1
	East  = 2
	South = 4
	West  = 8
)

var directions = []point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func hike(s byte, hikeMap Map) int {

	xMax, yMax := len(hikeMap[0]), len(hikeMap)
	shortestPath := xMax * yMax
	end := findPos(hikeMap, 'E')
	hikeMap[end.y][end.x] = 'z'
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			if hikeMap[y][x] != s {
				continue
			}
			start := point{x, y}
			hikeMap[start.y][start.x] = 'a'
			visits := map[point]byte{}
			paths := []map[point]byte{{start: 'a'}}
			for paths[len(paths)-1][end] == 0 {
				depth := len(paths) - 1
				curDepth := paths[depth]
				nextDepth := map[point]byte{}
				depth++
				paths = append(paths, nextDepth)
				for p := range curDepth {
					curLvl := hikeMap[p.y][p.x]
					for _, d := range directions {
						next := p.move(d)
						if next.x < 0 || next.y < 0 || next.x == xMax || next.y == yMax {
							continue
						}
						nextLvl := hikeMap[next.y][next.x]
						if nextLvl > curLvl+1 || curDepth[next] != 0 || visits[next] != 0 {
							continue
						}
						nextDepth[next] = nextLvl
					}
				}
				for p, c := range curDepth {
					visits[p] = c
				}
				if len(paths) > shortestPath {
					break
				}
			}
			if shortestPath > len(paths)-1 {
				shortestPath = len(paths) - 1
			}
		}
	}
	return shortestPath
}

func TestDay12Part1(t *testing.T) {
	hikeMap, err := loadFile("input_test.txt")
	require.NoError(t, err)
	pathLen := hike('S', hikeMap)
	assert.Equal(t, 31, pathLen)

	hikeMap, err = loadFile("input.txt")
	require.NoError(t, err)
	pathLen = hike('S', hikeMap)
	assert.Equal(t, 370, pathLen)
}

func TestDay12Part2(t *testing.T) {
	hikeMap, err := loadFile("input_test.txt")
	require.NoError(t, err)
	pathLen := hike('a', hikeMap)
	assert.Equal(t, 29, pathLen)

	hikeMap, err = loadFile("input.txt")
	require.NoError(t, err)
	pathLen = hike('a', hikeMap)
	assert.Equal(t, 370, pathLen)
}
