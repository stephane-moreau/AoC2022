package day5

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var (
	testInput = []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	largeInput = []string{
		"[N]             [R]             [C]",
		"[T] [J]         [S] [J]         [N]",
		"[B] [Z]     [H] [M] [Z]         [D]",
		"[S] [P]     [G] [L] [H] [Z]     [T]",
		"[Q] [D]     [F] [D] [V] [L] [S] [M]",
		"[H] [F] [V] [J] [C] [W] [P] [W] [L]",
		"[G] [S] [H] [Z] [Z] [T] [F] [V] [H]",
		"[R] [H] [Z] [M] [T] [M] [T] [Q] [W]",
		" 1   2   3   4   5   6   7   8   9 ",
		"",
		"move 3 from 9 to 7",
		"move 4 from 4 to 5",
		"move 2 from 4 to 6",
		"move 4 from 7 to 5",
		"move 3 from 7 to 3",
		"move 2 from 5 to 9",
		"move 5 from 6 to 3",
		"move 5 from 9 to 1",
		"move 3 from 8 to 4",
		"move 3 from 4 to 6",
		"move 8 from 1 to 8",
		"move 1 from 8 to 6",
		"move 2 from 8 to 2",
		"move 5 from 8 to 4",
		"move 1 from 8 to 1",
		"move 6 from 6 to 4",
		"move 1 from 7 to 9",
		"move 5 from 1 to 7",
		"move 1 from 1 to 2",
		"move 2 from 9 to 8",
		"move 6 from 4 to 9",
		"move 1 from 6 to 8",
		"move 3 from 2 to 7",
		"move 4 from 2 to 8",
		"move 4 from 9 to 3",
		"move 6 from 5 to 4",
		"move 7 from 8 to 1",
		"move 10 from 4 to 1",
		"move 12 from 1 to 5",
		"move 1 from 4 to 9",
		"move 1 from 2 to 3",
		"move 2 from 9 to 1",
		"move 1 from 9 to 3",
		"move 1 from 6 to 7",
		"move 1 from 9 to 1",
		"move 3 from 1 to 3",
		"move 9 from 5 to 9",
		"move 2 from 2 to 7",
		"move 2 from 7 to 4",
		"move 3 from 9 to 4",
		"move 7 from 5 to 7",
		"move 5 from 1 to 3",
		"move 2 from 4 to 5",
		"move 1 from 4 to 6",
		"move 1 from 6 to 9",
		"move 4 from 9 to 2",
		"move 12 from 7 to 9",
		"move 2 from 4 to 9",
		"move 6 from 5 to 9",
		"move 3 from 7 to 6",
		"move 12 from 9 to 6",
		"move 5 from 9 to 1",
		"move 1 from 7 to 6",
		"move 14 from 6 to 1",
		"move 20 from 3 to 5",
		"move 5 from 9 to 5",
		"move 3 from 2 to 8",
		"move 1 from 6 to 4",
		"move 1 from 9 to 2",
		"move 1 from 4 to 6",
		"move 1 from 2 to 6",
		"move 16 from 1 to 5",
		"move 1 from 2 to 1",
		"move 12 from 5 to 6",
		"move 1 from 8 to 4",
		"move 29 from 5 to 1",
		"move 5 from 6 to 9",
		"move 20 from 1 to 3",
		"move 4 from 1 to 3",
		"move 11 from 3 to 8",
		"move 1 from 4 to 3",
		"move 4 from 9 to 8",
		"move 7 from 1 to 8",
		"move 2 from 3 to 2",
		"move 2 from 6 to 7",
		"move 1 from 9 to 8",
		"move 10 from 3 to 5",
		"move 1 from 6 to 1",
		"move 1 from 7 to 2",
		"move 3 from 1 to 2",
		"move 6 from 2 to 4",
		"move 2 from 6 to 3",
		"move 4 from 6 to 5",
		"move 1 from 6 to 2",
		"move 1 from 2 to 9",
		"move 6 from 5 to 2",
		"move 1 from 9 to 3",
		"move 24 from 8 to 7",
		"move 1 from 4 to 8",
		"move 5 from 5 to 4",
		"move 1 from 4 to 8",
		"move 1 from 8 to 7",
		"move 2 from 8 to 9",
		"move 1 from 9 to 7",
		"move 6 from 2 to 4",
		"move 10 from 3 to 7",
		"move 3 from 5 to 3",
		"move 1 from 9 to 8",
		"move 3 from 3 to 8",
		"move 4 from 8 to 7",
		"move 1 from 4 to 6",
		"move 1 from 6 to 4",
		"move 13 from 4 to 3",
		"move 17 from 7 to 6",
		"move 1 from 6 to 3",
		"move 2 from 4 to 8",
		"move 3 from 7 to 5",
		"move 14 from 6 to 7",
		"move 1 from 5 to 9",
		"move 1 from 5 to 9",
		"move 2 from 6 to 7",
		"move 1 from 5 to 1",
		"move 1 from 1 to 6",
		"move 1 from 9 to 3",
		"move 29 from 7 to 4",
		"move 10 from 4 to 3",
		"move 6 from 7 to 5",
		"move 1 from 6 to 5",
		"move 1 from 9 to 7",
		"move 1 from 7 to 2",
		"move 4 from 3 to 2",
		"move 1 from 2 to 9",
		"move 1 from 8 to 5",
		"move 11 from 3 to 4",
		"move 24 from 4 to 7",
		"move 2 from 2 to 5",
		"move 10 from 3 to 2",
		"move 6 from 2 to 1",
		"move 5 from 4 to 7",
		"move 1 from 9 to 2",
		"move 3 from 5 to 1",
		"move 1 from 4 to 6",
		"move 4 from 2 to 3",
		"move 5 from 5 to 7",
		"move 2 from 5 to 3",
		"move 32 from 7 to 5",
		"move 16 from 5 to 1",
		"move 1 from 1 to 2",
		"move 3 from 2 to 9",
		"move 1 from 8 to 6",
		"move 3 from 7 to 6",
		"move 1 from 2 to 4",
		"move 5 from 6 to 8",
		"move 5 from 8 to 6",
		"move 2 from 9 to 3",
		"move 1 from 7 to 5",
		"move 9 from 5 to 4",
		"move 1 from 9 to 1",
		"move 2 from 3 to 1",
		"move 4 from 3 to 6",
		"move 1 from 3 to 8",
		"move 6 from 4 to 6",
		"move 6 from 5 to 9",
		"move 1 from 9 to 6",
		"move 1 from 5 to 1",
		"move 1 from 5 to 4",
		"move 1 from 3 to 6",
		"move 1 from 8 to 3",
		"move 1 from 4 to 2",
		"move 1 from 2 to 3",
		"move 17 from 6 to 4",
		"move 4 from 1 to 8",
		"move 3 from 9 to 6",
		"move 1 from 8 to 4",
		"move 1 from 9 to 7",
		"move 2 from 6 to 2",
		"move 1 from 7 to 8",
		"move 12 from 1 to 9",
		"move 8 from 9 to 2",
		"move 1 from 6 to 9",
		"move 6 from 2 to 8",
		"move 2 from 8 to 3",
		"move 18 from 4 to 9",
		"move 2 from 1 to 6",
		"move 1 from 6 to 5",
		"move 3 from 4 to 3",
		"move 7 from 3 to 8",
		"move 4 from 2 to 7",
		"move 1 from 4 to 6",
		"move 2 from 6 to 4",
		"move 13 from 9 to 6",
		"move 1 from 5 to 2",
		"move 5 from 9 to 3",
		"move 9 from 1 to 2",
		"move 1 from 1 to 8",
		"move 1 from 2 to 6",
		"move 3 from 7 to 6",
		"move 2 from 2 to 6",
		"move 9 from 8 to 6",
		"move 1 from 7 to 8",
		"move 1 from 8 to 7",
		"move 2 from 4 to 6",
		"move 5 from 3 to 6",
		"move 17 from 6 to 9",
		"move 7 from 8 to 4",
		"move 4 from 2 to 3",
		"move 17 from 6 to 2",
		"move 1 from 6 to 4",
		"move 1 from 7 to 8",
		"move 1 from 8 to 9",
		"move 24 from 9 to 6",
		"move 4 from 3 to 1",
		"move 1 from 1 to 5",
		"move 20 from 6 to 4",
		"move 4 from 6 to 9",
		"move 1 from 5 to 7",
		"move 2 from 4 to 2",
		"move 1 from 9 to 7",
		"move 25 from 4 to 3",
		"move 1 from 4 to 2",
		"move 2 from 1 to 6",
		"move 3 from 9 to 4",
		"move 2 from 4 to 7",
		"move 2 from 7 to 5",
		"move 1 from 4 to 2",
		"move 1 from 6 to 3",
		"move 1 from 1 to 5",
		"move 5 from 3 to 9",
		"move 1 from 5 to 6",
		"move 10 from 2 to 8",
		"move 9 from 2 to 5",
		"move 21 from 3 to 6",
		"move 1 from 7 to 6",
		"move 2 from 6 to 5",
		"move 5 from 9 to 7",
		"move 6 from 7 to 8",
		"move 19 from 6 to 9",
		"move 1 from 6 to 1",
		"move 8 from 8 to 1",
		"move 1 from 6 to 1",
		"move 2 from 8 to 5",
		"move 5 from 9 to 2",
		"move 6 from 8 to 2",
		"move 2 from 9 to 7",
		"move 9 from 9 to 4",
		"move 7 from 2 to 4",
		"move 1 from 6 to 4",
		"move 14 from 5 to 9",
		"move 1 from 1 to 8",
		"move 1 from 7 to 9",
		"move 4 from 2 to 9",
		"move 16 from 4 to 6",
		"move 3 from 2 to 8",
		"move 1 from 6 to 2",
		"move 2 from 8 to 9",
		"move 1 from 8 to 7",
		"move 1 from 8 to 3",
		"move 3 from 2 to 7",
		"move 1 from 3 to 9",
		"move 8 from 9 to 3",
		"move 4 from 7 to 8",
		"move 1 from 5 to 4",
		"move 4 from 6 to 3",
		"move 1 from 4 to 2",
		"move 9 from 3 to 8",
		"move 10 from 9 to 5",
		"move 8 from 6 to 7",
		"move 13 from 8 to 4",
		"move 8 from 5 to 2",
		"move 3 from 6 to 3",
		"move 7 from 9 to 6",
		"move 7 from 7 to 2",
		"move 2 from 4 to 6",
		"move 5 from 6 to 2",
		"move 3 from 1 to 5",
		"move 5 from 5 to 8",
		"move 4 from 6 to 2",
		"move 4 from 1 to 8",
		"move 15 from 2 to 6",
		"move 11 from 4 to 9",
		"move 12 from 6 to 8",
		"move 1 from 6 to 9",
		"move 5 from 3 to 7",
		"move 2 from 2 to 6",
		"move 6 from 7 to 1",
		"move 3 from 1 to 3",
		"move 1 from 4 to 1",
		"move 1 from 3 to 9",
		"move 1 from 3 to 9",
		"move 1 from 7 to 6",
		"move 1 from 3 to 2",
		"move 4 from 2 to 6",
		"move 4 from 2 to 7",
		"move 1 from 2 to 6",
		"move 4 from 1 to 6",
		"move 12 from 6 to 7",
		"move 2 from 6 to 1",
		"move 8 from 9 to 6",
		"move 1 from 7 to 4",
		"move 14 from 8 to 1",
		"move 8 from 1 to 5",
		"move 1 from 3 to 9",
		"move 5 from 9 to 5",
		"move 1 from 8 to 9",
		"move 1 from 9 to 2",
		"move 1 from 9 to 3",
		"move 5 from 8 to 3",
		"move 12 from 5 to 4",
		"move 1 from 9 to 2",
		"move 6 from 7 to 3",
		"move 7 from 3 to 2",
		"move 1 from 5 to 1",
		"move 1 from 8 to 3",
		"move 2 from 1 to 3",
		"move 2 from 6 to 9",
		"move 5 from 6 to 5",
		"move 5 from 1 to 7",
		"move 4 from 4 to 1",
		"move 7 from 2 to 8",
		"move 4 from 3 to 8",
		"move 1 from 9 to 3",
		"move 1 from 9 to 5",
		"move 4 from 1 to 8",
		"move 10 from 7 to 9",
		"move 1 from 6 to 7",
		"move 2 from 8 to 6",
		"move 6 from 4 to 2",
		"move 5 from 3 to 1",
		"move 2 from 6 to 3",
		"move 2 from 7 to 1",
		"move 5 from 2 to 5",
		"move 2 from 7 to 1",
		"move 7 from 5 to 7",
		"move 2 from 5 to 6",
		"move 2 from 5 to 3",
		"move 3 from 2 to 9",
		"move 9 from 9 to 3",
		"move 1 from 6 to 4",
		"move 3 from 3 to 1",
		"move 9 from 8 to 2",
		"move 6 from 3 to 6",
		"move 8 from 7 to 9",
		"move 4 from 9 to 8",
		"move 14 from 1 to 5",
		"move 1 from 9 to 2",
		"move 1 from 1 to 5",
		"move 2 from 3 to 6",
		"move 12 from 5 to 3",
		"move 2 from 2 to 8",
		"move 7 from 6 to 2",
		"move 12 from 2 to 8",
		"move 2 from 6 to 2",
		"move 6 from 9 to 6",
		"move 1 from 1 to 2",
		"move 1 from 9 to 3",
		"move 2 from 5 to 9",
		"move 1 from 9 to 2",
		"move 1 from 9 to 4",
		"move 1 from 3 to 2",
		"move 2 from 6 to 7",
		"move 2 from 6 to 9",
		"move 5 from 4 to 2",
		"move 14 from 3 to 9",
		"move 15 from 9 to 4",
		"move 1 from 7 to 4",
		"move 10 from 8 to 6",
		"move 1 from 5 to 9",
		"move 2 from 9 to 5",
		"move 10 from 8 to 1",
		"move 1 from 7 to 4",
		"move 5 from 1 to 2",
		"move 2 from 1 to 5",
		"move 3 from 4 to 6",
		"move 4 from 5 to 8",
		"move 5 from 8 to 6",
		"move 14 from 2 to 9",
		"move 2 from 6 to 7",
		"move 3 from 2 to 9",
		"move 3 from 1 to 7",
		"move 1 from 7 to 3",
		"move 3 from 7 to 1",
		"move 1 from 3 to 6",
		"move 1 from 7 to 6",
		"move 1 from 8 to 9",
		"move 2 from 1 to 4",
		"move 1 from 1 to 2",
		"move 16 from 9 to 4",
		"move 7 from 4 to 8",
		"move 5 from 8 to 1",
		"move 2 from 8 to 3",
		"move 2 from 1 to 7",
		"move 13 from 6 to 7",
		"move 2 from 2 to 3",
		"move 4 from 7 to 4",
		"move 6 from 4 to 5",
		"move 4 from 7 to 6",
		"move 3 from 1 to 2",
		"move 2 from 2 to 6",
		"move 3 from 3 to 8",
		"move 5 from 5 to 3",
		"move 2 from 9 to 6",
		"move 3 from 3 to 7",
		"move 1 from 8 to 1",
		"move 22 from 4 to 8",
		"move 1 from 4 to 3",
		"move 9 from 6 to 3",
		"move 1 from 2 to 1",
		"move 4 from 3 to 4",
		"move 2 from 4 to 5",
		"move 1 from 1 to 7",
		"move 4 from 3 to 7",
		"move 2 from 6 to 1",
		"move 1 from 6 to 7",
		"move 18 from 8 to 7",
		"move 2 from 6 to 5",
		"move 2 from 3 to 4",
		"move 1 from 5 to 4",
		"move 30 from 7 to 6",
		"move 2 from 1 to 3",
		"move 18 from 6 to 8",
		"move 12 from 6 to 4",
		"move 13 from 4 to 9",
		"move 2 from 3 to 8",
		"move 1 from 6 to 2",
		"move 3 from 7 to 2",
		"move 1 from 1 to 2",
		"move 2 from 5 to 9",
		"move 8 from 8 to 1",
		"move 1 from 7 to 8",
		"move 7 from 1 to 3",
		"move 2 from 4 to 9",
		"move 1 from 1 to 6",
		"move 4 from 2 to 1",
		"move 16 from 8 to 1",
		"move 1 from 2 to 6",
		"move 2 from 4 to 8",
		"move 2 from 5 to 1",
		"move 4 from 3 to 7",
		"move 3 from 7 to 1",
		"move 1 from 6 to 8",
		"move 1 from 8 to 9",
		"move 1 from 7 to 3",
		"move 6 from 3 to 5",
		"move 1 from 3 to 8",
		"move 1 from 6 to 9",
		"move 16 from 9 to 5",
		"move 4 from 5 to 3",
		"move 15 from 5 to 1",
		"move 1 from 5 to 8",
		"move 3 from 9 to 8",
		"move 9 from 8 to 5",
		"move 6 from 5 to 1",
		"move 4 from 5 to 6",
		"move 2 from 6 to 4",
		"move 1 from 6 to 4",
		"move 1 from 8 to 4",
		"move 3 from 3 to 6",
		"move 3 from 6 to 8",
		"move 1 from 6 to 8",
		"move 21 from 1 to 9",
		"move 4 from 8 to 5",
		"move 3 from 5 to 7",
		"move 2 from 5 to 1",
		"move 2 from 4 to 8",
		"move 2 from 8 to 2",
		"move 2 from 7 to 8",
		"move 1 from 7 to 9",
		"move 1 from 8 to 7",
		"move 5 from 1 to 8",
		"move 1 from 7 to 8",
		"move 4 from 8 to 4",
		"move 2 from 4 to 5",
		"move 1 from 2 to 7",
		"move 1 from 2 to 7",
		"move 2 from 7 to 6",
		"move 2 from 6 to 9",
		"move 1 from 4 to 9",
		"move 1 from 3 to 4",
		"move 16 from 1 to 5",
		"move 16 from 5 to 7",
		"move 2 from 5 to 4",
		"move 14 from 9 to 6",
		"move 5 from 4 to 3",
		"move 3 from 3 to 6",
		"move 5 from 1 to 4",
		"move 2 from 4 to 7",
		"move 7 from 9 to 4",
		"move 2 from 9 to 7",
		"move 10 from 6 to 9",
		"move 8 from 4 to 6",
		"move 1 from 8 to 4",
		"move 1 from 1 to 9",
		"move 14 from 6 to 3",
		"move 10 from 3 to 2",
		"move 3 from 7 to 8",
		"move 6 from 3 to 1",
		"move 2 from 7 to 9",
		"move 5 from 7 to 9",
		"move 10 from 9 to 1",
		"move 2 from 4 to 3",
		"move 1 from 2 to 1",
		"move 16 from 1 to 4",
		"move 1 from 6 to 1",
		"move 2 from 3 to 9",
		"move 3 from 8 to 5",
		"move 8 from 7 to 1",
		"move 3 from 5 to 9",
		"move 7 from 4 to 6",
		"move 7 from 1 to 5",
		"move 2 from 8 to 3",
		"move 1 from 7 to 8",
	}
)

