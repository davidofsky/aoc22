package elf

type Elf struct {
	Calories      []int
	TotalCalories int
}

func (e Elf) GetTotalCalories() int {
	total := 0
	for _, c := range e.Calories {
		total = total + c
	}
	return total
}
