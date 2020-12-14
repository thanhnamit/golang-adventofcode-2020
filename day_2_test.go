package golang_adventofcode_2020

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay2PasswordPhilosophyPolicy1(t *testing.T) {
	re := countValidPassword("./input/day_2_sample.txt", isValidPolicy1)
	assert.Equal(t, 2, re)

	re = countValidPassword("./input/day_2_input.txt", isValidPolicy1)
	assert.Equal(t, 556, re)
}

func TestDay2PasswordPhilosophyPolicy2(t *testing.T) {
	re := countValidPassword("./input/day_2_sample.txt", isValidPolicy2)
	assert.Equal(t, 1, re)

	re = countValidPassword("./input/day_2_input.txt", isValidPolicy2)
	assert.Equal(t, 605, re)
}

func TestMakePassword(t *testing.T) {
	re := makePassword("1-3 a: abcde")
	assert.Equal(t, 1, re.min)
	assert.Equal(t, 3, re.max)
	assert.Equal(t, 'a', re.char)
	assert.Equal(t, "abcde", re.password)
}

type Password struct {
	char rune
	min int
	max int
	password string
}

func countValidPassword(file string, f func(p Password) bool) int {
	passwords := readFile(file)
	cnt := 0
	for _, v := range passwords {
		if f(v) {
			cnt++
		}
	}
	return cnt
}

func isValidPolicy1(p Password) bool {
	cnt := 0
	ln := len(p.password)

	for i := 0; i < ln; i++ {
		if rune(p.password[i]) == p.char {
			cnt++
		}
	}

	if cnt >= p.min && cnt <= p.max {
		return true
	}

	return false
}

func isValidPolicy2(p Password) bool {
	first := rune(p.password[p.min - 1])
	second := rune(p.password[p.max - 1])

	if first == p.char && second != p.char {
		return true
	}

	if first != p.char && second == p.char {
		return true
	}

	return false
}

func readFile(file string) []Password {
	var lines []Password
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v", err)
		os.Exit(1)
	}

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		l := scan.Text()
		lines = append(lines, makePassword(l))
	}
	return lines
}

func makePassword(line string) Password {
	arr := strings.Split(line, ":")

	pwd := strings.Trim(arr[1]," ")
	policyParts := strings.Split(arr[0], " ")

	char := policyParts[1]
	rangePart := strings.Split(policyParts[0], "-")

	min, _ := strconv.Atoi(rangePart[0])
	max, _ := strconv.Atoi(rangePart[1])
	
	return Password{
		char:     []rune(char)[0],
		min:      min,
		max:      max,
		password: pwd,
	}
}