package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(file)), "\n")

	calc := func(time, distance []string) int {
		result := 1
		for i := range time {
			t, _ := strconv.ParseFloat(time[i], 64)
			d, _ := strconv.ParseFloat(distance[i], 64)
			b := math.Sqrt(math.Pow(t, 2) - 4*d)
			result *= int(math.Ceil((t+b)/2) - math.Floor((t-b)/2) - 1)
		}
		return result
	}

	fmt.Println("part one: ", calc(strings.Fields(split[0])[1:], strings.Fields(split[1])[1:]))
	fmt.Println("part two: ", calc([]string{strings.Join(strings.Fields(split[0])[1:], "")}, []string{strings.Join(strings.Fields(split[1])[1:], "")}))
}
