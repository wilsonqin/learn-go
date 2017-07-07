package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func walk_helper(t *tree.Tree, ch chan int) {
    if t != nil {
        walk_helper(t.Left, ch)
        ch <- t.Value
        walk_helper(t.Right, ch)
    }
    return
}

// Use Left Traversal of Tree, Ascending Value Order
func Walk(t *tree.Tree, ch chan int) {
    walk_helper(t, ch)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)

    for {
        v1, ok1 := <-c1
        v2, ok2 := <-c2
        if (ok1 != ok2) || (v1 != v2) {
            return false
        }

        if ok1 == true {
            break
        }
    }
    return true
}

func RunBinaryTrees() {
    // Test Binary Tree Walks in order
    ch := make(chan int)
    go Walk(tree.New(1), ch)

    for a := range ch{
       fmt.Println(a)
    }

    // Test same function works
    t1a := tree.New(1)
    t1b := tree.New(1)
    t2 := tree.New(2)

    fmt.Println("t1a == t1b", Same(t1a, t1b))
    fmt.Println("t1a == t2", Same(t1a, t2))
    fmt.Println("t2 == t1b", Same(t2, t1b))
}