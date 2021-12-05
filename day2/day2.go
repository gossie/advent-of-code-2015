package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type dimension struct {
	length int
	width  int
	height int
}

func readData(filename string) []dimension {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	var dimensions []dimension
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		factors := strings.Split(scanner.Text(), "x")
		l, err := strconv.Atoi(factors[0])
		if err != nil {
			fmt.Println("err", err)
			panic("length cannot be parsed")
		}

		w, err := strconv.Atoi(factors[1])
		if err != nil {
			panic("width cannot be parsed")
		}

		h, err := strconv.Atoi(factors[2])
		if err != nil {
			panic("height cannot be parsed")
		}

		dimensions = append(dimensions, dimension{length: l, width: w, height: h})
	}
	return dimensions
}

func WrappingPaper(file string) int {
	squareFeet := 0
	for _, d := range readData(file) {
		side1 := d.length * d.width
		side2 := d.width * d.height
		side3 := d.length * d.height

		squareFeet += 2*side1 + 2*side2 + 2*side3 + int(math.Min(math.Min(float64(side1), float64(side2)), float64(side3)))
	}
	return squareFeet
}

func Ribbon(file string) int {
	ribbon := 0
	for _, d := range readData(file) {
		length1 := 0
		length2 := 0

		if d.length < d.width {
			length1 = d.length
			if d.width < d.height {
				length2 = d.width
			} else {
				length2 = d.height
			}
		} else {
			length1 = d.width
			if d.length < d.height {
				length2 = d.length
			} else {
				length2 = d.height
			}
		}

		ribbon += 2*length1 + 2*length2 + d.length*d.width*d.height
	}
	return ribbon
}
