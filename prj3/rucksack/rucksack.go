package rucksack

import (
	"io/ioutil"
	"log"
	"strings"
)

type Group = struct {
	rucksacks []Rucksack
}

type Rucksack = struct {
	Items string
}

// read puzzle file and return list of rounds
func GetRucksacks(location string) []Rucksack {

	content, err := ioutil.ReadFile(location)

	if err != nil {
		log.Fatal(err)
	}

	puzzle := strings.Split(string(content), "\n")

	var rucksacks []Rucksack
	for _, itemList := range puzzle {
		newRucksack := Rucksack{
			Items: itemList,
		}

		rucksacks = append(rucksacks, newRucksack)
	}

	return rucksacks
}

func CreateGroups(rucksacks []Rucksack) []Group {
	groupAmount := len(rucksacks) / 3
	var groups []Group
	rIndex := 0
	for i := 0; i < groupAmount; i++ {
		var groupRucksacks []Rucksack
		groupRucksacks = rucksacks[rIndex : rIndex+3]
		groups = append(groups, Group{
			rucksacks: groupRucksacks,
		})
		rIndex += 3
	}
	return groups
}

func GetBadge(group Group) string {
	badge := ""
	for _, item := range group.rucksacks[0].Items {
		b1 := false
		b2 := false

		for _, indexItem := range group.rucksacks[1].Items {
			if indexItem == item {
				b1 = true
				break
			}
		}
		for _, indexItem := range group.rucksacks[2].Items {
			if indexItem == item {
				b2 = true
				break
			}
		}
		if b1 && b2 {
			badge = string(item)
			break
		}
	}
	return badge
}

func ItemPriority(item string) int {
	lowercaseAlphabet := "abcdefghijklmnopqrstuvwxyz"
	uppercaseAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	index := (strings.Index(lowercaseAlphabet, item))
	if index != -1 {
		return index + 1
	}

	index = (strings.Index(uppercaseAlphabet, item))
	if index != -1 {
		return index + 27
	}

	return 0
}

func ReoccuringItem(splitted [2]string) string {
	for _, i1 := range splitted[0] {
		for _, i2 := range splitted[1] {
			if i1 == i2 {
				return string(i1)
			}
		}
	}
	return ""
}

func SplitItems(items string) [2]string {
	var splitted [2]string

	splitted[0] = items[0 : len(items)/2]
	splitted[1] = items[len(items)/2:]

	return splitted
}
