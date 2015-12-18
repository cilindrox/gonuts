gonuts
======

Playing around with a bit of Golang

[![Build Status](https://travis-ci.org/cilindrox/gonuts.svg)](https://travis-ci.org/cilindrox/gonuts)


## Overview

This repo is contains implementations to exercises proposed in [An Introduction to Programming in Go][book] which I've found interesting to share.


## Chapter 5 - Control Structures

`fizzbuzz.go`

2. Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)

```go
for i := 0; i < 100; i++ {
  if i%3 == 0 {
    fmt.Println(i)
  }
}
```

3. Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".

```go
for i := 0; i < 100; i++ {
  fmt.Println(FizzBuzz(i))
}
```


## Chapter 6 - Arrays, Slices and Maps

`min.go`

4. Write a program that finds the smallest number in this list:

```go
x := []int{
  48, 96, 86, 68,
  57, 82, 63, 70,
  37, 34, 83, 27,
  19, 97, 9, 17,
}

fmt.Println(Min(x, x[0]))
```

## Chapter 7 - Functions

`half.go`

2. Write a function which takes an integer and halves it and returns `true` if it was even or `false` if it was odd. For example, `half(1)` should return `(0, false)` and `half(2)` should return `(1, true)`.

3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.

```go
func greatest(numbers ...int) int {
  max := numbers[0]
  for _, number := range numbers {
    if number > max {
      max = number
    }
  }

  return max
}
```

4. Write a `makeOddGenerator` function that generates odd numbers.

5. Write a Fibonacci sequence function that can find `fib(n)`.


## Chapter 8 - Pointers

`swap.go`

5. Write a program that can swap two integers (`x := 1; y := 2; swap(&x, &y)` should give you `x=2` and `y=1`).

```go
func swap(x, y *int) {
  tmp := new(int)
  *tmp = *x
  *x = *y
  *y = *tmp
}
```



[book]: https://www.golang-book.com/books/intro
