package main

// Half takes an integer and returns true if it is even or false if it was odd.
func Half(n int) (int, bool) {
	half := n / 2
	even := n%2 == 0
	return half, even
}

// OddGenerator returns a function that generates odd numbers.
func OddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// Fibonacci sequence
func Fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
