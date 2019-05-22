package ymdhmstimeparse

import (
	"time"
)

func toDigit(ch rune) (v int, ok bool) {
	if (ch >= '0') && (ch <= '9') {
		return int(ch - '0'), true
	}
	return 0, false
}

func nextInt(inputValue []rune, inputLen int, currentIndex, maxDigit int) (result, nextIndex int) {
	nextIndex = currentIndex
	digitCount := 0
	for (nextIndex < inputLen) && (digitCount < maxDigit) {
		if v, ok := toDigit(inputValue[nextIndex]); ok {
			result = (result * 10) + v
			digitCount++
		} else if digitCount > 0 {
			digitCount = maxDigit // having nextIndex increase before leave loop
		}
		nextIndex++
	}
	return
}

// Parse given string in YYYY-MM-DD hh:mm:ss form
func Parse(value string, loc *time.Location) time.Time {
	// var year, month, day, hour, minute, second int
	inputValue := []rune(value)
	inputLen := len(value)
	year, nextIndex := nextInt(inputValue, inputLen, 0, 4)
	if year < 50 {
		year = 2000 + year
	} else if year < 100 {
		year = 1900 + year
	} else if year < 1000 {
		year = 1000 + year
	}
	month, nextIndex := nextInt(inputValue, inputLen, nextIndex, 2)
	day, nextIndex := nextInt(inputValue, inputLen, nextIndex, 2)
	hour, nextIndex := nextInt(inputValue, inputLen, nextIndex, 2)
	minute, nextIndex := nextInt(inputValue, inputLen, nextIndex, 2)
	second, nextIndex := nextInt(inputValue, inputLen, nextIndex, 2)
	if nil == loc {
		loc = time.UTC
	}
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, loc)
}
