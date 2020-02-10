# Go-Pracitce
The notes of learning GO.

## How to run
- `go run <script_name>.go`
- build:
    - `go build <script_name>.go`
    - `./<script_name>`
- Install package:
    - `go get <url>`

## About GO
- 由 Google 開發的類似 C 語言的一種靜態、編譯，自帶垃圾回收和原生支持併發 (goroutine) 的編程語言。
- 適用於：
    - 高併發網路服務
- 特色：
    - goroutine：可理解為一種虛擬 thread
    - channel：多個 goroutine 進行通信的通道
    - producer [goroutine] - ( channel )[send data] -> customer [goroutine]
    - Go 沒有類！！然而，仍然可以在結構體類型上定義方法。
- 應用：
    - [Docker](https://github.com/docker)
    - [golang](https://github.com/golang/go) 
    - [kubernetes (k8s)](https://github.com/kubernetes/kubernetes)
    - [beego](https://github.com/astaxie/beego)(web framework)
    - [martini](https://github.com/go-martini/martini)(web framework)
    - [gin](https://github.com/gin-gonic/gin)
- [A Tour of GO](https://tour.golang.org/list)

## Basic Knowledge
- [package](https://golang.org/search?q=Package#Global_pkg/cmd/cgo)
- [goenv](https://github.com/syndbg/goenv/blob/master/INSTALL.md)
    - goenv aims to be as simple as possible and follow the already established successful version management model of pyenv and rbenv.
- [govendor](https://github.com/kardianos/govendor)
    - The Vendor Tool for Go
    - Install govendor: `go get -u github.com/kardianos/govendor`
    - `govendor sync`
- [fmt](https://golang.org/pkg/fmt/)(print) 
    - fmt.Printf
        - Print Formatter
        - this function allows you to format numbers, variables and strings into the first string parameter you give it
    - fmt.Println
        - Print Line
        - This cannot format anything, it simply takes a string, prints it and append a newline character, \n
    - fmt.Print
        - Print
        - same thing as Println() however it will NOT append a newline character
- for loop:
    ```
    // self-added operator: <var>++
    for i:= 0; i<10; i++ {
        // code
    }

    // The init and post statements are optional.
    for ; sum < 1000; {
        sum += sum
    }

    // while
    sum := 1 // should be used in func
	for sum < 1000 {
		sum += sum
	}
    ```
- if & else:
    ```
    // the expression need not be surrounded by parentheses ( )
    // but the braces { } are required.
    if condition {
        // code
    }

    // Variables declared inside an if short statement
    // are also available inside any of the else blocks.
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // can't use v here, though
    ``` 
- switch:
    ```
    switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
                fmt.Println("Linux.")
        default:
            // freebsd, openbsd,
            // plan9, windows...
            fmt.Printf("%s.", os)
    }
    ```
- swap:
    ```
    a, b = b, a
    ```
- anonymous variable (匿名變量):
    - no additional memory is allocated  
    ```
    func Get Data() (int, int){
        return 1, 2
    }

    // _ is anonymous variable
    // := can not used out of func
    a, _ := Get Data()
    ```
- variable:
    - format:
        - standard: var <var_name> <var_type> = \<expression\>  
            e.g., ```var money int = 1```
        - lazy: var <var_name> = \<expression\>  
            e.g., ```var money = 1```
        - lazier: <var_name> := \<expression\>  
            e.g., ```money := 1```
            - note: ```:=``` only assign value to the variable which has not been assigned before.
                e.g., ```money := 1``` than ```money := 2``` will raise error

    - type: 
        - int / unit (ASCII):
            - int8 / unit8 (byte)
            - int16 (C: short) / unit16
            - int32 (rune, Unicode) / unit32
            - int64 (C: long) / unit64
            - uintptr
        - string
        - array (fixed size)
            ```
            func main() {
                var a [2]string
                a[0] = "Hello"
                a[1] = "World"
                fmt.Println(a[0], a[1])
                // Hello World
                fmt.Println(a)
                // [Hello World]
            }
            ```
        - [slices](https://gobyexample.com/slices) [] (default: nil):
            - slices are typed only by the elements they contain
            - slices are like references to arrays (does not store any data)
            - use ```make``` to create empty slices   
                - ```
                  b := make([]int, 0, 5) // len(b)=0, cap(b)=5
                  b = b[:cap(b)] // len(b)=5, cap(b)=5
                  b = b[1:]      // len(b)=4, cap(b)=4
                  ```
            - with elements: ```primes := [6]int{2, 3, 5, 7, 11, 13}```
            - Unlike python, you can't use -1 as index ```primes[-1] // error```
        - float:
            - float32 (max: 3.4e38)
            - float64 (max: 1.8e308) 
        - complex:
            - complex64
                - is the set of all complex numbers with float32 real and imaginary parts.
            - complex128
                - is the set of all complex numbers with float64 real and imaginary parts.
        - bool
            - true
            - false
- const:
    - format:
        - const \<const_name\> = \<expression\>
          e.g., ```const World = "世界"```
        - can not use ```:=```
- func ()
    ```
    func swap(x, y string) (string, string) {
        return y, x
    }
    a, b := swap("hello", "world")
    // result: "world hello"
    ```
- struct (結構體)
    - A struct is a collection of fields.
    ```
    type Vertex struct {
        X int
        Y int
    }

    func main() {
        fmt.Println(Vertex{1, 2})
        // {1 2}

        v := Vertex{1, 2}
        v.X = 4
        fmt.Println(v.X)
        // 4 

        // pointer
        p := Vertex{1, 2}
        q := &p
        q.X = 1e2
        fmt.Println(p)
        // {100, 2}

        // pointer
        v := new(Vertex)
        fmt.Println(v)
        // &{0 0}

        v.X, v.Y = 11, 9
        fmt.Println(v)
        // &{11 9} 
    }
    ```
    - Go does not have class-object architecture
    ```
    type Vertex struct {
        X, Y float64
    }

    func (v *Vertex) Scale(f float64) {
        v.X = v.X * f
        v.Y = v.Y * f
    }

    func (v *Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func main() {
        v := &Vertex{3, 4}
        v.Scale(5)
        fmt.Println(v, v.Abs())
    }
    ```
- map
    ```
    type Vertex struct {
        Lat, Long float64
    }

    // map
    var m = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},
        "Google":    {37.42202, -122.08408},
    }

    func main() {
        fmt.Println(m)
        // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

        delete(m, "Google")
        v, ok := m["Google"]
        fmt.Println("The value:", v, "Present?", ok)
        // The value: {0 0} Present? false
    }
    ```
- goroutine
    ```go f(x, y, z)```
- channel
    - Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
    ```
    ch <- v    // Send v to channel ch.
    v := <-ch  // Receive from ch, and assign value to v.
    ```
    - A sender can close a channel to indicate that no more values will be sent.
    ```
    func fibonacci(n int, c chan int) {
        x, y := 0, 1
        for i := 0; i < n; i++ {
            c <- x
            x, y = y, x+y
        }
        close(c)
    }

    func main() {
        c := make(chan int, 10)
        go fibonacci(cap(c), c)
        for i := range c {
            fmt.Println(i)
        }
    }
    ```
- [Pointers](https://tour.golang.org/moretypes/1)
    - A pointer holds the memory address of a value
    - The type *T is a pointer to a T value. Its zero value is nil
    - The & operator generates a pointer to its operand
    - The * operator denotes the pointer's underlying value.
- [Mutable & Immutable Objects](https://stackoverflow.com/questions/8018081/which-types-are-mutable-and-immutable-in-the-google-go-language)
    - Mutable Objects:
        - arrays, slices, maps, channels, closures which are capturing at least 1 variable from the outer scope
    - Immutable Objects: 
        - interfaces, booleans, numeric values (including values of type int), strings, pointers, function pointers, and closures which can be reduced to function pointers, structs having a single field

## More information
- [Simplicity is complicated](https://talks.golang.org/2015/simplicity-is-complicated.slide)
- [Getting to Go: The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote)

## Reference
- [Go語言從入門到進階實戰（視頻教學版）](https://www.books.com.tw/products/CN11547747)
- [Println vs Printf vs Print in Go](https://stackoverflow.com/questions/53879154/println-vs-printf-vs-print-in-go)
- [gobyexample](https://gobyexample.com/)
