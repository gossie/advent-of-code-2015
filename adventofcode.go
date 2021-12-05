package main

import (
	"fmt"

	"github.com/gossie/adventofcode2015/day1"
	"github.com/gossie/adventofcode2015/day2"
)

func main() {
	fmt.Println("Performing tasks of day 1")
	fmt.Println("Day 1, task 1: ", day1.Floor("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.PositionWhenBasementWasEntered("day1/day1.txt"))

	fmt.Println("Performing tasks of day 2")
	fmt.Println("Day 2, task 1: ", day2.WrappingPaper("day2/day2.txt"))
	fmt.Println("Day 2, task 2: ", day2.Ribbon("day2/day2.txt"))
}
