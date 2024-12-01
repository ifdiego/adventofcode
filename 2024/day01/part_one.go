package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    var left []int;
    var right []int;
    
	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Println("ERR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        lvalue, _ := strconv.Atoi(fields[0])
        rvalue, _ := strconv.Atoi(fields[1])
        left = append(left, lvalue)
        right = append(right, rvalue)
	}

    sort.Ints(left)
    sort.Ints(right)

    distance := 0
    for i := range(left) {
        if left[i] >= right[i] {
            distance += left[i] - right[i]
        } else {
            distance += right[i] - left[i]
        }
    }

    fmt.Println(distance)
}
