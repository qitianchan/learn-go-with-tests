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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793
		assertCorrectMessage(t, got, want)
	})
}
