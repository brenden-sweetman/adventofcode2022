package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getTopCal() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	calList := list.New()
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			calList.PushBack(sum)
			sum = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			sum = sum + i
		}
	}
	maxCal := 0
	for e := calList.Front(); e != nil; e = e.Next() {
		currentValue := e.Value.(int)
		if currentValue > maxCal {
			maxCal = currentValue
		}
	}
	fmt.Printf("Top: %v\n", maxCal)
}
