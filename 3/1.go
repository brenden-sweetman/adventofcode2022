package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day31() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	prioritySum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Parse input
		sack := scanner.Text()
		fmt.Printf("Sack: %s  ", sack)
		fmt.Printf("Length: %d  ", len(sack))
		comp1 := sack[:len(sack)/2]
		comp2 := sack[len(sack)/2:]
		fmt.Printf("Comp1 %s  ", comp1)
		fmt.Printf("Comp2 %s  ", comp2)

		// get shared item
		var sharedItem byte
		for i := 0; i < len(comp1); i++ {
			for o := 0; o < len(comp2); o++ {
				if comp1[i] == comp2[o] {
					sharedItem = comp1[i]
				}
			}
		}
		fmt.Printf("Shared Item: %s  ", string(sharedItem))

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
