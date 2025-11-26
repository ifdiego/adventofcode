package main

import (
	"fmt"
	"github.com/ifdiego/adventofcode/2023/days"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <day>")
		return
	}
	day := os.Args[1]
	input, err := os.ReadFile(fmt.Sprintf("inputs/day%s.txt", day))
	if err != nil {
		panic(err)
	}
	switch day {
	case "01":
		days.Day01(string(input))
	case "02":
		days.Day02(string(input))
	case "04":
		days.Day04(string(input))
	case "06":
		days.Day06(string(input))
	case "08":
		days.Day08(string(input))
	default:
		fmt.Printf("day %s not yet implemented\n", day)
	}
}
