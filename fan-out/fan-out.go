package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	fileCh := read("random.csv")

	split1 := split("split 1", fileCh)
	split2 := split("split 2", fileCh)
	split3 := split("split 3", fileCh)

	for {
		if split1 == nil && split2 == nil && split3 == nil {
			break
		}

		select {
		case _, ok := <-split1:
			if !ok {
				split1 = nil
			}
		case _, ok := <-split2:
			if !ok {
				split2 = nil
			}
		case _, ok := <-split3:
			if !ok {
				split3 = nil
			}
		}
	}

	fmt.Println("all done: good night")
}

func split(worker string, ch <-chan []string) chan struct{} {
	out := make(chan struct{})

	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}

		close(out)
	}()

	return out
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
