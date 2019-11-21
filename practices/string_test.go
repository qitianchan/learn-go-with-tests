package main;

import "testing"
import "strings"

func assetCorrectMessage(got, expected int, t *testing.T) {
	if got != expected {
		t.Errorf("expected '%d', got '%d'", expected, got)
	}
} 

func TestCompare(t *testing.T) {
	t.Run("compare equal", func (t *testing.T) {
		got := strings.Compare("a", "a")
		expected := 0
		assetCorrectMessage(got, expected, t)
	})

	t.Run("test a < b", func(t *testing.T) {
		got := strings.Compare("a", "b")
		expected := -1
		assetCorrectMessage(got, expected, t)
	})
	
	t.Run("test a > b", func(t *testing.T) {
		got := strings.Compare("b", "a")
		expected := 1
		assetCorrectMessage(got, expected, t)
	})
}



func BenchmarkCampare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Compare("abdc", "acbd")	
	}
}