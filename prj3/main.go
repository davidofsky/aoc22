package main

import (
	"fmt"
	"prj3/rucksack"
)

func main() {
	fmt.Print("Enter the file location: ")
	var puzzleLocation string
	fmt.Scanln(&puzzleLocation)

	rucksacks := rucksack.GetRucksacks(puzzleLocation)

	groups := rucksack.CreateGroups(rucksacks)

	sum := 0
	for _, g := range groups {
		fmt.Println(g)
		badge := rucksack.GetBadge(g)
		priority := rucksack.ItemPriority(badge)
		sum += priority

	}
	fmt.Println(sum)

	// sum := 0
	// for _, r := range rucksacks {
	// 	splitItems := rucksack.SplitItems(r.Items)
	// 	fmt.Println(splitItems)
	// 	rItem := rucksack.ReoccuringItem(splitItems)
	// 	fmt.Println(`Reoccuring Item: ` + rItem)
	// 	sum += rucksack.ItemPriority(rItem)
	// }

	// fmt.Println(sum)

}
