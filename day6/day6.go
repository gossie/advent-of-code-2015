package day6

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type rectangle struct {
	start point
	end   point
}

type action struct {
	instruction int
	rect        rectangle
}

func createRectangle(suffix string) rectangle {
	coordinates := strings.Split(suffix, " through ")
	start := strings.Split(coordinates[0], ",")
	startX, _ := strconv.Atoi(start[0])
	startY, _ := strconv.Atoi(start[1])

	end := strings.Split(coordinates[1], ",")
	endX, _ := strconv.Atoi(end[0])
	endY, _ := strconv.Atoi(end[1])

	return rectangle{start: point{x: startX, y: startY}, end: point{x: endX, y: endY}}
}

func readData(filename string) []action {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	actions := make([]action, 0, 300)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "turn on"):
			actions = append(actions, action{instruction: 0, rect: createRectangle(line[8:])})
		case strings.HasPrefix(line, "turn off"):
			actions = append(actions, action{instruction: 1, rect: createRectangle(line[9:])})
		case strings.HasPrefix(line, "toggle"):
			actions = append(actions, action{instruction: 2, rect: createRectangle(line[7:])})
		}
	}
	return actions
}

func Lights(file string) int {
	actions := readData(file)
	field := make([][]bool, 0, 1000)
	for i := 0; i <= 999; i++ {
		field = append(field, make([]bool, 1000, 1000))
	}

	for _, action := range actions {
		for y := action.rect.start.y; y <= action.rect.end.y; y++ {
			for x := action.rect.start.x; x <= action.rect.end.x; x++ {
				switch action.instruction {
				case 0:
					field[y][x] = true
				case 1:
					field[y][x] = false
				case 2:
					field[y][x] = !field[y][x]
				}
			}
		}
	}

	turnedOnLights := 0
	for _, row := range field {
		for _, bulp := range row {
			if bulp {
				turnedOnLights++
			}
		}
	}
	return turnedOnLights
}

func Brightness(file string) int {
	actions := readData(file)
	field := make([][]int, 0, 1000)
	for i := 0; i <= 999; i++ {
		field = append(field, make([]int, 1000, 1000))
	}

	for _, action := range actions {
		for y := action.rect.start.y; y <= action.rect.end.y; y++ {
			for x := action.rect.start.x; x <= action.rect.end.x; x++ {
				switch action.instruction {
				case 0:
					field[y][x]++
				case 1:
					field[y][x] = int(math.Max(float64(field[y][x]-1), 0.0))
				case 2:
					field[y][x] += 2
				}
			}
		}
	}

	brightness := 0
	for _, row := range field {
		for _, b := range row {
			brightness += b
		}
	}
	return brightness
}
