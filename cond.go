package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

func RunCond(){
    cond := sync.NewCond(new(sync.Mutex))
    quitChan := make(chan string)
    var i int
    go func() {
        cond.L.Lock()
        for i != 7 {
            fmt.Printf("got %v :(\n", i)
            cond.Wait()
        }
        fmt.Println("got 7!")
        cond.L.Unlock()
        quitChan <- "yes quit"
    }()
    for range time.Tick(time.Second) {
        cond.L.Lock()
        i = rand.Intn(10)
        cond.L.Unlock()
        cond.Broadcast()
        if i == 7 {
            msg := <-quitChan
            fmt.Println(msg)
            break
        }
    }
}