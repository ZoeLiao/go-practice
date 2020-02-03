package main

import "fmt"

// duck typing
type Sort interface {
    swap()
}

type bubbleSort struct {
    array []int
}

func (b bubbleSort) swap() {
    for i:=0; i<len(b.array)-1; i++ {
        for j:=0; j<len(b.array)-1-i; j++ {
            if b.array[j] > b.array[j+1] {
                b.array[j], b.array[j+1] = b.array[j+1], b.array[j]
            }
        }
        fmt.Println(b.array)
    }
}

func f(sort Sort){
    fmt.Println("before:", sort)
    sort.swap()
    fmt.Println("after:", sort)
}

func main() {
    var n int
    fmt.Println("enter the length of array")
    fmt.Scan(&n)

    array := make([]int, n)
    fmt.Printf("enter %d subjects:\n", n)
    for i:=0; i<n; i++ {
        fmt.Scan(&array[i])
    }

    sort := bubbleSort{array: array}
    fmt.Println("Bubble Sort")
    f(sort)
}
