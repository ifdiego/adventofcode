package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	result := 0

	file, err := os.ReadFile("day06/input.txt")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	t := strings.Split(string(file), "\n")[0]
	d := strings.Split(string(file), "\n")[1]
	regex := regexp.MustCompile("[0-9]+")
	times := strings.Join(regex.FindAllString(t, -1), "")
	distances := strings.Join(regex.FindAllString(d, -1), "")

	length, _ := strconv.Atoi(times)
	distance, _ := strconv.Atoi(distances)
	for i := 1; i <= length; i++ {
		milliseconds := length - i
		millimeter := milliseconds * i
		if millimeter > distance {
			result += 1
		}
	}

	fmt.Println(result)
}
