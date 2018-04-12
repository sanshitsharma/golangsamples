package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type tContext struct {
	shutdown chan struct{}
	wg       *sync.WaitGroup
}

var ctx *tContext

func threadRoutine(id string) {
	for signal := false; !signal; {
		select {
		case res, ok := <-ctx.shutdown:
			if ok {
				fmt.Println(res)
			} else {
				fmt.Println("channel closed")
				signal = true
			}
		case <-time.After(1 * time.Second):
			fmt.Println("timeout " + id)
		}
	}

	ctx.wg.Done()
	fmt.Println(`exiting thread ` + id + ` routine`)
}

func blockingCall() {
	fmt.Println("Just checking something..")

	for i := 1; i <= 2; i++ {
		ctx.wg.Add(1)
		id := strconv.FormatInt(int64(i), 10)
		go threadRoutine(id)
	}

	fmt.Println("Returning from blocking call")
}

func main() {
	ctx = &tContext{
		shutdown: make(chan struct{}, 1),
		wg:       new(sync.WaitGroup),
	}

	blockingCall()
	fmt.Println("Threads spawned")
	time.Sleep(4 * time.Second)
	fmt.Println("Woke Up.. Closing channel..")
	close(ctx.shutdown)

	ctx.wg.Wait()
	fmt.Println("exiting execution")
}
