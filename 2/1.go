package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func topScore() {
	// A X ROCK
	// B Y PAPER
	// C Z SCISSORS
	pickMap := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
	roundScoreMap := map[string]int8{
		"AA": 3,
		"AB": 6,
		"AC": 0,
		"BA": 0,
		"BB": 3,
		"BC": 6,
		"CA": 6,
		"CB": 0,
		"CC": 3,
	}
	pickScoreMap := map[string]int8{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scoreSum := 0
	for scanner.Scan() {
		picks := strings.Split(scanner.Text(), " ")
		pick1 := picks[0]
		pick2 := pickMap[picks[1]]
		score := roundScoreMap[fmt.Sprintf("%s%s", pick1, pick2)] + pickScoreMap[pick2]
		scoreSum += int(score)
	}
	fmt.Printf("Total Score: %d\n", scoreSum)
}
