package main

import (
    // 格式化
    "fmt"
    // 隨機數
    "math/rand"
    "time"
)

func producer(header string, channel chan<-string) {
    // 無線循環
    for {
        channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
        time.Sleep(time.Second)
    }
}

func customer(channel <-chan string) {
    for {
        // 從 channel 取出數據 （阻塞直到 channel 返回數據）
        message := <-channel
        fmt.Println(message)
    }
}

func main() {
    channel := make(chan string)
    go producer("cat", channel)
    go producer("dog", channel)
    customer(channel)
}
