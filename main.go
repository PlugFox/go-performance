package main

import (
	"fmt"
	"runtime"
)

// Fizz Buzz
func fizzBuzz(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return fmt.Sprint(i)
	}
}

func main() {
	runtime.SetBlockProfileRate(1)
	for i := 1; i <= 1000; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
