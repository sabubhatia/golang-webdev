package main

import (
	"context"
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		log.Println("Waiting for GO routines to return...")
		wg.Wait()
		log.Println("Exiting..")		
	}()
	chi := gen(ctx)
	for v := range chi {
		log.Println("Val: ", v)
		if v > 5 {
			log.Println("Cancelling..")
			cancel()
			return
		}
	}
}

func gen(ctx context.Context) <-chan int {
	chi := make(chan int)

	val := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				if ctx.Err() != nil && ctx.Err() == context.Canceled {
					log.Println("Done()")
					close(chi)
					return
				}
				log.Fatal("Error:", ctx.Err())
			case chi <- val:
				val++
			}
		}
	}()

	return chi
}
