package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	
	if expected != sum {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}