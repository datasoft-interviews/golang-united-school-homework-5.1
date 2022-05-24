package square

import (
	"testing"
)

func TestEnd(t *testing.T) {
	input := Square{Point{1, 1}, 2}
	want := Point{3, 3}
	if got := input.End(); got != want {
		t.Errorf("End() = %o, want %o", got, want)
	}
}

func TestArea(t *testing.T) {
	input := Square{Point{1, 1}, 3}
	want := uint(9)
	if got := input.Area(); got != want {
		t.Errorf("Area() = %o, want %o", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	input := Square{Point{1, 1}, 3}
	want := uint(12)
	if got := input.Perimeter(); got != want {
		t.Errorf("Perimeter() = %o, want %o", got, want)
	}
}
