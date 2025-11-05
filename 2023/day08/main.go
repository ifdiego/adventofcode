package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var instructions string
	nodes := make(map[string][2]string)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		instructions = scanner.Text()
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			continue
		}
		node := parts[0]
		edges := strings.Trim(parts[1], "()")
		edgeParts := strings.Split(edges, ", ")
		if len(edgeParts) != 2 {
			continue
		}
		nodes[node] = [2]string{edgeParts[0], edgeParts[1]}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	currentNode := "AAA"
	steps := 0
	instructionIndex := 0
	instructionsLength := len(instructions)

	for currentNode != "ZZZ" {
		currentInstruction := instructions[instructionIndex]

		if currentInstruction == 'L' {
			currentNode = nodes[currentNode][0]
		} else if currentInstruction == 'R' {
			currentNode = nodes[currentNode][1]
		}

		instructionIndex = (instructionIndex + 1) % instructionsLength
		steps++
	}

	fmt.Println(steps)
}
