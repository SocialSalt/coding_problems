package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type P struct {
	x                        int32
	y                        int64
	test                     int32
	this_is_a_very_long_name string
}

func primeSieve(n int, buffer int) chan int {
	c := make(chan int, buffer)
	isNumComposite := make([]bool, n)
	go func() {
		c <- 2
		for i := 3; i < n; i += 2 {
			if !isNumComposite[i] {
				c <- i
				for k := i * i; k < n; k += i {
					isNumComposite[k] = true
				}
			}
		}
		close(c)
	}()
	return c
}

func isPrimeSieve(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	lim := int(math.Sqrt(float64(n))) + 1
	isNumComposite := make([]bool, lim)
	for i := 3; i < lim; i += 2 {
		if !isNumComposite[i] {
			if n%i == 0 {
				return false
			}
			for k := i * i; k < lim; k += i {
				isNumComposite[k] = true
			}
		}
	}
	return true
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	lim := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < lim; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func timeIsPrimeSieve(N int, L int) {
	rand.Seed(1)
	start := time.Now()
	for i := 0; i < N; i++ {
		num := rand.Intn(L)
		isPrimeSieve(num)
	}
	fmt.Printf("Found primality of %v random numbers in %v seconds\n", N, time.Now().Sub(start).Seconds())
}

func timeIsPrime(N int, L int) {
	rand.Seed(1)
	start := time.Now()
	for i := 0; i < N; i++ {
		num := rand.Intn(L)
		isPrime(num)
	}
	fmt.Printf("Found primality of %v random numbers in %v seconds\n", N, time.Now().Sub(start).Seconds())
}

func printPrimes() {
	var n int = 1000
	var c chan int = primeSieve(n, 2)
	for i := range c {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")
}

func ternary[K string | int | float64](val bool, ifTrue K, ifFalse K) K {
	if val {
		return ifTrue
	} else {
		return ifFalse
	}
}

func testIsPrime() {
	fmt.Printf("Testing 0: %v\n", ternary(!isPrime(0), "Passed", "Failed"))
	fmt.Printf("Testing 1: %v\n", ternary(!isPrime(1), "Passed", "Failed"))
	fmt.Printf("Testing 2: %v\n", ternary(isPrime(2), "Passed", "Failed"))
	fmt.Printf("Testing 3: %v\n", ternary(isPrime(3), "Passed", "Failed"))
	fmt.Printf("Testing 4: %v\n", ternary(!isPrime(4), "Passed", "Failed"))
	fmt.Printf("Testing 5: %v\n", ternary(isPrime(5), "Passed", "Failed"))
	fmt.Printf("Testing 101: %v\n", ternary(isPrime(101), "Passed", "Failed"))
	fmt.Printf("Testing 121: %v\n", ternary(!isPrime(121), "Passed", "Failed"))
	fmt.Printf("Testing 2707: %v\n", ternary(isPrime(2707), "Passed", "Failed"))
}

func testIsPrimeSieve() {
	fmt.Printf("Testing 0: %v\n", ternary(!isPrimeSieve(0), "Passed", "Failed"))
	fmt.Printf("Testing 1: %v\n", ternary(!isPrimeSieve(1), "Passed", "Failed"))
	fmt.Printf("Testing 2: %v\n", ternary(isPrimeSieve(2), "Passed", "Failed"))
	fmt.Printf("Testing 3: %v\n", ternary(isPrimeSieve(3), "Passed", "Failed"))
	fmt.Printf("Testing 4: %v\n", ternary(!isPrimeSieve(4), "Passed", "Failed"))
	fmt.Printf("Testing 5: %v\n", ternary(isPrimeSieve(5), "Passed", "Failed"))
	fmt.Printf("Testing 101: %v\n", ternary(isPrimeSieve(101), "Passed", "Failed"))
	fmt.Printf("Testing 121: %v\n", ternary(!isPrimeSieve(121), "Passed", "Failed"))
	fmt.Printf("Testing 2707: %v\n", ternary(isPrimeSieve(2707), "Passed", "Failed"))
}

func main() {
	// printPrimes()
	// testIsPrime()
	// testIsPrimeSieve()
	timeIsPrimeSieve(1_000_000, 100_000_000)
	timeIsPrime(1_000_000, 1_000_000)
}
