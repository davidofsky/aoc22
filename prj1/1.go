package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"prj1/elf"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Starting first puzzle")
	puzzle := readFile("./1.puzzle")
	if puzzle == "" {
		fmt.Println("failed to read puzzle file")
	}

	// strings.fields splits a string by spaces or newlines
	puzzleArr := strings.Split(puzzle, "\n")
	elfs := generateElfs(puzzleArr)

	winnerElfs := retrieveElfByCals(elfs, 3)

	total := 0
	for _, e := range winnerElfs {
		total = total + e.TotalCalories
	}
	println(total)
}

func retrieveElfByCals(elfs []elf.Elf, top int) []elf.Elf {
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].TotalCalories < elfs[j].TotalCalories
	})

	var winners []elf.Elf
	for i := len(elfs) - top; i < len(elfs); i++ {
		winners = append(winners, elfs[i])
	}

	return winners
}

func generateElfs(puzzleArr []string) []elf.Elf {
	var totalElfs []elf.Elf
	var newElf elf.Elf
	for _, calories := range puzzleArr {
		if calories == "" {
			newElf.TotalCalories = newElf.GetTotalCalories()
			totalElfs = append(totalElfs, newElf)
			newElf = elf.Elf{
				Calories: []int{},
			}
		} else {
			cals_64, _ := strconv.ParseInt(calories, 10, 0)
			cals := int(cals_64)
			newElf.Calories = append(newElf.Calories, cals)
		}
	}
	return totalElfs
}

func readFile(location string) string {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
