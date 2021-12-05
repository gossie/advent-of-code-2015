package day1

import (
	"testing"
)

func TestFloor(t *testing.T) {
	floor := Floor("day1.txt")
	if floor != 138 {
		t.Fatalf("floor = %v", floor)
	}
}

func TestDetermineNumberOfLargerSumedMeasurements(t *testing.T) {
	position := PositionWhenBasementWasEntered("day1.txt")
	if position != 1771 {
		t.Fatalf("position = %v", position)
	}
}
