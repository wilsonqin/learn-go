package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func NewSqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    return 0, nil
}

func Run_sqrt_errors() {
    fmt.Println(NewSqrt(2))
    fmt.Println(NewSqrt(-2))
}