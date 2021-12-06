package day5

import (
	"testing"
)

func TestNiceWordsOldRules(t *testing.T) {
	numberOfNiceStrings := NiceWordsOldRules("day5.txt")
	if numberOfNiceStrings != 258 {
		t.Fatalf("numberOfNiceStrings = %v", numberOfNiceStrings)
	}
}

func TestNiceWordsNewRules(t *testing.T) {
	numberOfNiceStrings := NiceWordsNewRules("day5.txt")
	if numberOfNiceStrings != 53 {
		t.Fatalf("numberOfNiceStrings = %v", numberOfNiceStrings)
	}
}
