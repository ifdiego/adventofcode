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

func main() {
	total := 0

	file, err := os.Open("day01/input.txt")
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
	}

	fmt.Println(total)
}