type piles [][]byte

type move struct {
	num  int
	from int
	to   int
}

func (ps piles) apply(m move) {
	for i := 0; i < m.num; i++ {
		ps[m.to-1] = append(ps[m.to-1], ps[m.from-1][len(ps[m.from-1])-1])
		ps[m.from-1] = ps[m.from-1][:len(ps[m.from-1])-1]
	}
}

func (ps piles) applySlice(m move) {
	from := ps[m.from-1]
	from = from[len(from)-m.num:]
	ps[m.to-1] = append(ps[m.to-1], from...)
	ps[m.from-1] = ps[m.from-1][:len(ps[m.from-1])-m.num]
}

type moves []move

func parseMove(l string) move {
	parts := strings.Split(l, " ")
	var m move
	m.num, _ = strconv.Atoi(parts[1])
	m.from, _ = strconv.Atoi(parts[3])
	m.to, _ = strconv.Atoi(parts[5])
	return m
}

func parseInput(lines []string) (piles, moves) {
	numPiles := len(lines[0])/4 + 1
	p := make(piles, numPiles)
	numChar := 0
	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if l == "" {
			numChar = i - 1
			break
		}
	}
	m := numChar + 2
	for numChar > 0 {
		numChar--
		l := lines[numChar]
		for i := 0; i < numPiles; i++ {
			c := byte(0)
			if l[4*i] == '[' {
				c = l[4*i+1]
			}
			if c != 0 {
				p[i] = append(p[i], c)
			}
		}
	}
	var mvs moves
	for m < len(lines) {
		mvs = append(mvs, parseMove(lines[m]))
		m++
	}
	return p, mvs
}

