package main

import "fmt"

func main() {
    i, j := 5, 100

    p := &i         // point to i (5) (&: get address)
    fmt.Println(*p) // read i through the pointer (5)
    *p = 10         // set i through the pointer (*: get value)
    fmt.Println(i)  // see the new value of i (10)

    p = &j         // point to j (100)
    *p = *p / 25   // divide j through the pointer (100/25)
    fmt.Println(j) // see the new value of j (4)

    fmt.Printf("p type: %T\n", p)
    fmt.Printf("p address: %p\n", p)
}
