package number

import (
	"math"
	"strconv"
	"strings"
)

func isMultiple(n int) func(int) bool {
	return func(i int) bool {
		if i%n == 0 {
			return true
		}
		return false
	}
}

func isSpecifiedNum(n int) func(int) bool {
	return func(i int) bool {
		if i == n {
			return true
		}
		return false
	}
}

func isOdd(i int) bool {
	if i%2 == 1 {
		return true
	}
	return false
}

func isEven(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func isBetweenAandB(a, b int) (checker func(int) bool) {
	return func(i int) bool {
		if a <= i && i <= b {
			return true
		}
		return false
	}
}

func parseNumType(condition string, checkers []func(int) bool) ([]func(int) bool, bool) {
	switch condition {
	case "odd":
		return append(checkers, isOdd), true
	case "even":
		return append(checkers, isEven), true
	}
	return checkers, false
}

func parseRange(condition string, checkers []func(int) bool) ([]func(int) bool, bool) {
	nums := strings.Split(condition, "~")
	if len(nums) < 2 {
		return checkers, false
	}

	min, err := strconv.Atoi(nums[0])
	if err != nil {
		min = 0
	}
	max, err := strconv.Atoi(nums[len(nums)-1])
	if err != nil {
		max = math.MaxInt64
	}
	return append(checkers, isBetweenAandB(min, max)), true
}

func parseOneNum(condition string, checkers []func(int) bool) ([]func(int) bool, bool) {
	num, err := strconv.Atoi(condition)
	if err != nil {
		return checkers, false
	}
	return append(checkers, isSpecifiedNum(num)), true
}

func parseMultipleNum(condition string, checkers []func(int) bool) ([]func(int) bool, bool) {
	if strings.HasSuffix(condition, "n") {
		num, err := strconv.Atoi(condition[:len(condition)-1])
		if err == nil {
			return append(checkers, isMultiple(num)), true
		}
	}
	return checkers, false
}

func parseRanges(ranges []string) []func(int) bool {
	checkers := []func(int) bool{}
	var ok bool // To avoid 'checkerss, ok := ~'
	for _, condition := range ranges {
		checkers, ok = parseNumType(condition, checkers)
		if ok {
			continue
		}
		checkers, ok = parseRange(condition, checkers)
		if ok {
			continue
		}
		checkers, ok = parseOneNum(condition, checkers)
		if ok {
			continue
		}
		checkers, _ = parseMultipleNum(condition, checkers)
	}
	return checkers
}

func inRanges(i int, checkers []func(int) bool) bool {
	for _, cheker := range checkers {
		if cheker(i) {
			return true
		}
	}
	return false
}
