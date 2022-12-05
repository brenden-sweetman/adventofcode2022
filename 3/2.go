package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day32() {
	prioritySum := 0

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Printf("Input Length: %d", len(input))

	for i := 0; i <= len(input)-3; i += 3 {
		var sharedItem byte
		for s := 0; s < len(input[i]); s++ {
			if strings.Contains(input[i+1], string(input[i][s])) && strings.Contains(input[i+2], string(input[i][s])) {
				sharedItem = input[i][s]
			}
		}
		fmt.Printf("Shared Item: %s\n", string(sharedItem))
		// get priority of shared item
		var priority int
		// lowercase
		if sharedItem >= byte(97) {
			priority = int(sharedItem) - 96
		}
		if sharedItem <= byte(90) {
			priority = int(sharedItem) - 38
		}
		fmt.Printf("Priority: %d\n", priority)
		prioritySum += priority
	}
	fmt.Printf("Total Priority: %d\n", prioritySum)

}
