package golang_adventofcode_2020

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"strconv"
	"testing"
)

func TestDay1TwoSum(t *testing.T) {
	re := findTwoSum("./input/day_1_sample.txt")
	assert.Equal(t, 514579, re, "not equal")

	re = findTwoSum("./input/day_1_input.txt")
	assert.Equal(t, 1014171, re, "not equal")
}

func TestDay1ThreeSum(t *testing.T) {
	re := findThreeSum("./input/day_1_sample.txt")
	assert.Equal(t, 241861950, re, "not equal")

	re = findThreeSum("./input/day_1_input.txt")
	assert.Equal(t, 46584630, re, "not equal")
}

func findThreeSum(file string) int {
	nums := readInput(file)
	sort.Ints(nums)

	target := 2020
	n := len(nums)

	for i := 0; i < n - 1; i++ {
		newTarget := target - nums[i]
		start := i + 1
		end := n - 1

		for start < end {
			sum := nums[start] + nums[end]
			if sum == newTarget {
				return nums[start] * nums[end] * nums[i]
			} else if sum > newTarget {
				end--
			} else {
				start++
			}
		}
	}
	return 0
}

func findTwoSum(file string) int {
	nums := readInput(file)
	sort.Ints(nums)

	target := 2020
	n := len(nums)
	start := 0
	end := n - 1

	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			return nums[start] * nums[end]
		} else if sum > target {
			end--
		} else {
			start++
		}
	}

	return 0
}

func readInput(file string) []int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file : %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	nums  := make([]int, 0)
	for input.Scan() {
		n, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Number format err :%v", err)
			continue
		}
		nums = append(nums, n)
	}
	return nums
}

