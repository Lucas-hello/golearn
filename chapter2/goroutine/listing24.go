package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) //用时间作为随机数种子


}

func main() {
	tasks := make(chan string, taskLoad)
	wg6.Add(numberGoroutine)
	for gr := 1; gr <= numberGoroutine; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	close(tasks)
	wg6.Wait()

}

var wg6 sync.WaitGroup

const (
	numberGoroutine = 4
	taskLoad        = 10
)

func worker(tasks chan string, worker int) {
	defer wg6.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker : %d :Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker : %d : Started %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker : %d : Completed %s\n", worker, task)
	}


}
