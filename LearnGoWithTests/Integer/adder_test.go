package integer

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(4, 4)
	expected := 8

	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
}
