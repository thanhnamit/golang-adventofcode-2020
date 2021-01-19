package golang_adventofcode_2020

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
)

func TestBoardingPassSample(t *testing.T) {
	boardingPass := toBoardingPass("FBFBBFFRLR")
	assert.Equal(t, 44, boardingPass.row)
	assert.Equal(t, 5, boardingPass.col)
	assert.Equal(t, 357, boardingPass.seatId)
	assert.Equal(t, "FBFBBFFRLR", boardingPass.code)

	passes := readBoardingPassFile("./input/day_5_sample.txt")
	assert.Equal(t, 3, len(passes))
	assert.Equal(t, 820, findMax(passes))
}

func TestBoardingPassInput(t *testing.T) {
	passes := readBoardingPassFile("./input/day_5_input.txt")
	assert.Equal(t, 838, findMax(passes))
}

func TestFindMySeat(t *testing.T)  {
	passes := readBoardingPassFile("./input/day_5_input.txt")
	assert.Equal(t, 714, findMySeat(passes))
}

type boardingPass struct {
	code string
	row int
	col int
	seatId int
}

type bySeatId []boardingPass
func (x bySeatId) Len() int { return len(x) }
func (x bySeatId) Less(i, j int) bool { return x[i].seatId < x[j].seatId }
func (x bySeatId) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func findMySeat(passes []boardingPass) int {
	sort.Sort(bySeatId(passes))
	n := len(passes)
	var next int
	for i := 1; i < n; i++ {
		if passes[i].seatId - passes[i - 1].seatId == 2 {
			next = i
			break
		}
	}
	return passes[next].seatId - 1
}

func findMax(passes []boardingPass) int {
	max := 0
	for _, p := range passes {
		if p.seatId >= max {
			max = p.seatId
		}
	}
	return max
}

func readBoardingPassFile(fileName string) []boardingPass {
	var bList []boardingPass
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		if l != "" {
			bList = append(bList, toBoardingPass(l))
		}
	}
	return bList
}

func toBoardingPass(code string) boardingPass {
	row := calculateRow(code)
	col := calculateCol(code)
	return boardingPass{
		code:   code,
		row:    row,
		col:    col,
		seatId: row * 8 + col,
	}
}

func calculateRow(code string) int {
	start := 0
	end := 127
	for i := 0; i < 7; i++ {
		ch := code[i]
		mid := start + (end - start) / 2
		if ch == 'F' {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

func calculateCol(code string) int {
	start := 0
	end := 7
	for i := 7; i < 10; i++ {
		ch := code[i]
		mid := start + (end - start) / 2
		if ch == 'L' {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}