package day2

import (
	"testing"
)

func TestWrappingPaper(t *testing.T) {
	wrappingPaper := WrappingPaper("day2.txt")
	if wrappingPaper != 1588178 {
		t.Fatalf("wrappingPaper = %v", wrappingPaper)
	}
}

func TestRibbon(t *testing.T) {
	ribbon := Ribbon("day2.txt")
	if ribbon != 3783758 {
		t.Fatalf("ribbon = %v", ribbon)
	}
}
