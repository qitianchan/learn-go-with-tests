package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got: %.2f, want: %.2f", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := rectangle.Perimeter()
	want := 40.0
	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f, want %.2f", got, tt.want)
		}
	}
}
