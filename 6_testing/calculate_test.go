package calculate

import "testing"

func TestAddTwo(t *testing.T) {
	given := 2
	want := 4

	get := addTwo(given)

	if want != get {
		t.Errorf("given %d want %d but %d\n", given, want, get)
	}
}
