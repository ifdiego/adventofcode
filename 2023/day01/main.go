package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	total := 0
	total2 := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		regex := regexp.MustCompile("[0-9]+")
		elements := regex.FindAllString(scanner.Text(), -1)
		lines := strings.Join(elements, "")
		numbers := strings.Split(lines, "")

		var last = len(numbers) - 1
		calibrationValue, _ := strconv.Atoi(numbers[0] + numbers[last])
		total = total + calibrationValue

		lines2 := formatLetters(scanner.Text())
		numbers2 := strings.Split(lines2, "")

		var last2 = len(numbers2) - 1
		calibrationValue2, _ := strconv.Atoi(numbers2[0] + numbers2[last2])
		total2 = total2 + calibrationValue2
	}

	fmt.Println("part one: ", total)
	fmt.Println("part two: ", total2)
}
