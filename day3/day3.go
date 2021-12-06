package day3

import (
	"bufio"
	"os"
)

type point struct {
	x int
	y int
}

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

func Houses(file string) int {
	current := point{x: 0, y: 0}
	houses := make(map[point]int)
	houses[current] = 1
	for _, direction := range readData(file) {
		switch direction {
		case '^':
			current = point{x: current.x, y: current.y - 1}
			houses[current] = houses[current] + 1
		case '>':
			current = point{x: current.x + 1, y: current.y}
			houses[current] = houses[current] + 1
		case '<':
			current = point{x: current.x - 1, y: current.y}
			houses[current] = houses[current] + 1
		case 'v':
			current = point{x: current.x, y: current.y + 1}
			houses[current] = houses[current] + 1
		}
	}

	return len(houses)
}

func HousesWithRoboSanta(file string) int {
	santasPosition := point{x: 0, y: 0}
	roboSantasPosition := point{x: 0, y: 0}
	houses := make(map[point]int)
	houses[santasPosition] = 2
	for i, direction := range readData(file) {
		switch direction {
		case '^':
			if i%2 == 0 {
				santasPosition = point{x: santasPosition.x, y: santasPosition.y - 1}
				houses[santasPosition] = houses[santasPosition] + 1
			} else {
				roboSantasPosition = point{x: roboSantasPosition.x, y: roboSantasPosition.y - 1}
				houses[roboSantasPosition] = houses[roboSantasPosition] + 1
			}
		case '>':
			if i%2 == 0 {
				santasPosition = point{x: santasPosition.x + 1, y: santasPosition.y}
				houses[santasPosition] = houses[santasPosition] + 1
			} else {
				roboSantasPosition = point{x: roboSantasPosition.x + 1, y: roboSantasPosition.y}
				houses[roboSantasPosition] = houses[roboSantasPosition] + 1
			}
		case '<':
			if i%2 == 0 {
				santasPosition = point{x: santasPosition.x - 1, y: santasPosition.y}
				houses[santasPosition] = houses[santasPosition] + 1
			} else {
				roboSantasPosition = point{x: roboSantasPosition.x - 1, y: roboSantasPosition.y}
				houses[roboSantasPosition] = houses[roboSantasPosition] + 1
			}
		case 'v':
			if i%2 == 0 {
				santasPosition = point{x: santasPosition.x, y: santasPosition.y + 1}
				houses[santasPosition] = houses[santasPosition] + 1
			} else {
				roboSantasPosition = point{x: roboSantasPosition.x, y: roboSantasPosition.y + 1}
				houses[roboSantasPosition] = houses[roboSantasPosition] + 1
			}
		}
	}

	return len(houses)
}
