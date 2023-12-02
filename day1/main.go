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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		log.Println("before:", scanner.Text())
		regex := regexp.MustCompile("[0-9]+")
		replace := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		str := replace.Replace(scanner.Text())
		log.Println("after", str)
		elements := regex.FindAllString(str, -1)
		lines := strings.Join(elements, "")
		numbers := strings.Split(lines, "")
		log.Println("numbers:", numbers)

		var last = len(numbers) - 1
		cv, _ := strconv.Atoi(numbers[0] + numbers[last])
		log.Println("first:", numbers[0], "/last:", numbers[last], "/calibration value:", cv)
		total = total + cv
		log.Println("sum:", total)
	}
	fmt.Println("total", total)
}
