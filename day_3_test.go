package golang_adventofcode_2020

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	space = '.'
	tree = '#'
)

func TestTobogganSlope(t *testing.T) {
	forest := readForestFile("./input/day_3_sample.txt")
	assert.Equal(t, 7, countTree(forest, 3, 1))

	forest = readForestFile("./input/day_3_input.txt")
	assert.Equal(t, 184, countTree(forest, 3, 1))
}

func TestTobogganMultipleSlopes(t *testing.T) {
	forest := readForestFile("./input/day_3_input.txt")
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	assert.Equal(t, 2431272960, countMultiple(forest, slopes))
}

func countMultiple(forest [][]int, slopes [][]int) int {
	re := 1
	for _, sl := range slopes {
		re *= countTree(forest, sl[0], sl[1])
	}
	return re
}

/*
countTree returns a number of trees encountered for a specific slopes (sr = right, sd = down)
 */
func countTree(forest [][]int, sr int, sd int) int {
	// current position
	cPos := 0
	cnt := 0

	rowCnt := len(forest)
	colCnt := len(forest[0])

	r := 0
	for r < rowCnt {
		cPos += sr
		if cPos >= colCnt {
			cPos = cPos % (colCnt - 1) - 1
		}

		r += sd

		// check
		if r < rowCnt && forest[r][cPos] == tree {
			cnt++
		}
	}

	return cnt
}

// we mark 1 for a tree, 0 for an open space
func readForestFile(fName string) [][]int {
	var forest [][]int
	f, err := os.Open(fName)
	must(err)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		txt := scan.Text()
		forest = append(forest, makeLine(txt))
	}
	return forest
}

func makeLine(l string) []int {
	var line []int
	for _, v := range l {
		if v == space {
			line = append(line, space)
		} else {
			line = append(line, tree)
		}
	}
	return line
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}