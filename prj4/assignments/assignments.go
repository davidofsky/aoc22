package assignments

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type AssignmentPair = struct {
	First  Assignment
	Second Assignment
}

type Assignment = struct {
	Min int
	Max int
}

func ReadPuzzle(location string) []AssignmentPair {
	content, err := os.ReadFile(location)

	if err != nil {
		log.Fatal(err)
	}

	var assignmentPairs []AssignmentPair
	for _, pairs := range strings.Split(string(content), "\n") {
		splitted := strings.Split(pairs, ",")

		min, _ := strconv.Atoi(strings.Split(splitted[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(splitted[0], "-")[1])

		assignment1 := Assignment{
			Min: min,
			Max: max,
		}

		min, _ = strconv.Atoi(strings.Split(splitted[1], "-")[0])
		max, _ = strconv.Atoi(strings.Split(splitted[1], "-")[1])
		assignment2 := Assignment{
			Min: min,
			Max: max,
		}

		assignmentPairs = append(assignmentPairs, AssignmentPair{
			First:  assignment1,
			Second: assignment2,
		})
	}

	return assignmentPairs
}
