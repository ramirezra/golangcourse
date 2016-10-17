package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	fmt.Printf("Factorial %d!\n", seed-1)
	prime := gen(seed)

	f := factorial(prime)

	for n := range f {
		fmt.Println(n)
	}
}

func gen(a int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < a; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	output := make(chan int)
	go func() {
		for i := range in {
			output <- fact(i)
		}
		close(output)
	}()
	return output
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total

}
