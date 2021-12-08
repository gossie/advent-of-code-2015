package day6

import "testing"

func TestLight(t *testing.T) {
	lights := Lights("day6.txt")
	if lights != 400410 {
		t.Fatalf("lights = %v", lights)
	}
}

func TestBrightness(t *testing.T) {
	brightness := Brightness("day6.txt")
	if brightness != 15343601 {
		t.Fatalf("brightness = %v", brightness)
	}
}
