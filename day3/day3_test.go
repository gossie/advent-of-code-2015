package day3

import (
	"testing"
)

func TestHouses(t *testing.T) {
	houses := Houses("day3.txt")
	if houses != 2081 {
		t.Fatalf("houses = %v", houses)
	}
}

func TestHousesWithRoboSanta(t *testing.T) {
	houses := HousesWithRoboSanta("day3.txt")
	if houses != 2341 {
		t.Fatalf("houses = %v", houses)
	}
}
