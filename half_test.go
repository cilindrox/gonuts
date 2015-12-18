package main

import "testing"

var halfTests = []struct {
	in   int
	out1 int
	out2 bool
}{
	{1, 0, false},
	{2, 1, true},
}

func TestHalf(t *testing.T) {
	for _, tt := range halfTests {
		h, e := Half(tt.in)
		if h != tt.out1 || e != tt.out2 {
			t.Errorf("Half(%d) => %d %t, want %d %t", tt.in, h, e, tt.out1, tt.out2)
		}
	}
}

var oddTests = []struct {
	in  int
	out uint
}{
	{1, 1},
	{2, 3},
	{3, 5},
}

func TestOddGenerator(t *testing.T) {
	nextOdd := OddGenerator()
	for _, tt := range oddTests {
		if i := nextOdd(); i != tt.out {
			t.Errorf("OddGenerator [%d] => %d, want %d", tt.in, i, tt.out)
		}
	}
}

var fibTests = []struct {
	in  uint
	out uint
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
}

func TestFibonacci(t *testing.T) {
	for _, tt := range fibTests {
		if fib := Fibonacci(tt.in); fib != tt.out {
			t.Errorf("Fibonacci(%d) => %d, want %d", tt.in, fib, tt.out)
		}
	}
}
