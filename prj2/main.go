package main

import (
	"fmt"
	"prj2/rps"
)

func main() {
	fmt.Print("Enter the file location: ")
	var puzzleLocation string
	fmt.Scanln(&puzzleLocation)
	rounds := rps.GetRounds(puzzleLocation)

	fmt.Print("which question are you on (1 or 2): ")
	var question string
	fmt.Scanln(&question)

	score := 0
	if question == "1" {
		for _, round := range rounds {
			score += getScore(round)
		}
	} else {
		for _, round := range rounds {
			score += assumeScore(round)
		}
	}

	fmt.Println(score)
}

// part two
func assumeScore(round rps.Round) int {
	if round.Mine.Move == "X" { // lose
		if round.Their.Score == 1 {
			return 3
		} else {
			return round.Their.Score - 1
		}
	} else if round.Mine.Move == "Y" { // draw
		return round.Their.Score + 3
	} else { // win
		if round.Their.Score == 3 {
			return 1 + 6
		} else {
			return round.Their.Score + 1 + 6
		}
	}
}

// part one
func getScore(round rps.Round) int {
	if round.Mine.Score == round.Their.Score {
		return round.Mine.Score + 3
	}

	if round.Mine.Score == 1 && round.Their.Score == 3 {
		return round.Mine.Score + 6
	} else if round.Mine.Score == round.Their.Score+1 {
		return round.Mine.Score + 6
	}

	return round.Mine.Score
}
