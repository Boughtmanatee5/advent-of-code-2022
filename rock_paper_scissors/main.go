package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Win  = "Z"
	Draw = "Y"
	Lose = "X"

	Rock    = "A"
	Paper   = "B"
	Scissor = "C"

	RockScore    = 1
	PaperScore   = 2
	ScissorScore = 3

	WinScore  = 6
	DrawScore = 3
	LoseScore = 0
)

var (
	RockMap = map[string]string{
		Win:  Paper,
		Draw: Rock,
		Lose: Scissor,
	}

	PaperMap = map[string]string{
		Win:  Scissor,
		Draw: Paper,
		Lose: Rock,
	}

	ScissorMap = map[string]string{
		Win:  Rock,
		Draw: Scissor,
		Lose: Paper,
	}

	PlaysMap = map[string]map[string]string{
		Rock:    RockMap,
		Paper:   PaperMap,
		Scissor: ScissorMap,
	}
)

func main() {
	input, readErr := os.OpenFile("./input", os.O_RDONLY, os.ModePerm)
	if readErr != nil {
		fmt.Printf("error reading file %s", readErr)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	// solution := partOne(scanner)
	solution := partTwo(scanner)

	fmt.Printf("answer %d\n", solution)
}

func partOne(s *bufio.Scanner) int {
	var totalScore int
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}

		plays := strings.Split(line, " ")

		switch plays[1] {
		case "X":
			switch plays[0] {
			case "A":
				totalScore += 4
			case "B":
				totalScore += 1
			case "C":
				totalScore += 7
			}
		case "Y":
			switch plays[0] {
			case "A":
				totalScore += 8
			case "B":
				totalScore += 5
			case "C":
				totalScore += 2
			}
		case "Z":
			switch plays[0] {
			case "A":
				totalScore += 3
			case "B":
				totalScore += 9
			case "C":
				totalScore += 6
			}
		}
	}

	return totalScore
}

func partTwo(s *bufio.Scanner) int {
	var totalScore int
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			continue
		}

		plays := strings.Split(line, " ")

		var outcomeScore int
		switch plays[1] {
		case Win:
			outcomeScore = WinScore
		case Draw:
			outcomeScore = DrawScore
		case Lose:
			outcomeScore = LoseScore
		}

		playMapping := PlaysMap[plays[0]]
		play := playMapping[plays[1]]
		var playScore int
		switch play {
		case Rock:
			playScore = RockScore
		case Paper:
			playScore = PaperScore
		case Scissor:
			playScore = ScissorScore
		}

		totalScore += playScore + outcomeScore
	}

	return totalScore
}
