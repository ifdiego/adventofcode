package main

import (
	"fmt"
	"github.com/ifdiego/adventofcode/2023/days"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run . <day>")
		// log.Println("ERROR:", err)
		// os.Exit(1)
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
	default:
		fmt.Print("day %s not yet implemented\n", day)
	}
}
