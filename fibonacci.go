package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	a := 0
	b := 1
	return func() int {
		i = a
		a = a + b
		b = i
		return i
	}

}

//run and print Fibonacci for 10 times.
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
