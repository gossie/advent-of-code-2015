package day4

import (
	"testing"
)

func TestMine(t *testing.T) {
	proofOfWork := Mine("00000")
	if proofOfWork != 254575 {
		t.Fatalf("proofOfWork = %v", proofOfWork)
	}
}
