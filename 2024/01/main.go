package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	file, _ := os.ReadFile("input.txt")

	var list1, list2 []int
	counts := map[int]int{}
	for _, s := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		var n1, n2 int
		fmt.Sscanf(s, "%d   %d", &n1, &n2)
		list1, list2 = append(list1, n1), append(list2, n2)
		counts[n2]++
	}

	slices.Sort(list1)
	slices.Sort(list2)

	part1, part2 := 0, 0
	for i := range list1 {
		part1 += abs(list2[i] - list1[i])
		part2 += list1[i] * counts[list1[i]]
	}
	fmt.Println("part one: ", part1)
	fmt.Println("part two: ", part2)
}
