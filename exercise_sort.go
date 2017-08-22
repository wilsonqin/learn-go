package main

import (
	"fmt"
	"sort"
)

func RunSortSearch() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

func RunSortSearch2() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 56

	i := sort.Search(len(a), func(i int) bool { return a[i] > x })
	if (i > 0) && !(a[i-1] < x) {
		fmt.Printf("Start Scan for %d at index %d in %v\n", x, i-1, a)
	} else {
		fmt.Printf("Start Scan for %d at index %d in %v\n", x, i, a)
	}
}