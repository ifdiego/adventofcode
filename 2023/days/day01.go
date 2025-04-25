package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func formatLetters(s string) string {
	newString := ""
	searches := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for char := range s {
		for key, value := range searches {
			if strings.HasPrefix(s[char:], key) {
				newString = newString + value
			}
		}
	}
	return newString
}

func sumCalibrationValue(lines []string, extract func(string) string) int {
	total := 0

	for _, line := range lines {
		digits := extract(line)
		numbers := strings.Split(digits, "")

		if len(numbers) == 0 {
			continue
		}

		var last = len(numbers) - 1
		calibrationValue, err := strconv.Atoi(numbers[0] + numbers[last])
		if err == nil {
			total += calibrationValue
		}
	}

	return total
}

func extractDigitsWithRegex(line string) string {
	regex := regexp.MustCompile("[0-9]+")
	elements := regex.FindAllString(line, -1)
	return strings.Join(elements, "")
}

func Day01(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	partOne := sumCalibrationValue(lines, extractDigitsWithRegex)
	fmt.Printf("part one: %v\n", partOne)

	partTwo := sumCalibrationValue(lines, formatLetters)
	fmt.Printf("part two: %v\n", partTwo)
}
