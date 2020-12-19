package golang_adventofcode_2020

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPassportProcessing(t *testing.T) {
	passports := readPassportFile("./input/day_4_sample.txt")
	assert.Equal(t, 2, countValid(passports))

	passports = readPassportFile("./input/day_4_input.txt")
	assert.Equal(t, 206, countValid(passports))
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func countValid(pList []passport) int {
	cnt := 0
	for _, p := range pList {
		if isValid(p) {
			cnt++
		}
	}
	return cnt
}

func readPassportFile(fileName string) []passport {
	var pList []passport
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var acc []string
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			pList = append(pList, toPassport(acc))
			acc = make([]string, 0)
		} else {
			acc = append(acc, l)
		}
	}

	// last one
	pList = append(pList, toPassport(acc))
	return pList
}

func toPassport(arr []string) passport {
	m := map[string]string{
		"byr": "",
		"iyr": "",
		"eyr": "",
		"hgt": "",
		"hcl": "",
		"ecl": "",
		"pid": "",
		"cid": "",
	}

	for _, line := range arr {
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			subs := strings.Split(token, ":")
			key := strings.TrimSpace(subs[0])
			m[key] = subs[1]
		}
	}

	return passport{
		byr: m["byr"],
		iyr: m["iyr"],
		eyr: m["eyr"],
		hgt: m["hgt"],
		hcl: m["hcl"],
		ecl: m["ecl"],
		pid: m["pid"],
		cid: m["cid"],
	}
}

func isValid(p passport) bool {
	blk := ""
	if p.byr == blk || p.iyr == blk || p.eyr == blk || p.hgt == blk || p.hcl == blk || p.ecl == blk || p.pid == blk {
		return false
	}
	return true
}