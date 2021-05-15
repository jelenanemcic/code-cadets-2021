package tasks

import (
	"context"
	"log"
	"sync"
)

func RunTasks(tasks ...Task) {
	// run each task in separate goroutine
	// wait for all tasks to finish
	//
	// when first task finishes, signal to the other goroutines that application should stop

	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range tasks {
		go func(i int) {
			defer wg.Done()
			defer cancel()
			tasks[i].Start(ctx)
		}(i)
	}

	// kad dođe signal -> signalHandler se gasi -> radi se cancel nad kontekstom -> zbog toga se gasi feed pa onda i ostali taskovi

	wg.Wait()
	log.Print("all tasks finished")

}

type Task interface {
	Start(ctx context.Context) error
}
