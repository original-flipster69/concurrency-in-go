package main

import (
	"fmt"
)

func main() {

	out := square(generate(4, 6, 1, 2, 3))

	for i := range out {
		fmt.Println("done: ", i)
	}

}

func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range nums {
			fmt.Println("gen: ", i)
			out <- i
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			fmt.Println("sqr: ", i)
			out <- i * i
		}
		close(out)
	}()
	return out
}
