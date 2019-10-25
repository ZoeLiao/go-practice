# Go-Pracitce
The notes of learning GO.

## How to run
- `go run <script_name>`

## About GO
- 由 Google 開發的類似 C 語言的一種靜態、編譯，自帶垃圾回收和原生支持併發 (goroutine) 的編程語言。
- 適用於：
    - 高併發網路服務
- 特色：
    - goroutine：可理解為一種虛擬 thread
    - channel：多個 goroutine 進行通信的通道
    - producer [goroutine] - ( channel )[send data] -> customer [goroutine]
- 應用：
    - [Docker](https://github.com/docker)
    - [golang](https://github.com/golang/go) 
    - [kubernetes (k8s)](https://github.com/kubernetes/kubernetes)
    - [beego](https://github.com/astaxie/beego)(web framework)
    - [martini](https://github.com/go-martini/martini)(web framework)

## Basic Knowledge
- [package](https://golang.org/search?q=Package#Global_pkg/cmd/cgo)
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
    ```
- if:
    ```
    if condition {
        // code
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
        - int / unit (byte, ASCII):
            - int8 / unit8
            - int16 (C: short) / unit16
            - int32 / unit32
            - int64 (C: long) / unit64
        - rune (Unicode, UTF-8, int32)
        - string
        - [slices](https://gobyexample.com/slices) []:
            - slices are typed only by the elements they contain
            - use ```make``` to create empty slices   
                e.g., ```s := make([]string, 3)```
            - 
        - float:
            - float32 (max: 3.4e38)
            - float64 (max: 1.8e308) 
        - bool
            - true
            - false
        - func ()
        - struct
        - map
        - channel
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

## Reference
- [Go語言從入門到進階實戰（視頻教學版）](https://www.books.com.tw/products/CN11547747)
- [Println vs Printf vs Print in Go](https://stackoverflow.com/questions/53879154/println-vs-printf-vs-print-in-go)
