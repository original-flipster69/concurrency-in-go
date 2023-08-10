package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {

	out1 := read("random.csv")
	out2 := read("random2.csv")

	out := merge(out1, out2)

	exit := make(chan struct{})

	go func() {
		for i := range out {
			fmt.Println(i)
		}
		close(exit)
	}()
	<-exit

	fmt.Println("All done... good night")
}

func read(file string) <-chan []string {
	f, err := os.Open(file)
	if err != nil {
		panic("fuck")
	}

	ch := make(chan []string)
	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)
				return
			}

			ch <- record
		}
	}()

	return ch
}

func merge(cs ...<-chan []string) <-chan []string {
	var wg sync.WaitGroup
	out := make(chan []string)

	send := func(c <-chan []string) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	return out
}
