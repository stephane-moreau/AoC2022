package day7

import (
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

type sizer interface {
	size() int
}

type dir struct {
	name    string
	content map[string]sizer
}

func (d dir) size() int {
	s := 0
	for _, f := range d.content {
		s += f.size()
	}
	return s
}

type file struct {
	name string
	sz   int
}

func (f file) size() int {
	return f.sz
}

func loadFile(fName string) dir {
	d := dir{
		"/",
		map[string]sizer{},
	}

	content, err := os.ReadFile(fName)
	if err != nil {
		panic(err)
	}

	stack := []dir{d}
	for _, c := range strings.Split(string(content), "\n") {
		c = strings.TrimSpace(c)
		cur := stack[len(stack)-1]
		if c == "$ cd .." {
			stack = stack[:len(stack)-1]
		} else if strings.HasPrefix(c, "$ cd ") {
			directory := c[5:]
			if directory == "/" {
				continue
			}
			stack = append(stack, cur.content[directory].(dir))
		} else if c == "$ ls" {
			continue
		} else {
			parts := strings.Split(c, " ")
			if parts[0] == "dir" {
				cur.content[parts[1]] = dir{parts[1], make(map[string]sizer)}
			} else {
				s, _ := strconv.Atoi(parts[0])
				cur.content[parts[1]] = file{parts[0], s}
			}
		}
	}
	return d
}

func TestDay7_Phase1(t *testing.T) {
	fs := loadFile("test.txt")
	s := 0
	dirs := []dir{fs}
	for i := 0; i < len(dirs); i++ {
		d := dirs[i]
		if sz := d.size(); sz < 100000 {
			s += sz
		}
		for _, c := range d.content {
			if _, isDir := c.(dir); !isDir {
				continue
			}
			dirs = append(dirs, c.(dir))
		}
	}

	if s != 95437 {
		t.Fail()
	}

	fs = loadFile("input.txt")
	s = 0
	dirs = []dir{fs}
	for i := 0; i < len(dirs); i++ {
		d := dirs[i]
		if sz := d.size(); sz < 100000 {
			s += sz
		}
		for _, c := range d.content {
			if _, isDir := c.(dir); !isDir {
				continue
			}
			dirs = append(dirs, c.(dir))
		}
	}

	if s != 1077191 {
		t.Fail()
	}
}

const (
	diskSpace     = 70_000_000
	requiredSpace = 30_000_000
)

func TestDay7_Phase2(t *testing.T) {
	fs := loadFile("test.txt")
	toFree := requiredSpace - (diskSpace - fs.size())

	dirs := []dir{fs}
	minFree := math.MaxInt
	for i := 0; i < len(dirs); i++ {
		d := dirs[i]
		if sz := d.size(); sz >= toFree {
			if minFree > sz {
				minFree = sz
			}
		}
		for _, c := range d.content {
			if _, isDir := c.(dir); !isDir {
				continue
			}
			dirs = append(dirs, c.(dir))
		}
	}

	if minFree != 24933642 {
		t.Fail()
	}

	fs = loadFile("input.txt")
	toFree = requiredSpace - (diskSpace - fs.size())
	dirs = []dir{fs}
	minFree = math.MaxInt
	for i := 0; i < len(dirs); i++ {
		d := dirs[i]
		if sz := d.size(); sz >= toFree {
			if minFree > sz {
				minFree = sz
			}
		}
		for _, c := range d.content {
			if _, isDir := c.(dir); !isDir {
				continue
			}
			dirs = append(dirs, c.(dir))
		}
	}
	// 9510609 too high
	if minFree != 5649896 {
		t.Fail()
	}
}
