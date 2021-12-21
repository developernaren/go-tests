package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T)  {

	expected := 5
	a := 2
	b := 3

	got := Add(a, b)

	if got != expected {
		panic(fmt.Sprintf("expected %d but got %d", expected, got))
	}
}
