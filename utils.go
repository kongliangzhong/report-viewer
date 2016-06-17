package main

import (
    "fmt"
)

func Zip(a, b []int) ([][]int, error) {
    if len(a) != len(b) {
        return nil, fmt.Errorf("zip: arguments must be of same length")
    }

    r := make([][]int, len(a), len(a))

    for i, e := range a {
        r[i] = []int{e, b[i]}
    }

    return r, nil
}

func mkSlice(args ...interface{}) []interface{} {
    return args
}
