package main

import (
	"fmt"

	"github.com/gossie/adventofcode2015/day1"
	"github.com/gossie/adventofcode2015/day2"
	"github.com/gossie/adventofcode2015/day3"
	"github.com/gossie/adventofcode2015/day4"
	"github.com/gossie/adventofcode2015/day5"
)

func main() {
	fmt.Println("Performing tasks of day 1")
	fmt.Println("Day 1, task 1: ", day1.Floor("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.PositionWhenBasementWasEntered("day1/day1.txt"))

	fmt.Println("\nPerforming tasks of day 2")
	fmt.Println("Day 2, task 1: ", day2.WrappingPaper("day2/day2.txt"))
	fmt.Println("Day 2, task 2: ", day2.Ribbon("day2/day2.txt"))

	fmt.Println("\nPerforming tasks of day 3")
	fmt.Println("Day 3, task 1: ", day3.Houses("day3/day3.txt"))
	fmt.Println("Day 3, task 2: ", day3.HousesWithRoboSanta("day3/day3.txt"))

	fmt.Println("\nPerforming tasks of day 4")
	fmt.Println("Day 4, task 1: ", day4.Mine("00000"))
	fmt.Println("Day 4, task 2: ", day4.Mine("000000"))

	fmt.Println("\nPerforming tasks of day 5")
	fmt.Println("Day 5, task 1: ", day5.NiceWordsOldRules("day5/day5.txt"))
	fmt.Println("Day 5, task 2: ", day5.NiceWordsNewRules("day5/day5.txt"))
}
