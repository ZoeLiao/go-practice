package main

import "fmt"

type Sort interface {
    swap()
}

type bubbleSort struct {
    array []int
}

func (b bubbleSort) swap() {
    for i:=0; i<len(b.array)-1; i++ {
        for j:=1; j<len(b.array); j++ {
            if b.array[i] > b.array[j] {
                b.array[i], b.array[j] = b.array[j], b.array[i]
            }
        }
    }
}

func f(sort Sort){
    fmt.Println("before:", sort)
    sort.swap()
    fmt.Println("after:", sort)
}

func main() {
    array := make([]int, 5)
    fmt.Println("enter marks of 5 subjects")
    for i:=0; i<5; i++ {
        fmt.Scan(&array[i])
    }
    sort := bubbleSort{array: array}
    f(sort)
}
