package main

import (
	"fmt"
	"prj4/assignments"
	"strconv"
)

func main() {

	var location string
	fmt.Print("Enter the path to the puzzle: ")
	fmt.Scanln(&location)

	pairs := assignments.ReadPuzzle(location)

	count := 0
	count2 := 0
	for _, p := range pairs {
		if (p.First.Min >= p.Second.Min && p.First.Max <= p.Second.Max) ||
			(p.Second.Min >= p.First.Min && p.Second.Max <= p.First.Max) {
			count++
		}

		if !(p.First.Max < p.Second.Min || p.Second.Max < p.First.Min) {
			count2++
		}
	}
	fmt.Println("Answer one " + strconv.Itoa(count))
	fmt.Println("Answer two " + strconv.Itoa(count2))

}
