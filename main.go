package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func workerGenerator(worker int) <-chan int {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- worker*1000 + i
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(ch)
	}()

	return ch
}
func merge(cs ...<-chan int) <-chan int {
	ch := make(chan int)
	wg := new(sync.WaitGroup)

	for _, c := range cs {
		wg.Add(1)

		localC := c
		go func() {
			defer wg.Done()

			for in := range localC {
				ch <- in
			}
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func worke(id int, bread <-chan int, result chan<- int) {

	for i := range bread {
		fmt.Println("Worker id ", id, " make bread ", i)
		time.Sleep(time.Second)
		//fmt.Println("===========Second===========")
		fmt.Println("Worker id ", id, " done bread ", i)
		result <- i
	}
}
func main() {
	bread := make(chan int, 5)
	result := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go worke(i, bread, result)
	}

	for i := 0; i < 5; i++ {
		bread <- i
	}
	close(bread)
	time.Sleep(time.Second * 3)

	mergeA := merge(workerGenerator(1), workerGenerator(2), workerGenerator(3))
	for v := range mergeA {
		fmt.Println(v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Bread cooked", <-result)
	}
}
