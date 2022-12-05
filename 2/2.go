package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func outcomeScore() {
	// A X ROCK
	// B Y PAPER
	// C Z SCISSORS
	drawMap := map[string]string{
		"A": "A",
		"B": "B",
		"C": "C",
	}
	loseMap := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
	winMap := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
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
		outcome := picks[1]
		pick2 := ""
		if outcome == "X" {
			pick2 = loseMap[pick1]
		} else if outcome == "Y" {
			pick2 = drawMap[pick1]
		} else {
			pick2 = winMap[pick1]
		}
		score := roundScoreMap[fmt.Sprintf("%s%s", pick1, pick2)] + pickScoreMap[pick2]
		scoreSum += int(score)
	}
	fmt.Printf("Total Score by outcome: %d\n", scoreSum)
}
