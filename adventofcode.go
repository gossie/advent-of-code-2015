package main

import (
	"fmt"

	"github.com/gossie/adventofcode2015/day1"
)

func main() {
	fmt.Println("Performing tasks of day 1")
	fmt.Println("Day 1, task 1: ", day1.Floor("day1/day1.txt"))
	fmt.Println("Day 1, task 2: ", day1.PositionWhenBasementWasEntered("day1/day1.txt"))
}
