package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(12, 6)
	want := 6
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(6, 6)
	want := 36
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDivide(t *testing.T) {
	got := Divide(49, 7)
	want := 7
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestModulus(t *testing.T) {
	got := Modulus(7, 6)
	want := 1
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
