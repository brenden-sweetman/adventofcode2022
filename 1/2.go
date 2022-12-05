package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getTop3Cals() {
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
	top3 := [...]int{0, 0, 0}
	for e := calList.Front(); e != nil; e = e.Next() {
		currentValue := e.Value.(int)
		dif := 0
		minDifIndex := -1
		for i := 0; i < len(top3); i++ {
			dif = currentValue - top3[i]
			if dif > 0 && minDifIndex == -1 {
				minDifIndex = i
			} else if dif > 0 && dif < (currentValue-top3[minDifIndex]) {
				minDifIndex = i
			} else {
				continue
			}
		}
		if minDifIndex != -1 {
			top3[minDifIndex] = currentValue
		}
	}
	fmt.Printf("Top 3: %v\n", top3)
	sum = 0
	for i := 0; i < len(top3); i++ {
		sum = sum + top3[i]
	}
	fmt.Printf("Sum: %v\n", sum)
}
