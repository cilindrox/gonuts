package main

import "testing"

var fizzbuzzTests = []struct {
	in  int
	out string
}{
	{1, "1"},
	{3, "Fizz"},
	{5, "Buzz"},
	{75, "FizzBuzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range fizzbuzzTests {
		s := FizzBuzz(tt.in)
		if s != tt.out {
			t.Errorf("Fizzbuzz(%d) => %s, want %s", tt.in, s, tt.out)
		}
	}
}
