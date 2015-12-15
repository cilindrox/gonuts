package main

import "testing"

func TestMin(t *testing.T) {

	x := []int{4, 8, 15, 16, 23, 42}
	i := x[0]
	m := Min(x, i)
	e := 4

	if m != e {
		t.Errorf("Min(%d, %d) => %d, want %d", x, i, m, e)
	}
}

func TestMinLast(t *testing.T) {

	x := []int{99, 87, 1024, 1}
	i := x[0]
	m := Min(x, i)
	e := 1

	if m != e {
		t.Errorf("Min(%d, %d) => %d, want %d", x, i, m, e)
	}
}

func TestMinBookSample(t *testing.T) {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	i := x[0]
	m := Min(x, i)
	e := 9

	if m != e {
		t.Errorf("Min(%d, %d) => %d, want %d", x, i, m, e)
	}
}
