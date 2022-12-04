package rps

import (
	"log"
	"os"
	"strings"
)

type Move = struct {
	Move  string
	Score int
}

type Round = struct {
	Their Move
	Mine  Move
}

// read puzzle file and return list of rounds
func GetRounds(location string) []Round {

	content, err := os.ReadFile(location)

	if err != nil {
		log.Fatal(err)
	}

	puzzle := strings.Split(string(content), "\n")

	var rounds []Round
	for _, set := range puzzle {
		var theirScore int
		var myScore int

		splitSet := strings.Split(set, " ")

		switch splitSet[0] {
		case "A":
			theirScore = 1
		case "B":
			theirScore = 2
		case "C":
			theirScore = 3
		}

		switch splitSet[1] {
		case "X":
			myScore = 1
		case "Y":
			myScore = 2
		case "Z":
			myScore = 3
		}

		newRound := Round{
			Their: Move{Move: splitSet[0], Score: theirScore},
			Mine:  Move{Move: splitSet[1], Score: myScore},
		}

		rounds = append(rounds, newRound)
	}
	return rounds
}
