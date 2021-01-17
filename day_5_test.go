package golang_adventofcode_2020

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
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

type boardingPass struct {
	code string
	row int
	col int
	seatId int
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