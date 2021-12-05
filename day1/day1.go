package day1

import (
	"bufio"
	"os"
)

func readData(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func Floor(file string) int {
	floor := 0
	for _, direction := range readData(file) {
		if direction == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func PositionWhenBasementWasEntered(file string) int {
	floor := 0
	for index, direction := range readData(file) {
		if direction == '(' {
			floor++
		} else {
			floor--
			if floor == -1 {
				return index + 1
			}
		}
	}
	panic("never entered basement")
}