func TestDay5Phase1(t *testing.T) {
	ps, mvs := parseInput(testInput)
	for _, m := range mvs {
		ps.apply(m)
	}
	res := ""
	for _, p := range ps {
		res = res + fmt.Sprintf("%c", p[len(p)-1])
	}
	if res != "CMZ" {
		t.Fail()
	}

	ps, mvs = parseInput(largeInput)
	for _, m := range mvs {
		ps.apply(m)
	}
	res = ""
	for _, p := range ps {
		res = res + fmt.Sprintf("%c", p[len(p)-1])
	}
	if res != "PTWLTDSJV" {
		t.Fail()
	}
}

func TestDay5Phase2(t *testing.T) {
	ps, mvs := parseInput(testInput)
	for _, m := range mvs {
		ps.applySlice(m)
	}
	res := ""
	for _, p := range ps {
		res = res + fmt.Sprintf("%c", p[len(p)-1])
	}
	if res != "MCD" {
		t.Fail()
	}

	ps, mvs = parseInput(largeInput)
	for _, m := range mvs {
		ps.applySlice(m)
	}
	res = ""
	for _, p := range ps {
		res = res + fmt.Sprintf("%c", p[len(p)-1])
	}
	if res != "WZMFVGGZP" {
		t.Fail()
	}
}
