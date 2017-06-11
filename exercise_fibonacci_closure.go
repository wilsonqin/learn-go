package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    prev := 0
    cur := 0
    next := 1
    
    
    return func() int {
        prev = cur
        cur = next
        next = prev + cur
        return prev
    }
}

func Run_fibonacci() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
