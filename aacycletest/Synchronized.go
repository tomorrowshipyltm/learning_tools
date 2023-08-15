package main

import (
	"fmt"
	"sync"
	"time"
)

type BaseHandler interface {
	run()
}

type SynchroHandler struct{}

type ChanHandler struct{}

func (syn SynchroHandler) run() {
	//实现interface方法
	var na = "hello"
	var wg sync.WaitGroup
	wg.Add(2)
	go func(name string) {
		defer wg.Done()
		fmt.Println("method1: ", name)
	}(na)

	go func(name string) {
		defer wg.Done()
		name += " world"
		fmt.Println("method2: ", name)
	}(na)

	wg.Wait()
	fmt.Println("success finished")
}

func (h ChanHandler) run() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "send ch1"
	}()
	go func() {
		ch2 <- "send ch2"
	}()

	// select只能指向一次，所有需要for循环
	for i := 0; i < 3; i++ {
		select {
		case re1 := <-ch1:
			{
				fmt.Println("receive from ch1: ", re1)
			}
		case re2 := <-ch2:
			{
				fmt.Println("receive from ch2: ", re2)
			}
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
		}

	}
}
