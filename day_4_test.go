package golang_adventofcode_2020

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestPassportProcessing(t *testing.T) {
	passports := readPassportFile("./input/day_4_sample.txt")
	assert.Equal(t, 6, countValid(passports))

	passports = readPassportFile("./input/day_4_input.txt")
	assert.Equal(t, 123, countValid(passports))
}

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

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
		if isValid(p) == true {
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
			m[key] = strings.TrimSpace(subs[1])
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
	if !isValidByr(p.byr) || !isValidEcl(p.ecl) || !isValidPid(p.pid) || !isValidHcl(p.hcl) || !isValidIyr(p.iyr) || !isValidHgt(p.hgt) || !isValidEyr(p.eyr) {
		return false
	}

	return true
}

func isValidPid(pid string) bool {
	if pid == "" {
		return false
	}

	m, err := regexp.MatchString(`^\d{9}$`, pid)
	if err != nil {
		return false
	}

	return m
}

func isValidEcl(ecl string) bool {
	if ecl == "" {
		return false
	}

	for _, v := range eyeColors {
		if v == ecl {
			return true
		}
	}
	return false
}

func isValidHcl(hcl string) bool {
	if hcl == "" {
		return false
	}

	m, err := regexp.MatchString(`^#[a-f0-9]{6}$`, hcl)
	if err != nil {
		return false
	}
	return m
}

func isValidHgt(hgt string) bool {
	if hgt == "" {
		return false
	}

	re := regexp.MustCompile(`^(\d+)((cm)|(in))$`)
	matches := re.FindAllStringSubmatch(hgt, -1)
	if matches == nil || len(matches) != 1 {
		return false
	}
	size, _ := strconv.Atoi(matches[0][1])
	m := matches[0][2]

	if m == "cm" && (size < 150 || size > 193) {
		return false
	}

	if m == "in" && (size < 59 || size > 76) {
		return false
	}

	return true
}

func isValidByr(byr string) bool {
	return isValidYear(byr, 1920, 2002)
}

func isValidIyr(iyr string) bool {
	return isValidYear(iyr, 2010, 2020)
}

func isValidEyr(eyr string) bool {
	return isValidYear(eyr, 2020, 2030)
}

func isValidYear(val string, min int, max int) bool {
	if val == "" {
		return false
	}

	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	return num >= min && num <= max
}
