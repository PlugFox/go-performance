package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	// Fizz Buzz
	if fizzBuzz(3) != "Fizz" {
		t.Error("Expected Fizz, got ", fizzBuzz(3))
	}

	if fizzBuzz(5) != "Buzz" {
		t.Error("Expected Buzz, got ", fizzBuzz(5))
	}

	if fizzBuzz(15) != "FizzBuzz" {
		t.Error("Expected FizzBuzz, got ", fizzBuzz(15))
	}

	if fizzBuzz(1) != "1" {
		t.Error("Expected 1, got ", fizzBuzz(1))
	}

	if fizzBuzz(2) != "2" {
		t.Error("Expected 2, got ", fizzBuzz(2))
	}

}
