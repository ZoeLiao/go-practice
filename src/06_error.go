// reference https://gobyexample.com/errors
package main

import (
    "errors"
    "fmt"
)


type argError struct {
    desc string
    arg int
}

func (e *argError) Error() string {
    return fmt.Sprintf("%s: %d", e.desc, e.arg)
}

func get_remainder(num1 int, num2 int) (int, error) {
    // errors.New constructs a basic error value with the given error message
    if num1 < 0 {
        return -1, errors.New("Error1: can't work with negative number")
    }
    // use custom types as errors by implementing the Error() method on them
    if num2 < 0 {
        return -1,  &argError{"Error2: can't work with negative number", num2}
    }

    return num1 % num2, nil
}


func main() {
    var n1, n2 int
    fmt.Println("enter the dividend:")
    fmt.Scan(&n1)
    fmt.Println("enter the divisor:")
    fmt.Scan(&n2)

    if r, e := get_remainder(n1, n2); e != nil {
        fmt.Println(e)
    } else {
        fmt.Println("The reminder is:", r)
    }
}
