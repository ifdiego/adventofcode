package days

import (
	"fmt"
	"strings"
)

func Day08(input string) {
	var instructions string
	nodes := make(map[string][2]string)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions = lines[0]

	for _, line := range lines[1:] {
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

	fmt.Println("part one: ", steps)
}
