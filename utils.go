package main

import (
    "fmt"
    "strconv"
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

func ZipF64(a, b []float64) ([][]float64, error) {
    if len(a) != len(b) {
        return nil, fmt.Errorf("zip: arguments must be of same length")
    }

    r := make([][]float64, len(a), len(a))

    for i, e := range a {
        r[i] = []float64{e, b[i]}
    }

    return r, nil
}

func mkSlice(args ...interface{}) []interface{} {
    return args
}

func StrArrToF64Arr(sArr []string) (f64Arr []float64, err error) {
    for _, s := range sArr {
        f64, ferr := strconv.ParseFloat(s, 64)
        if ferr != nil {
            err = ferr
            return
        }
        f64Arr = append(f64Arr, f64)
    }
    return
}
